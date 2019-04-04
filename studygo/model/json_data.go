package model

import (
	"fmt"
	"io/ioutil"
	"os"
)

func create_struct(create_sql string) {

}

func read_data(file_name string) string {
	fl, err := os.Open(file_name)
	if err != nil {
		fmt.Println(file_name, err)
		return ""
	}
	data, err := ioutil.ReadAll(fl)
	if err != nil {
		fmt.Println(file_name, err)
		return ""
	}
	return string(data)

}
