# structure

Program that prints the structure of current working directory, written in Go.

## Quick start

* Use build binaries from bin/ (seek.exe for windows, seek for linux).
* Build it yourself:
  * **Dependency**: need to have Go installed
  * go build -o seek seek.go
  * **[optional]** install the binary to local user with
  ```sh
  go install .
  ```

To build binary (64-bit) from linux to windows, run

```console
GOOS=windows GOARCH=amd64 go build -o seek.exe seek.go
```

## Usage

**structure [FLAGS]**

### FLAGS

**-help**

    Print this help.

**-ignore**=PATTERN

    REGEXP_PATTERN that we want to ignore. (default "\\.git")

**-depth**=int

    The depth of directory structure recursion, -1 is exhaustive recursion. (default -1)

**-tree**

    Format the directory structure in a tree format.


## Examples

* Print help for program `structure`
    ```sh
    structure --help
    ```
* Print current working directory structure to stdout
    ```sh
    structure
    ```
* Print current working directory structure to stdout, ignoring directories and files specified in pattern \<ignore>.
    ```sh
    structure -ignore <ignore>
    structure -ignore=<ignore>
    structure --ignore <ignore>
    structure --ignore=<ignore>
    ```
* Print current working directory structure to stdout, until subdirectory depth is 2 (including).
    ```sh
    structure -depth=2
    structure --depth=2
    ```
* Print the current working directory structure in a tree format
    ```sh
    structure --tree
    ```

### Example output

```
/path/to/directory
/path/to/directory/.ignore
/path/to/directory/README.md
/path/to/directory/bin
/path/to/directory/bin/structure
/path/to/directory/bin/structure.exe
/path/to/directory/go.mod
/path/to/directory/go.sum
/path/to/directory/structure.go
```

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

Meelis Utt
