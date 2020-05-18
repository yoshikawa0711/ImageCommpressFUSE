package getfile

import (
	"reflect"
	"os"
	"path/filepath"
	"testing"
)

func TestGetFiles(t *testing.T) {
	pwd, _ := os.Getwd()

	patterns := []struct {
		filepath string
		expected []string
	}{
		{pwd + "/../test/test005", []string{"test006.png"}},
		{pwd + "/../test/", []string{"test001.jpg", "test002.png", "test003.gif", "test004.txt", "test005/test006.png"}},
	}

	for idx, pattern := range patterns {
		actual := GetFiles(pattern.filepath)
		_AddPaths(pattern.filepath, pattern.expected)

		if !(reflect.DeepEqual(actual, pattern.expected)) {
			t.Errorf("pattern %d: want %v, actual %v", idx, pattern.expected, actual)
		}
	}

}

func _AddPaths(prepath string, files []string) {
	for idx, file := range files {
		files[idx] = filepath.Join(prepath, file)
	}
}
