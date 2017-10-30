package fib
//import "fmt"


func run(c,quit chan int)  {
	x,y := 0,1
	for {
	select {
		case c <- x:
			x,y = y,x+y
		case <-quit:
			return
	}
	}
}

func Fib2(n int) int{
	var res int
	c,quit := make(chan int),make(chan int)
	go func() {
		for i:=0;i<=n;i++{
			if i==n{
				res = <-c
			}
			<-c
		}
		quit <- 0
	}()
	run(c,quit)
	return res
}

