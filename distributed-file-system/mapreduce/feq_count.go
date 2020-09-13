package mapreduce

import (
	"fmt"
	"github.com/avis408/distributed-systems/distributed-file-system/dto"
	"strconv"
	"strings"
)

type FreqCountMapper struct {

}

func (f FreqCountMapper) Map(fileName string, contentByLine []string) []dto.KeyValuePair {
	var arr []dto.KeyValuePair
	for _, content := range contentByLine {
		values := strings.Split(content, " ")
		arr = append(arr, dto.KeyValuePair{
			Key: values[0],
			Value: values[1],
		})
	}
	return arr
}

type FreqCountReducer struct {

}

func (f FreqCountReducer) Reduce(key dto.Key, values []dto.Value) []string {
	sum := 0
	for _, v := range values {
		val, _ := strconv.Atoi(v.(string))
		sum += val
	}
	return []string{fmt.Sprintf("%v %v", key.(string), sum)}
}
