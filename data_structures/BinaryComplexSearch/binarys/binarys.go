package binarys

import (
    "sort"
    "errors"
    "math"
)

// Errors definition
var ListArgumentNullError = errors.New("binarys: nil was passed as method argument")
var SearchEmptyListError = errors.New("binarys: can't search elements: indexes are not initialized in object");

type BinaryComplexSearch struct {
    innerArray []complex128
    arrayIndex []index
}

func (b *BinaryComplexSearch) SetList(plist []complex128) error {
    if (plist == nil) { return ListArgumentNullError; }

    b.innerArray = plist
    b.updateIndex()

    return nil
}

func (b *BinaryComplexSearch) GetList() []complex128 {
    return b.innerArray
}

func (b *BinaryComplexSearch) SearchElement(element complex128) (int, error) {
    if (b.arrayIndex == nil || len(b.arrayIndex) == 0) { return -1, SearchEmptyListError; }

    left, right := 0, len(b.arrayIndex) - 1
    mid := (left + right) / 2
    offset := mid

    for {
        if b.arrayIndex[mid].Value == element { return b.arrayIndex[mid].Index, nil }
        if math.Abs(float64(left - right)) == 0 { return -1, nil }

        isElementLower := compareComplex(element, b.arrayIndex[mid].Value, func (n1 float64, n2 float64) bool {
            return n1 < n2;
        })

        if isElementLower {
            right -= offset
        } else {
            left += offset
        }

        mid = (left + right) / 2
        offset = map[bool]int{true: 1, false: offset / 2}[offset <= 1]
    }
}

func (b *BinaryComplexSearch) updateIndex() {
    // Generate []index for []complex128
    indexes := []index{}

    for i, v := range b.innerArray {
        indexes = append(indexes, index{Index: i, Value: v})
    }

    // Sort formed indexes
    sorter := arrayIndexSorter{Indexes: indexes}
    sort.Sort(&sorter)

    b.arrayIndex = sorter.Indexes
}

// Inner type for array sorting
type arrayIndexSorter struct {
    Indexes []index
}

func (s *arrayIndexSorter) Len() int { return len(s.Indexes); }
func (s *arrayIndexSorter) Swap(i, j int) { s.Indexes[i], s.Indexes[j] = s.Indexes[j], s.Indexes[i]; }

func (s *arrayIndexSorter) Less(i, j int) bool {
    return compareComplex(s.Indexes[i].Value, s.Indexes[j].Value, func (n1 float64, n2 float64) bool {
        return n1 < n2;
    })
}

// Inner type for indexes
type index struct {
    Value complex128
    Index int
}

// Inner function for comparing complex numbers sum with predicate
func compareComplex(complex1 complex128, complex2 complex128, predicate func (float64, float64) bool) bool {
    reali, imagi := real(complex1), imag(complex1)
    realj, imagj := real(complex2), imag(complex2)

    return predicate(reali + imagi, realj + imagj)
}
