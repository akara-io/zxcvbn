package feedback

import (
	"github.com/akara-io/zxcvbn/match"
	"github.com/akara-io/zxcvbn/scoring"
	"strings"
)

// Feedback represents feedback to improve a password
type Feedback struct {
	Warning     string
	Suggestions []string
}

// New returns an initialised Feedback struct
func New() *Feedback {
	return &Feedback{
		Warning:     "",
		Suggestions: []string{},
	}
}

// Warn returns a Feedback with Warning property set to s
func (f *Feedback) Warn(s string) *Feedback {
	f.Warning = s
	return f
}

// Suggest returns a Feedback with s added to the Suggestions list
func (f *Feedback) Suggest(s string) *Feedback {
	f.Suggestions = append(f.Suggestions, s)
	return f
}

// SuggestFirst returns a Feedback with s inserted as the first Suggestion.
func (f *Feedback) SuggestFirst(s string) *Feedback {
	f.Suggestions = append([]string{s}, f.Suggestions...)
	return f
}

// IsZero returns true if the Feedback is zero value
func (f *Feedback) IsZero() bool {
	return f.Warning == "" && (f.Suggestions == nil || len(f.Suggestions) == 0)
}

var defaultFeedback = Feedback{
	Warning: "",
	Suggestions: []string{
		"Use a few words, avoid common phrases",
		"No need for symbols, digits, or uppercase letters",
	},
}

// GetFeedback returns feedback on a password based on its score and sequence of matches
func GetFeedback(score int, sequence []*match.Match) Feedback {
	// Starting feedback
	if len(sequence) == 0 {
		return defaultFeedback
	}

	// No feedback if store is good or great
	if score > 2 {
		return *(New())
	}

	// Tie feedback to the longest match for longer sequences
	longestMatch := sequence[0]
	for _, m := range sequence[1:] {
		if len(m.Token) > len(longestMatch.Token) {
			longestMatch = m
		}
	}
	feedback := getMatchFeedback(longestMatch, len(sequence) == 1)
	extraFeedback := "Add another word or two. Uncommon words are better."
	if feedback != nil {
		feedback = feedback.SuggestFirst(extraFeedback)
	} else {
		feedback = New().Suggest(extraFeedback)
	}
	return *feedback
}

func getMatchFeedback(match *match.Match, isSoleMatch bool) *Feedback {
	var f *Feedback

	switch match.Pattern {
	case "dictionary":
		return getDictionaryMatchFeedback(match, isSoleMatch)

	case "spatial":
		if match.Turns == 1 {
			f = New().Warn("Straight rows of keys are easy to guess")
		} else {
			f = New().Warn("Short keyboard patterns are easy to guess")
		}
		f = f.Suggest("Use a longer keyboard pattern with more turns")

	case "repeat":
		if len(match.BaseToken) == 1 {
			f = New().Warn(`Repeats like "aaa" are easy to guess`)
		} else {
			f = New().Warn(`Repeats like "abcabcabc" are only slightly harder to guess than "abc"`)
		}
		f = f.Suggest("Avoid repeated words and characters")

	case "sequence":
		f = New().Warn("Sequences like abc or 6543 are easy to guess").
			Suggest("Avoid sequences")

	case "regex":
		if match.RegexName == "recent_year" {
			f = New().Warn("Recent years are easy to guess").
				Suggest("Avoid recent years").
				Suggest("Avoid years that are associated with you")
		}

	case "date":
		f = New().Warn("Dates are often easy to guess").
			Suggest("Avoid dates and years that are associated with you")

	default:
		f = nil
	}

	return f
}

func getDictionaryMatchFeedback(match *match.Match, isSoleMatch bool) *Feedback {
	f := New()

	if match.DictionaryName == "passwords" {
		if isSoleMatch && !match.L33t && !match.Reversed {
			if match.Rank <= 10 {
				f = f.Warn("This is a top-10 common password")
			} else if match.Rank <= 100 {
				f = f.Warn("This is a top-100 common password")
			} else {
				f = f.Warn("This is a very common password")
			}
		} else if match.Guesses <= 10000 {
			f = f.Warn("This is similar to a commonly used password")
		}
	} else if match.DictionaryName == "english_wikipedia" {
		if isSoleMatch {
			f = f.Warn("A word by itself is easy to guess")
		}
	} else if contains(match.DictionaryName, []string{"surnames", "male_names", "female_names"}) {
		if isSoleMatch {
			f = f.Warn("Names and surnames by themselves are easy to guess")
		} else {
			f = f.Warn("Common names and surnames are easy to guess")
		}
	}

	word := match.Token
	if scoring.ReStartUpper.MatchString(word) {
		f = f.Suggest("Capitalization doesn't help very much")
	} else if scoring.ReAllUpper.MatchString(word) && (strings.ToLower(word) != word) {
		f = f.Suggest("All-uppercase is almost as easy to guess as all-lowercase")
	}

	if match.Reversed && len(match.Token) >= 4 {
		f = f.Suggest("All-uppercase is almost as easy to guess as all-lowercase")
	}

	if match.L33t {
		f = f.Suggest("Predictable substitutions like '@' instead of 'a' don't help very much")
	}

	return f
}

func contains(find string, in []string) bool {
	for _, s := range in {
		if s == find {
			return true
		}
	}
	return false
}
