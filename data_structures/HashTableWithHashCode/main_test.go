package main

import (
    "testing"
    "./hashtable"
)

// Hashable type for strings
type StringHash string

func (str StringHash) HashCode() int64 {
    var sum int64

    for _, r := range str {
        sum += int64(r)
    }

    return sum
}

func (str StringHash) Compare(cmp interface{}) bool {
    cmpstr, ok := cmp.(StringHash)
    if !ok { return false; }

    return cmpstr == str
}

func TestHashTablePutAndGet(t *testing.T) {
    table := hashtable.CreateHashTable("No name", -1)

    // Putting elements into hash table
    table.PutElement(StringHash("George"), "Grillman")
    table.PutElement(StringHash("Michael"), "Bazovsky")
    table.PutElement(StringHash("Jill"), "Valentine")

    // Lets get some elements
    elem, _ := table.GetElement(StringHash("Michael"))
    elemStr, ok := elem.(string)

    if !ok || elemStr != "Bazovsky" {
        t.Errorf("TestHashTablePutAndGet: Excepted Bazovsky, got %v\n", elemStr)
    }

    elem, ok = table.GetElement(StringHash("Harry"))

    if ok {
        t.Error("TestHashTablePutAndGet: There is not such key as Harry, but function GetElement returned true as second parameter")
    }

    elemStr, ok = elem.(string)
    if !ok || elemStr != "No name" {
        t.Errorf("TestHashTablePutAndGet: Excepted No name, got %v\n", elemStr)
    }
}
