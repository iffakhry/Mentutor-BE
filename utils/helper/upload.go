package helper

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"os"
	"path/filepath"

	// "os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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

func UploadFotoProfile(c echo.Context) (string, error) {

	file, fileheader, err := c.Request().FormFile("images")
	if err != nil {
		log.Print(err)
		return "", err
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
			Bucket:      aws.String("mentutor"),                                   // bucket's name
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

func UploadProfileProduct(c echo.Context) (string, error) {

	file, fileheader, err := c.Request().FormFile("product_picture")
	if err != nil {
		log.Print(err)
		return "", err
	}

	randomStr := String(20)

	// godotenv.Load(".env")

	s3Config := &aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials("AKIATMRW76KPVGWL6COQ", "l4bJPSj+OnXuQX1AlGb/TEiyTnfyHhoG5nDS1wKu", ""),
	}
	s3Session := session.New(s3Config)

	uploader := s3manager.NewUploader(s3Session)

	input := &s3manager.UploadInput{
		Bucket:      aws.String("mentutor"),                                   // bucket's name
		Key:         aws.String("product/" + randomStr + "-" + fileheader.Filename), // files destination location
		Body:        file,                                                           // content of the file
		ContentType: aws.String("image/jpg"),                                        // content type
	}
	res, err := uploader.UploadWithContext(context.Background(), input)

	// RETURN URL LOCATION IN AWS
	return res.Location, err
}
