package task2

import "testing"

func TestParseInt(t *testing.T) {

	testsCases := []struct {
		Name      string
		InputStr  string
		Expect    int
		ExpectErr bool
	}{
		{
			Name:      "12_ok",
			InputStr:  "12",
			Expect:    12,
			ExpectErr: false,
		},
		{
			Name:      "-44_ok",
			InputStr:  "-44",
			Expect:    -44,
			ExpectErr: false,
		},
		{
			Name:      "1w2_invalid_input",
			InputStr:  "1w2",
			Expect:    0,
			ExpectErr: true,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := ParseInt(tt.InputStr)

			if (err != nil) != tt.ExpectErr {
				t.Errorf("ParseInt(%s) got error %v, want error %v", tt.InputStr, err, tt.ExpectErr)
			}
			if result != tt.Expect {
				t.Errorf("ParseInt(%s) got result %v, want %v", tt.InputStr, result, tt.Expect)
			}
		})
	}
}
