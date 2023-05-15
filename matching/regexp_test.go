package matching

import (
	"github.com/akara-io/zxcvbn/match"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegexpMatching(t *testing.T) {
	rm := regexpMatch{regexes: defaultRegexpMatch}
	assert.Equal(t, []*match.Match{
		{
			Pattern:   "regex",
			Token:     "1922",
			I:         0,
			J:         3,
			RegexName: "recent_year",
		},
	},
		rm.Matches("1922"),
	)

	assert.Equal(t, []*match.Match{
		{
			Pattern:   "regex",
			Token:     "2017",
			I:         0,
			J:         3,
			RegexName: "recent_year",
		},
	},
		rm.Matches("2017"),
	)
}
