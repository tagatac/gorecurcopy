// Package gorecurcopy provides recursive copying in Go (golang) with a
// minimum of extra packages. Original concept by Oleg Neumyvakin
// (https://stackoverflow.com/users/1592008/oleg-neumyvakin) and modified by:
// Dirk Avery (2019)
// David Tagatac (2022)
package gorecurcopy

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

//go:generate mockgen -destination=mock_gorecurcopy/mock_gorecurcopy.go github.com/tagatac/gorecurcopy Copier

type (
	Copier interface {
		// CopyDirectory recursively copies a src directory to a destination.
		CopyDirectory(src, dst string) error
		// Copy copies a src file to a dst file where src and dst are regular files.
		Copy(src, dst string) error
		// CopySymLink copies a symbolic link from src to dst.
		CopySymLink(src, dst string) error
	}

	copier struct {
		afero.Fs
	}
)

func NewCopier(fs afero.Fs) Copier {
	return copier{Fs: fs}
}

func (c copier) CopyDirectory(src, dst string) error {
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		sourcePath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dst, entry.Name())

		fileInfo, err := c.Fs.Stat(sourcePath)
		if err != nil {
			return err
		}

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if err := c.createDir(destPath, 0755); err != nil {
				return err
			}
			if err := c.CopyDirectory(sourcePath, destPath); err != nil {
				return err
			}
		case os.ModeSymlink:
			if err := c.CopySymLink(sourcePath, destPath); err != nil {
				return err
			}
		default:
			if err := c.Copy(sourcePath, destPath); err != nil {
				return err
			}
		}

		isSymlink := entry.Mode()&os.ModeSymlink != 0
		if !isSymlink {
			if err := os.Chmod(destPath, entry.Mode()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (c copier) Copy(src, dst string) error {
	sourceFileStat, err := c.Fs.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := c.Fs.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := c.Fs.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}

func (c copier) exists(path string) bool {
	if _, err := c.Fs.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func (c copier) createDir(dir string, perm os.FileMode) error {
	if c.exists(dir) {
		return nil
	}

	if err := c.Fs.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}

func (c copier) CopySymLink(src, dst string) error {
	osfs, ok := c.Fs.(afero.OsFs)
	if !ok {
		return errors.New("copying symlinks is only possible in a OS filesystem")
	}
	link, err := osfs.ReadlinkIfPossible(src)
	if err != nil {
		return err
	}
	return os.Symlink(link, dst)
}
