package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

var filePath string
var pingCount int
var outputFile string

func init() {
	flag.StringVar(&filePath, "file", "", "List of IP addresses")
	flag.StringVar(&outputFile, "out", "", "Output file")
	flag.IntVar(&pingCount, "count", 1, "Count of ping requests")
	flag.Parse()

	_, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("File does not exists")
		os.Exit(1)
	}
}

func ping(ip string, count int) (bool, probing.Statistics) {

	// parsed_ip := net.ParseIP(ip)
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	pinger, err := probing.NewPinger(ip)

	defer cancel()

	if err != nil {
		fmt.Println("Cannot init pinger")
		os.Exit(3)
	}

	pinger.Count = count

	err = pinger.RunWithContext(ctx)

	if err != nil {
		fmt.Println("Unreachable")
		return false, *pinger.Statistics()
	}

	fmt.Println(err)

	return true, *pinger.Statistics()
}

func main() {

	if outputFile == "" || filePath == "" {
		flag.Usage()
		fmt.Println("Out/input file is empty")
		os.Exit(4)
	}
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Cannot read file")
		os.Exit(2)
	}

	defer file.Close()

	out, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)

	if err != nil {
		fmt.Println("Cannot create file")
		os.Exit(5)
	}

	defer out.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "scan error")
	}

	for scanner.Scan() {
		line := scanner.Text()
		state, stats := ping(line, pingCount)

		if state == true {
			reachable := fmt.Sprintf("REACHABLE: %s packet_sent: %d packet_recv: %d packet_loss: %f \n", stats.Addr, stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
			_, err := out.WriteString(reachable)
			if err != nil {
				fmt.Println("Cannot append to file")
				os.Exit(6)
			}
		} else {
			unreachable := fmt.Sprintf("UNREACHABLE: %s packet_sent: %d packet_recv: %d packet_loss: %f \n", stats.Addr, stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
			_, err := out.WriteString(unreachable)
			if err != nil {
				fmt.Println("Cannot append to file")
				os.Exit(6)
			}
		}
	}
}
