package main

const MidPort = ":50050"

func main() {
	midServiceHandler, listConn := CreateMidServiceHandler()
	defer listConn[0].Close()
	defer listConn[1].Close()
	StartMidServer(midServiceHandler, MidPort)
}
