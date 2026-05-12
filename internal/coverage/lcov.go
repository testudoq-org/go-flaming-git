package coverage

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

type fileLines map[string]map[int]int

// ConvertProfileToLCOV converts a Go coverage profile into LCOV format.
func ConvertProfileToLCOV(in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)
	linesByFile := make(fileLines)

	lineNo := 0
	for scanner.Scan() {
		lineNo++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if lineNo == 1 {
			if !strings.HasPrefix(line, "mode:") {
				return fmt.Errorf("invalid coverage profile header: %q", line)
			}
			continue
		}

		if err := collectLineCoverage(linesByFile, line); err != nil {
			return fmt.Errorf("parse coverage line %d: %w", lineNo, err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("read coverage profile: %w", err)
	}

	return writeLCOV(out, linesByFile)
}

func collectLineCoverage(linesByFile fileLines, line string) error {
	fields := strings.Fields(line)
	if len(fields) != 3 {
		return fmt.Errorf("expected 3 fields, got %d", len(fields))
	}

	pathAndRange := fields[0]
	countRaw := fields[2]

	count, err := strconv.Atoi(countRaw)
	if err != nil {
		return fmt.Errorf("invalid execution count %q: %w", countRaw, err)
	}

	parts := strings.SplitN(pathAndRange, ":", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid location %q", pathAndRange)
	}

	file := parts[0]
	rangePart := parts[1]
	rangeParts := strings.SplitN(rangePart, ",", 2)
	if len(rangeParts) != 2 {
		return fmt.Errorf("invalid range %q", rangePart)
	}

	start, err := parseLineNumber(rangeParts[0])
	if err != nil {
		return fmt.Errorf("invalid start location %q: %w", rangeParts[0], err)
	}
	end, err := parseLineNumber(rangeParts[1])
	if err != nil {
		return fmt.Errorf("invalid end location %q: %w", rangeParts[1], err)
	}

	if end < start {
		return fmt.Errorf("invalid range %d-%d", start, end)
	}

	if _, ok := linesByFile[file]; !ok {
		linesByFile[file] = make(map[int]int)
	}

	for n := start; n <= end; n++ {
		if count > linesByFile[file][n] {
			linesByFile[file][n] = count
		}
	}

	return nil
}

func parseLineNumber(segment string) (int, error) {
	parts := strings.SplitN(segment, ".", 2)
	if len(parts) != 2 {
		return 0, fmt.Errorf("missing column separator")
	}

	lineNo, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}

	return lineNo, nil
}

func writeLCOV(out io.Writer, linesByFile fileLines) error {
	files := make([]string, 0, len(linesByFile))
	for file := range linesByFile {
		files = append(files, file)
	}
	sort.Strings(files)

	for _, file := range files {
		if _, err := fmt.Fprintln(out, "TN:"); err != nil {
			return err
		}
		if _, err := fmt.Fprintf(out, "SF:%s\n", file); err != nil {
			return err
		}

		lineNos := make([]int, 0, len(linesByFile[file]))
		for lineNo := range linesByFile[file] {
			lineNos = append(lineNos, lineNo)
		}
		sort.Ints(lineNos)

		for _, lineNo := range lineNos {
			if _, err := fmt.Fprintf(out, "DA:%d,%d\n", lineNo, linesByFile[file][lineNo]); err != nil {
				return err
			}
		}

		if _, err := fmt.Fprintln(out, "end_of_record"); err != nil {
			return err
		}
	}

	return nil
}
