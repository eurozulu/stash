package stash

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

var PageDelimiter = []byte("\n---\n--\n")

type fileStashPage struct {
	filepath   string
	count      int
	bufferSize int
}

func (fp *fileStashPage) Count() int {
	if fp.count == 0 {
		if c, err := fp.readFileCount(); err != nil {
			log.Printf("Failed to refresh page count  %v", err)
		} else {
			fp.count = c
		}
	}
	return fp.count
}

func (fp *fileStashPage) Get(offset int) ByteValue {
	result := fp.GetRange(offset, 1)
	if len(result) == 0 {
		return nil
	}
	return result[0]
}

func (fp *fileStashPage) GetRange(offset, length int) []ByteValue {
	result, err := fp.readFileRange(offset, length)
	if err != nil {
		log.Printf("Failed to read byte data at offset %d  %v", offset, err)
		return nil
	}
	return result
}

func (fp *fileStashPage) PutRange(v ...ByteValue) int {
	panic("not implemented")
}

func (fp fileStashPage) writeFileRange(data []ByteValue) (int, error) {
	file, err := os.Open(fp.filepath)
	if err != nil {
		return 0, fmt.Errorf("Error reading file: %v", err)
	}
	defer file.Close()
	return 0, err
}

func (fp fileStashPage) readFileRange(offset, length int) ([]ByteValue, error) {
	if offset >= fp.Count() {
		return nil, fmt.Errorf("offset %d is out of range of page size %d", offset, fp.count)
	}
	file, err := os.Open(fp.filepath)
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %v", err)
	}
	defer file.Close()

	var split = func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		index := bytes.Index(data, PageDelimiter)
		if index < 0 {
			return 0, nil, nil
		}
		return index + len(PageDelimiter), data[0:index], nil
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(split)
	if fp.bufferSize > 0 {
		scanner.Buffer(make([]byte, fp.bufferSize), bufio.MaxScanTokenSize)
	}

	result := make([]ByteValue, length)
	var count int
	for scanner.Scan() {
		entry := scanner.Bytes()
		if count >= offset {
			index := count - offset
			result[index] = entry
			if index+1 >= length {
				// the end of the range
				break
			}
		}
		count++
	}
	return result, nil
}

func (fp *fileStashPage) readFileCount() (int, error) {
	f, err := os.Open(fp.filepath)
	if err != nil {
		return 0, fmt.Errorf("Error reading file %s: %v", fp.filepath, err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	if !scanner.Scan() {
		return 0, fmt.Errorf("failed to read initial count in filepage  %v", scanner.Err())
	}
	return strconv.Atoi(scanner.Text())
}
