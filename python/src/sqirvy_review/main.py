import sys
from pathlib import Path
sys.path.append(str(Path(__file__).parent.parent))

from common.anthropic import review

def main():
    if len(sys.argv) != 2:
        print("Usage: sqirvy-review <filename>", file=sys.stderr)
        sys.exit(1)

    md, time, input_tokens, output_tokens = review(sys.argv[1])
    sys.stderr.write(f"Cached API call time: {time:.2f} seconds\n")
    sys.stderr.write(f"Cached API call input tokens: {input_tokens}\n")
    sys.stderr.write(f"Cached API call output tokens: {output_tokens}\n")
    print(md)

if __name__ == "__main__":
    main()
