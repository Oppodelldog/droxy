package proxyfile

import (
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// NewClonesStrategy creates a new FileCreationStrategy that produces clones of droxy command
func NewClonesStrategy() FileCreationStrategy {
	return &ClonesStrategy{}
}

//ClonesStrategy contains the implementation of creating clones of droxy execuable
type ClonesStrategy struct{}

//CreateProxyFile creates a clone of the given commandBinaryFilePath to commandNameFilePath
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
	defer func() {
		err = sf.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()
	if err = os.Remove(cleanDst); err != nil && !os.IsNotExist(err) {
		return err
	}
	df, err := os.OpenFile(cleanDst, os.O_CREATE|os.O_WRONLY, 0766)
	if err != nil {
		return err
	}
	defer func() {
		err = df.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	_, err = io.Copy(df, sf)

	return err

}
