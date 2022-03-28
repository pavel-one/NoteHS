package ImageService

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gabriel-vasile/mimetype"
	"image"
	"image/color"
	"mime/multipart"
	"os"
)

func SaveImageWithBuff(buffer []byte, name string, id uint, extension string) (string, error) {
	src, _, err := image.Decode(bytes.NewReader(buffer))
	if err != nil {
		return "", err
	}

	return saveImage(src, name, id, ".png")
}

func SaveImageWithForm(file multipart.FileHeader, name string, id uint) (string, error) {
	r, _ := file.Open()

	mime, err := mimetype.DetectReader(r)
	if err != nil {
		return "", err
	}

	r, _ = file.Open()
	src, err := imaging.Decode(r)
	if err != nil {
		return "", err
	}

	return saveImage(src, name, id, mime.Extension())
}

func saveImage(src image.Image, name string, id uint, extension string) (string, error) {
	name = getHashName(name)

	dir, err := getAndCreateImageDir(id)
	if err != nil {
		return "", err
	}

	path := dir + getHashName(name) + extension

	rgba := imaging.Fill(src, 350, 200, imaging.Center, imaging.Lanczos)

	dst := imaging.New(350, 200, color.NRGBA{})
	dst = imaging.Paste(dst, rgba, image.Pt(0, 0))
	err = imaging.Save(dst, path)

	if err != nil {
		return "", err
	}

	return path, nil
}

func getAndCreateImageDir(id uint) (string, error) {
	dir := fmt.Sprintf("storage/screenshot/%v/", id)

	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return "", err
	}

	return dir, nil
}

func getHashName(name string) string {
	hasher := md5.New()
	hasher.Write([]byte(name))

	return hex.EncodeToString(hasher.Sum(nil))
}
