package commpressimage

import (
	"os"
	"testing"
)

func TestImageCommpress(t *testing.T) {
	patterns := []struct {
		filename string
		expected bool
	}{
		{"test001.jpg", true},
		{"test002.png", false},
		{"test003.gif", false},
		{"test004.txt", false},
	}

	pwd, _ := os.Getwd()

	for idx, pattern := range patterns {
		actual := ImageCommpress(pwd + "/../test/" + pattern.filename)
		if pattern.expected != actual {
			t.Errorf("pattern %d: want %t, actual %t", idx, pattern.expected, actual)
		}
	}
}

