package limitreader

import "io"

// LmReader ...
type LmReader struct {
	s io.Reader // обвёрнутый ридер
	n int64     // ограничитель
}

// LimitReader ...
func LimitReader(r io.Reader, n int64) io.Reader {
	return &LmReader{r, n} // возвращает обработанный ридер
}

func (r *LmReader) Read(b []byte) (n int, err error) {
	if r.n <= 0 {
		return 0, io.EOF // всё прочитали => EOF
	}

	if int64(len(b)) > r.n {
		b = b[0:r.n] // обрезали входную строку до ограничителя
	}
	n, err = r.s.Read(b) // прочитали обрезанную строку b в r.s
	r.n -= int64(n)      // почистили структуру
	return
}
