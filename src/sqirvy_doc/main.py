import sys

from common.anthropic import doc

def main():
    if len(sys.argv) != 2:
        sys.stderr.write("Usage: sqirvy-doc <filename>\n")
        sys.exit(1)
    md, time, input_tokens, output_tokens = doc(sys.argv[1])
    sys.stderr.write(f"Cached API call time: {time:.2f} seconds\n")
    sys.stderr.write(f"Cached API call input tokens: {input_tokens}\n")
    sys.stderr.write(f"Cached API call output tokens: {output_tokens}\n")
    print(md)

if __name__ == "__main__":
    main()

