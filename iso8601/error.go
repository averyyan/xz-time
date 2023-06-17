package ISO8601

import (
	"errors"
	"fmt"
)

var (

	// ErrBadFormat is returned when parsing fails
	ErrBadFormat = errors.New("bad format string")

	// ErrNoMonth is raised when a month is in the format string
	ErrNoMonth = errors.New("no months allowed")

	// ErrZoneCharacters indicates an incorrect amount of characters was passed to ParseISOZone.
	ErrZoneCharacters = errors.New("ISO8601: Expected between 3 and 6 characters for zone information")

	// ErrInvalidZone indicates an invalid timezone per the standard that doesn't violate any specific
	// character parsing rules.
	ErrInvalidZone = errors.New("ISO8601: Specified zone is invalid")

	// ErrRemainingData indicates that there is extra data after a `Z` character.
	ErrRemainingData = errors.New("ISO8601: Unexepected remaining data after `Z`")

	// ErrNotString indicates that a non string type was passed to the UnmarshalJSON method of `Time`.
	ErrNotString = errors.New("ISO8601: Invalid json type (expected string)")

	// ErrPrecision indicates that there was too much precision (characters) given to parse
	// for the fraction of a second of the input time.
	ErrPrecision = errors.New("ISO8601: Too many characters in fraction of second precision")
)

func newUnexpectedCharacterError(c byte) error {
	return &UnexpectedCharacterError{Character: c}
}

// UnexpectedCharacterError indicates the parser scanned a character that was not expected at that time.
type UnexpectedCharacterError struct {
	Character byte
}

func (e *UnexpectedCharacterError) Error() string {
	return fmt.Sprintf("ISO8601: Unexpected character `%c`", e.Character)
}

// RangeError indicates that a value is not in an expected range.
type RangeError struct {
	Value   string
	Element string
	Min     int
	Max     int
	Given   int
}

func (e *RangeError) Error() string {
	return fmt.Sprintf("iso8601: Cannot parse %q: %s %d is not in range %d-%d", e.Value, e.Element, e.Given, e.Min, e.Max)
}
