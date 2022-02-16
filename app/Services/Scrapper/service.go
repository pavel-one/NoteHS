package Scrapper

import (
	"context"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"os"
)

type Url struct {
	Title       string
	Url         string
	Screen      string
	Description string
}

func GetUrlInfo(url string, filename string) (Url, error) {
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
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.CaptureScreenshot(&buf),
		chromedp.Title(&title),
	)

	if err != nil {
		return Url{}, err
	}

	screenshot, err := saveScreenshot(buf, filename)
	if err != nil {
		return Url{}, err
	}

	return Url{
		Title:  title,
		Url:    url,
		Screen: screenshot,
	}, nil
}

func saveScreenshot(buffer []byte, name string) (string, error) {

	dir := "storage/screenshot/"
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
