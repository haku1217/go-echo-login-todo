package main

func main() {
	router := newRouter()
	router.Start(":4000")
}
