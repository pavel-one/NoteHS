package Scrapper

import (
	"app/Services/ImageService"
	"context"
	"github.com/chromedp/chromedp"
	"log"
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

	screenshot, err := ImageService.SaveImageWithBuff(buf, filename, userId, ".png")
	if err != nil {
		return Url{}, err
	}

	return Url{
		Title:  title,
		Url:    url,
		Screen: screenshot,
	}, nil
}
