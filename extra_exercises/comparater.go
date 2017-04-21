// Sample program to compare files between two or more directory
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	args := os.Args[1:]
	dirList := make([][]string, 0)
	for _, arg := range args {
		dir := list(arg)
		dirList = append(dirList, dir[:])
	}
	compare(dirList[0], dirList[1])
}

func compare(dirA, dirB []string) {
	l := len(dirA)
	longest := ""
	if len(dirA) > len(dirB) {
		longest = "A"
		l = len(dirB)
	}


	for i := 0; i < len(dirA); i++ {
		for j := 0; j < len(dirB); j++ {
			if dirA[i] == dirB[j] {
				fmt.Printf("Matched: \t%s", dirA[i])
			}
		}
	}
	
	for i := 0; i < l; i++ {
		
		fmt.Printf("%s\t\t\t%s\n", dirA[i], dirB[i])
		
	}

	fn := func(dir []string) {
		for _, file := range dir[l:] {
			fmt.Printf("\t\t\t%s\n", file)
		}
	}
	if longest == "A" {
		fn(dirA)
	} else {
		fn(dirB)
	}
}

func list(fName string) []string {

	cmd := exec.Command("find", fName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	dirs := strings.Fields(string(out))

	r := regexp.MustCompile(`\w+\..*`)

	files := make([]string, 0)

	for i, f := range dirs {
		if i == 0 {
			files = append(files, f)
		}
		if file := r.FindString(f); len(file) > 0 {
			files = append(files, file)
		}
	}
	return files

}
