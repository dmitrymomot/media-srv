package resizer

import (
	"image"
	"os"
	"testing"
)

func TestResize(t *testing.T) {
	txt, err := os.Open("./testdata/test.txt")
	if err != nil {
		panic(err)
	}
	defer txt.Close()

	t.Run("not image", func(t *testing.T) {
		got, err := Resize(txt, 100, 100)
		if err == nil {
			t.Errorf("Resize() error = %v, wantErr %v", err, image.ErrFormat)
			return
		}
		if got != nil {
			t.Errorf("Resize() = %v, wantErr %v", got, image.ErrFormat)
			return
		}
	})

	img, err := os.Open("./testdata/image.png")
	if err != nil {
		panic(err)
	}
	defer img.Close()

	t.Run("success", func(t *testing.T) {
		got, err := Resize(img, 100, 100)
		if err != nil {
			t.Errorf("Resize() error = %v", err)
			return
		}
		if got == nil {
			t.Errorf("Resize() = %v, want io.ReadSeeker interface", got)
		}
	})

}
