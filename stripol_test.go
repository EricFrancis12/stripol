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

	i := New(leftDelim, rightDelim)

	t.Run("New", func(t *testing.T) {
		assert.Equal(t, leftDelim, i.leftDelim)
		assert.Equal(t, leftDelim, i.LeftDelim())

		assert.Equal(t, rightDelim, i.rightDelim)
		assert.Equal(t, rightDelim, i.RightDelim())

		ld, rd := i.Delims()
		assert.Equal(t, leftDelim, ld)
		assert.Equal(t, rightDelim, rd)
	})

	t.Run("Basic substitution", func(t *testing.T) {
		i.RegisterVar("FAV_ANIMAL", "tigers")
		result := i.Eval("{{ FAV_ANIMAL }} are my favorite animal.")
		assert.Equal(t, "tigers are my favorite animal.", result)
	})

	t.Run("Variable not registered", func(t *testing.T) {
		result := i.Eval("{{ NON_EXISTENT_VAR }} is not registered.")
		assert.Equal(t, " is not registered.", result)
	})

	t.Run("Multiple variables", func(t *testing.T) {
		i.RegisterVar("FAV_ANIMAL", "tigers")
		i.RegisterVar("SECOND_FAV", "lions")
		result := i.Eval("My favorite animals are {{ FAV_ANIMAL }} and {{ SECOND_FAV }}.")
		assert.Equal(t, "My favorite animals are tigers and lions.", result)
	})

	t.Run("No delimiters", func(t *testing.T) {
		result := i.Eval("No delimiters here.")
		assert.Equal(t, "No delimiters here.", result)
	})

	t.Run("Empty string", func(t *testing.T) {
		result := i.Eval("")
		assert.Equal(t, "", result)
	})

	t.Run("Empty variable value", func(t *testing.T) {
		i.RegisterVar("EMPTY_VAR", "")
		result := i.Eval("{{ EMPTY_VAR }} should be empty.")
		assert.Equal(t, " should be empty.", result)
	})

	t.Run("Extra spaces", func(t *testing.T) {
		i.RegisterVar("SPACED_VAR", "with spaces")
		result := i.Eval("Variable {{ SPACED_VAR }} has extra spaces.")
		assert.Equal(t, "Variable with spaces has extra spaces.", result)
	})

	t.Run("Variable with spaces around", func(t *testing.T) {
		i.RegisterVar(" VAR_WITH_SPACES ", "trimmed")
		value, ok := i.data["VAR_WITH_SPACES"]
		assert.True(t, ok)
		assert.Equal(t, "trimmed", value)

		result := i.Eval("The variable {{  VAR_WITH_SPACES  }} should be trimmed.")
		assert.Equal(t, "The variable trimmed should be trimmed.", result)
	})

	t.Run("Nested delimiters", func(t *testing.T) {
		i.RegisterVar("NESTED_VAR", "value")
		result := i.Eval("{{ {{ NESTED_VAR }} }} should not match.")
		assert.Equal(t, " value  should not match.", result)
	})

	t.Run("Set Left Delimiter", func(t *testing.T) {
		i.SetLeftDelim("[[")
		assert.Equal(t, "[[", i.LeftDelim())
	})

	t.Run("Set Right Delimiter", func(t *testing.T) {
		i.SetRightDelim("]]")
		assert.Equal(t, "]]", i.RightDelim())
	})

	t.Run("Set Both Delimiters", func(t *testing.T) {
		i.SetDelims("<<", ">>")
		assert.Equal(t, "<<", i.LeftDelim())
		assert.Equal(t, ">>", i.RightDelim())
	})

	t.Run("Get Delimiters", func(t *testing.T) {
		left, right := i.Delims()
		assert.Equal(t, "<<", left)
		assert.Equal(t, ">>", right)
	})

	t.Run("Register Multiple Variables", func(t *testing.T) {
		vars := map[string]string{
			"FIRST":  "one",
			"SECOND": "two",
		}
		i.RegisterVars(vars)
		assert.Equal(t, "one", i.data["FIRST"])
		assert.Equal(t, "two", i.data["SECOND"])
	})

	t.Run("Set Data", func(t *testing.T) {
		newData := map[string]string{
			"NEW_VAR": "new value",
		}
		i.SetData(newData)
		assert.Equal(t, "new value", i.data["NEW_VAR"])
		assert.Equal(t, 1, len(i.data))
	})

	t.Run("Reset Data", func(t *testing.T) {
		i.RegisterVar("TEMP_VAR", "temp value")
		i.ResetData()
		assert.Empty(t, i.data)
	})
}
