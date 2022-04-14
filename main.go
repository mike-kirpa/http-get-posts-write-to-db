package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {

	//getAllPosts()
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 1; i <= 100; i++ {
		go func(i int) {
			getPostById(i)
			defer wg.Done()
		}(i)
	}
	wg.Wait()

}

func getPostById(id int) {
	resp, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	// write whole the body
	err = ioutil.WriteFile(fmt.Sprintf("posts/%d.txt", id), body, 0644)
	if err != nil {
		panic(err)
	}
}
