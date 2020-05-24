// Упражнение 7.8 - алгоритм подсмотрел у https://github.com/renatofq
package main

import (
	"container/list"
)

type trackSort struct {
	t        []*Track   // список тректов
	sortList *list.List // В sortList хранится список из структур sortProperty. т.е список полей и функций сравнения.
}

// элементы этого типа добавляются в список sortList
// содержит в себе имя поля и функцию сравнения для него
type sortProperty struct {
	name string                 // Заголовок, по которому проводится сортировка
	less func(i, j *Track) bool // Средство сравнения для заголовка. Истинно, если i меньше j.
}

func (ts *trackSort) Len() int      { return len(ts.t) }
func (ts *trackSort) Swap(i, j int) { ts.t[i], ts.t[j] = ts.t[j], ts.t[i] }

// определяю как сравнивать элементы
func (ts *trackSort) Less(i, j int) bool {
	for field := ts.sortList.Front(); field != nil; field = field.Next() { // для каждого sortProperty
		sortProp := field.Value.(*sortProperty)                                   // берётся его значение из листа (тип *sortProperty)
		if !sortProp.less(ts.t[i], ts.t[j]) && !sortProp.less(ts.t[j], ts.t[i]) { // если элементы равны
			continue
		} else {
			return sortProp.less(ts.t[i], ts.t[j]) // в остальных случаях вернуть реальный результат
		}
	}
	return false
}

// Метод для сортировки по полю, учитывает предыдущие ключи.
// В цикле по списку ищется элемемент с полем name соответсвующим fieldName.
// Если он найден, то весь элемент перемещается в начало списка.
func (ts *trackSort) SortBy(fieldName string) {
	l := ts.sortList
	for field := l.Front(); field != nil; field = field.Next() {
		if prop := field.Value.(*sortProperty); prop.name == fieldName {
			l.MoveToFront(field)
			break
		}
	}
}

// определяет функции сравнения для каждого поля, хранит результат в *list.List
func defLessFuncs(tracks []*Track) *trackSort {
	// нужно создать начальную структуру из списка треков и путой лист,
	// который будет далее наполнятся для каждого поля типа Track
	ts := &trackSort{t: tracks, sortList: list.New()}

	// добавляем в лист функции сравнения для каждого поля
	ts.sortList.PushBack(&sortProperty{
		name: "Title", // поле
		less: func(x, y *Track) bool { // функция
			return x.Title < y.Title
		},
	})

	ts.sortList.PushBack(&sortProperty{
		name: "Artist",
		less: func(x, y *Track) bool {
			return x.Artist < y.Artist
		},
	})

	ts.sortList.PushBack(&sortProperty{
		name: "Album",
		less: func(x, y *Track) bool {
			return x.Album < y.Album
		},
	})

	ts.sortList.PushBack(&sortProperty{
		name: "Year",
		less: func(x, y *Track) bool {
			return x.Year < y.Year
		},
	})

	ts.sortList.PushBack(&sortProperty{
		name: "Length",
		less: func(x, y *Track) bool {
			return x.Length < y.Length
		},
	})
	return ts
}
