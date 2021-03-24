package go_insect

import (
	"go.etcd.io/etcd/clientv3"
)

var GlobalEtcdCLI = &clientv3.Client{}

var GlobalEtcdAddress = "http://0.0.0.0:2379"
var GlobalEtcdTTL int64 = 10

var INSECT_SERVER_NAME = ""
var INSECT_SERVER_URL = ""
var INSECT_SERVER_PORT int = 0
