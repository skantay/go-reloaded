# go-reloader

This tool is designed to perform various modifications to a given text based on a set of rules. The tool receives the name of an input file containing text that requires modifications and the name of an output file to store the modified text.

## Usage

```bash
go run . input.txt output.txt
```

Replace `input.txt` with the name of your input file and `output.txt` with the desired output file.

## Modifications

1. **Hexadecimal Conversion:**
   - **Requirement:** The word preceding `(hex)` should be a valid hexadecimal number.
     - Example: `"1E (hex) files were added"` -> `"30 files were added"`

2. **Binary Conversion:**
   - **Requirement:** The word preceding `(bin)` should be a valid binary number.
     - Example: `"It has been 10 (bin) years"` -> `"It has been 2 years"`

3. **Uppercase Conversion:**
   - **Requirement:** The word preceding `(up)` should be a valid word.
     - Example: `"Ready, set, g!o (up)!"` -> `"Ready, set, GO!"`

4. **Lowercase Conversion:**
   - **Requirement:** The word preceding `(low)` should be a valid word.
     - Example: `"I should stop SHOUTING (low)"` -> `"I should stop shouting"`

5. **Capitalize Conversion:**
   - **Requirement:** The word preceding `(cap)` should be a valid word.
     - Example: `"Welcome to the Brooklyn bridge (cap)"` -> `"Welcome to the Brooklyn Bridge"`
   - **Additional Requirement:** If a number is specified, it should be a positive integer, and it applies to the specified number of words.
     - Example: `"This is so exciting (up, 2)"` -> `"This is SO EXCITING"`

6. **Punctuation Formatting:**
   - **Requirement:** Ensure that punctuation marks (., ,, !, ?, :, ;) are correctly spaced.
     - Example: `"I was sitting over there ,and then BAMM !!"` -> `"I was sitting over there, and then BAMM!!"`
   - **Additional Requirement:** Handle groups of punctuation like ... or !?.
     - Example: `"I was thinking ... You were right"` -> `"I was thinking... You were right"`

7. **Single Quote Formatting:**
   - **Requirement:** Ensure correct formatting of single quotes (').
     - Example: `"I am exactly how they describe me: ' awesome '"` -> `"I am exactly how they describe me: 'awesome'"`
   - **Additional Requirement:** Handle cases with more than one word between the two single quotes.
     - Example: `"As Elton John said: ' I am the most well-known homosexual in the world '"` -> `"As Elton John said: 'I am the most well-known homosexual in the world'"`

8. **Indefinite Article Correction:**


## Example

Consider the following command:

```bash
go run . input.txt output.txt
```