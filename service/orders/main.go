package main

func main(){

	grpcServer := NewGRPCServer(":8081")
	grpcServer.Run()
}