package helper

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"mime/multipart"
	"strings"

	"os"
	"path/filepath"

	// "os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
)

// CREATE RANDOM STRING

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func autoGenerate(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return autoGenerate(length, charset)
}

// UPLOAD FOTO PROFILE TO AWS S3

func UploadFotoProfile(file multipart.File, fileheader *multipart.FileHeader) (string, error) {

	size := fileheader.Size
	if size > 5*1024*1024 {
		log.Print("INI ERROR  size")
		return "", errors.New("file size is too large")
	}

	fileExtension := filepath.Ext(fileheader.Filename)

	if fileExtension == ".jpeg" || fileExtension == ".jpg" || fileExtension == ".png" {
		randomStr := String(20)

		godotenv.Load("config.env")

		s3Config := &aws.Config{
			Region:      aws.String("ap-southeast-1"),
			Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_KEY"), os.Getenv("AWS_SECRET"), ""),
		}

		s3Session := session.New(s3Config)

		uploader := s3manager.NewUploader(s3Session)

		input := &s3manager.UploadInput{
			Bucket:      aws.String("mentutor"),                                         // bucket's name
			Key:         aws.String("profile/" + randomStr + "-" + fileheader.Filename), // files destination location
			Body:        file,                                                           // content of the file
			ContentType: aws.String("image/jpg"),                                        // content type
		}
		res, err := uploader.UploadWithContext(context.Background(), input)

		// RETURN URL LOCATION IN AWS
		return res.Location, err
	}
	return "", errors.New("file not an image")
}

func UploadFileTugas(file multipart.File, fileheader *multipart.FileHeader) (string, error) {

	size := fileheader.Size
	if size > 5*1024*1024 {
		log.Print("INI ERROR  size")
		return "", errors.New("file size is too large")
	}

	fileExtension := filepath.Ext(fileheader.Filename)

	if fileExtension == ".xlsx" || fileExtension == ".docx" || fileExtension == ".pdf" || fileExtension == ".pptx" {
		randomStr := String(20)

		godotenv.Load("config.env")

		s3Config := &aws.Config{
			Region:      aws.String("ap-southeast-1"),
			Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_KEY"), os.Getenv("AWS_SECRET"), ""),
		}

		s3Session := session.New(s3Config)

		uploader := s3manager.NewUploader(s3Session)

		input := &s3manager.UploadInput{
			Bucket: aws.String("mentutor"),                                           // bucket's name
			Key:    aws.String("task-file/" + randomStr + "-" + fileheader.Filename), // files destination location
			Body:   file,                                                             // content of the file                                   // content type
		}
		res, err := uploader.UploadWithContext(context.Background(), input)

		// RETURN URL LOCATION IN AWS
		return res.Location, err
	}
	return "", errors.New("file not an image")
}

func UploadGambarTugas(file multipart.File, fileheader *multipart.FileHeader) (string, error) {

	size := fileheader.Size
	if size > 5*1024*1024 {
		log.Print("INI ERROR  size")
		return "", errors.New("file size is too large")
	}

	fileExtension := filepath.Ext(fileheader.Filename)

	if fileExtension == ".jpeg" || fileExtension == ".jpg" || fileExtension == ".png" {
		randomStr := String(20)

		godotenv.Load("config.env")

		s3Config := &aws.Config{
			Region:      aws.String("ap-southeast-1"),
			Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_KEY"), os.Getenv("AWS_SECRET"), ""),
		}

		s3Session := session.New(s3Config)

		uploader := s3manager.NewUploader(s3Session)

		input := &s3manager.UploadInput{
			Bucket:      aws.String("mentutor"),                                             // bucket's name
			Key:         aws.String("task-images/" + randomStr + "-" + fileheader.Filename), // files destination location
			Body:        file,                                                               // content of the file
			ContentType: aws.String("images/jpg"),                                           // content type
		}
		res, err := uploader.UploadWithContext(context.Background(), input)

		// RETURN URL LOCATION IN AWS
		return res.Location, err
	}
	return "", errors.New("file not an image")
}

func UploadStatusImages(file multipart.File, fileheader *multipart.FileHeader) (string, error) {

	size := fileheader.Size
	if size > 5*1024*1024 {
		log.Print("INI ERROR  size")
		return "", errors.New("file size is too large")
	}

	fileExt := filepath.Ext(fileheader.Filename)
	fileExtension := strings.ToLower(fileExt)
	if fileExtension == ".jpeg" || fileExtension == ".jpg" || fileExtension == ".png" {
		randomStr := String(20)

		godotenv.Load("config.env")

		s3Config := &aws.Config{
			Region:      aws.String("ap-southeast-1"),
			Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_KEY"), os.Getenv("AWS_SECRET"), ""),
		}

		s3Session := session.New(s3Config)

		uploader := s3manager.NewUploader(s3Session)

		input := &s3manager.UploadInput{
			Bucket:      aws.String("mentutor"),                                               // bucket's name
			Key:         aws.String("status-images/" + randomStr + "-" + fileheader.Filename), // files destination location
			Body:        file,                                                                 // content of the file
			ContentType: aws.String("images/jpg"),                                             // content type
		}
		res, err := uploader.UploadWithContext(context.Background(), input)

		// RETURN URL LOCATION IN AWS
		return res.Location, err
	}
	return "", errors.New("file not an image")
}

func UploadFileSubmission(file multipart.File, fileheader *multipart.FileHeader) (string, error) {

	size := fileheader.Size
	if size > 5*1024*1024 {
		log.Print("INI ERROR  size")
		return "", errors.New("file size is too large")
	}

	fileExtension := filepath.Ext(fileheader.Filename)

	if fileExtension == ".pdf" || fileExtension == ".docx" || fileExtension == ".xlsx" || fileExtension == ".pptx" {
		randomStr := String(20)

		godotenv.Load("config.env")

		s3Config := &aws.Config{
			Region:      aws.String("ap-southeast-1"),
			Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_KEY"), os.Getenv("AWS_SECRET"), ""),
		}

		s3Session := session.New(s3Config)

		uploader := s3manager.NewUploader(s3Session)

		input := &s3manager.UploadInput{
			Bucket: aws.String("mentutor"),                                                 // bucket's name
			Key:    aws.String("submission-file/" + randomStr + "-" + fileheader.Filename), // files destination location
			Body:   file,                                                                   // content of the file                                   // content type
		}
		res, err := uploader.UploadWithContext(context.Background(), input)

		// RETURN URL LOCATION IN AWS
		return res.Location, err
	}
	return "", errors.New("file not an image")
}
