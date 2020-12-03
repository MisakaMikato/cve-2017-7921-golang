package gip

import (
	"testing"
)

func TestCheckIPFormat(t *testing.T) {
	testStr1 := "ccc192.168.1.123"
	testStr2 := "255.255.255.256"
	testStr3 := "not ip format"
	testStr4 := "192.168.1.1"
	testStr5 := "192.168.1.0/31"
	testStr6 := "192.168.1.0/255"

	cases := []struct {
		origData []byte
		excepted bool
	}{
		{[]byte(testStr1), false},
		{[]byte(testStr2), false},
		{[]byte(testStr3), false},
		{[]byte(testStr4), true},
		{[]byte(testStr5), true},
		{[]byte(testStr6), false},
	}

	for _, c := range cases {
		result := CheckIPFormat(c.origData)
		if result != c.excepted {
			t.Fatalf(
				"CheckIPFormat function failed, origData: %s, execpted:%t, result:%t",
				c.origData,
				c.excepted,
				result,
			)
		}
	}
	t.Log("CheckIPFormat function pass.")
}
