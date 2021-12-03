package hashtable

import (
	"fmt"
	"testing"
)

func TestAddToHashTable(t *testing.T) {
	hashTable := New()
	for i := 0; i < 2000; i++ {
		err := hashTable.Insert(i)
		if err != nil {
			t.Errorf(fmt.Sprintf("Unexpected error occured while adding element: %s.", err.Error()))
		}
	}

	for i := 0; i < 2000; i++ {
		success, err := hashTable.Has(i)
		if !success {
			t.Errorf(fmt.Sprintf("Unexpected hashtable anwser. Hash table must contain: %d", i))
		}
		if err != nil {
			t.Errorf(fmt.Sprintf("Unexpected error while has operation occured: %s.", err.Error()))
		}
	}
	length := hashTable.Len()
	if length != 2000 {
		t.Errorf(fmt.Sprintf("Unexpected hashtable len %d but must be %d", length, 2000))
	}
}

func TestAddAndDeleteToHashTable(t *testing.T) {

	hashTable := New()
	for i := 0; i < 2000; i++ {
		err := hashTable.Insert(i)
		if err != nil {
			t.Errorf(fmt.Sprintf("Unexpected error occured while adding element: %s.", err.Error()))
		}
	}
	has, err := hashTable.Delete(0)
	if !has {
		t.Errorf(fmt.Sprintf("Unexpected behavior occured while deleting element."))
	}
	if err != nil {
		t.Errorf(fmt.Sprintf("Unexpected error occured while deleting element: %s.", err.Error()))
	}

	for i := 1; i < 2000; i++ {
		success, err := hashTable.Has(i)
		if !success {
			t.Errorf(fmt.Sprintf("Unexpected hashtable anwser. Hash table must contain: %d", i))
		}
		if err != nil {
			t.Errorf(fmt.Sprintf("Unexpected error occured: %s.", err.Error()))
		}
	}

	success, err := hashTable.Has(0)
	if success {
		t.Errorf(fmt.Sprintf("Unexpected hashtable anwser. Hash table must not contain: 0"))
	}

}
