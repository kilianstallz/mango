package download

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func DownloadGoVersion(version string) error {
	// get os and arch
	goos := runtime.GOOS
	fmt.Println(goos)
	arch := runtime.GOARCH
	fmt.Println(arch)

	// get the download url
	fpath := fmt.Sprintf("https://golang.org/dl/%s.%s-%s.tar.gz", version, goos, arch)

	// download the file
	out, err := os.Create("go_tmp.tar.gz")
	if err != nil {
		return err
	}

	resp, err := http.Get(fpath)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	// move the file to the right place
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	// extract the file
	err = Untar(path.Join(home, "go", version), "./go_tmp.tar.gz")
	if err != nil {
		return err
	}

	// delete the temp file
	out.Close()
	err = os.Remove("./go_tmp.tar.gz")
	if err != nil {
		return err
	}

	return nil
}

func Untar(dst string, srcFile string) error {
	r, err := os.Open(srcFile)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()

		switch {

		// if no more files are found return
		case err == io.EOF:
			return nil

		// return any other error
		case err != nil:
			return err

		// if the header is nil, just skip it (not sure how this happens)
		case header == nil:
			continue
		}

		// the target location where the dir/file should be created
		target := filepath.Join(dst, header.Name)

		// the following switch could also be done using fi.Mode(), not sure if there
		// a benefit of using one vs. the other.
		// fi := header.FileInfo()

		// check the file type
		switch header.Typeflag {

		// if its a dir and it doesn't exist create it
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}

		// if it's a file create it
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			// copy over contents
			if _, err := io.Copy(f, tr); err != nil {
				return err
			}

			// manually close here after each file operation; defering would cause each file close
			// to wait until all operations have completed.
			f.Close()
		}
	}
}
