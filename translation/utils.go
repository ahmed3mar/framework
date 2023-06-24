package translation

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/spf13/cast"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func filePathWalkDir(root string) ([]string, error) {
	var files []string

	if _, err := os.Stat(root); os.IsNotExist(err) {
		return files, nil
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			// only take json files
			if filepath.Ext(path) != ".json" {
				return nil
			}

			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func getDirsInPath(root string) ([]string, error) {
	var dirs []string

	if _, err := os.Stat(root); os.IsNotExist(err) {
		return dirs, nil
	}

	// Get all files in the directory.
	files, err := ioutil.ReadDir(root)
	if err != nil {
		return dirs, err
	}

	// Iterate over the files and filter out directories.
	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		// Add the directory to the list of directories.
		dirs = append(dirs, filepath.Join(root, file.Name()))
	}

	return dirs, nil
}

func parseTranslation(str string, replace http.Json) string {
	for k, v := range replace {
		str = strings.ReplaceAll(str, ":"+cast.ToString(k), cast.ToString(v))
	}
	return str
}