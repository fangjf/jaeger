// Code generated by "esc -pkg frontend -o examples/hotrod/services/frontend/gen_assets.go -prefix examples/hotrod/services/frontend/web_assets examples/hotrod/services/frontend/web_assets"; DO NOT EDIT.

package frontend

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/index.html": {
		name:    "index.html",
		local:   "examples/hotrod/services/frontend/web_assets/index.html",
		size:    3552,
		modtime: 1592364558,
		compressed: `
H4sIAAAAAAAC/9RXX1PbSBJ/96fom82d5DOSMIZAjOUtDmcJ2cuSM5Ct3FYeRqO2NEaaUWZGtlmK7341
+mNkSK72cdcPMNPd092//mtPUpNn0x7AJEdDgaVUaTQhuby+8k5Ojt54Q/LEFTTHkKw4rgupDAEmhUFh
QrLmsUnDGFecoVdd9oALbjjNPM1ohuHQ39+DnG54XuZdUqlRVXcaZRju18ZSpLE9AEwMNxlO30kzv5qB
B3Meo4YrATPMqYgnQc2vZTVTvDCgFQtJakyhx0HAZIz+8muJ6t5nMg/qozfyh/7Qz7nwl5pMJ0H9tNGT
cXEHCrOQaHOfoU4RDYFU4eJJb043LBZ+JKXRRtHCXqz+LSEY+SP/OGBaP9Eqg0xrAlwYTBQ39yHRKR2d
HHr/+vSZ8+vLn/DnYXyRv5+f3d2z8t3Zu3kyOrjKb9l6fSzFaP45Tg4/0cHH/PpG/x78/PpkFcVvl+lh
SYApqbVUPOEiJFRIcZ/LUpP/E5w/CmL5HMPymxBu2NHlf3i0f3D8dXW/vP6weLe8+kD/fbcof/20+e/m
9qM4f392nB3k57/+cllcvMkvzmcn64tfLtnH2fHNhn4fwlOCGjA2L9OeX5Y8hgfIqUq48IwsxjA8Kjan
8NjzU2mUjL2oNEYKeICCxjEXyRgO9q0EK5WWagyFtEDU6a6S/W8pGadyhQoeXr5d8MygGkOkeJIagVq7
J0d/71sVPzQqMpl8x9MfDC++w6rABg1a2xlB2xqTSMb3TWpjvgKWUa1DYjuScoGqSfsutwoXzVCZ+q/H
xULa6MZ8tZVnaDG1V9uNQ9t/MPev/Jk/CdJhl3c4nWA+fdGWmE8nQXrYkey4oeSaPHFeQsi8PPZGYA86
914/k60LoKDiBdV+GiWRERAZUQGsDlEm2R3spJN8U0FMDfVYqY3MUYVkeDAi0zllKWaOhp8yqWgGM9Q8
EXoSWDeeIenG8s8ObvTmgExvlMzhPJVMZtRwVH95VMejIZm+pwUVqNHmSqMyf/1kHb0+JtOznP7ORQLn
crFAhLmk2qD6I+CeXy1OHofE8IJMzzPO7kAKaM1Vqx5oJFcIRoJUMSqgwKjyv6foac6RFnuGNH4+XoLu
fNmyJkE9z3rbTTXt9RalYIZLAQupcmpmpaL26sbNoQ8PPYAVVRBDCC0VAnCH+9UH/gnD+t/r/f5pI1sK
bjSE4ORcOJao0JRKwAdqUl/JUsRu3IdBLXfae+z17CuWcRTm9vZyBmFXtD5SEcvc7Tf2rC37JqPazPFr
idpUz/ZPe71XLqm2Fun79puXSz7LUsEao8aCo4HHY7vhlBTJlMCga3oAxK6DmtVv1O2WUt9nNpluGzwX
V6YO1I47g0EbD9XxcNeUZ63vvGrfLBTq9JwqCOGV+8olnSVH+n6hsEARu063maonHqOKVBtjxnVBDUtt
Mdd15fu/Kfw6Bmew9WjgfGlWiS0Tp++zlGexQuH2f9v/ss3otmhDwJXxDVUJGt+2j0bjt9xW2m5QVLYC
HpqadJYUE1ReRJOEJuiMwdGoNZcidJ7H39lrA1bxtq72AB6tBSaFlhn6mUzcxtLWzwgXUiGEMKMGfSHX
rk0hQBDArUb7rUKhMHA7vwSqIaIaC2pSW/1Al3TTGtONOsv8qHDBNxDCmotYrv1MsqoJfMu0PWxt7wh2
Ln8LgQQEfuzSxuA4lVOvfGvT7bAG4ARxk7Yft3OpilCbgAE4/xBSaKzIO72x14S7Ccq4PexV1BxNKuMx
OBdvb5yapEvGUOsxbCvZZnQPDG7MtaGm1P1tBm046MJUJdANbjOJOzmxOraManS0YyN8MWaoof7bm7Ot
eFv1dec6zY+SSTS1YCvpmeKrOgyTIJoCVYqvbIVzAZVMa2sADjTV3i2iur4yalCw+5rnVrC8unTsWHJy
/cVpXHq0kXrsn/Ye60LqfE2eBPUPu/8FAAD//xNRNP3gDQAA
`,
	},

	"/": {
		name:  "/",
		local: `examples/hotrod/services/frontend/web_assets`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"examples/hotrod/services/frontend/web_assets": {
		_escData["/index.html"],
	},
}
