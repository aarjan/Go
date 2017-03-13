package main

import (
	"fmt"
	"regexp"
)

func main() {
	/*
		Compile is the heart of the regexp-package.
		Every regular expression must be prepared with Compile or its sister-function MustCompile.
		The MustCompile-function behaves almost like Compile, but throws a panic if the regular expression cannot be compiled.
		Because any error in MustCompile leads to a panic, there is no need for returning an error code as second return value.
		But you should avoid the repeated compilation of a regular expression in a loop for performance reasons.
	*/
	s := "Enie  meenie miny moe"

	fmt.Println("Special repetition")

	// A word is a sequence of character of type '\w' ; 'space is also a character'
	r := regexp.MustCompile(`\w+`) // '+' symbol signifies a repetition
	// In '\w+' spaces '\s' are inclusive; but while explicitly using '\s', the regexp follows the specified pattern

	/*
		FindAllString() function returns all the strings that matches the case
		Providing a non negative arguments will limit the repetition
	*/
	fmt.Printf("%v\n", r.FindAllString(s, 2))

	/*
		FindStringIndex() returns an array with two enries.
		The first entry is the index (starting from 0, of course) where the regular expression matched.
		The second is the index in front of which the regexp ended.
	*/
	r = regexp.MustCompile(`\w+\s\w`)
	fmt.Printf("%v \n%v\n", r.FindAllString(s, -1), r.FindAllStringIndex(s, -1))

	/*
		But if there is a user applied input, there might be two spaces
		Its better to use '+'in practise, because of the uncertainty of user input
	*/
	r = regexp.MustCompile(`\w+\s+\w+`)
	fmt.Printf("%v \n%v\n", r.FindAllString(s, -2), r.FindAllStringIndex(s, -1))

	fmt.Println()
	fmt.Println("Anchor and Boundaries")

	// The caret symbol '^' denotes begin-of-line
	r = regexp.MustCompile(`^K`)
	fmt.Printf("%v\n", r.MatchString(s)) // Do we have 'K' in the beginning

	//The dollar symbol denotes 'end-of-line'
	r = regexp.MustCompile(`miny$`)
	fmt.Printf("%v\n", r.MatchString(s)) // false, not end of line

	// You can find a word boundry with '\b'.
	s = "How much wood would a woodchunk chuck in a Hollywood?"
	//   012345678901234567890123456789012345678901234567890
	//             10        20        30        40        50
	//            -1--         -2--                    -3--

	// Find words that start with "wood"
	r = regexp.MustCompile(`\bwood`) //	1	2
	fmt.Printf("%v\n", r.FindAllStringIndex(s, -1))

	// Find words that end with "wood"
	r = regexp.MustCompile(`wood\b`) //	1	3
	fmt.Printf("%v\n", r.FindAllStringIndex(s, -1))

	// Find words that start and end with "word"
	r = regexp.MustCompile(`\bwood\b`) //	1
	fmt.Printf("%v\n", r.FindAllStringIndex(s, -1))

	fmt.Println()
	fmt.Println("Character classes")
	/*
		Instead of literal character you can require a set(or a class) of characters in any location.
		Any of the characters in the square brackets will satisfy the regexp.
	*/

	r = regexp.MustCompile("H[ui]llo") //	match "Hullo" or "Hillo"
	fmt.Println(r.FindString("Hello regular expression. Hullo again"))

	// A negated class will reverses match of the class.
	r = regexp.MustCompile("H[^ui]llo") //	match except "Hullo" or "Hillo"
	fmt.Println(r.FindAllString("Hello regular expression. Hullo & Hillo again", -1))

	fmt.Println()
	fmt.Println("Alternatives")
	/*
		You can provide alternatives using the pipe-symbol '|' to allow two (or more) different possible matches.
		If you want to allow alternatives only in parts of the regular expression,
		you can use parentheses for grouping.
	*/

	t, _ := regexp.Compile(`Santa Clara|Santa Barbara`)
	s = "Clara was from Santa Barbara and Barbara was from Santa Clara"
	fmt.Printf("%v\n", t.FindAllStringIndex(s, -1))

	u, _ := regexp.Compile(`Santa (Clara|Barbara)`) // Equivalent
	v := "Clara was from Santa Barbara and Barbara was from Santa Clara"
	fmt.Printf("%v\n", u.FindAllStringIndex(v, -1))

}
