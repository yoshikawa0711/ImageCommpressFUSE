package getfile

import (
	"io/ioutil"
	"path/filepath"
)

func GetFiles(dirpath string) []string{
	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, GetFiles(filepath.Join(dirpath, file.Name()))...)
			continue
		}

		paths = append(paths, filepath.Join(dirpath, file.Name()))

	}

	return paths
}
