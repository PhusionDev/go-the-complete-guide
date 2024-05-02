# 2 - Go Essentials

## Main Package/Function

- The `main()` function tells the compiler the entry point to the application
- Not all packages require a `main()` function
  - Libraries are a good example
- Cannot have more than 1x `func main() {}` within a main package

## String Formatting

### Formatting Text

- `fmt.Print(<string>, <variable>)`
- `fmt.Printf(<string with %>, <vars>)`

### Formatting Integers

- `%b` base 2
- `%c` the character represented by the corresponding Unicode code point
- `%d` base 10
- `%o` base 8
- `%O` base 8 with 0o prefix
- `%q` a single-quoted character literal safely escaped with Go syntax
- `%x` base 16, with lower-case letters for a-f
- `%X` base 16, with upper-case letters for A-F
- `%U` Unicode format: U+1234; same as "U+%04X"

### Formatting Booleans

- `%t` the word true or false
