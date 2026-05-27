// Copyright 2015 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package xstrings

// Reverse a utf8 encoded string.
func Reverse(str string) string { _ = "STUB: not implemented"; return "" }

// Slice a string by rune.
//
// Start must satisfy 0 <= start <= rune length.
//
// End can be positive, zero or negative.
// If end >= 0, start and end must satisfy start <= end <= rune length.
// If end < 0, it means slice to the end of string.
//
// Otherwise, Slice will panic as out of range.
func Slice(str string, start, end int) string { _ = "STUB: not implemented"; return "" }

// Partition splits a string by sep into three parts.
// The return value is a slice of strings with head, match and tail.
//
// If str contains sep, for example "hello" and "l", Partition returns
//
//	"he", "l", "lo"
//
// If str doesn't contain sep, for example "hello" and "x", Partition returns
//
//	"hello", "", ""
func Partition(str, sep string) (head, match, tail string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

// LastPartition splits a string by last instance of sep into three parts.
// The return value is a slice of strings with head, match and tail.
//
// If str contains sep, for example "hello" and "l", LastPartition returns
//
//	"hel", "l", "o"
//
// If str doesn't contain sep, for example "hello" and "x", LastPartition returns
//
//	"", "", "hello"
func LastPartition(str, sep string) (head, match, tail string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

// Insert src into dst at given rune index.
// Index is counted by runes instead of bytes.
//
// If index is out of range of dst, panic with out of range.
func Insert(dst, src string, index int) string { _ = "STUB: not implemented"; return "" }

// Scrub scrubs invalid utf8 bytes with repl string.
// Adjacent invalid bytes are replaced only once.
func Scrub(str, repl string) string { _ = "STUB: not implemented"; return "" }

// No invalid byte.

// WordSplit splits a string into words. Returns a slice of words.
// If there is no word in a string, return nil.
//
// Word is defined as a locale dependent string containing alphabetic characters,
// which may also contain but not start with `'` and `-` characters.
func WordSplit(str string) []string { _ = "STUB: not implemented"; return nil }

// Still in word.
