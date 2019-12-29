package main

import (
	"awesomeProject3/Cache"
	"fmt"
	"net/http"
	"time"
)

var c  = Cache.NewCache()
func generate(c *Cache.Cache, t time.Duration) {
	for {
		time.Sleep(time.Second*t)
		c.GenerateItem()
	}
}
func main() {
	fmt.Println("start")
	go generate(c, 5)
	http.HandleFunc("/get", GetValue2)

	err := http.ListenAndServe(":8080", nil)
	if err!=nil {
		fmt.Println("Error start listen server")
	}
	}
func GetValue2 (w http.ResponseWriter, r *http.Request) {
	for i:=0;i<2;i++ {
		i,e :=c.GetItem()
		if e!=nil {
			w.Write([]byte(e.Error()))
			break
		} else {
			fmt.Fprintf(w, "%s"+"\n", i.Data)
		}
	}
}
