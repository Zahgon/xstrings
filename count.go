// Copyright 2015 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package xstrings

// Len returns str's utf8 rune length.
func Len(str string) int { _ = "STUB: not implemented"; return 0 }

// WordCount returns number of words in a string.
//
// Word is defined as a locale dependent string containing alphabetic characters,
// which may also contain but not start with `'` and `-` characters.
func WordCount(str string) int { _ = "STUB: not implemented"; return 0 }

// Still in word.

const minCJKCharacter = '\u3400'

// Checks r is a letter but not CJK character.
func isAlphabet(r rune) bool { _ = "STUB: not implemented"; return false }

// Quick check for non-CJK character.

// Common CJK characters.

// Rare CJK characters.

// Rare and historic CJK characters.

// Width returns string width in monotype font.
// Multi-byte characters are usually twice the width of single byte characters.
//
// Algorithm comes from `mb_strwidth` in PHP.
// http://php.net/manual/en/function.mb-strwidth.php
func Width(str string) int { _ = "STUB: not implemented"; return 0 }

// RuneWidth returns character width in monotype font.
// Multi-byte characters are usually twice the width of single byte characters.
//
// Algorithm comes from `mb_strwidth` in PHP.
// http://php.net/manual/en/function.mb-strwidth.php
func RuneWidth(r rune) int { _ = "STUB: not implemented"; return 0 }
