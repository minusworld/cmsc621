package dict3

import (
	// "fmt"
	"os"
	"errors"
	"encoding/json"
	"bufio"
	"strings"
)

type KeyRelationship struct {
	Key, Relationship string
}

type KeyRelationshipValue struct {
	Key, Relationship string
	Value interface{}
}

type DICT3 struct {
	Dict map[KeyRelationship]interface{}
	StorageFile string
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func (d *DICT3) Save() error {
	os.Remove(d.StorageFile)
	
	var newErr error
	newErr = nil
	
	fout, fErr := os.OpenFile(d.StorageFile, os.O_WRONLY | os.O_CREATE, 0660)
	defer fout.Close()
	if fErr != nil {
		return fErr
	}
	for kr := range d.Dict {
		//serialize as k,r,{json}
		jsonString, jsonMarshalErr := json.Marshal(d.Dict[kr])
		if jsonMarshalErr != nil {
			newErr = errors.New("Error marshaling JSON at " + kr.Key +","+ kr.Relationship)
		}
		line := kr.Key +"|"+ kr.Relationship +"|"+ string(jsonString)+"\n"
		_, writeErr := fout.Write([]byte(line))
		if writeErr != nil {
			newErr = writeErr
		}
	}
	return newErr
}

func (d *DICT3) Load(filename string) {
	fin, fErr := os.OpenFile(filename, os.O_RDONLY | os.O_CREATE, 0660)
	if fErr != nil {
		panic(fErr)
	}
	defer fin.Close()
	
	scanner := bufio.NewScanner(fin)
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, "|")
		jsonObj := new(interface{})
		jsonUnmarshalErr := json.Unmarshal([]byte(splits[2]), &jsonObj)
		if jsonUnmarshalErr != nil {
			panic(jsonUnmarshalErr)
		}
		d.Dict[KeyRelationship{Key: splits[0], Relationship: splits[1]}] = jsonObj
	}
}

func (d *DICT3) Lookup(key, relation string) (interface{}, error) {
	var err error
	kr := KeyRelationship{key, relation}
	value, ok := d.Dict[kr]
	if ok {
		err = nil
	} else {
		err = errors.New("KeyRelationshp [" + key + "," + relation + "] does not exist")
	}
	return value, err
}

func (d *DICT3) Insert(key string, relation string, value interface{}) error {
	defer d.Save()
	kr := KeyRelationship{key, relation}
	if _, ok := d.Dict[kr]; !ok {
		d.Dict[kr] = value 
		return nil
	} else {
		return errors.New("KeyRelationship [" + key + "," + relation + "] already exists")
	}
}

func (d *DICT3) InsertOrUpdate(key string, relation string, value interface{}) error {
	defer d.Save()
	kr := KeyRelationship{key, relation}
	d.Dict[kr] = value
	return nil
}

func (d *DICT3) Delete(key, relation string) error {
	defer d.Save()
	kr := KeyRelationship{key, relation}
	delete(d.Dict, kr)
	return nil
}

func (d *DICT3) ListKeys() (keys []string, err error) {
	for kr := range d.Dict {
		keys = append(keys, kr.Key)
	}
	return keys, nil
}

func (d *DICT3) ListIDs() (keyrels []string, err error) {
	for kr := range d.Dict {
		keyrels = append(keyrels, kr.Key+","+kr.Relationship)
	}
	return keyrels, nil
}

func Shutdown() {
	os.Exit(0)
}