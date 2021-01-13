package feedback

// SpatialFeedback creates for spatial match type
func SpatialFeedback(turns int) Feedback {
	var w = KeyPattern
	if turns == 1 {
		w = StraightRow
	}
	return Feedback{
		Warn:        w,
		Suggestions: []Suggestion{LongerKeyboardPattern},
	}
}

// DateFeedback creates for Date match type
func DateFeedback() Feedback {
	return Feedback{
		Warn:        WarnDates,
		Suggestions: []Suggestion{SuggestDates},
	}
}
