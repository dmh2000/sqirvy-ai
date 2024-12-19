import anthropic
import time
import sys

# create a new claude instance
# claude = anthropic.Claude(max_tokens_to_sample=1000, stop_sequences=["\n"])
client = anthropic.Anthropic()
MODEL_NAME = "claude-3-5-sonnet-20241022"

def fetch_code(fd):
    # read the file
    code = fd.read()
    if fd != sys.stdin:
        fd.close()
    # break into lines
    lines = (line.strip() for line in code.splitlines())
    # break into chunks
    chunks = (phrase.strip() for line in lines for phrase in line.split(" . "))
    # filter out empty chunks
    text = '\n'.join(chunk for chunk in chunks if chunk)
    return text


def fetch_messages(code):
    msg = [
        {
            "role": "user",
            "content": [
                {
                    "type": "text",
                    "text": "<code>" + code + "</code>",
                    "cache_control": {"type": "ephemeral"}
                },
                {
                    "type": "text",
                    "text": """
                    Analyze this program and write a description that would be suitable for a README file. 
                    Output is markdown. 
                    Limit the output to less than 1024 tokens.
                    """
                }
            ]
        }
    ]
    return msg

def make_non_cached_api_call(text):
    messages = fetch_messages(text)

    start_time = time.time()
    response = client.messages.create(
        model=MODEL_NAME,
        max_tokens=2048,
        messages=messages,
        extra_headers={"anthropic-beta": "prompt-caching-2024-07-31"}

    )
    end_time = time.time()

    return response, end_time - start_time


def make_cached_api_call(text):
    messages = fetch_messages(text)

    start_time = time.time()
    response = client.messages.create(
        model=MODEL_NAME,
        max_tokens=2048,
        messages=messages,
        extra_headers={"anthropic-beta": "prompt-caching-2024-07-31"}
    )
    end_time = time.time()

    return response, end_time - start_time

