package mapreduce

import "github.com/avis408/distributed-systems/distributed-file-system/dto"

type Mapper interface {
	Map(fileName string, contentByLine []string) []dto.KeyValuePair
}

type Reducer interface {
	Reduce(key dto.Key, values []dto.Value) []string
}
