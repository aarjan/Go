// Sample program to give a brief introduction of regular expressions

package main

import (
	"fmt"
	"regexp"
)

var r *regexp.Regexp
var regexpExpr []string = []string{

	`This `,   // The expression must be regularly matched in every character to return true
	`This my`, // Check this out

	// Character Classes

	`Th\ws`, // Represents any character from the class [A-Za-z0-9], mnemonic: 'word'

	`\d`, // Represents any numeric digit

	`\t`, // Represents TAB, other whitespaces: SPACE,CR,LF. Or more precisely [\t\n\r\f]

	`\S`, // Matches any non-white character; Character classes can be negated by using the uppercase '\W','\D','\S'. Thus '\D' is any character that is not a '\d'

	`\W`, // Checks if a string has anything that is not a word-char; eg: "hyphen" ;'-'

	`.st`, // Dot '.' matches any character

	`\\`, // Finding one backslash '\', or "\\\\"

	// It must be escaped twice in the regex and once in the string.

	`\.`, // Finding a literal dot, or `\.`

	`\$`, // Finding a literal dollar symbol
}

func main() {

	fmt.Println("Strings to be  matched:", "\\ This is my 1st regular expression exercise.")

	for _, expr := range regexpExpr {
		var err error
		r, err = regexp.Compile(expr)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("expression: %s ; matched:%v ; string:%s\n", expr, MatchString(), FindString())

	}

}

// Returns true if match is found
func MatchString() bool {
	return r.MatchString("\\ This is my 1st regular expression exercise.")
}

// Returns string after the first match
// If it doesn't finds the string that matches the regular expressions, it will return the empty string that might be a vaild match.
func FindString() string {
	return r.FindString("\\ This is my 1st regular expresssion exercise.")
}
