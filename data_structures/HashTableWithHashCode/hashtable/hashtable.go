package hashtable

import (
    "math"
)

// This interface allows to use value as key in out hashtable
type Hashable interface {
    HashCode() int64
    Compare(interface{}) bool
}

// Type of hashtable
type HashTable struct {
    bucket []*hashTableElement
    defaultValue interface{}
}

// Struct initializator
func CreateHashTable(defaultValue interface{}, bucketSize int) *HashTable {
    if bucketSize <= 0 {
        bucketSize = 32 // Default value
    }

    return &HashTable{bucket: make([]*hashTableElement, bucketSize, bucketSize), defaultValue: defaultValue}
}

func (h *HashTable) PutElement(key Hashable, value interface{}) {
    // Calculate relative hash in hashtable bucket
    absoluteHash := key.HashCode()
    relativeHash := int64(math.Mod(float64(absoluteHash), float64(len(h.bucket))))

    // Getting current element for collision check
    currentElement := h.bucket[relativeHash]

    if currentElement == nil {
        h.bucket[relativeHash] = &hashTableElement{SourceElement: key, Value: value}
        return
    }

    // Check type of collision: collision in relative hash or absolute?
    if absoluteHash == currentElement.SourceElement.HashCode() {
        currentElement.PushBack(&hashTableElement{SourceElement: key, Value: value})
    } else {
        h.rebuildBucket()
        h.PutElement(key, value)
    }
}

func (h *HashTable) GetElement(key Hashable) (interface{}, bool) {
    // Calculate relative Hash where to get element
    absoluteHash := key.HashCode()
    relativeHash := int64(math.Mod(float64(absoluteHash), float64(len(h.bucket))))

    currentElement := h.bucket[relativeHash]

    if currentElement == nil {
        return h.defaultValue, false
    }

    // Check, if there are a chained elements
    foundedValue := currentElement.FindValueByKey(key)

    if foundedValue == nil {
        return h.defaultValue, false
    } else {
        return foundedValue, true
    }
}

// Function thats perform rebuilding bucket when relative collision occurs
func (h *HashTable) rebuildBucket() {
    bucketSize := len(h.bucket)
    var newBucket []*hashTableElement

    for {
        bucketFilled := true
        bucketSize *= 2
        newBucket = make([]*hashTableElement, bucketSize, bucketSize)

        // Copy element from existing bucket to extended new, if relative collision occurs again,
        // we must trash newBucket and create new with x2 size
        for _, elem := range h.bucket {
            if elem == nil {
                continue
            }

            absoluteHash := elem.SourceElement.HashCode()
            relativeHash := int64(math.Mod(float64(absoluteHash), float64(len(newBucket))))

            currentElement := newBucket[relativeHash]

            if currentElement != nil {
                bucketFilled = false
                break // Collision detected, rebuild bucket again
            }

            newBucket[relativeHash] = elem
        }

        if bucketFilled {
            break
        }
    }

    h.bucket = newBucket
}

// Utility struct for workarounding hash collision problem
type hashTableElement struct {
    SourceElement Hashable
    Value interface{}
    NextElement *hashTableElement
}

func (e *hashTableElement) PushBack(elem *hashTableElement) {
    if e.NextElement != nil {
        e.NextElement.PushBack(elem)
        return
    }

    e.NextElement = elem
}

func (e *hashTableElement) FindValueByKey(key Hashable) interface{} {
    if e.SourceElement.Compare(key) {
        return e.Value
    }

    if e.NextElement != nil {
        return e.NextElement.FindValueByKey(key)
    }

    return nil
}
