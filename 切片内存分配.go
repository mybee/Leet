package main

import "fmt"

const SNOW_FLAKE_DICTIONARY = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func main() {
	base62_encode(3325698880595111250)
}

func base62_encode(id int64) string {

	dictionaryxxx := []byte(SNOW_FLAKE_DICTIONARY)
	idBytes := []byte{}
	for {
		if id <= 0 {
			break
		}
		i := id % 62
		//fmt.Println("this is i:", i, string(base[i:i+1]))
		//fmt.Println(string(dictionaryxxx))

		//v := dictionaryxxx[i : i+1]
		fmt.Println(i, ">>>", string(dictionaryxxx))
		fmt.Println(i, "--->", string(idBytes))
		a := []byte("sfe")
		idBytes = append(dictionaryxxx[i : i+1], a...)
		fmt.Println("====>", string(idBytes))
		id = (id - i) / 62
	}
	fmt.Println(string(idBytes))
	return string(idBytes)
}