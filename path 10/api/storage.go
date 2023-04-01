package main

import (
	"errors"
	"sync"
)

type Sotrudnik struct {
	Id         int    `json:"id"`
	Surname    string `json:"surname"`
	FirstName  string `json:"firstname"`
	Patronymic string `json:"patronymic"`
	YearsOld   int    `json:"yo"`
}

type Storage interface {
	Insert(e *Sotrudnik)
	Get(id int) (Sotrudnik, error)
	Update(id int, e Sotrudnik)
	Delete(id int)
}

type MemoryStorage struct {
	counter int
	data    map[int]Sotrudnik
	sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data:    make(map[int]Sotrudnik),
		counter: 1,
	}
}

func (m *MemoryStorage) Insert(s *Sotrudnik) {
	m.Lock()

	s.Id = m.counter
	m.data[s.Id] = *s

	m.counter++

	m.Unlock()
}

func (m *MemoryStorage) Get(id int) (Sotrudnik, error) {
	m.Lock()
	defer m.Unlock()

	sotr, ok := m.data[id]
	if !ok {
		return sotr, errors.New("Sotr not found")
	}

	return sotr, nil
}

func (m *MemoryStorage) Update(id int, s Sotrudnik) {
	m.Lock()
	m.data[id] = s
	m.Unlock()
}

func (m *MemoryStorage) Delete(id int) {
	m.Lock()
	delete(m.data, id)
	m.Unlock()
}
