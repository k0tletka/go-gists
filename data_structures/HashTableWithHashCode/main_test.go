package main

import (
    "testing"
    "github.com/k0tletka/go-gists/data_structures/HashTableWithHashCode/hashtable"
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

func TestHashTableRebuildBucket(t *testing.T) {
    table := hashtable.CreateHashTable(nil, 4)

    // Putting elements with hashcodes, thats will
    // produce same relative hash with bucketSize equals 4
    table.PutElement(StringHash("a"), 1)
    table.PutElement(StringHash("e"), 1)

    // Test to get element
    elem, ok := table.GetElement(StringHash("a"))
    if !ok {
        t.Error("TestHashTableRebuildBucket: Element with key a not found, while it must be available")
    }

    if elemNum, ok := elem.(int); !ok || elemNum != 1 {
        t.Errorf("TestHashTableRebuildBucket: Element with key a: Except 1, got %v\n", elemNum)
    }

    elem, ok = table.GetElement(StringHash("e"))
    if !ok {
        t.Error("TestHashTableRebuildBucket: Element with key e not found, while it must be available")
    }

    if elemNum, ok := elem.(int); !ok || elemNum != 1 {
        t.Errorf("TestHashTableRebuildBucket: Element with key e: Except 1, got %v\n", elemNum)
    }
}
