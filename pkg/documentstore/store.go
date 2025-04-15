package documentstore

type Store struct {
	collections map[string]Collection
}

func NewStore() *Store {
	return &Store{
		collections: map[string]Collection{},
	}
}

func (s *Store) CreateCollection(name string) (bool, *Collection) {
	// Створюємо нову колекцію і повертаємо `true` якщо колекція була створена
	// Якщо ж колекція вже створеня то повертаємо `false` та nil

	if _, ok := s.collections[name]; ok {
		return false, nil
	}
	collect := Collection{}
	s.collections[name] = collect
	return true, &Collection{}
}

func (s *Store) GetCollection(name string) (*Collection, bool) {
	if colect, ok := s.collections[name]; ok {
		return &colect, true
	}
	return nil, false
}

func (s *Store) DeleteCollection(name string) bool {
	if _, ok := s.collections[name]; ok {
		delete(s.collections, name)
		return true
	}
	return false
}
