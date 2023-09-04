package unixsock

import (
	"encoding/json"
	"fmt"
	"host-api-service/pkg/traveler"
	"net"
	"os"
	"strconv"

	"k8s.io/klog"
)

type NodeResponse struct {
	ClusterName string `json:"clusterName"`
	NodeName    string `json:"node"`
	GPU         string `json:"GPU"`
}

func CreateSocket() {

	socketPath := "/var/lib/.keti/api-service.sock"
	if err := os.Mkdir("/var/lib/.keti/", os.ModeDir); err != nil {
		fmt.Println(err)
	}
	// 소켓 파일이 이미 존재하면 삭제
	if err := os.Remove(socketPath); err != nil && !os.IsNotExist(err) {
		fmt.Println("Error removing socket file:", err)
		return
	}

	// UNIX 도메인 소켓 생성
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		fmt.Println("Error creating socket:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on", socketPath)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 클라이언트로부터 메시지 수신
	buffer := make([]byte, 1024)
	var data []byte

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			// 오류 처리
			break
		}
		data = append(data, buffer[:n]...)
	}

	fmt.Printf("Received: %s\n", data)

	nodeName := os.Getenv("HOSTNAME")
	clusterName := string(buffer)
	gpuCount := ""

	if traveler.IsGPUNode() {
		gpuCount = strconv.Itoa(traveler.GetGPUs())
	} else {
		gpuCount = "0"
	}

	responseStruct := &NodeResponse{
		ClusterName: clusterName,
		NodeName:    nodeName,
		GPU:         gpuCount,
	}

	// 클라이언트에 응답
	response, err := json.Marshal(responseStruct)
	if err != nil {
		klog.Errorln(err)
	}
	_, err = conn.Write(response)
	if err != nil {
		fmt.Println("Error writing to connection:", err)
		return
	}
}
