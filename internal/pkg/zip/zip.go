package zip

import (
	"archive/zip"
	"io/ioutil"
	"os"
	"path"
)

func Zip(folder string, savePath string) error {
	outFile, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	w := zip.NewWriter(outFile)

	err = addFiles(w, folder, "")
	if err != nil {
		return err
	}

	return w.Close()
}

func addFiles(w *zip.Writer, basePath, baseInZip string) error {
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			filePath := path.Join(basePath, file.Name())
			dat, err := ioutil.ReadFile(filePath)
			if err != nil {
				return err
			}

			// Add some files to the archive.
			filePath = path.Join(baseInZip, file.Name())
			f, err := w.Create(filePath)
			if err != nil {
				return err
			}
			_, err = f.Write(dat)
			if err != nil {
				return err
			}
		} else if file.IsDir() {
			// Recurse
			newBase := path.Join(basePath, file.Name())
			newBaseInZip := path.Join(baseInZip, file.Name())
			err := addFiles(w, newBase, newBaseInZip)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
