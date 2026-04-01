package examples

import "fmt"

/// A sample of the strategy pattern in Go
// the PhoneCameraApp struct uses a ShareStrategy interface to share photos
// - The ShareStrategy interface defines a family of algorithms (interchangeable behavior) for sharing photos
//   - Each strategy can be swapped in or out without modifying the PhoneCameraApp struct

func strategySample() {
	a := &PhoneCameraApp{}

	a.SetShareStrategy(&TextShareStrategy{})
	a.SharePhoto("some-photo.jpg")

	a.SetShareStrategy(&EmailShareStrategy{})
	a.SharePhoto("some-photo.jpg")

	a.SetShareStrategy(&SocialMediaShareStrategy{})
	a.SharePhoto("some-photo.jpg")
}

type PhoneCameraApp struct {
	shareStrategy ShareStrategy
}

func (a *PhoneCameraApp) TakePhoto() {
	fmt.Println("Taking photo")
}
func (a *PhoneCameraApp) EditPhoto() {
	fmt.Println("Editing photo")
}
func (a *PhoneCameraApp) SharePhoto(photo string) {
	if a.shareStrategy == nil {
		fmt.Println("No share strategy set")
		return
	}
	a.shareStrategy.SharePhoto(photo)
}
func (a *PhoneCameraApp) SetShareStrategy(strategy ShareStrategy) {
	a.shareStrategy = strategy
}

type ShareStrategy interface {
	SharePhoto(photo string)
}

type TextShareStrategy struct {
}

func (t *TextShareStrategy) SharePhoto(photo string) {
	println("Sharing photo via text")
}

type EmailShareStrategy struct {
}

func (e *EmailShareStrategy) SharePhoto(photo string) {
	println("Sharing photo via email")
}

type SocialMediaShareStrategy struct {
}

func (s *SocialMediaShareStrategy) SharePhoto(photo string) {
	println("Sharing photo via social media")
}
