package commpressimage

import (
//	"log"
	"os"
	"image/jpeg"

	"github.com/nfnt/resize"
)

func ImageCommpress(filepath string) bool {
	file, err := os.Open(filepath)
	if err != nil {
	//	log.Fatal(err)
		return false
	}

	img, err := jpeg.Decode(file)
	if err != nil {
	//	log.Fatal(err)
		return false
	}
	file.Close()

	m := resize.Resize(0, 0, img, resize.Lanczos3)
	out, err := os.Create(filepath + "resized.jpg")
	if err != nil  {
	//	log.Fatal(err)
		return false
	}
	defer out.Close()

	jpeg.Encode(out, m, nil)

	return true
}
