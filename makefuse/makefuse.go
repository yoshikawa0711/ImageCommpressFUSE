package makefuse

import (
	"flag"
	"log"
	"strconv"

	"github.com/hanwen/go-fuse/v2/fuse"
)

// test用に作成したファイルシステムの構造体
type TestFS struct {
	fuse.RawFileSystem
}

func NewTestFS() *TestFS {
	return &TestFS{
		RawFileSystem: fuse.NewDefaultRawFileSystem(),
	}
}

// ファイルの属性が取得された時に，呼び出される関数
func (fs *TestFS) GetAttr(cancel <-chan struct{}, input *fuse.GetAttrIn, out *fuse.AttrOut) (code fuse.Status) {
	log.Print("Get Attr.")
	out.Mode = fuse.S_IFDIR | 0755

	code = fuse.Status(0)
	return code
}

// ファイルが新しく作られた時に，呼び出される関数
func (fs *TestFS) Create(cancel <-chan struct{}, input *fuse.CreateIn, name string, out *fuse.CreateOut) (code fuse.Status) {
	log.Print(name + " is Created.")

	code = fuse.Status(0)
	return code
}

// ファイルが開かれた時に，呼び出される関数
func (fs *TestFS) Open(cancel <-chan struct{}, input *fuse.OpenIn, out *fuse.OpenOut) (status fuse.Status) {
	log.Print(strconv.FormatUint(input.InHeader.NodeId, 10) + " is Opened.")

	status = fuse.Status(0)
	return status
}

func (fs *TestFS) Read(cancel <-chan struct{}, input *fuse.ReadIn, buf []byte) (fuse.ReadResult, fuse.Status) {
	log.Print(strconv.FormatUint(input.InHeader.NodeId, 10) + " is Read.")
	
	status := fuse.Status(0)
	return fuse.ReadResultData(buf), status
}

// FUSEを作る本体部分
func makefuse() {
	// デバッグに関わる部分
	debag := flag.Bool("debag", false, "print debag data")
	flag.Parse()

	// 入力されている個数が正しいかを判断する
	if len(flag.Args()) < 1 {
		log.Fatal("Usage:\n makefuse MOUNTPOINT")
	}

	testfs := NewTestFS()
	mountpoint := flag.Arg(0)
	// FUSEのマウントに関わるオプションを設定する
	opts := &fuse.MountOptions{FsName: "TestFS", Name: "Test", Debug: true}
	opts.Debug = *debag

	fssrv, err := fuse.NewServer(testfs, mountpoint, opts)
	if err != nil {
		log.Panic(err)
	}

	fssrv.Serve()
}
