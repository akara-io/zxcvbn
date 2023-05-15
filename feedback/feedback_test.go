package feedback_test

import (
	"github.com/akara-io/zxcvbn"
	"github.com/akara-io/zxcvbn/feedback"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFeedback(t *testing.T) {
	// Test cases are taken from the zxcvbn examples at https://lowe.github.io/tryzxcvbn/
	tests := []struct {
		password     string
		wantFeedback feedback.Feedback
	}{
		{
			password: "zxcvbn",
			wantFeedback: feedback.Feedback{
				Warning:     "This is a top-100 common password",
				Suggestions: []string{"Add another word or two. Uncommon words are better."},
			},
		},
		{
			password: "qwER43@!",
			wantFeedback: feedback.Feedback{
				Warning: "Short keyboard patterns are easy to guess",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
					"Use a longer keyboard pattern with more turns",
				},
			},
		},
		{
			password: "Tr0ub4dour&3",
			wantFeedback: feedback.Feedback{
				Warning: "",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
					"Capitalization doesn't help very much",
					"Predictable substitutions like '@' instead of 'a' don't help very much",
				},
			},
		},
		{
			password: "correcthorsebatterystaple",
			wantFeedback: feedback.Feedback{
				Warning:     "",
				Suggestions: []string{},
			},
		},
		{
			password: "coRrecth0rseba++ery9.23.2007staple$",
			wantFeedback: feedback.Feedback{
				Warning:     "",
				Suggestions: []string{},
			},
		},
		{
			password: "p@ssword",
			wantFeedback: feedback.Feedback{
				Warning: "This is similar to a commonly used password",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
					"Predictable substitutions like '@' instead of 'a' don't help very much",
				},
			},
		},
		{
			password: "p@$$word",
			wantFeedback: feedback.Feedback{
				Warning: "This is similar to a commonly used password",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
					"Predictable substitutions like '@' instead of 'a' don't help very much",
				},
			},
		},
		{
			password: "123456",
			wantFeedback: feedback.Feedback{
				Warning: "This is a top-10 common password",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
				},
			},
		},
		{
			password: "123456789",
			wantFeedback: feedback.Feedback{
				Warning: "This is a top-10 common password",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
				},
			},
		},
		{
			password: "11111111",
			wantFeedback: feedback.Feedback{
				Warning: "This is a top-100 common password",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
				},
			},
		},
		{
			password: "zxcvbnm,./",
			wantFeedback: feedback.Feedback{
				Warning: "Straight rows of keys are easy to guess",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
					"Use a longer keyboard pattern with more turns",
				},
			},
		},
		{
			password: "love88",
			wantFeedback: feedback.Feedback{
				Warning: "This is similar to a commonly used password",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
				},
			},
		},
		{
			password: "angel08",
			wantFeedback: feedback.Feedback{
				Warning: "Common names and surnames are easy to guess",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
					"Predictable substitutions like '@' instead of 'a' don't help very much",
				},
			},
		},
		{
			password: "monkey13",
			wantFeedback: feedback.Feedback{
				Warning: "This is a very common password",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
				},
			},
		},
		{
			password: "iloveyou",
			wantFeedback: feedback.Feedback{
				Warning: "This is a top-100 common password",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
				},
			},
		},
		{
			password: "woaini",
			wantFeedback: feedback.Feedback{
				Warning: "This is a very common password",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
				},
			},
		},
		{
			password: "wang",
			wantFeedback: feedback.Feedback{
				Warning: "Names and surnames by themselves are easy to guess",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
				},
			},
		},
		{
			password: "tianya",
			wantFeedback: feedback.Feedback{
				Warning: "Common names and surnames are easy to guess",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
				},
			},
		},
		{
			password: "zhang198822",
			wantFeedback: feedback.Feedback{
				Warning: "Dates are often easy to guess",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
					"Avoid dates and years that are associated with you",
				},
			},
		},
		{
			password: "li4478",
			wantFeedback: feedback.Feedback{
				Warning: "",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
				},
			},
		},
		{
			password: "a6a4Aa8a",
			wantFeedback: feedback.Feedback{
				Warning: "",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
				},
			},
		},
		{
			password: "b6b4Bb8b",
			wantFeedback: feedback.Feedback{
				Warning: "",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
				},
			},
		},
		{
			password: "z6z4Zz8z",
			wantFeedback: feedback.Feedback{
				Warning: "",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
				},
			},
		},
		{
			password: "pässwörd",
			wantFeedback: feedback.Feedback{
				Warning: "",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
				},
			},
		},
		{
			password: "alpha bravo charlie delta",
			wantFeedback: feedback.Feedback{
				Warning:     "",
				Suggestions: []string{},
			},
		},
		{
			password: " a b c d e f g h i j k l m n o p q r s t u v w x y z 0 1 2 3 4 5 6 7 8 9",
			wantFeedback: feedback.Feedback{
				Warning:     "",
				Suggestions: []string{},
			},
		},
		{
			password: "!\"£$%^&*()",
			wantFeedback: feedback.Feedback{
				Warning: "Straight rows of keys are easy to guess",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
					"Use a longer keyboard pattern with more turns",
				},
			},
		},
		{
			password: "D0g..................",
			wantFeedback: feedback.Feedback{
				Warning: "Repeats like \"aaa\" are easy to guess",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
					"Avoid repeated words and characters",
				},
			},
		},
		{
			password: "abcdefghijk987654321",
			wantFeedback: feedback.Feedback{
				Warning: "Sequences like abc or 6543 are easy to guess",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
					"Avoid sequences",
				},
			},
		},
		{
			password: "1qaz2wsx3edc",
			wantFeedback: feedback.Feedback{
				Warning: "This is a very common password",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
				},
			},
		},
		{
			password: "temppass22",
			wantFeedback: feedback.Feedback{
				Warning: "This is similar to a commonly used password",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
				},
			},
		},
		{
			password: "briansmith",
			wantFeedback: feedback.Feedback{
				Warning: "Common names and surnames are easy to guess",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
				},
			},
		},
		{
			password: "briansmith4mayor",
			wantFeedback: feedback.Feedback{
				Warning:     "",
				Suggestions: []string{},
			},
		},
		{
			password: "password1",
			wantFeedback: feedback.Feedback{
				Warning: "This is a very common password",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
				},
			},
		},
		{
			password: "ScoRpi0ns",
			wantFeedback: feedback.Feedback{
				Warning: "",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
					"Predictable substitutions like '@' instead of 'a' don't help very much",
				},
			},
		},
		{
			password: "Rosebud",
			wantFeedback: feedback.Feedback{
				Warning: "This is a very common password",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
					"Capitalization doesn't help very much",
				},
			},
		},
		{
			password: "ROSEBUD",
			wantFeedback: feedback.Feedback{
				Warning: "This is a very common password",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
					"All-uppercase is almost as easy to guess as all-lowercase",
				},
			},
		},
		{
			password: "rosebuD",
			wantFeedback: feedback.Feedback{
				Warning: "This is a very common password",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
				},
			},
		},
		{
			password: "ros3bud99",
			wantFeedback: feedback.Feedback{
				Warning: "This is similar to a commonly used password",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
					"Predictable substitutions like '@' instead of 'a' don't help very much",
				},
			},
		},
		{
			password: "r0s3bud99",
			wantFeedback: feedback.Feedback{
				Warning: "This is similar to a commonly used password",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
					"Predictable substitutions like '@' instead of 'a' don't help very much",
				},
			},
		},
		{
			password: "R0$38uD99",
			wantFeedback: feedback.Feedback{
				Warning: "",
				Suggestions: []string{
					"Add another word or two. Uncommon words are better.",
					"Predictable substitutions like '@' instead of 'a' don't help very much",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.password, func(t *testing.T) {
			result := zxcvbn.PasswordStrength(tt.password, nil)
			feedback := feedback.GetFeedback(result.Score, result.Sequence)
			assert.Equal(t, tt.wantFeedback, feedback)
		})
	}
}
