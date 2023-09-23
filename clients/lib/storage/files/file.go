package files

import (
	"encoding/gob"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	e "telego/clients/lib/error"
	"telego/clients/lib/storage"
	"time"
)

type Storage struct {
	basePath string
}

const defaultPermission = 0774

var ErrNoSavedFiles = errors.New("no saved files")

func New(path string) Storage {
	return Storage{basePath: path}
}

func Save(page *storage.Page) (err error) {
	defer func() {err = e.Wrap("cant save page", err)}()

	// using for Windows systems
	fPath := filepath.Join(s.basePath, page.UserName)

	if err := os.MkdirAll(fPath, defaultPermission); err != nil{
		return err
	} 	
		
	fName, err := fileName(page)
	if err != nil {
		return err
	}

	fPath = filepath.Join(fPath, fName)

	file, err := os.Create(fPath)
	if err != nil {
		return err
	}

	defer func(){ _=file.Close()}()

	if err:= gob.NewEncoder(file).Encode(page);err != nil{
		return err
	}

	return nil
}

func (s Storage) PickRandom(userName string) (page *storage.Page,err error) {
	defer func() {err = e.Wrap("cant pick random page", err)}()

	path := filepath.Join(s.basePath, userName)

	files,err := os.ReadDir(path) 
	if err != nil {
		return nil, err
	}

	if len(files) == 0{
		return nil, ErrNoSavedFiles
	}

	seed := time.Now().UnixNano()

	rand.New(rand.NewSource(seed))
	

	n:= rand.Intn(len(files))

	file := files[n]

	return s.decodePage(filepath.Join(path, file.Name()))

}

func (s Storage) Remove(p  *storage.Page) error {
	fileName,err := fileName(p)
	if err != nil {
		return e.Wrap("cant remove file", err)
	}

	path := filepath.Join(s.basePath,p.UserName, fileName)

	msg := fmt.Sprintf("cant remove file %s", path)

	if err := os.Remove(path);err != nil {
		return e.Wrap(msg, err)
	}
	return nil
} 

func (s Storage) IsExist(p *storage.Page)(bool, error)  {
	fileName,err := fileName(p)
	if err != nil {
		return false, e.Wrap("cant check is file exists", err)
	}

	path := filepath.Join(s.basePath, p.UserName, fileName)

	switch _,err = os.Stat(path); {	
	case errors.Is(err, os.ErrNotExist): 
		return false, nil
	case err!= nil:
		msg := fmt.Sprintf("cant check is file %s exists", path)
		return false, e.Wrap(msg, err)
	}
	return true, nil
	
}

func (s Storage) decodePage(filePath string) (*storage.Page, error)  {
	f,err := os.Open(filePath)
	if err != nil {
		return nil, e.Wrap("cant decode page:", err)
	}
	defer func() {_=f.Close()}()

	var p storage.Page

	if err:= gob.NewDecoder(f).Decode(&p); err != nil {
		return nil, e.Wrap("Cant decode page", err)
	}

	return &p, nil
}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}