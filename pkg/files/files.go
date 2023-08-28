package files

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/mjevans93308/platformscience/util/localctx"
)

// ProcessFile builds the path to the current directory
// and uses that to open the supplied filename
// then reads from the file and processes the contents
func ProcessFile(myCtx *localctx.Localctx, filename string) ([]string, error) {
	var content []string
	currDir, err := os.Getwd()
	if err != nil {
		myCtx.Logger.Errorf("Error when getting current directory:%s", err)
		return nil, err
	}
	fullpath := filepath.Join(currDir, filename)
	file, err := os.Open(fullpath)
	if err != nil {
		myCtx.Logger.Errorf("Error when opening file at filepath: %s\n%s", fullpath, err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// strip whitespace from string before saving
		// count vowels and consonants now while we're in the goroutine
		// to save time downstream
		content = append(content, line)
	}

	if err := scanner.Err(); err != nil {
		myCtx.Logger.Errorf("Error when scanning file:\n%s", err)
		return nil, err
	}

	return content, nil
}
