package stripol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrIpol(t *testing.T) {
	const (
		leftDelim  string = "{{"
		rightDelim string = "}}"
	)

	s := New(leftDelim, rightDelim)

	t.Run("New", func(t *testing.T) {
		assert.Equal(t, leftDelim, s.leftDelim)
		assert.Equal(t, leftDelim, s.LeftDelim())

		assert.Equal(t, rightDelim, s.rightDelim)
		assert.Equal(t, rightDelim, s.RightDelim())

		ld, rd := s.Delims()
		assert.Equal(t, leftDelim, ld)
		assert.Equal(t, rightDelim, rd)
	})

	t.Run("Basic substitution", func(t *testing.T) {
		s.RegisterVar("FAV_ANIMAL", "tigers")
		result := s.Eval("{{ FAV_ANIMAL }} are my favorite animal.")
		assert.Equal(t, "tigers are my favorite animal.", result)
	})

	t.Run("Basic substitution with no spaces", func(t *testing.T) {
		s.RegisterVar("FAV_ANIMAL", "tigers")
		result := s.Eval("{{FAV_ANIMAL}} are my favorite animal.")
		assert.Equal(t, "tigers are my favorite animal.", result)
	})

	t.Run("Basic substitution with left space", func(t *testing.T) {
		s.RegisterVar("FAV_ANIMAL", "tigers")
		result := s.Eval("{{ FAV_ANIMAL}} are my favorite animal.")
		assert.Equal(t, "tigers are my favorite animal.", result)
	})

	t.Run("Basic substitution with left space", func(t *testing.T) {
		s.RegisterVar("FAV_ANIMAL", "tigers")
		result := s.Eval("{{FAV_ANIMAL }} are my favorite animal.")
		assert.Equal(t, "tigers are my favorite animal.", result)
	})

	t.Run("Variable not registered", func(t *testing.T) {
		result := s.Eval("{{ NON_EXISTENT_VAR }} is not registered.")
		assert.Equal(t, " is not registered.", result)
	})

	t.Run("Multiple variables", func(t *testing.T) {
		s.RegisterVar("FAV_ANIMAL", "tigers")
		s.RegisterVar("SECOND_FAV", "lions")
		result := s.Eval("My favorite animals are {{ FAV_ANIMAL }} and {{ SECOND_FAV }}.")
		assert.Equal(t, "My favorite animals are tigers and lions.", result)
	})

	t.Run("No delimiters", func(t *testing.T) {
		result := s.Eval("No delimiters here.")
		assert.Equal(t, "No delimiters here.", result)
	})

	t.Run("Empty string", func(t *testing.T) {
		result := s.Eval("")
		assert.Equal(t, "", result)
	})

	t.Run("Empty variable value", func(t *testing.T) {
		s.RegisterVar("EMPTY_VAR", "")
		result := s.Eval("{{ EMPTY_VAR }} should be empty.")
		assert.Equal(t, " should be empty.", result)
	})

	t.Run("Extra spaces", func(t *testing.T) {
		s.RegisterVar("SPACED_VAR", "with spaces")
		result := s.Eval("Variable {{ SPACED_VAR }} has extra spaces.")
		assert.Equal(t, "Variable with spaces has extra spaces.", result)
	})

	t.Run("Variable with spaces around", func(t *testing.T) {
		s.RegisterVar(" VAR_WITH_SPACES ", "trimmed")
		value, ok := s.data["VAR_WITH_SPACES"]
		assert.True(t, ok)
		assert.Equal(t, "trimmed", value)

		result := s.Eval("The variable {{  VAR_WITH_SPACES  }} should be trimmed.")
		assert.Equal(t, "The variable trimmed should be trimmed.", result)
	})

	t.Run("Nested delimiters", func(t *testing.T) {
		s.RegisterVar("NESTED_VAR", "value")
		result := s.Eval("{{ {{ NESTED_VAR }} }} should not match.")
		assert.Equal(t, " value  should not match.", result)
	})

	t.Run("Set Left Delimiter", func(t *testing.T) {
		s.SetLeftDelim("[[")
		assert.Equal(t, "[[", s.LeftDelim())
	})

	t.Run("Set Right Delimiter", func(t *testing.T) {
		s.SetRightDelim("]]")
		assert.Equal(t, "]]", s.RightDelim())
	})

	t.Run("Set Both Delimiters", func(t *testing.T) {
		s.SetDelims("<<", ">>")
		assert.Equal(t, "<<", s.LeftDelim())
		assert.Equal(t, ">>", s.RightDelim())
	})

	t.Run("Get Delimiters", func(t *testing.T) {
		left, right := s.Delims()
		assert.Equal(t, "<<", left)
		assert.Equal(t, ">>", right)
	})

	t.Run("Register Multiple Variables", func(t *testing.T) {
		vars := map[string]string{
			"FIRST":  "one",
			"SECOND": "two",
		}
		s.RegisterVars(vars)
		assert.Equal(t, "one", s.data["FIRST"])
		assert.Equal(t, "two", s.data["SECOND"])
	})

	t.Run("Set Data", func(t *testing.T) {
		newData := map[string]string{
			"NEW_VAR": "new value",
		}
		s.SetData(newData)
		assert.Equal(t, "new value", s.data["NEW_VAR"])
		assert.Equal(t, 1, len(s.data))
	})

	t.Run("Reset Data", func(t *testing.T) {
		s.RegisterVar("TEMP_VAR", "temp value")
		s.ResetData()
		assert.Empty(t, s.data)
	})
}
