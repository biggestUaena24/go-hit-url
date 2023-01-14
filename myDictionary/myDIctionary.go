package myDictionary

import "errors"

var errKeyNotFound = errors.New("No such key")
var errExist = errors.New("Already exist a the same key")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]

	if exists == true{
		return value, nil
	}
	return "", errKeyNotFound
}

func (d Dictionary) Add(key string, value string) error {
	if _,exist := d[key]; exist == true{
		return errExist
	}
	d[key] = value
	return nil
}

func (d Dictionary) Update (key string, value string) error {
	if _,exist := d[key]; exist == true{
		d[key] = value
		return nil
	}
	return errKeyNotFound
}

func (d Dictionary) Delete (key string) error {
	if _,exist := d[key]; exist == true{
		delete(d,key)
		return nil
	}
	return errKeyNotFound
}