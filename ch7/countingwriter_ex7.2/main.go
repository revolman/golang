package countingwrite

import (
	"io"
)

// определяем структуру, для которой создадим метод Write
type writer struct {
	w io.Writer // оригинальный io.Writer
	c int64     // счётчик байтов
}

// CountingWriter Должен возращать io.Writer и количество записанных байтов
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	wr := &writer{w, 0} // передаём в тип writer оригинальный io.Writer и счётчик 0
	return wr, &wr.c    // возвращаем значение типа wr. Теперь у него есть метод Write, значит он удовлетворяет контракту io.Writer
}

func (wr *writer) Write(b []byte) (int, error) { // сигнатура в соответствии с контрактом io.Writer
	n, err := wr.w.Write(b) // записываем выходные данные используя метод Write из контакта
	wr.c += int64(n)        // полученное число байтов добавляем в счётчик
	return n, err
}
