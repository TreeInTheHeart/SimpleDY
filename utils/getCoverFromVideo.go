package utils

import (
	"bytes"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"image"
	"image/jpeg"
	"io"
	"os"
)

// ReadFrameAsJpeg
/*从视频流中截取一帧并返回 需要在本地环境中安装ffmpeg并将bin添加到环境变量
 */
func readFrameAsJpeg(inFileName string) (io.Reader, error) {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	return buf, nil
}

//getCoverFromVideo
/*param 视频路径 存储路径
* 从视频路径的视频中截取一帧并保存到存储路径
 */
func GetCoverFromVideo(videoPath, coverPath string) error {
	f, err := os.OpenFile(coverPath, os.O_APPEND|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		return err
	}
	reader, err := readFrameAsJpeg(videoPath)
	if err != nil {
		return err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return err
	}
	err = jpeg.Encode(f, img, &jpeg.Options{Quality: 80})
	if err != nil {
		return err
	}
	return nil
}
