package main

import (
	"testing"
)

func TestLineExistsInComment(t *testing.T) {
	commentLines := []string{
		"/**",
		" * This is a test comment",
		" * @param {number} num - a number",
		" */",
	}

	lineToAdd := "@param {number} num - a number"

	if !lineExistsInComment(commentLines, lineToAdd) {
		t.Errorf("lineExistsInComment failed, expected true, got false")
	}

	lineToAdd = "@returns {number} - the result"

	if lineExistsInComment(commentLines, lineToAdd) {
		t.Errorf("lineExistsInComment failed, expected false, got true")
	}
}

func TestAppendCommentLine(t *testing.T) {
	lines := []string{
		"function add(a, b) {",
		"  /**",
		"   * This is a test comment",
		"   */",
		"   function add(a, b) {}",
		"}",
	}

	commentLines := []string{
		"  /**",
		"   * This is a test comment",
		"   */",
	}

	lineToAdd := "@param {number} a - first number"

	newLines := appendCommentLine(lines, commentLines, lineToAdd)

	expectedLines := []string{
		"function add(a, b) {",
		"  /**",
		"   * This is a test comment",
		"   * @param {number} a - first number",
		"   */",
		"   function add(a, b) {}",
		"}",
	}

	for i, line := range newLines {
		if line != expectedLines[i] {
			t.Errorf("appendCommentLine failed, expected %v, got %v", expectedLines, newLines)
			break
		}
	}
}
