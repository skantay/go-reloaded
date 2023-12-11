package tool

import (
	"testing"
)

func TestLow(t *testing.T) {
	t.Log("Given the need to test Low() function behavoir")
	{
		testID := 0
		t.Logf("\tTest %d:\t", testID)
		{
			text := "Ready, set, GO (low)!"
			want := "Ready, set, go!"
			got := Low(text, "")

			if want != got {
				t.Fatalf("\tGOOD\tGot: \"%s\" but wanted: \"%s\"", got, want)
			}
			t.Logf("\tBAD\tGot: \"%s\" matched with: \"%s\"", got, want)
		}

	}
}

func TestLowCount(t *testing.T) {
	t.Log("Given the need to test LowCount() function behavoir")
	{
		testID := 0
		t.Logf("\tTest %d:\t", testID)
		{
			text := "Ready, set, Go (low, 1)!"
			want := "Ready, set, go!"
			got := LowCount(text)

			if want != got {
				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
			}
			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
		}

		testID++
		t.Logf("\tTest %d:\t", testID)
		{
			text := "This Is So Exciting (low, 20)"
			want := "this is so exciting"
			got := LowCount(text)

			if want != got {
				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
			}
			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
		}

		testID++
		t.Logf("\tTest %d:\t", testID)
		{
			text := "This is So Exciting (low, 0)"
			want := "This is So Exciting"
			got := LowCount(text)

			if want != got {
				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
			}
			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
		}

		testID++
		t.Logf("\tTest %d:\t", testID)
		{
			text := "(low, 5)This is so exciting"
			want := "This is so exciting"
			got := LowCount(text)

			if want != got {
				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
			}
			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
		}
	}
}
