package imagecheck

import (
	"os"
	"testing"
)

func TestIsImage(t *testing.T) {
	patterns := []struct {
		filename string
		expected bool
	}{
		{"test001.jpg", true},
		{"test002.png", true},
		{"test003.gif", true},
		{"test004.txt", false},
	}

	pwd, _ := os.Getwd()

	for idx, pattern := range patterns {
		actual := IsImage(pwd + "/../test/" + pattern.filename)
		if pattern.expected != actual {
			t.Errorf("pattern %d: want %t, actual %t", idx, pattern.expected, actual)
		}
	}
}

