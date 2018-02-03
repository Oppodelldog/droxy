package proxyfile

import (
	"io"
	"os"
	"path/filepath"
)

func NewClonesStrategy() *ClonesStrategy {
	return &ClonesStrategy{}
}

type ClonesStrategy struct{}

func (s *ClonesStrategy) CreateProxyFile(commandBinaryFilePath, commandNameFilePath string) error {

	cleanSrc := filepath.Clean(commandBinaryFilePath)
	cleanDst := filepath.Clean(commandNameFilePath)
	if cleanSrc == cleanDst {
		return nil
	}
	sf, err := os.Open(cleanSrc)
	if err != nil {
		return err
	}
	defer sf.Close()
	if err := os.Remove(cleanDst); err != nil && !os.IsNotExist(err) {
		return err
	}
	df, err := os.OpenFile(cleanDst, os.O_CREATE|os.O_WRONLY, 0766)
	if err != nil {
		return err
	}
	defer df.Close()

	_, err = io.Copy(df, sf)

	return err

}
