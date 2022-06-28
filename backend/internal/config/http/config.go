package http

import (
	"openreplay/backend/internal/config/common"
	"openreplay/backend/internal/config/configurator"
	"openreplay/backend/pkg/env"
	"time"
)

type Config struct {
	common.Config
	HTTPHost          string        `env:"HTTP_HOST,default="`
	HTTPPort          string        `env:"HTTP_PORT,required"`
	HTTPTimeout       time.Duration `env:"HTTP_TIMEOUT,default=60s"`
	TopicRawWeb       string        `env:"TOPIC_RAW_WEB,required"`
	TopicRawIOS       string        `env:"TOPIC_RAW_IOSv"`
	BeaconSizeLimit   int64         `env:"BEACON_SIZE_LIMIT,required"`
	JsonSizeLimit     int64         `env:"JSON_SIZE_LIMIT,default=1e3"`
	FileSizeLimit     int64         `env:"FILE_SIZE_LIMIT,default=1e7"`
	AWSRegion         string        `env:"AWS_REGION,required"`
	S3BucketIOSImages string        `env:"S3_BUCKET_IOS_IMAGES,required"`
	Postgres          string        `env:"POSTGRES_STRING,required"`
	TokenSecret       string        `env:"TOKEN_SECRET,required"`
	UAParserFile      string        `env:"UAPARSER_FILE,required"`
	MaxMinDBFile      string        `env:"MAXMINDDB_FILE,required"`
	WorkerID          uint16
}

func New() *Config {
	cfg := &Config{WorkerID: env.WorkerID()}
	configurator.Process(cfg)
	return cfg
}
