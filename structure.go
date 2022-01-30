// Print directory structure
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/moledoc/walks"
)

// root is a variable to hold current working directory path.
// It is used so that we could access the root (starting directory) from our printer function without passing it as an argument.
var root string

// treeFormat is a boolean value to identify if we want to print directory structure in a tree format or not.
var treeFormat bool

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
	if treeFormat {
		fmt.Println(printed)
		return
	}
	fmt.Println(path)
}

// help is a function to print program help.
func help() {
	fmt.Println("USAGE\n\tstructure [FLAGS]")
	fmt.Println("DESCRIPTION\n\tProgram that prints the structure of current working directory, written in Go.")
	fmt.Println("FLAGS")
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {

	ignore := flag.String("ignore", "\\.git", "REGEXP_PATTERN that we want to ignore.")
	depth := flag.Int("depth", -1, "The depth of directory structure recursion, -1 is exhaustive recursion.")
	helpBool := flag.Bool("help", false, "Print this help.")
	tree := flag.Bool("tree", false, "Format the directory structure in a tree format.")
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if *helpBool {
		help()
	}

	treeFormat = *tree

	rootLoc, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	rootLoc = strings.Replace(rootLoc, "\\", "/", -1)
	root = rootLoc

	// set walks.Ignore
	walks.Ignore = regexp.MustCompile(*ignore)
	// 	walks.SetIgnore(*ignoreLoc)

	fmt.Printf("%s", rootLoc)
	if *tree {
		fmt.Print(":")
	}
	fmt.Print("\n")
	// Walk current working directory recursively.
	walks.WalkLinear(rootLoc, printer, printer, *depth, 0)
}
