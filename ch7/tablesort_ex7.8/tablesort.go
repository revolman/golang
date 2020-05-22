package tablesort

import (
	"container/list"
	"time"
)

// Track тип трека
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}



type trackSort struct {
	t        []*Track		// список тректов
	sortList *list.List // В sortList хранится список из структур sortProperty. т.е список полей и функций сравнения.
}

// элементы этого типа добавляются в список sortList
// содержит в себе имя поля и функцию сравнения для него
type sortProperty struct {
	name string										// Заголовок, по которому проводится сортировка
	less func(i, j *Track) bool		// Средство сравнения для заголовка. Истинно, если i меньше j.
}

func (ts *trackSort) Len() int { return len(ts.t) }
func (ts *trackSort) Swap(i, j int) { ts.t[i], ts.t[j] = ts.t[j], ts.t[i] }

// определяю как сравнивать элементы
func (ts *trackSort) Less(i, j int) bool {
	for field := ts.sortList.Front(); field != nil; field = field.Next() { // для каждого sortProperty
		sortProp := field.Value.(*sortProperty)  // берётся его значение из листа (тип *sortProperty)
		if !sortProp.less(ts.t[i], ts.t[j]) && !sortProp.less(ts.t[j], ts.t[i]) { // если элементы равны
			continue
		} else {
			return sortProp.less(ts.t[i], ts.t[j]) // в остальных случаях вернуть реальный результат
		}
	}
	return false
}

func ()  {
	
}
