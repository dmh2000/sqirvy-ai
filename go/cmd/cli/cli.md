# Password Generator CLI

A command-line tool for generating secure passwords with customizable options.

## Features

- Generates passwords with a minimum length of 24 characters
- Optional symbol inclusion
- Ability to generate multiple passwords at once
- Both short and long-form command flags
- Formatted output for better readability

## Usage

```bash
password-generator [-l length] [-c count] [-s] [-h]
```

### Flags

- `-l, --length`: Password length (minimum 24 characters, default: 24)
- `-c, --count`: Number of passwords to generate (default: 1)
- `-s, --symbols`: Include special symbols in the password
- `-h, --help`: Display help information

## Output

The program outputs each generated password in two formats:
1. Plain format
2. Formatted version for improved readability

Each password is separated by a line of dashes for clear visual distinction.

## Example

```bash
$ password-generator -l 24 -s -c 1

----------------------------------------
[Generated password appears here]
[Formatted version appears here]
```

## Notes

- The program enforces a minimum password length of 24 characters for security
- If both short and long flags are provided, the long flag takes precedence
- Errors during password generation or formatting will cause the program to exit with an error message
