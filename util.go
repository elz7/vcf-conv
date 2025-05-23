package main

import (
	"hash/fnv"
	"strings"
	"fmt"
)

func get_rand_name(phone string) (string, string) {
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
	}

	h := fnv.New64()
	h.Write([]byte(phone))
	s := h.Sum64()

	i1 := int(s/3)%len(fn)
	i2 := int(2*s/3)%len(ln)
	fmt.Println(i1,i2)
	return fn[i1], ln[i2]
}

func make_map(v, t string) map[string]string {
        ret := make(map[string]string, 0)
        arr := strings.Split(v, ";")

        for _, el := range arr {
		ret[el] = t
        }

        return ret
}
