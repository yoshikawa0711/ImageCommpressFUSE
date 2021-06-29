package main

import (
	"log"
	"time"

	"github.com/hanwen/go-fuse/v2/fuse"
	"github.com/hanwen/go-fuse/v2/fuse/nodefs"
)

type Root struct {
	nodefs.Node
}

func (root *Root) GetAttr(out *fuse.Attr, file nodefs.File, ctx *fuse.Context) fuse.Status {
	out.Mode = 0755
	out.Atime = uint64(time.Now().Unix())
	out.Mtime = uint64(time.Now().Unix())

	return fuse.OK
}

func main() {
	root := &Root{
		Node: nodefs.NewDefaultNode(),
	}

	opts := nodefs.NewOptions()

	s, _, err := nodefs.MountRoot("/mnt/fuse", root, opts)
	if err != nil {
		log.Panic(err)
	}

	s.Serve()
}
