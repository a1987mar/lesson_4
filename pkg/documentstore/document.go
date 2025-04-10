package documentstore

type DocumentFieldType string

const (
	DocumentFieldTypeString DocumentFieldType = "string"
	DocumentFieldTypeNumber DocumentFieldType = "number"
	DocumentFieldTypeBool   DocumentFieldType = "bool"
	DocumentFieldTypeArray  DocumentFieldType = "array"
	DocumentFieldTypeObject DocumentFieldType = "object"
)

type DocumentField struct {
	Type DocumentFieldType `json:"type"`
	Cfg  CollectionConfig  `json:"Cfg"`
}

type Document struct {
	Fields map[string]DocumentField `json:"fields"`
}
