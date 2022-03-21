package macvendor

import (
   "testing"
)

func Test(t *testing.T) {
   macvendor, err := New()
   if err != nil {
      t.Fatal(err)
   }

   testcases := []string{"CE:9F:22:01:02:03", "56:C6:51:01:02:03"}

   for _, tc := range testcases {
      _, err := macvendor.Lookup(tc)
      if err != nil {
         t.Error(err)
      }
   }
}
