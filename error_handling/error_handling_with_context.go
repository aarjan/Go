/*
	Error Handling by adding context to the error values
	Using github.com/pkg/errors package
	For more understanding https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/errors"
)

type Config struct {
	Dir  string `json:"dir"`
	User string `json:"user"`
}

func main() {
	byt, err := readConfig("config.json")
	if err != nil {
		log.Fatal(errors.Wrap(err, "could not read config"))
	}
	config := Config{}
	err = json.Unmarshal(byt, &config)
	if err != nil {
		log.Fatal(errors.Wrap(err, "could not decode message"))
	}
	fmt.Printf("%#v\n", config)
}

func readConfig(fname string) ([]byte, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, errors.Wrap(err, "open failed")
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "read failed")
	}
	return buf, nil
}
