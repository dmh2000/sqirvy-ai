import anthropic
import time
import sys
from sqirvy_ai.fetch import fetch_code

# create a new claude instance
# claude = anthropic.Claude(max_tokens_to_sample=1000, stop_sequences=["\n"])
client = anthropic.Anthropic()
MODEL_NAME = "claude-3-5-sonnet-20241022"
PROMPT_CACHE = "prompt-caching-2024-07-31"
TOKEN_LIMIT = 2048

def _review_message(code):
    """Create a message for reviewing a code file"""
    msg = [
        {
            "role": "user",
            "content": [
                {
                    "type": "text",
                    "text": code,
                    "cache_control": {"type": "ephemeral"}
                },
                {
                    "type": "text",
                    "text": """
                    Review this code for bugs, security issues, style and idiomatic structure.
                    Suggest improvements to the code.
                    Output is markdown.
                    """
                }
            ]
        }
    ]
    return msg

def _doc_message(code):
    """Create a message for documenting a code file"""
    msg = [
        {
            "role": "user",
            "content": [
                {
                    "type": "text",
                    "text": code,
                    "cache_control": {"type": "ephemeral"}
                },
                {
                    "type": "text",
                    "text": """
                    Analyze this program and write a description that would be suitable for a README file. 
                    Output is markdown. 
                    """
                }
            ]
        }
    ]
    return msg

def _call_claude(text,message):
    """call the Anthropic Claude API"""
    start_time = time.time()
    response = client.messages.create(
        model=MODEL_NAME,
        max_tokens=TOKEN_LIMIT,
        messages=message,
        extra_headers={"anthropic-beta": PROMPT_CACHE}
    )
    end_time = time.time()

    return response, end_time - start_time

def _execute(fname, message):
    """Execute for a given filename and message"""

    text = fetch_code(fname)

    cached_response, cached_time = _call_claude(text,message)

    if len(cached_response.content) == 0:
        sys.stderr.write("No code found in response")
        sys.exit(1)

    doc_text = cached_response.content[0].text
    return doc_text, cached_time, cached_response.usage.input_tokens, cached_response.usage.output_tokens

def doc(fname):
    """Document a code file"""

    try :
        text = fetch_code(fname)
        msg = _doc_message(text)
        return _execute(fname,msg)
    except FileNotFoundError:
        raise

def review(fname):
    """Review a code file"""

    try :
        text = fetch_code(fname)
        msg = _review_message(text)
        return _execute(fname,msg)
    except FileNotFoundError:
        raise

