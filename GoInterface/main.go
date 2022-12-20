package main

import (
	"fmt"
	"math"
)

// interface는 메서드들의 집합체
// 하나의 사용자 정의 타입이 interface를 구현하기 위해서는 단순히 그 인터페이스가 갖는 모든 메서드들을 구현하면 된다.
type Shape interface {
	area() float64
	perimeter() float64
}

// Rect 정의
type Rect struct {
	width, height float64
}

// Circle 정의
type Circle struct {
	radius float64
}

// Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) area() float64 { return r.width * r.height }
func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Circle 타입에 대한 Shape 인터페이스 구현
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r := Rect{10., 20.}
	c := Circle{10}

	showArea(r, c)
}

func showArea(shapes ...Shape) {
	// 함수가 파라미터로 인터페이스를 받아들이는 경우에 유
	for _, s := range shapes {
		a := s.area() //인터페이스 메서드 호출
		fmt.Printf("[AREA] %f\n", a)
		b := s.perimeter()
		fmt.Printf("[PERIMETER] %f\n", b)
	}
}

// http://golang.site/go/article/18-Go-%EC%9D%B8%ED%84%B0%ED%8E%98%EC%9D%B4%EC%8A%A4
// Empty interface (=인터페이스 타입 이라고도 불림)
// 메서드를 전혀 갖지 않는 빈 인터페이스 (모든 Type을 나타내기 위해 빈 인터페이스를 사용)
// C/C++에서는 void*와 같다고 볼 수 있음.

// Type Assertion
// Interface type의 x와 타입 T에 대하여 x.(T)로 표현했을 때,
// 이는 x가 nil이 아니며, x는 T 타입에 속한다는 점을 확인(assert)하는 것으로 이러한 표현을 "Type Assertion"이라 부른다.
//
//func main() {
//	var a interface{} = 1
//
//	i := a       // a와 i 는 dynamic type, 값은 1
//	j := a.(int) // j는 int 타입, 값은 1
//
//	println(i)  // 포인터주소 출력
//	println(j)  // 1 출력
//}
