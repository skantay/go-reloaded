package tool

// import (
// 	"testing"
// )

// func TestHex(t *testing.T) {
// 	t.Log("Given the need to test Hex() function behavoir")
// 	{
// 		testID := 0
// 		t.Logf("\tTest %d:\t", testID)
// 		{
// 			text := "1E (hex) files were added"
// 			want := "30 files were added"
// 			got := Hex(text)

// 			if want != got {
// 				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
// 			}
// 			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
// 		}

// 		testID++
// 		t.Logf("\tTest %d:\t", testID)
// 		{
// 			text := "1EEEEE (hex) files were added"
// 			want := "2027246 files were added"
// 			got := Hex(text)

// 			if want != got {
// 				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
// 			}
// 			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
// 		}

// 		testID++
// 		t.Logf("\tTest %d:\t", testID)
// 		{
// 			text := "(hex) files were added"
// 			want := "(hex) files were added"
// 			got := Hex(text)

// 			if want != got {
// 				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
// 			}
// 			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
// 		}

// 		testID++
// 		t.Logf("\tTest %d:\t", testID)
// 		{
// 			text := "files were added (hex)"
// 			want := "files were 712173"
// 			got := Hex(text)

// 			if want != got {
// 				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
// 			}
// 			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
// 		}

// 		testID++
// 		t.Logf("\tTest %d:\t", testID)
// 		{
// 			text := "files were addezd (hex)"
// 			want := "files were 0"
// 			got := Hex(text)

// 			if want != got {
// 				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
// 			}
// 			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
// 		}
// 	}
// }
