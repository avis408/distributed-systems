package main

import "github.com/avis408/distributed-systems/distributed-file-system/mapreduce"

func main() {

	fileName := "input_1.txt"
	outFileName := "output_1.txt"
	mapreduce.RunMapReduce(fileName, outFileName, mapreduce.FreqCountMapper{}, mapreduce.FreqCountReducer{})
}