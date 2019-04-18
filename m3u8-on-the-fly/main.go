package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/grafov/m3u8"
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
	p, _, err := m3u8.DecodeFrom(rc, true)
	playlist := p.(*m3u8.MediaPlaylist)
	fmt.Printf("%v", playlist)
}

func main() {
	readM3u8()
}
