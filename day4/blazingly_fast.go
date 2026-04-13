package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"io"
	"log"
	"log/syslog"
	"os"
)

func WriteSyslog(state string) {
	switch state {
	case "success":
		logWriter, err := syslog.New(syslog.LOG_INFO, "filebackupper")
		if err != nil {
			log.Fatal("Failed to write to syslog: ", err)
		}

		logWriter.Info("Successfully backupped files")
	case "fail":
		logWriter, err := syslog.New(syslog.LOG_ALERT, "filebackupper")
		if err != nil {
			log.Fatal("Failed to write to syslog: ", err)
		}

		logWriter.Info("Failed to backup files")
	}
}

func CreateSHA256Hash(src, dst string) ([]byte, []byte) {

	srcFile, err := os.Open(src)
	defer srcFile.Close()

	if err != nil {
		log.Fatal("FAILED TO COPY FILE:", err)
	}

	dstFile, err := os.Open(dst)
	defer dstFile.Close()

	if err != nil {
		log.Fatal("FAILED TO COPY FILE:", err)
	}

	srcBytes, err := io.ReadAll(srcFile)

	if err != nil {
		log.Fatal("FAILED TO READ FILE:", err)
	}

	dstBytes, err := io.ReadAll(dstFile)

	if err != nil {
		log.Fatal("FAILED TO READ FILE:", err)
	}

	srcHasher := sha256.New()
	_, err = srcHasher.Write(srcBytes)
	if err != nil {
		log.Fatal("FAILED TO CREATE SRCHASH:", err)
	}

	dstHasher := sha256.New()
	_, err = srcHasher.Write(dstBytes)
	if err != nil {
		log.Fatal("FAILED TO CREATE DSTHASH:", err)
	}

	return srcHasher.Sum(nil), dstHasher.Sum(nil)
}

func compareHash(x, y []byte) bool {

	if len(x) != len(y) {
		return false
	}

	return subtle.ConstantTimeCompare(x, y) == 1
}

func CopyFile(src, dst string) {

	srcFile, err := os.Open(src)

	if err != nil {
		log.Fatal("FAILED TO COPY FILE:", err)
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		log.Fatal("FAILED TO CREATE FILE:", err)
	}

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatal("FAILED TO COPY FILE:", err)
	}

	err = dstFile.Sync()
	if err != nil {
		log.Fatal("FAILED TO SYNC FILE:", err)
	}

	_, err = os.Stat(dst)
	if err != nil {

		log.Fatal("FAILED TO STAT FILE:", err)
		os.Exit(1)
	}

	srcFile.Close()
	dstFile.Close()

}
func main() {

	var src string = "/media/nfs/SH3.iso"
	var dst string = "/srv/SH3.iso"

	if _, err := os.Stat(dst); err != nil {
		if compareHash(CreateSHA256Hash(src, dst)) {
			os.Exit(0)
		}
	}

	CopyFile(src, dst)

	if compareHash(CreateSHA256Hash(src, dst)) {
		WriteSyslog("success")
	} else {
		WriteSyslog("fail")
	}
}
