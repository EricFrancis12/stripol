package stripol

import (
	"strings"
)

// StrIpol is a structure that holds data for string interpolation, including delimiters and variable data.
type StrIpol struct {
	data       map[string]string
	leftDelim  string
	rightDelim string
}

// New creates a new StrIpol instance with specified left and right delimiters.
func New(leftDelim, rightDelim string) *StrIpol {
	return &StrIpol{
		data:       make(map[string]string),
		leftDelim:  leftDelim,
		rightDelim: rightDelim,
	}
}

// SetLeftDelim sets the left delimiter for the StrIpol instance.
func (s *StrIpol) SetLeftDelim(leftDelim string) {
	s.leftDelim = leftDelim
}

// SetRightDelim sets the right delimiter for the StrIpol instance.
func (s *StrIpol) SetRightDelim(rightDelim string) {
	s.rightDelim = rightDelim
}

// SetDelims sets both the left and right delimiters for the StrIpol instance.
func (s *StrIpol) SetDelims(leftDelim, rightDelim string) {
	s.SetLeftDelim(leftDelim)
	s.SetRightDelim(rightDelim)
}

// LeftDelim returns the current left delimiter.
func (s *StrIpol) LeftDelim() string {
	return s.leftDelim
}

// RightDelim returns the current right delimiter.
func (s *StrIpol) RightDelim() string {
	return s.rightDelim
}

// Delims returns both the left and right delimiters.
func (s *StrIpol) Delims() (leftDelim, rightDelim string) {
	return s.LeftDelim(), s.RightDelim()
}

// RegisterVar adds a single variable with its value to the StrIpol instance.
func (s *StrIpol) RegisterVar(name, value string) {
	s.data[strings.Trim(name, " ")] = value
}

// RegisterVars adds multiple variables and their values to the StrIpol instance.
func (s *StrIpol) RegisterVars(vars map[string]string) {
	for name, value := range vars {
		s.RegisterVar(name, value)
	}
}

// SetData replaces the existing variable data with the provided map.
func (s *StrIpol) SetData(data map[string]string) {
	s.data = data
}

// Reset clears all registered variables from the StrIpol instance. Left and right delims are left unchanged.
func (s *StrIpol) ResetData() {
	s.data = make(map[string]string)
}

// Eval evaluates the given string, replacing any variables found within the delimiters with their registered values.
func (s *StrIpol) Eval(str string) string {
	partsA := strings.Split(str, s.leftDelim)
	for i, a := range partsA {
		if strings.Contains(a, s.rightDelim) {
			partsB := strings.Split(a, s.rightDelim)
			trimmed := strings.Trim(partsB[0], " ")
			value, ok := s.data[trimmed]
			if ok {
				partsB[0] = value
			} else {
				partsB[0] = ""
			}
			partsA[i] = strings.Join(partsB, "")
		}
	}
	return strings.Join(partsA, "")
}
