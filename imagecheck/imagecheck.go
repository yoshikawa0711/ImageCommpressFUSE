package imagecheck

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func IsImage(filepath string) bool {
	f, _ := os.Open(filepath)
	defer f.Close()

	_, _, err := image.DecodeConfig(f)

	// 画像フォーマットではない場合はエラーが出る
	if err != nil {
		return false
	}
	return true
}
