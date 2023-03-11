package main

import "fmt"

type sotrudnik struct {
	id         int
	surname    string
	firstName  string
	patronymic string
	yearsOld   int
}

type storage interface {
	insert(e sotrudnik) error
	get(id int) (sotrudnik, error)
	delete(id int) error
}

type memoryStorage struct {
	data map[int]sotrudnik
}

func NewMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data: make(map[int]sotrudnik),
	}
}

func (m *memoryStorage) insert(s sotrudnik) error {
	m.data[s.id] = s

	return nil
}

func (m *memoryStorage) get(id int) (sotrudnik, error) {
	return m.data[id], nil
}

func (m *memoryStorage) delete(id int) error {
	delete(m.data, id)
	return nil
}

func main() {
	s := *new(storage)

	fmt.Printf("%v\n", s == nil)
	fmt.Printf("%T\n", s)

	s = NewMemoryStorage()

	fmt.Printf("%v\n", s == nil)
	fmt.Printf("%T\n", s)
}
