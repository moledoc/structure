# structure

Program that prints the structure of current working directory written in Go.

## Usage

**structure [ SUBCOMMAND | FLAG ]**

### SUBCOMMAND

**help**

    Print this help.

### FLAG
**-ignore string**

    File, where each line represents one directory or file that is ignored. If when .adoignore exist in the current directory, this flag is not necessary. (default ".adoignore")

**-depth int**

    The depth of directory structure recursion, -1 is exhaustive recursion. (default -1)


## Examples

* Print help for program `structure`
    ```sh
    structure help
    ```
* Print current working directory structure to stdout
    ```sh
    structure
    ```
* Print current working directory structure to stdout, ignoring directories and files specified in \<ignore> file.
    ```sh
    structure -ignore <ignore>
    structure -ignore=<ignore>
    structure --ignore <ignore>
    structure --ignore=<ignore>
    ```
- Print current working directory structure to stdout, until subdirectory depth is 2 (including).
    ```sh
    structure -depth=2
    structure --depth=2
    ```

### Example output

```
/path/to/directory:
|.ignore
| README.md
| bin
| | structure
| | structure.exe
| go.mod
| go.sum
| structure.go
```

## Author

Written by
Meelis Utt