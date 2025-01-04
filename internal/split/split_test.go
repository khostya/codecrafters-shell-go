package split

import "testing"

func TestSplit(t *testing.T) {
	type test struct {
		input string
		want  []string
	}

	tests := []test{
		{"echo abc", []string{"echo", "abc"}},
		{"echo abc '123   123'", []string{"echo", "abc", "123   123"}},
		{"echo abc \"123   123\"", []string{"echo", "abc", "123   123"}},
		{"echo abc '123   123''123'", []string{"echo", "abc", "123   123123"}},
		{"echo abc \"123   123\"\"123\"", []string{"echo", "abc", "123   123123"}},
		{"echo 'test hello'", []string{"echo", "test hello"}},
		{"echo \"test hello\"", []string{"echo", "test hello"}},
		{"echo shell     hello", []string{"echo", "shell", "hello"}},
		{"echo \"test' hello\"", []string{"echo", "test' hello"}},
		{"echo \"before\\   after\"", []string{"echo", "before\\   after"}},
		{"echo world\\ \\ \\ \\ \\ \\ script", []string{"echo", "world      script"}},
		{"echo \\'\\\"world shell\\\"\\'", []string{"echo", "'\"world", "shell\"'"}},
		{`echo "hello'world'\\n'script"`, []string{"echo", `hello'world'\n'script`}},
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
