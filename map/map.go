package _map

import (
	"errors"
	"fmt"
)

type Map struct {
	Dictionary map[int]int
}

func NewMap() *Map{
	return &Map{
		Dictionary: make(map[int]int)}
}

func (m *Map) Add(value int) {
	if m.Dictionary[value] > 0{
		m.Dictionary[value] += 1
	}else{
		m.Dictionary[value] = 1
	}
}

func (m *Map) ContainsKey(key int) bool {
	return m.Dictionary[key] > 0
}

func (m* Map) GetValue(key int) (int, error) {
	if m.ContainsKey(key) == false{
		return 0, errors.New("key is not found")
	}
	return m.Dictionary[key], nil
}

func (m *Map) Delete(key int) error{
	if m.ContainsKey(key) == false{
		return errors.New("key is not found")
	}
	delete(m.Dictionary, key)
	return nil
}

func (m *Map) PrintMapsKeyValues() {
	for key, value :=range m.Dictionary{
		fmt.Printf("Key: %v and value: %d\n", key, value)
	}
}

