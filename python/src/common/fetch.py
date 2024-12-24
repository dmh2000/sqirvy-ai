import sys

def fetch_code(fname):
    # read the file
    try :
        with open(fname) as fd:
            code = fd.read()
            # break into lines
            lines = (line.strip() for line in code.splitlines())
            # break into chunks
            chunks = (phrase.strip() for line in lines for phrase in line.split(" . "))
            # filter out empty chunks
            text = '\n'.join(chunk for chunk in chunks if chunk)
    except FileNotFoundError:
        print(f"File {fname} not found", file=sys.stderr)
        sys.exit(1)
    return text

