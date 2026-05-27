// Copyright 2015 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package xstrings

// ExpandTabs can expand tabs ('\t') rune in str to one or more spaces dpending on
// current column and tabSize.
// The column number is reset to zero after each newline ('\n') occurring in the str.
//
// ExpandTabs uses RuneWidth to decide rune's width.
// For example, CJK characters will be treated as two characters.
//
// If tabSize <= 0, ExpandTabs panics with error.
//
// Samples:
//
//	ExpandTabs("a\tbc\tdef\tghij\tk", 4) => "a   bc  def ghij    k"
//	ExpandTabs("abcdefg\thij\nk\tl", 4)  => "abcdefg hij\nk   l"
//	ExpandTabs("z中\t文\tw", 4)           => "z中 文  w"
func ExpandTabs(str string, tabSize int) string { _ = "STUB: not implemented"; return "" }

// LeftJustify returns a string with pad string at right side if str's rune length is smaller than length.
// If str's rune length is larger than length, str itself will be returned.
//
// If pad is an empty string, str will be returned.
//
// Samples:
//
//	LeftJustify("hello", 4, " ")    => "hello"
//	LeftJustify("hello", 10, " ")   => "hello     "
//	LeftJustify("hello", 10, "123") => "hello12312"
func LeftJustify(str string, length int, pad string) string { _ = "STUB: not implemented"; return "" }

// RightJustify returns a string with pad string at left side if str's rune length is smaller than length.
// If str's rune length is larger than length, str itself will be returned.
//
// If pad is an empty string, str will be returned.
//
// Samples:
//
//	RightJustify("hello", 4, " ")    => "hello"
//	RightJustify("hello", 10, " ")   => "     hello"
//	RightJustify("hello", 10, "123") => "12312hello"
func RightJustify(str string, length int, pad string) string { _ = "STUB: not implemented"; return "" }

// Center returns a string with pad string at both side if str's rune length is smaller than length.
// If str's rune length is larger than length, str itself will be returned.
//
// If pad is an empty string, str will be returned.
//
// Samples:
//
//	Center("hello", 4, " ")    => "hello"
//	Center("hello", 10, " ")   => "  hello   "
//	Center("hello", 10, "123") => "12hello123"
func Center(str string, length int, pad string) string { _ = "STUB: not implemented"; return "" }

func writePadString(output *stringBuilder, pad string, padLen, remains int) {
	_ = "STUB: not implemented"
	return
}
