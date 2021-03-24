package go_insect

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	uuid "github.com/satori/go.uuid"
	"go.etcd.io/etcd/clientv3"
)

func initServerIP() {
	if INSECT_SERVER_URL != "" {
		return
	}
	address, _ := net.InterfaceAddrs()
	for _, addr := range address {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				INSECT_SERVER_URL = ipnet.IP.String()
			}
		}
	}
}

func initEtcdCli() {
	cfg := clientv3.Config{
		Endpoints:   []string{GlobalEtcdAddress},
		DialTimeout: 5 * time.Second,
	}
	cli, err := clientv3.New(cfg)
	if err != nil {
		GlobalEtcdCLI = nil
	}
	GlobalEtcdCLI = cli
}

func putWithLease() {

	key := fmt.Sprintf("%s_%s", INSECT_SERVER_NAME, uuid.NewV4().String())
	value := INSECT_SERVER_URL
	if INSECT_SERVER_PORT != 0 {
		value = fmt.Sprintf("%s:%d", INSECT_SERVER_URL, INSECT_SERVER_PORT)
	}

	lease := clientv3.NewLease(GlobalEtcdCLI)
	leaseResp, err := lease.Grant(context.TODO(), GlobalEtcdTTL)
	if err != nil {
		fmt.Println(err)
		return
	}
	leaseID := leaseResp.ID
	_, err = GlobalEtcdCLI.KV.Put(context.TODO(), key, value, clientv3.WithLease(leaseID))
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("lease-", key, "-", GlobalEtcdTTL, "sec")
}

func EtcdProxy() {
	if INSECT_SERVER_NAME != "" {
		initEtcdCli()
		initServerIP()
		for {
			putWithLease()
			time.Sleep(time.Duration(GlobalEtcdTTL) * time.Second)
		}
	}
}
