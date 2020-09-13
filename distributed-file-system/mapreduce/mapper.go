package mapreduce

import (
	"github.com/avis408/distributed-systems/distributed-file-system/dal"
	"github.com/avis408/distributed-systems/distributed-file-system/dto"
)

func ExecuteMapFunction(mapper Mapper, fileName string) dto.KeyValueListMap {

	mp := make(dto.KeyValueListMap)
	ch := make(chan dto.KeyValueListMap, len(shardNames))
	for _, shard := range shardNames {
		go func(mapper Mapper, shard, fileName string) {
			ch <- readDataFromShard(mapper, shard, fileName)
		}(mapper, shard, fileName)
	}
	for range shardNames {
		t := <- ch
		for k, val := range t {
			_, ok := mp[k]; if !ok {
				var v []dto.Value
				mp[k] = v
			}
			mp[k] = append(mp[k], val...)
		}
	}
	return mp
}

func ExecuteReduceFunction(reducer Reducer, data dto.KeyValueListMap, outFileName string) {
	//count := 0
	var result []string
	for k, v := range data {
		result = append(result, reducer.Reduce(k, v)...)
	}
	dal.WriteFileLines(outFileName, result)
}

func RunMapReduce(fileName, outFileName string, mapper Mapper, reducer Reducer) {
	ExecuteReduceFunction(reducer, ExecuteMapFunction(mapper, fileName), outFileName)
}

func readDataFromShard(mapper Mapper, shard, fileName string) dto.KeyValueListMap {
	mp := make(dto.KeyValueListMap)
	content, _ := dal.GetFileLines(shard, fileName)
	intermediateData := mapper.Map(fileName, content)
	for _, d := range intermediateData {
		_, ok := mp[d.Key]; if !ok {
			var v []dto.Value
			mp[d.Key] = v
		}
		mp[d.Key] = append(mp[d.Key], d.Value)
	}
	return mp
}
