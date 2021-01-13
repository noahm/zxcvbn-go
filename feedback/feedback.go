package feedback

// Warning to the user about unsafe passwords
type Warning string

const (
	// StraightRow Straight rows of keys are easy to guess
	StraightRow Warning = "straightRow"
	// KeyPattern Short keyboard patterns are easy to guess
	KeyPattern Warning = "keyPattern"
	// SimpleRepeat Repeated characters like "aaa" are easy to guess
	SimpleRepeat Warning = "simpleRepeat"
	// ExtendedRepeat Repeated character patterns like "abcabcabc" are easy to guess
	ExtendedRepeat Warning = "extendedRepeat"
	// WarnSequences Common character sequences like "abc" are easy to guess
	WarnSequences Warning = "sequences"
	// WarnRecentYears Recent years are easy to guess
	WarnRecentYears Warning = "recentYears"
	// WarnDates Dates are easy to guess
	WarnDates Warning = "dates"
	// TopTen This is a heavily used password
	TopTen Warning = "topTen"
	// TopHundred This is a frequently used password
	TopHundred Warning = "topHundred"
	// Common This is a commonly used password
	Common Warning = "common"
	// SimilarToCommon This is similar to a commonly used password
	SimilarToCommon Warning = "similarToCommon"
	// WordByItself Single words are easy to guess
	WordByItself Warning = "wordByItself"
	// NamesByThemselves Single names or surnames are easy to guess
	NamesByThemselves Warning = "namesByThemselves"
	// CommonNames Common names and surnames are easy to guess
	CommonNames Warning = "commonNames"
)

// Suggestion to the user on ways to improve a password
type Suggestion string

const (
	// L33t Avoid predictable letter substitutions like '@' for 'a'
	L33t Suggestion = "l33t"
	// ReverseWords Avoid reversed spellings of common words
	ReverseWords Suggestion = "reverseWords"
	// AllUppercase Capitalize some, but not all letters
	AllUppercase Suggestion = "allUppercase"
	// Capitalization Capitalize more than the first letter
	Capitalization Suggestion = "capitalization"
	// SuggestDates Avoid dates and years that are associated with you
	SuggestDates Suggestion = "dates"
	// SuggestRecentYears Avoid recent years
	SuggestRecentYears Suggestion = "recentYears"
	// AssociatedYears Avoid years that are associated with you
	AssociatedYears Suggestion = "associatedYears"
	// SuggestSequences Avoid common character sequences
	SuggestSequences Suggestion = "sequences"
	// Repeated Avoid repeated words and characters
	Repeated Suggestion = "repeated"
	// LongerKeyboardPattern Use longer keyboard patterns and change typing direction multiple times
	LongerKeyboardPattern Suggestion = "longerKeyboardPattern"
	// AnotherWord Add more words that are less common
	AnotherWord Suggestion = "anotherWord"
	// UseWords Use multiple words, but avoid common phrases
	UseWords Suggestion = "useWords"
	// NoNeed You can create strong passwords without using symbols, numbers, or uppercase letters
	NoNeed Suggestion = "noNeed"
)

// Feedback info for the user on why a score is low and how to improve it
type Feedback struct {
	Warn        Warning
	Suggestions []Suggestion
}
