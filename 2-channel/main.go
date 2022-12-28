package main

import "github.com/20k/2-channel/demo"

func main() {
	demo.SendReceive()
	demo.Blocking()
	demo.RangeDemo()
	demo.SelectDemo()
	demo.Timeout()
	demo.Timer()
	demo.Timer2()
	demo.Ticker()
	demo.Close()
	demo.Close2()
	demo.Close3()
	demo.Close4()
	demo.Producer()
}
