import sys

from builder import make_cached_api_call, fetch_code


def main():

    text = fetch_code(sys.stdin)

    cached_response, cached_time = make_cached_api_call(text)

    sys.stderr.write(f"Cached API call time: {cached_time:.2f} seconds\n")
    sys.stderr.write(f"Cached API call input tokens: {cached_response.usage.input_tokens}\n")
    sys.stderr.write(f"Cached API call output tokens: {cached_response.usage.output_tokens}\n")

    if len(cached_response.content) == 0:
        sys.stderr.write("No code found in response")
        sys.exit(1)

    code = cached_response.content[0].text
    print(code)


if __name__ == "__main__":
    main()

