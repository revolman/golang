package bytecounter

import (
	"bufio"
	"bytes"
)

// WordCounter накапливает количество байтов
type WordCounter int

// LineCounter ...
type LineCounter int

// Возвращение значения в данном случае это желание соответствовать
// контракту io.Writer
func (wc *WordCounter) Write(p []byte) (int, error) {
	var counter int
	scaner := bufio.NewScanner(bytes.NewBuffer(p))
	scaner.Split(bufio.ScanWords)
	for scaner.Scan() {
		counter++
	}
	*wc = WordCounter(counter)

	return counter, scaner.Err()
}

func (lc *LineCounter) Write(p []byte) (int, error) {
	var counter int
	scaner := bufio.NewScanner(bytes.NewBuffer(p))
	scaner.Split(bufio.ScanLines)
	for scaner.Scan() {
		counter++
	}
	*lc = LineCounter(counter)

	return counter, scaner.Err()
}
