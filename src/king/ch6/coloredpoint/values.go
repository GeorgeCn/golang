package main

import (
	"fmt"
	"time"
)

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	distanceFromP := p.Distance    // 方法变量
	fmt.Println(distanceFromP(q))  // "5"
	var origin Point               // {0, 0}
	fmt.Println(distanceFromP(origin))

	scaleP := p.ScaleBy           // 方法变量
	scaleP(2)                     // p 变成 (2, 4)
	scaleP(3)

	type Rocket struct {/*-------*/}
	func (r *Rocket) Launch() {/*--------*/}

	r := new(Rocket)
	time.AfterFunc(10 * time.Second, func() {
		r.Launch()
	})
	//如果使用方法变量则可以更简洁
	time.AfterFunc()
}
