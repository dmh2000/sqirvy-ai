import os
import re
import sys
import argparse


def extract_files(content, output_dir):
    """
    Reads an input, splits it into sections based on a delimiter,
    and writes each section to a separate file under the specified output directory.
    """
    # Split the content by the delimiter
    sections = re.split(r"\*{8}(.*)\n(.*)", content)

    # Iterate through the sections and write to separate files
    i = 1
    while i < len(sections):
        file_name = sections[i].strip()
        file_content = sections[i + 1].strip()
        print(f"file_name: {file_name} file_content: {file_content}")

        # Create the directory if it doesn't exist
        output_path = os.path.join(output_dir, file_name)
        print(f"output_path: {output_path}")
        output_file_dir = os.path.dirname(output_path)
        print(f"output_dir: {output_file_dir}")
        if output_file_dir and not os.path.exists(output_file_dir):
            os.makedirs(output_file_dir, exist_ok=True)

        # Create the file
        with open(output_path, "w") as outfile:
            outfile.write(file_content)

        print(f"Created file: {output_path}")
        i += 3


if __name__ == "__main__":
    parser = argparse.ArgumentParser(
        description="Extract files from a single input stream."
    )
    parser.add_argument(
        "output_dir", help="The directory to output the extracted files to."
    )
    args = parser.parse_args()

    print(args)
    content = sys.stdin.read()
    print(f"content: {content}")
    extract_files(content, args.output_dir)
