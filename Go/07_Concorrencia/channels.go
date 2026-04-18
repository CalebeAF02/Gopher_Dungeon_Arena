package goguide

// SendValue envia um valor inteiro para o canal.
func SendValue(ch chan int, value int) {
	ch <- value
}

// ReceiveValue recebe um valor inteiro do canal.
func ReceiveValue(ch chan int) int {
	return <-ch
}

// BufferedChannelExample cria um canal com buffer para enviar valores sem bloqueio imediato.
func BufferedChannelExample() chan int {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	return ch
}
