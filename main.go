package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s -i <ktx path> -o <png path>\n", os.Args[0])
		flag.PrintDefaults()
	}

	var (
		ktxPath,
		pngPath string
	)

	flag.StringVar(&ktxPath, "i", "", "Path to input ktx")
	flag.StringVar(&pngPath, "o", "", "Path to destination png")

	flag.Parse()
	if ktxPath == "" {
		fmt.Fprintf(os.Stderr, "Missing <ktx path>\n\n")
		flag.Usage()
		os.Exit(1)
	}

	execPath := filepath.Join(os.TempDir(), execFilename)
	file, err := os.OpenFile(execPath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0777)
	handleError(err)
	_, err = file.Write(execBytes)
	file.Close()
	handleError(err)

	file, err = os.Open(ktxPath)
	handleError(err)
	finfo, err := file.Stat()
	switch {
	case err != nil:
		handleError(err)
	case finfo.IsDir():
		if pngPath == "" {
			pngPath = ktxPath
		}
		for {
			files, err := file.Readdir(10)
			if err != nil {
				break
			}
			for _, f := range files {
				if !f.IsDir() {
					if idx := strings.LastIndex(f.Name(), ".ktx"); idx > 0 {
						kp := filepath.Join(ktxPath, f.Name())
						pp := filepath.Join(pngPath, f.Name()[:idx]+".png")
						convert(execPath, kp, pp)
					}
				}
			}
		}
	default:
		convert(execPath, ktxPath, pngPath)
	}
	file.Close()
	os.Remove(execPath)
}

func convert(execPath, ktxPath, pngPath string) error {
	tmpPvrPath := filepath.Join(os.TempDir(), "tmp.pvr")
	args := []string{"-i", ktxPath, "-o", tmpPvrPath, "-f", "r8g8b8a8", "-d"}
	if pngPath != "" {
		args = append(args, pngPath)
	}
	cmd := exec.Command(execPath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	os.Remove(tmpPvrPath)
	return err
}

func handleError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
