
package main

import (
        "log"
        "fmt"
        //"time"
        "os"
        "path/filepath"
	//"path"
	//"strings"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
)

const (
        imgpath = "d:\\image"
        interval = 1
)

var imglist [100]string
var pos     int
var num     int


func main() {
        
	mw := new(MyMainWindow)
	//var openAction *walk.Action

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Walk Image Viewer Example",
		Size:    Size{800, 600},
		Layout:  HBox{},
		Children: []Widget{
			ImageView{
				AssignTo: &(mw.imageView.iv),
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	} 

        mw.SetFullscreen(true)

        mw.initialize();

        mw.Run()
}

type MyImageView struct {
	iv *walk.ImageView 
}

type MyMainWindow struct {
	*walk.MainWindow
        imageView    MyImageView
	prevFilePath string
}

func (mw *MyMainWindow) initialize() error {
    pos, num = 0, 1
    getfiles()
    pos = 1
    mw.imageView.openImage()
    return nil
}

func (mw *MyMainWindow) continue() error {
    pos, num = 0, 1
    getfiles()
    pos = 1
    mw.imageView.openImage()
    return nil
}

func getfiles() error {
	if err := filepath.Walk(imgpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			if info == nil {
				return filepath.SkipDir
			}
		}

		if info.IsDir() || num>100 {
		} else {
                    imglist[num] = path
                    num += 1
                }

		return nil
	}); err != nil {
		return err
	}

	return nil
}

/*
func (mw *MyMainWindow) openAction_Triggered() {
	if err := mw.displayFiles(imgpath); err != nil {
		log.Print(err)
	}
}
*/

func (iv *MyImageView) openImage() error {

	img, err := walk.NewImageFromFile(imglist[pos])
	if err != nil {
		return err
	}

	var succeeded bool
	defer func() {
		if !succeeded {
			img.Dispose()
		}
	}()


	if err := iv.iv.SetImage(img); err != nil {
		return err
	}

        iv.iv.CustomWidget.Invalidate()

	succeeded = true

	return nil
}

func (iv *MyImageView) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
    fmt.Println("aaa")
    switch msg {
        case win.WM_LBUTTONDOWN:
            pos += 1
            iv.openImage()
        case win.WM_KEYDOWN:
            switch wParam {
                case win.VK_LEFT: fallthrough
                case win.VK_UP:
                    pos -= 1
                    iv.openImage()
                case win.VK_RIGHT: fallthrough
                case win.VK_DOWN:
                    pos += 1
                    iv.openImage()
            }
    }
    
    return iv.iv.WidgetBase.WndProc(hwnd, msg, wParam, lParam)
}

