package tool

// import (
// 	"testing"
// )

// func TestBin(t *testing.T) {
// 	t.Log("Given the need to test Bin() function behavoir")
// 	{
// 		testID := 0
// 		t.Logf("\tTest %d:\t", testID)
// 		{
// 			text := "It has been 10 (bin) years"
// 			want := "It has been 2 years"
// 			got := Bin(text)

// 			if want != got {
// 				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
// 			}
// 			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
// 		}

// 		testID++
// 		t.Logf("\tTest %d:\t", testID)
// 		{
// 			text := "It has been 1A0 (bin) years"
// 			want := "It has been 0 years"
// 			got := Bin(text)

// 			if want != got {
// 				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
// 			}
// 			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
// 		}

// 		testID++
// 		t.Logf("\tTest %d:\t", testID)
// 		{
// 			text := "It has been 01, (bin) years"
// 			want := "It has been 01, (bin) years"
// 			got := Bin(text)

// 			if want != got {
// 				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
// 			}
// 			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
// 		}

// 		testID++
// 		t.Logf("\tTest %d:\t", testID)
// 		{
// 			text := "(bin) years"
// 			want := "(bin) years"
// 			got := Bin(text)

// 			if want != got {
// 				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
// 			}
// 			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
// 		}

// 		testID++
// 		t.Logf("\tTest %d:\t", testID)
// 		{
// 			text := "years (bin)"
// 			want := "0"
// 			got := Bin(text)

// 			if want != got {
// 				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
// 			}
// 			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
// 		}

// 		testID++
// 		t.Logf("\tTest %d:\t", testID)
// 		{
// 			text := "(bin)"
// 			want := "(bin)"
// 			got := Bin(text)

// 			if want != got {
// 				t.Fatalf("\t%s\tGot: \"%s\" but wanted: \"%s\"", failed, got, want)
// 			}
// 			t.Logf("\t%s\tGot: \"%s\" matched with: \"%s\"", success, got, want)
// 		}
// 	}
// }
