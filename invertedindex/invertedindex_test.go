package invertedindex

import (
	"fmt"
	"reflect"
	"testing"
)

var wordListTest = []string{
	"new", "HOme", "sales", "top", "forecasts",
	"home", "sales", "rise", "in", "July",
	"increase", "in", "home", "SALES", "in",
	"July", "new", "home", "sales", "rise", "July",
}

func TestPreprocessing(t *testing.T) {
	wordList := wordListTest

	expectedList := []string{
		"new", "home", "sales", "top", "forecasts",
		"home", "sales", "rise", "in", "july",
		"increase", "in", "home", "sales", "in",
		"july", "new", "home", "sales", "rise", "july",
	}

	actualList := Preprocessing(wordList)

	if !reflect.DeepEqual(expectedList, actualList) {
		t.Fatalf("\nExpected:%v \nGot:%v", expectedList, actualList)
	}
}

func TestPreprocessing_NoWordList(t *testing.T) {
	wordList := make([]string, 0)

	expectedList := make([]string, 0)

	actualList := Preprocessing(wordList)

	if !reflect.DeepEqual(expectedList, actualList) {
		t.Fatalf("\nExpected:%v \nGot:%v", expectedList, actualList)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	wordList := Preprocessing(wordListTest)

	expectedList := []string{
		"new", "home", "sales", "top", "forecasts",
		"rise", "in", "july",
		"increase",
	}

	actualList := RemoveDuplicates(wordList)

	if !reflect.DeepEqual(expectedList, actualList) {
		t.Fatalf("\nExpected:%v \nGot:%v", expectedList, actualList)
	}
}

func TestRemoveDuplicates_NoWordList(t *testing.T) {
	wordList := make([]string, 0)

	expectedList := make([]string, 0)

	actualList := RemoveDuplicates(wordList)

	if !reflect.DeepEqual(expectedList, actualList) {
		t.Fatalf("\nExpected:%v \nGot:%v", expectedList, actualList)
	}
}

func TestTokenize(t *testing.T) {
	doc := "new home sales top forecasts NEW"

	expectedList := []string{
		"new", "home", "sales", "top", "forecasts",
	}

	actualList := Tokenize(doc)

	if !reflect.DeepEqual(expectedList, actualList) {
		t.Fatalf("\nExpected:%v \nGot:%v", expectedList, actualList)
	}
}

func TestTokenize_NoDoc(t *testing.T) {
	var doc string

	expectedList := []string{}

	actualList := Tokenize(doc)

	if !reflect.DeepEqual(expectedList, actualList) {
		t.Fatalf("\nExpected:%v \nGot:%v", expectedList, actualList)
	}
}

func TestGenerateInvertedIndexWithPreExistingIds(t *testing.T) {
	input := make(map[int][]string, 0)

	input[1] = []string{"1001", "1002"}
	input[23] = []string{"1001", "1003"}

	expected := InvertedIndex{
		HashMap: nil,
		Items: []*InvertedIndexEntry{{
			Term:            "1001",
			Frequency:       2,
			DocumentListing: []int{1, 23},
		}, {
			Term:            "1002",
			Frequency:       1,
			DocumentListing: []int{1},
		}, {
			Term:            "1003",
			Frequency:       1,
			DocumentListing: []int{23},
		}},
	}

	actual := GenerateInvertedIndexWithPreExistingIds(input)

	if !reflect.DeepEqual(actual.Items, expected.Items) {
		fmt.Printf("Expected %+v, actual %+v\n", expected.Items, actual.Items)
		t.Fail()
	}
}
