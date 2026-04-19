package goguide

// SelectExample demonstra o uso de select para escolher entre múltiplos canais.
func SelectExample(ch1, ch2 chan string) string {
	select {
	case msg := <-ch1:
		return msg
	case msg := <-ch2:
		return msg
	default:
		return "nenhuma mensagem"
	}
}

// SelectWithTimeout usa select para receber ou tempo limite.
func SelectWithTimeout(ch chan string, timeout chan bool) string {
	select {
	case msg := <-ch:
		return msg
	case <-timeout:
		return "tempo esgotado"
	}
}
