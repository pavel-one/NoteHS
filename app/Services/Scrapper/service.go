package Scrapper

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Url struct {
	Title       string
	Url         string
	Screen      string
	Description string
}

func GetUrlInfo(url string, filename string, userId uint) (Url, error) {
	var buf []byte
	var title string

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoFirstRun,
		chromedp.NoSandbox,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Headless,
		chromedp.DisableGPU,
		chromedp.WindowSize(1280, 720),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))

	defer cancel()

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
	)

	if err != nil {
		return Url{}, err
	}

	time.Sleep(3 * time.Second)

	err = chromedp.Run(
		ctx,
		chromedp.CaptureScreenshot(&buf),
		chromedp.Title(&title),
	)

	if err != nil {
		return Url{}, err
	}

	screenshot, err := saveScreenshot(buf, filename, userId)
	if err != nil {
		return Url{}, err
	}

	return Url{
		Title:  title,
		Url:    url,
		Screen: screenshot,
	}, nil
}

func saveScreenshot(buffer []byte, name string, userId uint) (string, error) {

	dir := fmt.Sprintf("storage/screenshot/%v/", userId)
	path := dir + name + ".png"

	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(path, buffer, 0644)
	if err != nil {
		return "", err
	}

	return path, nil
}
