package utils

import (
	"errors"
	"encoding/hex"
	"github.com/sirupsen/logrus"
	"os"
)

func StringToMac(stringMac string) ([]byte, error) {
	var (
		byteMac = make([]byte, 6)
		byteBuf []byte
		err error
	)

	if len(stringMac) != 17 {
		return nil, errors.New("MAC length is incorrect")
	}

	for i := 0; i < 6; i += 1 {
		byteBuf, err = hex.DecodeString(string(stringMac[3 * i : 3 * i + 2]))

		if err != nil {
			return nil, err
		}

		if len(byteBuf) != 1 {
			return nil, errors.New("MAC is incorrect")
		}

		byteMac[i] = byteBuf[0]
	}

	return byteMac, nil
}

func InitLogger(logPath string) (*logrus.Logger, error) {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	if logPath != "" {
		file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logger.SetOutput(os.Stdout)
			return logger, err
		}
		logger.SetOutput(file)
	}

	return logger, nil
}