//+------------------------------------------------+
//|                                                |
//|                                                |
//|                                                |
//|           ######    ######  ########           |
//|          ##    ##  ##    ## ##                 |
//|          ##        ##       ##                 |
//|          ##   #### ##       ######             |
//|          ##    ##  ##       ##                 |
//|          ##    ##  ##    ## ##                 |
//|           ######    ######  ########           |
//|                                                |
//|                                                |
//|                                                |
//+------------------------------------------------+
package main

import (
	"fmt"
	"gce/gnet"
)

func main() {
	//1 创建一个server 使用gce的api
	s := gnet.NewServer("gce v0.1")
	//2 启动server
	s.Serve()
	fmt.Println("hello world")
}