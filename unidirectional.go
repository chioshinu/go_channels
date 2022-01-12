package main

func sendData(sendch chan<- int) {
	sendch <- 10
}

// func unidiractional() {
// 	sendch := make(chan<- int)
// 	go sendData(sendch)
// 	fmt.Println(<-sendch)
// }
