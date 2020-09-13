package dto

type Key interface {

}

type Value interface {

}

type KeyValuePair struct {
	Key Key
	Value Value
}

type KeyValueListMap map[Key][]Value