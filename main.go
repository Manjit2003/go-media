package main

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func main() {

	if err := ffmpeg.Input("examples/example1.mp4").Output("examples/output.m3u8", ffmpeg.KwArgs{
		"b:v":              "1M",
		"g":                "60",
		"format":           "hls",
		"hls_time":         "2",
		"hls_segment_size": "500000",
		"hls_list_size":    "0",
		"hls_flags":        "delete_segments",
	}).ErrorToStdOut().Run(); err != nil {
		panic(err)
	}

}
