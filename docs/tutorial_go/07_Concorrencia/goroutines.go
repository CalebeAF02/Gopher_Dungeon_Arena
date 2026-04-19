package goguide

// RunInGoroutine mostra como iniciar uma goroutine para executar trabalho em paralelo.
func RunInGoroutine(done chan bool) {
	go func() {
		// Trabalho assíncrono executado dentro da goroutine.
		done <- true
	}()
}

// LaunchWorker inicia uma goroutine que envia um valor para o canal quando termina.
func LaunchWorker(done chan string, name string) {
	go func() {
		done <- "worker " + name + " concluído"
	}()
}
