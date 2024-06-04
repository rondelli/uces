package main

import (
	"encoding/json"
	"fmt"
	bolt "go.etcd.io/bbolt"
	"log"
	"strconv"
)

type Alumne struct {
	Legajo   int
	Nombre   string
	Apellido string
}

//...

//FUNC OMIT

//...

func main() {
	db, err := bolt.Open("mi-uces.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cristina := Alumne{1, "Cristina", "Kirchner"}
	data, err := json.Marshal(cristina)
	if err != nil {
		log.Fatal(err)
	}

	CreateUpdate(db, "alumne", []byte(strconv.Itoa(cristina.Legajo)), data)

	resultado, err := ReadUnique(db, "alumne", []byte(strconv.Itoa(cristina.Legajo)))

	fmt.Printf("%s\n", resultado)
}

//CREATE OMIT

func CreateUpdate(db *bolt.DB, bucketName string, key []byte, val []byte) error {
	// abre transacción de escritura
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	b, _ := tx.CreateBucketIfNotExists([]byte(bucketName))

	err = b.Put(key, val)
	if err != nil {
		return err
	}

	// cierra transacción
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

//READUNIQUE OMIT

func ReadUnique(db *bolt.DB, bucketName string, key []byte) ([]byte, error) {
	var buf []byte

	// abre una transacción de lectura
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		buf = b.Get(key)
		return nil
	})

	return buf, err
}

// END OMIT
