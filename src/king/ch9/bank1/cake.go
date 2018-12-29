package main

type Cake struct {
	state string
}

func baker(cooked chan<- *Cake) {
	for{
		cake := new(Cake)
		cake.state = "cooked"
		cooked <-cake // baker 不在访问cake变量
	}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state = "iced"
		iced <-cake
	}
}
