package tool

import "testing"

const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestUp(t *testing.T) {
	t.Log("Given the need to test Up() function behavoir")
	{
		testID := 0
		t.Logf("\tTest %d:\t", testID)
		{
			text := "Ready, set, go (up)!"
			want := "Ready, set, GO!"
			got := Up(text)

			if want != got {
				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
			}
			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
		}

		testID++
		t.Logf("\tTest %d:\t", testID)
		{
			text := "This is a (up) (up) test (up) case"
			want := "This is A (up) TEST case"
			got := Up(text)

			if want != got {
				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
			}
			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
		}

		testID++
		t.Logf("\tTest %d:\t", testID)
		{
			text := "Coding is fun, (up) right?"
			want := "Coding is FUN, right?"
			got := Up(text)

			if want != got {
				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
			}
			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
		}

		testID++
		t.Logf("\tTest %d:\t", testID)
		{
			text := "The quick brown fox jumps over the lazy dog (up)on the hill"
			want := "The quick brown fox jumps over the lazy DOGon the hill"
			got := Up(text)

			if want != got {
				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
			}
			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
		}

		testID++
		t.Logf("\tTest %d:\t", testID)
		{
			text := "The quick brown fox jumps over the lazy dog(up)on the hill"
			want := "The quick brown fox jumps over the lazy DOGon the hill"
			got := Up(text)

			if want != got {
				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
			}
			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
		}

		testID++
		t.Logf("\tTest %d:\t", testID)
		{
			text := "The quick brown fox jumps over the lazy dog(up) on the hill"
			want := "The quick brown fox jumps over the lazy DOG on the hill"
			got := Up(text)

			if want != got {
				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
			}
			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
		}
	}
}

func TestLow(t *testing.T) {
	t.Log("Given the need to test Up() function behavoir")
	{
		testID := 0
		t.Logf("\tTest %d:\t", testID)
		{
			text := "Ready, set, go (up)!"
			want := "Ready, set, GO!"
			got := Up(text)

			if want != got {
				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
			}
			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
		}

		testID++
		t.Logf("\tTest %d:\t", testID)
		{
			text := "This is a (up) (up) test (up) case"
			want := "This is A (up) TEST case"
			got := Up(text)

			if want != got {
				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
			}
			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
		}

		testID++
		t.Logf("\tTest %d:\t", testID)
		{
			text := "Coding is fun, (up) right?"
			want := "Coding is FUN, right?"
			got := Up(text)

			if want != got {
				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
			}
			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
		}

		testID++
		t.Logf("\tTest %d:\t", testID)
		{
			text := "The quick brown fox jumps over the lazy dog (up)on the hill"
			want := "The quick brown fox jumps over the lazy DOGon the hill"
			got := Up(text)

			if want != got {
				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
			}
			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
		}

		testID++
		t.Logf("\tTest %d:\t", testID)
		{
			text := "The quick brown fox jumps over the lazy dog(up)on the hill"
			want := "The quick brown fox jumps over the lazy DOGon the hill"
			got := Up(text)

			if want != got {
				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
			}
			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
		}

		testID++
		t.Logf("\tTest %d:\t", testID)
		{
			text := "The quick brown fox jumps over the lazy dog(up) on the hill"
			want := "The quick brown fox jumps over the lazy DOG on the hill"
			got := Up(text)

			if want != got {
				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
			}
			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
		}
	}
}
