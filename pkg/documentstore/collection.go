package documentstore

import (
	"fmt"
)

type Collection struct {
	documents map[string]Document
	config    CollectionConfig
}

type CollectionConfig struct {
	PrimaryKey string `json:"cgg"`
}

func (s *Collection) Put(doc Document) {
	// Потрібно перевірити що документ містить поле `{cfg.PrimaryKey}` типу `string`
	keyf, ok := doc.Fields[s.config.PrimaryKey]
	if !ok {
		fmt.Println("Error: Document must contain a 'key' field of type string")
		return
	}

	if keyf.Type != DocumentFieldTypeString {
		fmt.Println("Error: Key field must be of type string")
		return
	}

	if s.documents == nil {
		s.documents = map[string]Document{}
	}
	s.documents[s.config.PrimaryKey] = doc
}

func (s *Collection) Get(key string) (*Document, bool) {
	if doc, exists := s.documents[key]; exists {
		return &doc, true
	}
	return nil, false
}

func (s *Collection) Delete(key string) bool {
	if _, exists := s.documents[key]; exists {
		delete(s.documents, key)
		return true
	}
	return false
}

func (s *Collection) List() []Document {
	sList := make([]Document, 0, len(s.documents))
	for _, v := range s.documents {
		sList = append(sList, v)
	}
	return sList
}
