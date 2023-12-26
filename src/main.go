package main

import "github.com/yangas1/usdtpay/bootstrap"

func main() {
	defer func() {
		if err := recover(); err != nil {
			println("[Start Server Err!!!] ", err)
		}
	}()
	bootstrap.Start()
}
