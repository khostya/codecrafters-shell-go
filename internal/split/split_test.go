package split

import "testing"

func TestSplit(t *testing.T) {
	type test struct {
		input string
		want  []string
	}

	tests := []test{
		{"abc", []string{"abc"}},
		{"abc '123   123'", []string{"abc", "123   123"}},
		{"abc '123   123''123'", []string{"abc", "123   123123"}},
		{"echo 'test hello'", []string{"echo", "test hello"}},
		{"echo shell     hello", []string{"echo", "shell", "hello"}},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			t.Parallel()

			got := Split(tc.input)
			if len(got) != len(tc.want) {
				t.Errorf("Split(%q) = %v; want %v", tc.input, got, tc.want)
				return
			}

			for i := range got {
				if got[i] != tc.want[i] {
					t.Errorf("Split(%q) = %v; want %v", tc.input, got, tc.want)
					break
				}
			}
		})
	}
}
