package main

import (
	"fmt"
	"lesson4/pkg/documentstore"
)

func main() {

	k1 := documentstore.CollectionConfig{
		PrimaryKey: "ID1",
	}
	k2 := documentstore.CollectionConfig{
		PrimaryKey: "ID2",
	}
	k3 := documentstore.CollectionConfig{
		PrimaryKey: "ID3",
	}
	k4 := documentstore.CollectionConfig{
		PrimaryKey: "ID4",
	}
	k41 := documentstore.CollectionConfig{
		PrimaryKey: "ID4",
	}

	d1 := documentstore.Document{Fields: make(map[string]documentstore.DocumentField)}
	d1.Fields["key"] = documentstore.DocumentField{
		Type: documentstore.DocumentFieldTypeString,
		Cfg:  k1,
	}

	d2 := documentstore.Document{Fields: make(map[string]documentstore.DocumentField)}
	d2.Fields["key"] = documentstore.DocumentField{
		Type: documentstore.DocumentFieldTypeString,
		Cfg:  k2,
	}

	d21 := documentstore.Document{Fields: make(map[string]documentstore.DocumentField)}
	d21.Fields["key"] = documentstore.DocumentField{
		Type: documentstore.DocumentFieldTypeString,
		Cfg:  k2,
	}

	d3 := documentstore.Document{Fields: make(map[string]documentstore.DocumentField)}
	d3.Fields["key"] = documentstore.DocumentField{
		Type: documentstore.DocumentFieldTypeString,
		Cfg:  k3,
	}

	d4 := documentstore.Document{Fields: make(map[string]documentstore.DocumentField)}
	d4.Fields["key"] = documentstore.DocumentField{
		Type: documentstore.DocumentFieldTypeString,
		Cfg:  k4,
	}

	d4_1 := documentstore.Document{Fields: make(map[string]documentstore.DocumentField)}
	d4_1.Fields["key"] = documentstore.DocumentField{
		Type: documentstore.DocumentFieldTypeString,
		Cfg:  k41,
	}

	sd := documentstore.NewStore()
	sd.CreateCollection("db_1", &k1)
	sd.CreateCollection("db_2", &k2)
	sd.CreateCollection("db_3", &k3)
	sd.CreateCollection("db_4", &k4)
	sd.DeleteCollection("db_3")

	s1, _ := sd.GetCollection("db_4")
	s1.Put(d4_1)
	s1.Put(d3)
	s1.Put(d4)
	sd.StoreMap["db_4"] = *s1
	fmt.Println(s1)

	s2, _ := sd.GetCollection("db_2")
	s2.Put(d2)
	s2.Put(d21)
	s2.Get("ID2")

	sd.StoreMap["db_2"] = *s2

}
