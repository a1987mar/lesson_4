package documentstore

type Store struct {
	StoreMap map[string]Collection `json:"storeMap"`
}

func NewStore() *Store {
	return &Store{
		StoreMap: map[string]Collection{},
	}
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (bool, *Collection) {
	// Створюємо нову колекцію і повертаємо `true` якщо колекція була створена
	// Якщо ж колекція вже створеня то повертаємо `false` та nil

	if _, ok := s.StoreMap[name]; ok {
		return false, nil
	}
	s.StoreMap[name] = Collection{
		Colect:     make(map[string]Document),
		PrimaryKey: *cfg,
	}
	return true, &Collection{}
}

func (s *Store) GetCollection(name string) (*Collection, bool) {
	if colect, ok := s.StoreMap[name]; ok {
		return &colect, true
	}
	return nil, false
}

func (s *Store) DeleteCollection(name string) bool {
	if _, ok := s.StoreMap[name]; ok {
		delete(s.StoreMap, name)
		return true
	}
	return false
}
