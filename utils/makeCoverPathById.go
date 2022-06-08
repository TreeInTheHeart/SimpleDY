package utils

import (
	"SimpleDY/global"
	"strconv"
)

//MakeCoverPathById
/*通过id生成cover并返回Cover的地址
 */
func MakeCoverPathById(id uint64) string {
	CoverName := strconv.FormatUint(id, 10) + ".jpeg"
	CoverPath := global.VideoCoverPath + CoverName
	return CoverPath
}
