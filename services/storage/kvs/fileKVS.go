package kvs

import (
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	log "github.com/jeanphorn/log4go"
	"github.com/nikita-tomilov/gotsdb/utils"
	"io/ioutil"
	"os"
	"sync"
)

type FileKVS struct {
	Path string `summer.property:"kvs.fileKVSPath|/tmp/gotsdb/kvs"`
	lock sync.Mutex
}

func (f *FileKVS) createKey(s []byte) string {
	return base58.Encode(s)
}

func (f *FileKVS) getKey(s string) []byte {
	return base58.Decode(s)
}

func (f *FileKVS) toFilename(key []byte) string {
	return f.Path + "/" + f.createKey(key)
}

func (f *FileKVS) InitStorage() {
	os.MkdirAll(f.Path, os.ModePerm)
	log.Warn("FILE-BASED KVS storage initialized at %s", f.Path)
}

func (f *FileKVS) CloseStorage() {
	//nothing here
}

func (f *FileKVS) Save(key []byte, value []byte) {
	f.lock.Lock()
	log.Warn("Request on setting value on key %s", string(key))
	err := ioutil.WriteFile(f.toFilename(key), value, 0644)
	utils.Check(err)
	f.lock.Unlock()
}

func (f *FileKVS) KeyExists(key []byte) bool {
	f.lock.Lock()
	fname := f.toFilename(key)
	ok := utils.FileExists(fname)
	f.lock.Unlock()
	return ok
}

func (f *FileKVS) Retrieve(key []byte) []byte {
	isPresent := f.KeyExists(key)
	log.Warn("Request on getting value on key %s", string(key))
	if !isPresent {
		return nil
	}
	fname := f.toFilename(key)
	f.lock.Lock()
	val, err := ioutil.ReadFile(fname)
	utils.Check(err)
	f.lock.Unlock()
	return val
}

func (f *FileKVS) Delete(key []byte) {
	f.lock.Lock()
	fname := f.toFilename(key)
	utils.DeleteFile(fname)
	f.lock.Unlock()
}

func (f *FileKVS) GetAllKeys() [][]byte {
	f.lock.Lock()
	files := utils.GetFileNames(f.Path)
	keys := make([][]byte, len(files))
	for i, file := range files {
		keys[i] = f.getKey(file)
	}
	f.lock.Unlock()
	return keys
}

func (f *FileKVS) String() string {
	return fmt.Sprintf("Simple FS-based KVS at path %s", f.Path)
}