package memo

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

// request представляет собой сообщение,
// требующее применения Func к key.
type request struct {
	key      string
	response chan<- result // клиенту нужен только result
}

// Memo кэширует результат выполняемой функции
type Memo struct{ requests chan request }

// New возвращает f с запоминанием.
// В последствии клиенты должны вызывать Close/
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

// Get получает значение функции.
func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

// Close закрывает канал запросов.
func (memo *Memo) Close() { close(memo.requests) }

// Управляющая go-подпрограмма.
func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // call f(key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key)
	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition.
	<-e.ready
	// Send the result to the client.
	response <- e.res
}
