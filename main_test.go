package main

import "testing"

func TestAge(t *testing.T) {
    result := Age(17)
    if result != 16 {
        t.Errorf("Ожидалось 16, а получено %d", result)
    }
}
