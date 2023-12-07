# go-reloader

This tool is designed to perform various modifications to a given text based on a set of rules. The tool receives the name of an input file containing text that requires modifications and the name of an output file to store the modified text.

## Usage

```bash
go run . input.txt output.txt
```

Replace `input.txt` with the name of your input file and `output.txt` with the desired output file.

## Flag Rules

1. **Flag Introduction:**
   - A flag is a key phrase used to modify the text.
   - Flags:
     1. `(hex)`
     2. `(bin)`
     3. `(up)`
     4. `(low)`
     5. `(cap)`
  
2. **Flag Application:**
   - The modification starts from the closest word preceding the flag.
     - Example: `"1E (hex) files were added"` → `"30 files were added"`

   - A flag can be valid only if it is separated by white spaces from both sides. Otherwise, the flag is invalid, and the program treats it as regular text.
     - Valid example: `"1E (hex) files were added"` → `"30 files were added"`
     - Invalid examples:
        - `"1E(hex)files were added"` → `"1E(hex)files were added"`
        - `"1E(hex) files were added"` → `"1E(hex) files were added"`
        - `"1E (hex)files were added"` → `"1E (hex)files were added"`
  
   - If perceding symbol is `\n` then flag won't change anything.

3. **Specific Flags with Second Argument:**
   - Specific flags, namely `(up)`, `(cap)`, and `(low)`, have a second argument.
   - The second argument must be a valid positive number.
     - Example: `"asdf files were (cap, 3) added"` → `"Asdf Files Were added"`

4. **Flag Composition:**
   - Flags cannot contain another flag.
   - Flags cannot contain ANY unnecessary symbols such as `,` `!` `\n` etc.

5. **Flag Wrapping:**
   - Flags can be wrapped in multiple parentheses.

These rules serve as a guide for creating a program that performs text modifications based on specified flags. Ensure adherence to these rules for consistent and accurate text transformations.

## Example

Consider the following command:

```bash
go run . input.txt output.txt
```