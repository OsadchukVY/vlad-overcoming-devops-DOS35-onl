package main

import (
	"io"
	"log"
	"log/syslog"
	"os"
	"os/exec"
	"strings"
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

func CalculateHashes(src, dst string) (string, string) {

	src_out, err := exec.Command("sha256sum", src).Output()

	if err != nil {
		log.Fatal(err)
	}

	src_hash := strings.Split(string(src_out), " ")

	dst_out, err := exec.Command("sha256sum", dst).Output()

	if err != nil {
		log.Fatal(err)
	}

	out_hash := strings.Split(string(dst_out), " ")
	return src_hash[0], out_hash[0]
}

func CompareHashes(src, dst string) bool {
	src_hash, out_hash := CalculateHashes(src, dst)
	if strings.Compare(src_hash, out_hash) != 0 {
		return false
	}

	return true
}

func CopyFile(src, dst string) bool {

	srcFile, err := os.Open(src)

	if err != nil {
		WriteSyslog("no file in mounted directory")
		return false
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		WriteSyslog("can't create file in destination directory")
		return false
	}

	defer srcFile.Close()
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		WriteSyslog("can't copy file in destination directory")
		return false
	}

	err = dstFile.Sync()
	if err != nil {
		WriteSyslog("can't sync file in destination directory")
		return false

	}

	_, err = os.Stat(dst)
	if err != nil {
		WriteSyslog("can't stat file in destination directory")
		return false
	}

	return true
}

func main() {

	var src string = "/media/nfs/SH3.iso"
	var dst string = "/srv/SH3.iso"

	if _, err := os.Stat(dst); err == nil {
		if CompareHashes(src, dst) {
			os.Exit(0)
		}
	}

	if !CopyFile(src, dst) {
		os.Exit(0)
	}

	if CompareHashes(src, dst) {
		WriteSyslog("success")
	} else {
		WriteSyslog("fail")
	}
}
