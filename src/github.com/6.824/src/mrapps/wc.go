package main

//
// a word-count application "plugin" for MapReduce.
//co
// go build -buildmode=plugin wc.go
//

import "6.824/mr"
import "unicode"
import "strings"
import "strconv"

// The map function is called once for each file of input. The first
// argument is the name of the input file, and the second is the
// file's complete contents. You should ignore the input file name,
// and look only at the contents argument. The return value is a slice
// of key/value pairs.
func Map(filename string, contents string) []mr.KeyValue {
	// function to detect word separators.
	ff := func(r rune) bool { return !unicode.IsLetter(r) } // 定义非字母则返回true的函数

	// split contents into an array of words.
	words := strings.FieldsFunc(contents, ff) // 用于在contents中从非字母字符处切片

	kva := []mr.KeyValue{} // 定义空切片
	for _, w := range words {
		kv := mr.KeyValue{w, "1"} // 将单词编写为键值对
		kva = append(kva, kv)     // 插入kva数组中
	}
	return kva
}

// The reduce function is called once for each key generated by the
// map tasks, with a list of all the values created for that key by
// any map task.
func Reduce(key string, values []string) string {
	// return the number of occurrences of this word.
	return strconv.Itoa(len(values)) // 将存values的数组长度转换为字符串
}