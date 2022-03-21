package macvendor

import (
   "fmt"
   "testing"
)

func Test(t *testing.T) {
   macvendor := New()
   fmt.Println(macvendor.Lookup("CE:9F:22:01:02:03"))
   fmt.Println(macvendor.Lookup("56:C6:51:01:02:03"))
}
