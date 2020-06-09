package proxyfile

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func copyFile(src, dst string) error {
	sf, err := os.Open(src)
	if err != nil {
		return err
	}

	defer func() {
		err = sf.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	if err = os.Remove(dst); err != nil && !os.IsNotExist(err) {
		return err
	}

	df, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, 0766)
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
