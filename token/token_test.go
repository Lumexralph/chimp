package token

import "testing"

func TestLookupIdentifier(t *testing.T) {
	cases := []struct{
		ident string
		want Type
	}{
		{"fn", FUNCTION},
		{"let", LET},
		{"var", IDENT},
	}

	for _, tc := range cases {
		t.Run("Lookup the type of the token", func(t *testing.T){
			got := LookupIdentifier(tc.ident)

			if got != tc.want {
				t.Errorf("LookupIdentifier(%v) got %v, want %v", tc.ident, got, tc.want)
			}
		})
	}
}