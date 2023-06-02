package main

import (
	"log"

	"github.com/yangbo-1997/goav/avcodec"
	"github.com/yangbo-1997/goav/avdevice"
	"github.com/yangbo-1997/goav/avfilter"
	"github.com/yangbo-1997/goav/avformat"
	"github.com/yangbo-1997/goav/avutil"
	"github.com/yangbo-1997/goav/swresample"
	"github.com/yangbo-1997/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
