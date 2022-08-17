package main

import "github.com/phial3/go-notes/etcd/watch"

func main() {
	watch.WatchConfig("/registry/configs")
}
