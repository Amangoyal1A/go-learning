package main

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var inMemoryData string // In-memory data to be synced across nodes
func main() {
	// Initialize etcd client
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"}, // etcd server
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal("Failed to connect to etcd:", err)
	}
	defer cli.Close()
	// Start watching for changes
	go watchEtcdChanges(cli)
	// Simulate updating data
	for {
		var newData string
		fmt.Print("Enter new data to update: ")
		fmt.Scanln(&newData)
		updateEtcdData(cli, newData)
	}
}

// updateEtcdData updates the data in etcd and triggers sync
func updateEtcdData(cli *clientv3.Client, data string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := cli.Put(ctx, "my_app/in_memory_data", data)
	if err != nil {
		log.Println("[ERROR] Failed to update etcd:", err)
		return
	}
	fmt.Println("[SUCCESS] Data updated in etcd")
}

// watchEtcdChanges listens for data changes and updates in-memory data
func watchEtcdChanges(cli *clientv3.Client) {
	fmt.Println("[INFO] Watching for data changes...")
	rch := cli.Watch(context.Background(), "my_app/in_memory_data")
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("[WATCH] %s : %s\n", ev.Type, ev.Kv.Value)
			inMemoryData = string(ev.Kv.Value)
			fmt.Println("[SYNCED] In-memory data updated to:", inMemoryData)
		}
	}
}
