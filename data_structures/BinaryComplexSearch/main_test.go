package main

import (
    "testing"
    "./binarys"
)

func initBinarySearcher() *binarys.BinaryComplexSearch {
    complexList := []complex128{
        4 + 1i,
        6 + 3i,
        11 + 2i,
        -2 + 3i,
        -5 - 2i,
        7 + 12i,
        5 + 25i,
        1i,
        2,
        8 + 8i,
    }

    bSearcher := &binarys.BinaryComplexSearch{}
    _ = bSearcher.SetList(complexList)

    return bSearcher
}


func TestBinarySearch1(t *testing.T) {
    bSearcher := initBinarySearcher()
    var got int

    got, _ = bSearcher.SearchElement(11 + 2i)
    if got != 2 {
        t.Errorf("TestBinarySearch: Except 2, got %d\n", got)
    }

    got, _ = bSearcher.SearchElement(7 + 12i)
    if got != 5 {
        t.Errorf("TestBinarySearch: Except 5, got %d\n", got)
    }

    got, _ = bSearcher.SearchElement(11i)
    if got != -1 {
        t.Errorf("TestBinarySearch: Except -1, got %d\n", got)
    }
}
