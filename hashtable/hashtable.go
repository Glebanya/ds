package hashtable

import (
	"bytes"
	"encoding/gob"
)

const hashTableSize uint8 = 255

var rand = make([]uint8, hashTableSize)

func getBytes(item interface{}) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(item); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func getHash(item interface{}) (uint8, error) {
	bytes, err := getBytes(item)
	if err != nil {
		return uint8(0), err
	}

	hash := uint8(0)
	for _, byte := range bytes {
		hash = rand[hash^uint8(byte)]
	}

	return hash, nil
}

func init() {
	for i := uint8(0); i < hashTableSize; i++ {
		rand[i] = i
	}
}

type (
	HashMap struct {
		hashMap [hashTableSize][]*interface{}
		len     int
	}
)

func New() *HashMap {

	hashMapArray := [hashTableSize][]*interface{}{}

	return &HashMap{
		hashMap: hashMapArray,
		len:     0,
	}

}

func (hmap *HashMap) Insert(item *interface{}) error {
	key, err := getHash(item)
	if err != nil {
		return err
	}

	if hmap.hashMap[key] == nil {
		hmap.hashMap[key] = make([]*interface{}, 2)
	}

	hmap.hashMap[key] = append(hmap.hashMap[key], item)
	hmap.len++

	return nil
}

func (hmap *HashMap) Has(item *interface{}) (bool, error) {
	key, err := getHash(item)
	if err != nil {
		return false, err
	}

	if hmap.hashMap[key] == nil {
		return false, nil
	}

	for _, v := range hmap.hashMap[key] {
		if v == item {
			return true, nil
		}
	}

	return false, nil
}

func (hmap *HashMap) Delete(item *interface{}) (bool, error) {
	key, err := getHash(item)
	if err != nil {
		return false, err
	}

	if hmap.hashMap[key] == nil {
		return false, nil
	}

	for i, v := range hmap.hashMap[key] {
		if v == item {
			n := len(hmap.hashMap[key])
			if n == 1 {
				hmap.hashMap[key] = nil
			} else {
				hmap.hashMap[key] = append(hmap.hashMap[key][:i], hmap.hashMap[key][i+1:]...)
			}

			hmap.len--

			return true, nil
		}
	}

	return false, nil
}

func (hmap *HashMap) Len() int {
	return hmap.len
}
