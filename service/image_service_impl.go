package service

import (
	"boilerplate/model"
	"boilerplate/repository"
	"fmt"
	"io"

	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/klauspost/compress/zip"
)

type imageService struct {
	repo repository.ImageRepository
}

func NewImageService(repo repository.ImageRepository) ImageService {
	return &imageService{repo: repo}
}

func (is *imageService) UploadImage(file multipart.File, header *multipart.FileHeader) error {
	// Unzip file zip
	zipReader, err := zip.NewReader(file, header.Size)
	if err != nil {
		return err
	}

	// Iterate through the files in the archive
	for _, f := range zipReader.File {
		name := f.Name
		if !isValidFileFormat(name) {
			return fmt.Errorf("invalid file format %s, only jpeg, jpg, and png are allowed", name)
		}

		// Open the file inside the archive
		rc, err := f.Open()
		if err != nil {
			return err
		}

		// Get the file size
		size := f.FileInfo().Size()

		// Generate a unique url for the image
		url := fmt.Sprintf("http://localhost:3000/images/%s", name)

		// Save the image to the server
		imagePath := filepath.Join("public/images", name)
		file, err := os.Create(imagePath)
		if err != nil {
			return err
		}
		_, err = io.Copy(file, rc)
		if err != nil {
			return err
		}
		file.Close()

		// Save the image information to the database
		image := &model.Image{
			Name: name,
			Size: size,
			URL:  url,
		}
		err = is.repo.Save(image)
		if err != nil {
			return err
		}

		// Close the file inside the archive
		rc.Close()
	}

	return nil
}

func isValidFileFormat(name string) bool {
	// check file format
	extension := filepath.Ext(name)
	if extension != ".jpeg" && extension != ".jpg" && extension != ".png" {
		return false
	}
	return true
}

func (is *imageService) GetImageByName(name string) (*model.Image, error) {
	return is.repo.GetImageByName(name)
}
