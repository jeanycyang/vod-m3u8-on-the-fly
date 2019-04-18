package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"cloud.google.com/go/storage"
	"github.com/grafov/m3u8"
	"google.golang.org/api/option"
)

// GCPCreds is GCP Credentials
type GCPCreds struct {
	ClientEmail string `json:"client_email"`
	PrivateKey  string `json:"private_key"`
}

var gcpCreds GCPCreds

func readM3u8() string {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("./key.json"))
	if err != nil {
		panic(err)
	}
	rc, err := client.Bucket("vod-m3u8").Object("vod2.mov/hls.m3u8").NewReader(ctx)
	if err != nil {
		panic(err)
	}
	defer rc.Close()
	p, _, err := m3u8.DecodeFrom(rc, true)
	playlist := p.(*m3u8.MediaPlaylist)

	for i := range playlist.Segments {
		if playlist.Segments[i] == nil {
			break
		}
		playlist.Segments[i].URI = signURL(playlist.Segments[i].URI)
	}
	// Package grafov/m3u8's bug, it will automatically
	// add default key as the first seq's key
	// so we need to sign "the first seq's key" instead of default key
	playlist.Segments[0].Key.URI = signURL(playlist.Key.URI)
	return playlist.String()
}

func signURL(fileName string) string {
	url, err := storage.SignedURL("vod-m3u8", fileName, &storage.SignedURLOptions{
		GoogleAccessID: gcpCreds.ClientEmail,
		PrivateKey:     []byte(gcpCreds.PrivateKey),
		Method:         "GET",
		Expires:        time.Now().Add(48 * time.Hour),
	})
	if err != nil {
		panic(err)
	}
	return url
}

func main() {
	keyFile, _ := ioutil.ReadFile("./key.json")
	json.Unmarshal(keyFile, &gcpCreds)
	signedM3u8 := readM3u8()
	fmt.Print(signedM3u8)
}
