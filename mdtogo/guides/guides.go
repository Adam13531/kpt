package guides

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func ParseGuideDocs(files []string) []guide {
	var guides []guide
	for _, path := range files {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		guide := parse(path, string(b))

		guides = append(guides, guide)
	}
	return guides
}

var (
	mdLink = regexp.MustCompile(`^\[.*]:.*$`)
)

func parse(path, content string) guide {
	pathDir := filepath.Dir(path)
	_, name := filepath.Split(pathDir)

	var lines []string
	scanner := bufio.NewScanner(bytes.NewBufferString(content))
	inFormat := false
	indent := false
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "---") {
			inFormat = !inFormat
			continue
		}

		if inFormat {
			continue
		}

		if mdLink.MatchString(line) {
			continue
		}

		if strings.HasPrefix(line, "```") {
			indent = !indent
			continue
		}

		if indent {
			line = "  " + line
		}

		line = strings.ReplaceAll(line, "`", "` + \"`\" + `")

		lines = append(lines, line)
	}

	return guide{
		Name:    name,
		Content: strings.Join(lines, "\n"),
	}
}

func Write(guides []guide, dest, license string) error {
	for _, guide := range guides {
		out := []string{license, `
// Code generated by "mdtogo"; DO NOT EDIT.
package ` + filepath.Base(dest) + "\n"}

		out = append(out, guide.String())

		if _, err := os.Stat(dest); err != nil {
			_ = os.Mkdir(dest, 0700)
		}

		o := strings.Join(out, "\n")
		fileName := fmt.Sprintf("%s.go", guide.Name)
		err := ioutil.WriteFile(filepath.Join(dest, fileName), []byte(o), 0600)
		if err != nil {
			return err
		}
	}
	return nil
}

type guide struct {
	Name    string
	Content string
}

func (g guide) String() string {
	name := strings.Title(g.Name)
	name = strings.ReplaceAll(name, "-", "")

	return fmt.Sprintf("var %sGuide = `%s`", name, g.Content)
}