// Copyright 2015 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package xstrings

import (
	"math/rand"
)

// ToCamelCase is to convert words separated by space, underscore and hyphen to camel case.
//
// Some samples.
//
//	"some_words"      => "someWords"
//	"http_server"     => "httpServer"
//	"no_https"        => "noHttps"
//	"_complex__case_" => "_complex_Case_"
//	"some words"      => "someWords"
//	"GOLANG_IS_GREAT" => "golangIsGreat"
func ToCamelCase(str string) string { _ = "STUB: not implemented"; return "" }

// ToPascalCase is to convert words separated by space, underscore and hyphen to pascal case.
//
// Some samples.
//
//	"some_words"      => "SomeWords"
//	"http_server"     => "HttpServer"
//	"no_https"        => "NoHttps"
//	"_complex__case_" => "_Complex_Case_"
//	"some words"      => "SomeWords"
//	"GOLANG_IS_GREAT" => "GolangIsGreat"
func ToPascalCase(str string) string { _ = "STUB: not implemented"; return "" }

func toCamelCase(str string, isBig bool) string { _ = "STUB: not implemented"; return "" }

// leading connector will appear in output.

// A special case for a string contains only 1 rune.

// ToSnakeCase can convert all upper case characters in a string to
// snake case format.
//
// Some samples.
//
//	"FirstName"    => "first_name"
//	"HTTPServer"   => "http_server"
//	"NoHTTPS"      => "no_https"
//	"GO_PATH"      => "go_path"
//	"GO PATH"      => "go_path"  // space is converted to underscore.
//	"GO-PATH"      => "go_path"  // hyphen is converted to underscore.
//	"http2xx"      => "http_2xx" // insert an underscore before a number and after an alphabet.
//	"HTTP20xOK"    => "http_20x_ok"
//	"Duration2m3s" => "duration_2m3s"
//	"Bld4Floor3rd" => "bld4_floor_3rd"
func ToSnakeCase(str string) string { _ = "STUB: not implemented"; return "" }

// ToKebabCase can convert all upper case characters in a string to
// kebab case format.
//
// Some samples.
//
//	"FirstName"    => "first-name"
//	"HTTPServer"   => "http-server"
//	"NoHTTPS"      => "no-https"
//	"GO_PATH"      => "go-path"
//	"GO PATH"      => "go-path"  // space is converted to '-'.
//	"GO-PATH"      => "go-path"  // hyphen is converted to '-'.
//	"http2xx"      => "http-2xx" // insert an underscore before a number and after an alphabet.
//	"HTTP20xOK"    => "http-20x-ok"
//	"Duration2m3s" => "duration-2m3s"
//	"Bld4Floor3rd" => "bld4-floor-3rd"
func ToKebabCase(str string) string { _ = "STUB: not implemented"; return "" }

func camelCaseToLowerCase(str string, connector rune) string { _ = "STUB: not implemented"; return "" }

// nothing.

// consider number as a part of previous word.
// e.g. "Bld4Floor" => "bld4_floor"

// if there are some lower case letters following a number,
// add connector before the number.
// e.g. "HTTP2xx" => "http_2xx"

func isConnector(r rune) bool { _ = "STUB: not implemented"; return false }

type wordType int

const (
	invalidWord wordType = iota
	numberWord
	upperCaseWord
	alphabetWord
	connectorWord
	punctWord
	otherWord
)

func nextWord(str string) (wt wordType, word, remaining string) {
	_ = "STUB: not implemented"
	return *new(wordType), "", ""
}

// it's a bit complex when dealing with a case like "HTTPStatus".
// it's expected to be splitted into "HTTP" and "Status".
// Therefore "S" should be in remaining instead of word.

func nextValidRune(str string, prev rune) (r rune, size int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func toLower(buf *stringBuilder, wt wordType, str string, connector rune) {
	_ = "STUB: not implemented"
	return
}

// SwapCase will swap characters case from upper to lower or lower to upper.
func SwapCase(str string) string { _ = "STUB: not implemented"; return "" }

// FirstRuneToUpper converts first rune to upper case if necessary.
func FirstRuneToUpper(str string) string { _ = "STUB: not implemented"; return "" }

// FirstRuneToLower converts first rune to lower case if necessary.
func FirstRuneToLower(str string) string { _ = "STUB: not implemented"; return "" }

// Shuffle randomizes runes in a string and returns the result.
// It uses default random source in `math/rand`.
func Shuffle(str string) string { _ = "STUB: not implemented"; return "" }

// ShuffleSource randomizes runes in a string with given random source.
func ShuffleSource(str string, src rand.Source) string { _ = "STUB: not implemented"; return "" }

// Successor returns the successor to string.
//
// If there is one alphanumeric rune is found in string, increase the rune by 1.
// If increment generates a "carry", the rune to the left of it is incremented.
// This process repeats until there is no carry, adding an additional rune if necessary.
//
// If there is no alphanumeric rune, the rightmost rune will be increased by 1
// regardless whether the result is a valid rune or not.
//
// Only following characters are alphanumeric.
//   - a - z
//   - A - Z
//   - 0 - 9
//
// Samples (borrowed from ruby's String#succ document):
//
//	"abcd"      => "abce"
//	"THX1138"   => "THX1139"
//	"<<koala>>" => "<<koalb>>"
//	"1999zzz"   => "2000aaa"
//	"ZZZ9999"   => "AAAA0000"
//	"***"       => "**+"
func Successor(str string) string { _ = "STUB: not implemented"; return "" }

// Needs to add one character for carry.

// Reserve enough space for write.

// No alphanumeric character. Simply increase last rune's value.
