import os
import socket
import threading
from datetime import datetime
from http import HTTPStatus
from urllib.parse import unquote

class WebServer:
    def __init__(self, host='localhost', port=8080, root_dir='./public'):
        self.host = host
        self.port = port
        self.root_dir = os.path.abspath(root_dir)
        self.socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.socket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        self.mime_types = {
            '.html': 'text/html',
            '.css': 'text/css',
            '.js': 'text/javascript',
            '.jpg': 'image/jpeg',
            '.jpeg': 'image/jpeg',
            '.png': 'image/png',
            '.gif': 'image/gif',
            '.txt': 'text/plain'
        }

    def start(self):
        self.socket.bind((self.host, self.port))
        self.socket.listen(5)
        print(f"Server started at http://{self.host}:{self.port}")
        
        while True:
            client_socket, client_address = self.socket.accept()
            thread = threading.Thread(target=self.handle_client,
                                   args=(client_socket, client_address))
            thread.start()

    def handle_client(self, client_socket, client_address):
        try:
            request = client_socket.recv(1024).decode('utf-8')
            if not request:
                return

            request_line = request.split('\n')[0]
            method, path, _ = request_line.split()
            
            if method != 'GET':
                self.send_error(client_socket, HTTPStatus.METHOD_NOT_ALLOWED)
                return

            file_path = self.get_file_path(path)
            if not os.path.exists(file_path):
                self.send_error(client_socket, HTTPStatus.NOT_FOUND)
                return

            if not self.is_file_accessible(file_path):
                self.send_error(client_socket, HTTPStatus.FORBIDDEN)
                return

            self.serve_file(client_socket, file_path)

        except Exception as e:
            print(f"Error handling client {client_address}: {e}")
            self.send_error(client_socket, HTTPStatus.INTERNAL_SERVER_ERROR)
        
        finally:
            client_socket.close()

    def get_file_path(self, path):
        path = unquote(path)
        if path == '/':
            path = '/index.html'
        return os.path.join(self.root_dir, path.lstrip('/'))

    def is_file_accessible(self, file_path):
        return os.path.commonpath([file_path]) == os.path.commonpath([self.root_dir])

    def get_content_type(self, file_path):
        _, ext = os.path.splitext(file_path)
        return self.mime_types.get(ext.lower(), 'application/octet-stream')

    def serve_file(self, client_socket, file_path):
        try:
            file_size = os.path.getsize(file_path)
            content_type = self.get_content_type(file_path)
            
            headers = [
                'HTTP/1.1 200 OK',
                f'Date: {datetime.utcnow().strftime("%a, %d %b %Y %H:%M:%S GMT")}',
                'Server: Python Web Server',
                f'Content-Length: {file_size}',
                f'Content-Type: {content_type}',
                'Connection: close',
                '',
                ''
            ]
            
            client_socket.send('\r\n'.join(headers).encode())

            with open(file_path, 'rb') as f:
                while True:
                    data = f.read(8192)
                    if not data:
                        