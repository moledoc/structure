// Print directory structure
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"gitlab.com/utt_meelis/walks"
)

// help prints info about the usage of this program
func help() {
	fmt.Println("structure [ SUBCOMMAND | FLAG ]")
	fmt.Println("SUBCOMMAND")
	fmt.Printf("  help\n")
	fmt.Printf("  \tPrint this help.\n")
	fmt.Println("FLAG")
	flag.PrintDefaults()
}

// printer prints given path with defined format
func printer(path string) {
	pathLoc := strings.Replace(path, root, "", 1)
	var printed string
	id := 0
	for i, char := range pathLoc {
		if string(char) == "/" {
			id = i
			printed += "| "
		}
	}
	printed += string(pathLoc[id+1:])
	fmt.Println(printed)
}

var root string

func main() {

	ignoreLoc := flag.String("ignore", ".ignore", "File, where each line represents one directory or file that is ignored. If when .ignore exist in the current directory, this flag is not necessary.")
	depth := flag.Int("depth", -1, "The depth of directory structure recursion, -1 is exhaustive recursion.")
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	for _, arg := range os.Args {
		if strings.Contains(arg, "help") {
			help()
			os.Exit(0)
		}
	}

	rootLoc, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	rootLoc = strings.Replace(rootLoc, "\\", "/", -1)
	root = rootLoc

	// set walks.Ignore
	walks.SetIgnore(*ignoreLoc)

	fmt.Printf("%s:\n", rootLoc)
	// Walk current working directory recursively.
	walks.WalkLinear(rootLoc, printer, printer, *depth, 0)
}
