package main

import (
	ui "afg/imagepool/proto"
	"afg/imagepool/stuffs"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

const (
	ST_FOLDER = "hdd" // папка хранения файлов
	ADR_1     = "localhost:8080"
	ADR_2     = "localhost:8085"
	MAX_1     = 10
	MAX_2     = 100
)

type ImagePoolServer struct {
	ui.ImagePoolServiceServer
	folder string
}

type ImageListServer struct {
	ui.ImageListServiceServer
	folder string
}

func CreateServerExchanger() (s *ImagePoolServer) {
	s = new(ImagePoolServer)
	exp, err := os.Executable()
	stuffs.ErrorExit(err)
	s.folder = filepath.Dir(exp) + "\\" + ST_FOLDER + "\\"
	if os.MkdirAll(s.folder, os.ModePerm) != nil {
		stuffs.ErrorExit(os.Mkdir(s.folder, os.ModePerm))
		log.Printf("Server storage created in [%s] ", s.folder)
	} else {
		log.Printf("Server storage opened in [%s] ", s.folder)
	}
	return s
}

func CreateServerList(ex *ImagePoolServer) (s *ImageListServer) {
	s = new(ImageListServer)
	s.folder = ex.folder
	return s
}

// сервис из протофайла
func (s *ImagePoolServer) Exchanger(stream ui.ImagePoolService_ExchangerServer) (err error) {
	msg_client, err := stream.Recv()
	stuffs.ErrorExit(err)
	switch msg_client.GetCmd() {
	case ui.Cmd_C_WAIT:
	case ui.Cmd_C_PUSH:
		pushFile(s, stream, msg_client)
	case ui.Cmd_C_GET:
		getFile(s, stream, msg_client)
	default:
	}
	msg_server := &ui.Response{Message: "Done", State: ui.State_S_DONE}
	stuffs.ErrorExit(stream.Send(msg_server))
	return err
}

// сервис из протофайла
func (s *ImageListServer) List(stream ui.ImageListService_ListServer) (err error) {
	log.Print("New Req")
	msg_client, err := stream.Recv()
	stuffs.ErrorExit(err)
	switch msg_client.GetCmd() {
	case ui.Cmd_C_LIST:
		listFile(s, stream, msg_client)
	default:
	}
	msg_server := &ui.Response{Message: "Done", State: ui.State_S_DONE}
	stuffs.ErrorExit(stream.Send(msg_server))
	return err
}

// загрузить файл на сервер
func pushFile(s *ImagePoolServer, stream ui.ImagePoolService_ExchangerServer, msg_client *ui.Request) error {
	fileName := msg_client.GetMessage()
	log.Printf("Load file [%s]: new request", fileName)
	file, err := os.Create(s.folder + fileName)
	stuffs.ErrorExit(err)
	msg_server := &ui.Response{State: ui.State_S_READY}
	stuffs.ErrorExit(stream.Send(msg_server))
	for {
		msg_client, err := stream.Recv()
		stuffs.ErrorExit(err)
		if msg_client.GetState() == ui.State_S_ERROR {
			log.Printf("Load file [%s]: error from client [%s]", fileName, msg_client.GetMessage())
			file.Close()
			break
		}
		if msg_client.GetState() == ui.State_S_DONE || err == io.EOF {
			log.Printf("Load file [%s]: complete", fileName)
			file.Close()
			break
		} else {
			file.Write(msg_client.GetBytes())
		}
	}
	return err
}

// выгрузить файл из сервера клиенту
func getFile(s *ImagePoolServer, stream ui.ImagePoolService_ExchangerServer, msg_client *ui.Request) error {
	fileName := msg_client.GetMessage()
	log.Printf("Download file [%s]: new request", fileName)
	file, err := os.Open(s.folder + fileName)
	if err != nil {
		msg_server := &ui.Response{Message: fmt.Sprint(err), State: ui.State_S_ERROR}
		stuffs.ErrorExit(stream.Send(msg_server))
		log.Printf("Download file [%s]: file dont found", fileName)
		return err
	}
	log.Printf("Download file[%s]: file founded", fileName)
	msg_server := &ui.Response{Message: fileName, State: ui.State_S_READY}
	stuffs.ErrorExit(stream.Send(msg_server))
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			log.Printf("Download file[%s]: file sended", fileName)
			file.Close()
			break
		}
		msg_server := &ui.Response{Bytes: buffer[:n]}
		stuffs.ErrorExit(stream.Send(msg_server))
	}
	return err
}

// список файлов в хранилище
func listFile(s *ImageListServer, stream ui.ImageListService_ListServer, msg_client *ui.Request) error {
	log.Println("Send file list: new request")
	files, err := ioutil.ReadDir(s.folder)

	stuffs.ErrorExit(err)
	msg_server := &ui.Response{Message: "No | filename | created | modified", State: ui.State_S_READY}
	stuffs.ErrorExit(stream.Send(msg_server))
	fileCount := 0
	for _, f := range files {
		fileCount++
		// возможны проблемы с кроссплатформенностью методов syscall
		// но видимо, перехватить панику таким способом не получится
		// может через горутину, хз
		tNs := f.Sys().(*syscall.Win32FileAttributeData).CreationTime.Nanoseconds()
		if recover() != nil {
			tNs = int64(0)
		}
		fileCreated := fmt.Sprint(time.Unix(0, tNs).Date())
		fileModyfied := fmt.Sprint(f.ModTime().Date())
		msg_server := &ui.Response{Message: fmt.Sprintf("%d | %s | %s | %s ", fileCount, f.Name(), fileCreated, fileModyfied)}
		stuffs.ErrorExit(stream.Send(msg_server))
	}
	log.Printf("Send file list: done")
	return err
}

func main() {
	listener_1, err := net.Listen("tcp", ADR_1)
	stuffs.ErrorExit(err)
	listener_2, err := net.Listen("tcp", ADR_2)
	stuffs.ErrorExit(err)

	server_1 := grpc.NewServer(grpc.MaxConcurrentStreams(MAX_1))
	server_2 := grpc.NewServer(grpc.MaxConcurrentStreams(MAX_2))

	instance_1 := CreateServerExchanger()
	instance_2 := CreateServerList(instance_1)

	ui.RegisterImagePoolServiceServer(server_1, instance_1)
	ui.RegisterImageListServiceServer(server_2, instance_2)

	go func() { stuffs.ErrorExit(server_1.Serve(listener_1)) }()
	stuffs.ErrorExit(server_2.Serve(listener_2))

	// два отдельных порта нужно для ограничения подключения на одном до 10, на другом до 100
	// как это сделать с привязкой к методу не знаю пока (из протофайла который, если бы их там было несколько)
}
