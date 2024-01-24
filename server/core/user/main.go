package main

const UserPort = ":50051"

func main() {
	userServiceHandler := NewUserServiceHandler()
	StartUserServer(userServiceHandler, UserPort)
}
