package main

import (
	"sync"
)

func ping(args []Value) Value {
	if len(args) == 0 {
		return Value{typ: "string", str: "PONG"}
	}
	return Value{typ: "string", str: args[0].bulk}
}

var SETs = map[string]string{}
var SETsMutex = sync.RWMutex{}

func set(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "Wrong number of args supplied to SET command"}
	}

	key := args[0].bulk
	val := args[1].bulk

	SETsMutex.Lock()
	SETs[key] = val
	SETsMutex.Unlock()

	return Value{typ: "string", str: "OK"}
}

func get(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "Wrong number of args supplied to GET command"}
	}

	SETsMutex.RLock()
	val, ok := SETs[args[0].bulk]
	SETsMutex.RUnlock()

	if !ok {
		return Value{typ: "null"}
	}

	return Value{typ: "bulk", bulk: val}
}

var HSETs = map[string]map[string]string{}
var HSETsMutex = sync.RWMutex{}

func hset(args []Value) Value {
	if len(args) != 3 {
		return Value{typ: "error", str: "Wrong number of args supplied to HSET command"}
	}

	hash := args[0].bulk
	key := args[1].bulk
	val := args[2].bulk

	_, ok := HSETs[hash]
	if !ok {
		HSETs[hash] = map[string]string{}
	}

	HSETsMutex.Lock()
	HSETs[hash][key] = val
	HSETsMutex.Unlock()

	return Value{typ: "string", str: "OK"}
}

func hget(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "Wrong number of args supplied to HGET command"}
	}

	hash := args[0].bulk
	key := args[1].bulk

	HSETsMutex.RLock()
	val, ok := HSETs[hash][key]
	HSETsMutex.RUnlock()

	if !ok {
		return Value{typ: "null"}
	}

	return Value{typ: "bulk", bulk: val}

}

func mapToValueArray(m map[string]string) []Value {
	var arr []Value
	for k, v := range m {
		arr = append(arr, Value{typ: "bulk", bulk: k})
		arr = append(arr, Value{typ: "bulk", bulk: v})
	}
	return arr
}

func hgetall(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "Wrong number of args supplied to HGETALL command"}
	}

	hash := args[0].bulk

	HSETsMutex.RLock()
	keys, ok := HSETs[hash]
	if !ok {
		return Value{typ: "null"}
	}
	HSETsMutex.RUnlock()

	return Value{typ: "array", array: mapToValueArray(keys)}

}

var Handlers = map[string]func([]Value) Value{
	"PING":    ping,
	"SET":     set,
	"GET":     get,
	"HSET":    hset,
	"HGET":    hget,
	"HGETALL": hgetall,
}
