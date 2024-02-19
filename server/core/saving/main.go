package main

const SavingPort = ":50052"

func main() {
	savingServiceHandler := NewSavingServiceHandler()
	StartSavingServer(savingServiceHandler, SavingPort)
}
