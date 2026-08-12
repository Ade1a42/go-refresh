package ai

import "testing"

func TestDetectLanguage(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{
			name: "pure english",
			text: "the cat is in the box and it is happy",
			want: "Language: English (100%)",
		},
		{
			name: "pure french",
			text: "le chat est dans la boite et il est content",
			want: "Language: French (100%)",
		},
		{
			name: "no signal words",
			text: "banana rocket wizard tomorrow",
			want: "Language: Unknown",
		},
		{
			name: "empty string",
			text: "",
			want: "Language: Unknown",
		},
		{
			name: "mixed signals from the task example",
			text: "the menu is in the café and the prices are à la carte",
			want: "Language: English (67%)",
		},
		{
			name: "tie goes to english",
			// 1 english signal ("the"), 1 french signal ("la") -> tie -> engCount >= frCount -> English
			text: "the la",
			want: "Language: English (50%)",
		},
		{
			name: "case insensitive matching",
			text: "THE AND OF TO IS",
			want: "Language: English (100%)",
		},
		{
			name: "whole word only, not substring",
			// "theatre" contains "the" but must not count as a signal
			text: "theatre incroyable spectacle",
			want: "Language: Unknown",
		},
		{
			name: "accented characters count as french even without signal words",
			text: "café château",
			want: "Language: French (100%)",
		},
		{
			name: "french wins by majority",
			// french signals: le, et, la, et (4) ; english signals: is (1) -> total 5, fr score = round(4*100/5) = 80
			text: "le chat et la maison sont ici et is",
			want: "Language: French (80%)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DetectLanguage(tt.text)
			if got != tt.want {
				t.Errorf("DetectLanguage(%q) = %q, want %q", tt.text, got, tt.want)
			}
		})
	}
}