/*
 * Charsets
 */

package pwdgen

// a map of possible charsets.
//var charsets map[string] string

func InitializeCharsets() {
//    charsets := make(map[string] string) {
//        "alpha":    "abcdefghijklmnopqrstuvwxyz",
//        "ALPHA":    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
//        "num":      "0123456789",
//        "specials": ",.!?#@/+-*$%&()",
//    }
}


// A way to get all the charsets built into the application.
func GetCharsetsNames() []string {
    return []string{"alpha", "ALPHA", "num", "specials"}
}

