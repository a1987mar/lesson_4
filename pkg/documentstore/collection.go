package documentstore

import (
	"fmt"
	"reflect"
)

type Collection struct {
	Colect     map[string]Document `json:"colect"`
	PrimaryKey CollectionConfig
}

type CollectionConfig struct {
	PrimaryKey string `json:"cgg"`
}

func (s *Collection) Put(doc Document) {
	// Потрібно перевірити що документ містить поле `{cfg.PrimaryKey}` типу `string`
	keyf, ok := doc.Fields["key"]
	if !ok {
		fmt.Println("Error: Document must contain a 'key' field of type string")
		return
	}

	keys := reflect.TypeOf(keyf.Cfg.PrimaryKey).Kind()
	if keys != reflect.String {
		fmt.Println("Error: 'key' field value must be a string")
		return
	}

	s.Colect[keyf.Cfg.PrimaryKey] = doc
}

func (s *Collection) Get(key string) (*Document, bool) {
	if doc, exists := s.Colect[key]; exists {
		return &doc, true
	}
	return nil, false
}

func (s *Collection) Delete(key string) bool {
	if _, exists := s.Colect[key]; exists {
		delete(s.Colect, key)
		return true
	}
	return false
}

func (s *Collection) List() []Document {
	sList := make([]Document, 0, len(s.Colect))
	for _, v := range s.Colect {
		sList = append(sList, v)
	}
	return sList
}
