package storage

type Sotrudnik struct {
	id         int
	surname    string
	firstName  string
	patronymic string
	yearsOld   int
}

type Storage interface {
	Insert(e Sotrudnik) error
	Get(id int) (Sotrudnik, error)
	Delete(id int) error
}

type MemoryStorage struct {
	data map[int]Sotrudnik
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[int]Sotrudnik),
	}
}

func (m *MemoryStorage) Insert(s Sotrudnik) error {
	m.data[s.id] = s

	return nil
}

func (m *MemoryStorage) Get(id int) (Sotrudnik, error) {
	return m.data[id], nil
}

func (m *MemoryStorage) Delete(id int) error {
	delete(m.data, id)
	return nil
}
