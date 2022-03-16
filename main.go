package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	filename := "input.csv"
	input, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error open file", err)
	} else {
		read := csv.NewReader(input)
		count := 0
		count2 := 1
		count3 := 0
		for {
			record, err := read.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			map_list := make(map[string]string)
			for _, value := range record {
				count += 1
				if count > 3 {
					if count2 == 1 {
						map_list["organizacion"] = value
						count2 += 1
					} else if count2 == 2 {
						map_list["usuario"] = value
						count2 += 1
					} else if count2 == 3 {
						map_list["rol"] = value
						count2 = 1
						count3 += 1
						id_n := strconv.Itoa(count3)
						map_list["id"] = id_n
					}
				}

			}
			for k, i := range map_list {
				fmt.Println("______________")
				fmt.Println(k)
				fmt.Println(i)

			}
			fmt.Println(map_list)
		}
	}
}
