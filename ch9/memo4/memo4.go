package memo

import "sync"

// Func тип функции, результат который нужно кэшировать
type Func func(key string) (interface{}, error)

// тип результата функции
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // Закрывается, когда res готов
}

// Memo кэширует результат выполняемой функции
type Memo struct {
	f     Func
	mu    sync.Mutex // защита cache
	cache map[string]*entry
}

// New cоздаёт новый объект Memo
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

// Get *Memo - безопасен с точки зрения параллельности.
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e != nil {
		// Это первый запрос данног ключа.
		// Эта go-подпрограмма становится ответственной за
		// вычисление значений и оповещение о готовности.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready) // Широковещательное оповещение о готовности
	} else {
		// Это повторный запрос данного ключа.
		memo.mu.Unlock()

		<-e.ready // Ожидание готовности
	}

	return e.res.value, e.res.err
}
