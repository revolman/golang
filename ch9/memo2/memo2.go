package memo

import "sync"

// Memo кэширует результат выполняемой функции
type Memo struct {
	f     Func
	mu    sync.Mutex // защита cache
	cache map[string]result
}

// Func тип функции, результат который нужно кэшировать
type Func func(key string) (interface{}, error)

// тип результата функции
type result struct {
	value interface{}
	err   error
}

// New cоздаёт новый объект Memo
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// Get *Memo - безопасен с точки зрения параллельности.
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	if !ok { // если в отображении нет такого значения, то его нужно записать
		res.value, res.err = memo.f(key) // получить значение функции, которая содержится в объекте
		memo.cache[key] = res            // записать результат в кэш
	}
	memo.mu.Unlock()
	return res.value, res.err
}
