package main

import (
	"fmt"
	"hash/fnv"
	"path/filepath"
	"strings"
)

func randName(phone string) (string, string) {
	fn := []string{
		"Rose",
		"Fire",
		"Lake",
		"River",
		"Moon",
		"Train",
		"Sun",
		"Needle",
		"Fries",
		"Step",
		"Never",
	}
	ln := []string{
		"Bush",
		"Smith",
		"Cock",
		"Bycicle",
		"Bitcoin",
		"Woo",
		"Knife",
		"Dry",
		"Monolith",
		"Railway",
		"Frog",
		"Feet",
	}

	h := fnv.New64()
	h.Write([]byte(phone))
	s := h.Sum64()

	i1 := int(s/3) % len(fn)
	i2 := int(2*s/3) % len(ln)
	fmt.Println(i1, i2)
	return fn[i1], ln[i2]
}

func makeMap(v, t string) map[string]string {
	ret := make(map[string]string, 0)
	arr := strings.SplitSeq(v, ";")

	for el := range arr {
		ret[el] = t
	}

	return ret
}

func getOrNone[T any](slice []T, index int) (T, bool) {
	if index < len(slice) {
		return slice[index], true
	}
	var zero T
	return zero, false
}

func fileNameWithoutExtension(name string) string {
	name = filepath.Base(name)

	if name == "." || name == "" {
		return "phones"
	}

	ext := filepath.Ext(name)
	if ext == "" {
		return name
	}

	return name[:len(name)-len(ext)]
}
