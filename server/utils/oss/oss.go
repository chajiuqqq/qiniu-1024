package oss

import (
	"bytes"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
	"os"
)

var ossClient *OssClient

type OssClient struct {
	mac    *qbox.Mac
	bucket string
	conf   *storage.Config
	domain string
}
type Config struct {
	AK     string
	SK     string
	Bucket string
	domain string
}

func NewOssClient(conf *Config) *OssClient {
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuadongZheJiang2
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = true

	return &OssClient{
		mac:    qbox.NewMac(conf.AK, conf.SK),
		bucket: conf.Bucket,
		conf:   &cfg,
		domain: conf.domain,
	}
}
func (c *OssClient) FileUpload(localFile string, name string) (string, error) {
	bucket := c.bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	upToken := putPolicy.UploadToken(c.mac)
	resumeUploader := storage.NewResumeUploaderV2(c.conf)
	ret := storage.PutRet{}
	recorder, err := storage.NewFileRecorder(os.TempDir())
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("create file recorder failed, err:%w", err)
	}
	putExtra := storage.RputV2Extra{
		Recorder: recorder,
	}
	err = resumeUploader.PutFile(context.Background(), &ret, upToken, name, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("upload file to oss failed, err:%w", err)
	}
	return ret.Key, nil
}

func (c *OssClient) ByteUpload(data []byte, name string) (string, error) {
	bucket := c.bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	upToken := putPolicy.UploadToken(c.mac)
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(c.conf)
	ret := storage.PutRet{}
	dataLen := int64(len(data))

	err := formUploader.Put(context.Background(), &ret, upToken, name, bytes.NewReader(data), dataLen, nil)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("upload bytes to oss failed, err:%w", err)
	}
	return ret.Key, nil
}

func (c *OssClient) ResourceUrl(key string) string {
	publicAccessURL := storage.MakePublicURL(c.domain, key)
	return publicAccessURL
}
func (c *OssClient) GetResource(key string) ([]byte, error) {
	bm := storage.NewBucketManager(c.mac, &storage.Config{})

	// err 和 resp 可能同时有值，当 err 有值时，下载是失败的，此时如果 resp 也有值可以通过 resp 获取响应状态码等其他信息
	resp, err := bm.Get(c.bucket, key, &storage.GetObjectInput{
		DownloadDomains: []string{
			c.domain, // 当前仅支持配置一个，不配置时，使用源站域名进行下载，会对下载的 URL 进行签名
		},
		PresignUrl: true, // 下载 URL 是否进行签名，源站域名或者私有空间需要配置为 true
	})
	if err != nil || resp == nil {
		return nil, fmt.Errorf("download file from oss failed, err:%w", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("download file from oss failed, err:%w", err)
	}
	return body, nil
}
