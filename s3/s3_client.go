package s3

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/concourse/s3-resource"
	"github.com/pivotal-cf/go-pivnet/logger"
)

type Client struct {
	bucket          string

	logger logger.Logger
	stderr io.Writer

	s3client s3resource.S3Client
}

type NewClientConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
	RegionName      string
	Bucket          string

	Logger            logger.Logger
	Stderr            io.Writer
	SkipSSLValidation bool
}

func NewClient(config NewClientConfig) *Client {
	endpoint := ""
	disableSSL := config.SkipSSLValidation

	awsConfig := s3resource.NewAwsConfig(
		config.AccessKeyID,
		config.SecretAccessKey,
		config.SessionToken,
		config.RegionName,
		endpoint,
		disableSSL,
		config.SkipSSLValidation,
	)

	s3client := s3resource.NewS3Client(
		config.Stderr,
		awsConfig,
		false,
	)

	return &Client{
		bucket:          config.Bucket,
		stderr:          config.Stderr,
		logger:          config.Logger,
		s3client:        s3client,
	}
}

func (c Client) Upload(fileGlob string, to string, sourcesDir string) error {
	matches, err := filepath.Glob(filepath.Join(sourcesDir, fileGlob))

	if err != nil {
		return err
	}

	if len(matches) == 0 {
		return fmt.Errorf("no matches found for pattern: '%s'", fileGlob)
	}

	if len(matches) > 1 {
		return fmt.Errorf(
			"more than one match found for pattern: '%s': %v",
			fileGlob,
			matches,
		)
	}

	localPath := matches[0]
	remotePath := filepath.Join(to, filepath.Base(localPath))

	options := s3resource.NewUploadFileOptions()

	c.logger.Info(fmt.Sprintf(
		"Uploading %s to s3://%s/%s",
		localPath,
		c.bucket,
		remotePath,
	))

	_, err = c.s3client.UploadFile(
		c.bucket,
		remotePath,
		localPath,
		options,
	)
	if err != nil {
		return err
	}

	// the s3client does not append a new-line to its output
	fmt.Fprintln(c.stderr)

	c.logger.Info(fmt.Sprintf(
		"Successfully uploaded '%s' to 's3://%s/%s'",
		localPath,
		c.bucket,
		remotePath,
	))

	return nil
}
