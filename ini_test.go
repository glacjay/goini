package ini

import (
	"testing"
)

var (
	dict Dict
	err  error
)

func init() {
	dict, err = Load("example.ini")
}

func TestLoad(t *testing.T) {
	if err != nil {
		t.Error("Example: load error:", err)
	}
}

func TestGetBool(t *testing.T) {
	b, found := dict.GetBool("pizza", "ham")
	if !found || !b {
		t.Error("Example: parse error for key ham of section pizza.")
	}
	b, found = dict.GetBool("pizza", "mushrooms")
	if !found || !b {
		t.Error("Example: parse error for key mushrooms of section pizza.")
	}
	b, found = dict.GetBool("pizza", "capres")
	if !found || b {
		t.Error("Example: parse error for key capres of section pizza.")
	}
	b, found = dict.GetBool("pizza", "cheese")
	if !found || b {
		t.Error("Example: parse error for key cheese of section pizza.")
	}
}

func TestGetStringIntAndDouble(t *testing.T) {
	str, found := dict.GetString("wine", "grape")
	if !found || str != "Cabernet Sauvignon" {
		t.Error("Example: parse error for key grape of section wine.")
	}
	i, found := dict.GetInt("wine", "year")
	if !found || i != 1989 {
		t.Error("Example: parse error for key year of section wine.")
	}
	str, found = dict.GetString("wine", "country")
	if !found || str != "Spain" {
		t.Error("Example: parse error for key grape of section wine.")
	}
	d, found := dict.GetDouble("wine", "alcohol")
	if !found || d != 12.5 {
		t.Error("Example: parse error for key grape of section wine.")
	}
}

func TestGetNotExist(t *testing.T) {
	_, found := dict.GetString("not", "exist")
	if found {
		t.Error("There is no key exist of section not.")
	}
}

func TestGetSections(t *testing.T) {
	sections := dict.GetSections()
	if len(sections) != 3 {
		t.Error("The number of sections is wrong:", len(sections))
	}
	for _, section := range sections {
		if section != "" && section != "pizza" && section != "wine" {
			t.Errorf("Section '%s' should not be exist.", section)
		}
	}
}
