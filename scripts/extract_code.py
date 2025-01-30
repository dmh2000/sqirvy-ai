import re
import os
import uuid
import sys
from pathlib import Path

def extract_code_blocks(markdown_file):
    """Extract code blocks from markdown file and save them as separate files."""
    
    # Read the markdown content
    with open(markdown_file, 'r') as f:
        content = f.read()
    
    # Regular expression to match code blocks with optional language specification
    # Matches: ```language\ncode\n``` or ```\ncode\n```
    pattern = r'```(\w+)?\n(.*?)\n```'
    matches = re.finditer(pattern, content, re.DOTALL)
    
    # Create output directory based on markdown filename
    output_dir = Path(markdown_file).stem
    os.makedirs(output_dir, exist_ok=True)
    
    # Process each code block
    for match in matches:
        # Extract language and code content
        language = match.group(1)
        code = match.group(2)
        
        # Determine filename
        if language:
            # Use appropriate extension based on language
            extensions = {
                'javascript': '.js',
                'html': '.html',
                'css': '.css',
                'json': '.json',
                'python': '.py',
                # Add more mappings as needed
            }
            ext = extensions.get(language, f'.{language}')
            
            # Look for filename hints in the code (e.g., comments with filenames)
            filename_hint = re.search(r'[/\s]*([\w.-]+\.' + language + ')', code)
            if filename_hint:
                filename = filename_hint.group(1)
            else:
                filename = f'file{ext}'
        else:
            # Generate random filename for blocks without language specification
            filename = f'file_{uuid.uuid4().hex[:8]}.txt'
        
        # Create full file path
        filepath = os.path.join(output_dir, filename)
        
        # If file already exists, add number to filename
        base, ext = os.path.splitext(filepath)
        counter = 1
        while os.path.exists(filepath):
            filepath = f'{base}_{counter}{ext}'
            counter += 1
        
        # Write code to file
        with open(filepath, 'w') as f:
            f.write(code.strip())
        
        print(f'Created: {filepath}')

if __name__ == '__main__':
    if len(sys.argv) != 2:
        print('Usage: python extract_code.py <markdown_file>')
        sys.exit(1)
    
    markdown_file = sys.argv[1]
    if not os.path.exists(markdown_file):
        print(f'Error: File {markdown_file} not found')
        sys.exit(1)
    
    extract_code_blocks(markdown_file)
