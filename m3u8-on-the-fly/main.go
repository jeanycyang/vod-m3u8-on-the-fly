package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func readM3u8() {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("./key.json"))
	if err != nil {
		panic(err)
	}
	bkt := client.Bucket("vod-m3u8")
	rc, err := bkt.Object("vod1.mov/hls.m3u8").NewReader(ctx)
	if err != nil {
		panic(err)
	}
	defer rc.Close()
	data, err := ioutil.ReadAll(rc)
	if err != nil {
		panic(err)
	}
	m3u8 := string(data)
	fmt.Print(m3u8)
}

func main() {
	readM3u8()
}
