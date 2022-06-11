package utils

import (
	"SimpleDY/global"
	"strconv"
	"strings"
)

// MakeVideoPathById
/*通过视频ID生成视频存放路径  格式为 id.视频类型 eg: data/video/32.MP4
 */
func MakeVideoPathById(id uint64, filename string) string {
	s := strings.Split(filename, ".")
	videoName := strconv.FormatUint(id, 10) + "." + s[len(s)-1]
	return global.VideoLibPath + videoName
}
