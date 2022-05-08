package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	test()
	basic()
}

func basic() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match) // true

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach")) // true

	// 정규식과 일치하는 문자열 찾기 - 첫번째로 매칭되는 문자열 반환
	fmt.Println(r.FindString("peach punch")) // peach

	// 일치하는 텍스트의 첫 인덱스와 마지막 인덱스를 반환
	fmt.Println(r.FindStringIndex("peach punch")) // [0,5]

	// 전체 패턴 일치와 해당 일치의 부분 일치에 대한 정보를 모두 포함
	fmt.Println(r.FindStringSubmatch("peach punch")) // [peach ea]

	// 전체 일치와 부분 일치의 인덱스에 대한 정보를 반환
	fmt.Println(r.FindStringSubmatchIndex("peach punch")) // [0 5 1 3]

	// 정규식에 대해 모든 일치 항목들을 찾기
	fmt.Println(r.FindAllString("peach punch pinch", -1)) // [peach punch pinch]

	// 함수의 두번째 인자에 음이 아닌 정수를 전달하면 일치 항목의 갯수를 제한
	fmt.Println(r.FindAllString("peach punch pinch", 2)) // [peach punch]

	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1)) // [[0 5 1 3] [6 11 7 9] [12 17 13 15]]

	// 함수명에서 String을 없애고 인자로 []byte 전달하기
	fmt.Println(r.Match([]byte("peach"))) // true

	// 정규표현식으로 상수를 만들때 MustCompile 사용
	// 일반 Compile은 반환값이 2개라 상수로 사용할 수 없다.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r) // p([a-z]+)ch

	// 부분문자열을 다른값으로 바꾸기
	fmt.Println(r.ReplaceAllString("a peach pinch peace", "<fruit>")) // a <fruit> <fruit> peace

	// Func 변형을 사용하여 주어진 함수로 일치된 텍스트를 변환
	in := []byte("a peach pinch peace")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out)) // a PEACH PINCH peace
}

func test() {
	r, _ := regexp.Compile("[0-9]")
	pageList := r.FindAllString("1페이지 결과 1,534건", -1)[1:]
	totalResults, _ := strconv.Atoi(strings.Join(pageList, ""))
	itemsPerPage := 50
	pageNum := totalResults / itemsPerPage
	if totalResults%itemsPerPage != 0 {
		pageNum += 1
	}
	fmt.Println(pageNum)
}
