package storage

import "github.com/aws/aws-sdk-go/service/s3"

// S3Mock struct
type S3Mock struct {
	Error error
}

// PutObject ...
func (s *S3Mock) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	if s.Error != nil {
		return nil, s.Error
	}
	return &s3.PutObjectOutput{}, nil
}

// GetObject ...
func (s *S3Mock) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	if s.Error != nil {
		return nil, s.Error
	}
	return &s3.GetObjectOutput{}, nil
}

// DeleteObject ...
func (s *S3Mock) DeleteObject(input *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	if s.Error != nil {
		return nil, s.Error
	}
	return &s3.DeleteObjectOutput{}, nil
}
