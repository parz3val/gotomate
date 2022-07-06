package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/lijo-jose/gffmpeg/pkg/gffmpeg"
)
func getFF() gffmpeg.GFFmpeg {
	g,e:= gffmpeg.NewGFFmpeg("/opt/homebrew/bin/ffmpeg")
	if e!=nil {
		panic(e)
	}
	return g

}

func Mp4ToFrames(inFile, outDir string, fps int) error {
	
	// check if the directory exists
	_, err := ioutil.ReadDir(outDir)
	if err != nil {
		// create new directory if it doesn't exist
		err = os.MkdirAll(fmt.Sprintf("vault/%s",outDir), 0755)
	}

	empty := DirectoryEmpty(fmt.Sprintf("vault/%s",outDir))
	// if not empty 
	if !empty {
		return fmt.Errorf("Directory is not empty")
	}
	svc, e := New(getFF())
	if e != nil {
		return e
	}
	fmt.Printf("Extracting frames from %s to %s\n",  RelativeDirToAbsDir(fmt.Sprintf("vault/%s",outDir)), outDir)
	return svc.ExtractFrames(inFile, RelativeDirToAbsDir(fmt.Sprintf("vault/%s/",outDir)), fps)
}


type Service interface {
	ExtractFrames(inFile, outDir string, fps int) error
}

type service struct {
	ff gffmpeg.GFFmpeg
}

func New(ff gffmpeg.GFFmpeg) (Service, error) {
	return &service{ff: ff}, nil
}

func (svc *service) ExtractFrames(inFile, outDir string, fps int) error {
	bd := gffmpeg.NewBuilder()
	bd = bd.SrcPath(inFile).VideoFilters("fps=" + strconv.Itoa(fps)).DestPath(outDir + "/%03d.jpg")
	svc.ff = svc.ff.Set(bd)

	ret  := svc.ff.Start(nil)
	if ret.Err != nil {
		return ret.Err
	}
	return nil
}