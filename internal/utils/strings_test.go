package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMixCase(t *testing.T) {
	cases := []struct {
		name string
		text string
		want string
	}{
		{"empty", "", ""},
		{"cyrillic", "МИКРОФОН", "мИкРоФоН"},
		{"latin", "Search", "sEaRcH"},
		{"punctuation and digits", "А-1Б в!Г", "а-1Б в!Г"},
		{"no letters", "123 !", "123 !"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, MixCase(tc.text))
		})
	}
}
