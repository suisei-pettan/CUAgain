package dao

import (
	"log"
	"os"
	"path/filepath"
)

func TryGetAssetCache(path string) []byte {
	cacheFile, err := os.ReadFile("./cache" + path)
	if err != nil {
		return nil
	} else {
		return cacheFile
	}
}
func WriteAssetCache(path string, data []byte) error {
	// Ensure the directory structure exists
	cacheDir := "./cache"
	fullPath := filepath.Join(cacheDir, path)
	dir := filepath.Dir(fullPath)

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Println("error creating directories: ", err)
		return err
	}

	// Write the file
	err = os.WriteFile(fullPath, data, 0644)
	if err != nil {
		log.Println("write cache error: ", err)
		return err
	}

	return nil
}
