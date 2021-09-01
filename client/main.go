package main

import (
	ui "afg/imagepool/proto"
	"afg/imagepool/stuffs"

	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"
)

const ABOUT = `This is the Client for recieve and request files from Server
Usage:
	client <command> [arguments]
Commands:
	push   [filepath] [filename]  push file to server from [filepath]           (Example: client C:/images/ picture.png)
	get    [filepath] [filename]  get file from server and save on [filepath]   (Example: client C:/images/ picture.gif)
	list                          show list of file in server
`
const (
	ADR_1 = "localhost:8080"
	ADR_2 = "localhost:8085"

	PUSHF = "push"
	GETF  = "get"
	LISTF = "list"

	UNEXP  = "Unexpected command, run with '?' or 'help' for usage"
	MOREAR = "Need more arguments: <filepath> <name.type>"
)

func pushFile(c ui.ImagePoolServiceClient, filePath, fileName string) {
	file, err := os.Open(fmt.Sprintf("%s%s", filePath, fileName))
	stuffs.ErrorExit(err)
	defer file.Close()

	stream, err := c.Exchanger(context.Background())
	stuffs.ErrorExit(err)

	msg_client := &ui.Request{Message: fileName, Cmd: ui.Cmd_C_PUSH, State: ui.State_S_READY}
	stuffs.ErrorExit(stream.Send(msg_client)) // отправка

	msg_server, err := stream.Recv() // получение
	stuffs.ErrorExit(err)

	if msg_server.GetState() == ui.State_S_READY {
		reader := bufio.NewReader(file)
		buffer := make([]byte, 1024)
		for {
			n, err := reader.Read(buffer)
			if err == io.EOF {
				msg_client = &ui.Request{Cmd: ui.Cmd_C_WAIT, State: ui.State_S_DONE}
				stuffs.ErrorExit(stream.Send(msg_client))
				log.Printf("Client: file uploaded with name: %s", fileName)
				break
			}
			stuffs.ErrorExit(err)
			msg_client = &ui.Request{Bytes: buffer[:n]}
			stuffs.ErrorExit(stream.Send(msg_client))
		}
	} else {
		msg_client = &ui.Request{Message: "Server not ready", Cmd: ui.Cmd_C_WAIT, State: ui.State_S_ERROR}
		stuffs.ErrorExit(stream.Send(msg_client))
	}
}

func getFile(c ui.ImagePoolServiceClient, filePath, fileName string) {

	stream, err := c.Exchanger(context.Background())
	stuffs.ErrorExit(err)

	msg_client := &ui.Request{Message: fileName, Cmd: ui.Cmd_C_GET, State: ui.State_S_READY}
	stuffs.ErrorExit(stream.Send(msg_client))

	msg_server, err := stream.Recv() // получение
	stuffs.ErrorExit(err)

	if msg_server.GetState() == ui.State_S_READY {
		fileName := msg_server.GetMessage()
		file, err := os.Create(fmt.Sprint(filePath, fileName))
		stuffs.ErrorExit(err)
		for {
			msg_server, err := stream.Recv()
			stuffs.ErrorExit(err)
			if msg_server.GetState() == ui.State_S_ERROR {
				log.Printf("Request download [%s]: error from server [%s]", fileName, msg_server.GetMessage())
				file.Close()
				break
			}
			if msg_server.GetState() == ui.State_S_DONE || err == io.EOF {
				log.Printf("Request download [%s]: complete", fileName)
				file.Close()
				break
			} else {
				file.Write(msg_server.GetBytes())
			}
		}
	} else {
		msg_client = &ui.Request{Message: "Server not ready", Cmd: ui.Cmd_C_WAIT, State: ui.State_S_ERROR}
		stuffs.ErrorExit(stream.Send(msg_client))
		log.Fatalf("Request download [%s]: failed with error - %s", fileName, msg_server.GetMessage())
	}
}

func listFile(c ui.ImageListServiceClient) {
	stream, err := c.List(context.Background())
	stuffs.ErrorExit(err)
	msg_client := &ui.Request{Cmd: ui.Cmd_C_LIST, State: ui.State_S_READY}
	stuffs.ErrorExit(stream.Send(msg_client))
	msg_server, err := stream.Recv()
	stuffs.ErrorExit(err)

	if msg_server.GetState() == ui.State_S_READY {
		fileCount := 0
		log.Print(msg_server.GetMessage())
		for {
			msg_server, err := stream.Recv()
			stuffs.ErrorExit(err)
			if msg_server.GetState() == ui.State_S_ERROR {
				log.Printf("Request get list of all files: error from server [%s]", msg_server.GetMessage())
				break
			}
			if msg_server.GetState() == ui.State_S_DONE || err == io.EOF {
				log.Printf("Request get list of all files: complete, total files %d", fileCount)
				break
			} else {
				log.Println(msg_server.GetMessage())
				fileCount++
			}
		}
	} else {
		msg_client = &ui.Request{Message: "Server not ready", Cmd: ui.Cmd_C_WAIT, State: ui.State_S_ERROR}
		stuffs.ErrorExit(stream.Send(msg_client))
		log.Fatalf("Request get list of all files: failed with error - %s", msg_server.GetMessage())
	}
}

func main() {
	conn, err := grpc.Dial(ADR_1, grpc.WithInsecure())
	stuffs.ErrorExit(err)
	conn2, err := grpc.Dial(ADR_2, grpc.WithInsecure())
	stuffs.ErrorExit(err)

	clientExchanger := ui.NewImagePoolServiceClient(conn)
	clientList := ui.NewImageListServiceClient(conn2)

	if len(os.Args) < 2 {
		fmt.Println(UNEXP)
		os.Exit(1)
	}
	// можно получше
	switch os.Args[1] {
	case PUSHF:
		if len(os.Args) < 4 || os.Args[2] == "" || os.Args[3] == "" {
			log.Fatalln(MOREAR)
		} else {
			pushFile(clientExchanger, os.Args[2], os.Args[3])
		}
	case GETF:
		if len(os.Args) < 4 || os.Args[2] == "" || os.Args[3] == "" {
			log.Fatalln(MOREAR)
		} else {
			getFile(clientExchanger, os.Args[2], os.Args[3])
		}
	case LISTF:
		listFile(clientList)
	case "?", "help":
		fmt.Println(ABOUT)
	default:
		fmt.Println(UNEXP)
	}
}