package dal

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func GetFileLines(shardName, fileName string) ([]string, error) {

	file, err := ioutil.ReadFile(fmt.Sprintf("/Users/avinash/avinash/go/src/github.com/avis408/distributed-systems/distributed-file-system/resources/%v/%v",
		shardName, fileName))
	if err != nil {
		return nil, err
	}
	return strings.Split(string(file), "\n"), nil
}

func WriteFileLines(fileName string, data []string) error {

	err := ioutil.WriteFile(fmt.Sprintf("/Users/avinash/avinash/go/src/github.com/avis408/distributed-systems/distributed-file-system/resources/output/%v",
		fileName), []byte(strings.Join(data, "\n")), os.ModeAppend)
	if err != nil {
		return err
	}
	return nil
}
