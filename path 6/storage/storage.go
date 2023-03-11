package storage

type Sotrudnik struct {
	id         int
	surname    string
	firstName  string
	patronymic string
	yearsOld   int
}

type storage interface {
	insert(e Sotrudnik) error
	get(id int) (Sotrudnik, error)
	delete(id int) error
}

type memoryStorage struct {
	data map[int]Sotrudnik
}

func NewMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data: make(map[int]Sotrudnik),
	}
}

func (m *memoryStorage) insert(s Sotrudnik) error {
	m.data[s.id] = s

	return nil
}

func (m *memoryStorage) get(id int) (Sotrudnik, error) {
	return m.data[id], nil
}

func (m *memoryStorage) delete(id int) error {
	delete(m.data, id)
	return nil
}
