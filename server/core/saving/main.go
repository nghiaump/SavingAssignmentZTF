package main

const Port = ":50052"

func main() {
	savingServiceHandler := NewSavingServiceHandler()
	StartSavingServer(savingServiceHandler, Port)
}
