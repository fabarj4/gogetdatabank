package gogetdatabank

import "testing"

func TestDataBank(t *testing.T) {
	t.Run("Test Get from atm bersama", func(t *testing.T) {
		_, err := DataBank()
		if err != nil {
			t.Fatal(err)
		}
	})
}
