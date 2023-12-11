# go-reloaded

This tool is designed to perform various modifications to a given text based on a set of rules. The tool receives the name of an input file containing text that requires modifications and the name of an output file to store the modified text. If an output file does not exist then it will create by given name. 

## Usage

```bash
cd cmd
```

```bash
go run . input.txt result.txt
```

Replace `input.txt` with the name of your input file and `output.txt` with the desired output file.

## Argument filepath rules

1. **Valid files names**
   - The program accepts files included in this directory and outside by giving it the file path.
   - The program processes only .txt files. Otherwise it returns error.

   - The first file has to exist. Errors if not.
   - If the second file does not exist, then it can be created. So you can name the second file as you wish ;)

## Flag Rules

1. **Flag Introduction:**
   - A flag is a key phrase used to modify the text. Flags are found in the text and only valid flags are processed, invalid flags are treated as regular text
   - Flags:
     1. `(hex)`
     2. `(bin)`
     3. `(up)`
     4. `(low)`
     5. `(cap)`
  
2. **Flag Application:**
   - The modification starts from the closest word preceding the flag. Flag cancels out if there is NO WORD
     - Example: `"4F  ih(hex) files were added"` → `"30 files were added"`
   - If flag is not applied then it will be treated as regular text
   - When using invalid number with `(hex)` then it is treated as regular text
   - When using invalid number with `(bin)` then it is treated as regular text
   - Flags can be wrapped in multiple parentheses.
   - A word is a single unit of language that has meaning and can be spoken

   - Flags remove the preceding white spaces. Following white spaces are not removed

3. **Specific Flags with Second Argument:**
   - Specific flags, namely `(up)`, `(cap)`, and `(low)`, have the second argument.
   - The second argument must be a valid positive number.
     - Example: `"asdf files were (cap, 3) added"` → `"Asdf Files Were added"`

4. **Invalid flags:**
   - Flags cannot contain another flag.
   - Flags cannot contain ANY unnecessary symbols such as `.` `!` `\n` etc except SPACE ` `.

      - Otherwise FLAG IS TREATED AS REGULAR TEXT

These rules serve as a guide for creating the right input file. Ensure adherence to these rules for consistent and accurate text transformations.