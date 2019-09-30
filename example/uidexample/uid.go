package main

import (
	"fmt"
	"github.com/lanyutc/blue/uid"
)

func main() {
	//一定要在使用前调用
	uid.StartUidServe()

	go func() {
		for i := 0; i < 10000; i++ {
			uid, err := uid.GetUid()
			if err != nil {
				panic(err)
			}
			fmt.Println("goroutine1 uid:", uid)
		}
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			uid, err := uid.GetUid()
			if err != nil {
				panic(err)
			}
			fmt.Println("goroutine2 uid:", uid)
		}
	}()
	select {}
}
