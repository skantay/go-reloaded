package tool

// import (
// 	"testing"
// )

// func TestCap(t *testing.T) {
// 	t.Log("Given the need to test Cap() function behavoir")
// 	{
// 		testID := 0
// 		t.Logf("\tTest %d:\t", testID)
// 		{
// 			text := "Welcome to the Brooklyn bridge (cap)"
// 			want := "Welcome to the Brooklyn Bridge"
// 			got := Cap(text)

// 			if want != got {
// 				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
// 			}
// 			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
// 		}

// 		testID++
// 		t.Logf("\tTest %d:\t", testID)
// 		{
// 			text := "Welcome to the Brooklyn Bridge (cap)"
// 			want := "Welcome to the Brooklyn Bridge"
// 			got := Cap(text)

// 			if want != got {
// 				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
// 			}
// 			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
// 		}

// 		testID++
// 		t.Logf("\tTest %d:\t", testID)
// 		{
// 			text := "Welcome to the Brooklyn Bridge, (cap)"
// 			want := "Welcome to the Brooklyn Bridge, (cap)"
// 			got := Cap(text)

// 			if want != got {
// 				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
// 			}
// 			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
// 		}

// 		testID++
// 		t.Logf("\tTest %d:\t", testID)
// 		{
// 			text := "Welcome to the Brooklyn bridge(cap) BRIDG"
// 			want := "Welcome to the Brooklyn Bridge BRIDG"
// 			got := Cap(text)

// 			if want != got {
// 				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
// 			}
// 			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
// 		}
// 	}
// }

// func TestCapCount(t *testing.T) {
// 	t.Log("Given the need to test CapCount() function behavoir")
// 	{
// 		testID := 0
// 		t.Logf("\tTest %d:\t", testID)
// 		{
// 			text := "Ready, set, go (cap, 1)!"
// 			want := "Ready, set, Go!"
// 			got := CapCount(text)

// 			if want != got {
// 				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
// 			}
// 			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
// 		}

// 		testID++
// 		t.Logf("\tTest %d:\t", testID)
// 		{
// 			text := "This is so exciting (cap, 20)"
// 			want := "This Is So Exciting"
// 			got := CapCount(text)

// 			if want != got {
// 				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
// 			}
// 			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
// 		}

// 		testID++
// 		t.Logf("\tTest %d:\t", testID)
// 		{
// 			text := "This is so exciting (cap, 0)"
// 			want := "This is so exciting"
// 			got := CapCount(text)

// 			if want != got {
// 				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
// 			}
// 			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
// 		}

// 		testID++
// 		t.Logf("\tTest %d:\t", testID)
// 		{
// 			text := "(cap, 5)This is so exciting"
// 			want := "This is so exciting"
// 			got := CapCount(text)

// 			if want != got {
// 				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
// 			}
// 			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
// 		}
// 	}
// }
