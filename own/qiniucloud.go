package own

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type QiNiuCloud struct {
	Bucket    string `json:"bucket"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Address   string `json:"address"`
}

var (
	bucket    string
	accessKey string
	secretKey string
	address   string
)

func InitQiNiuCloud(cloud *QiNiuCloud) {
	bucket = cloud.Bucket
	accessKey = cloud.AccessKey
	secretKey = cloud.SecretKey
	address = cloud.Address
}

func UploadFile(file *multipart.FileHeader) (string, error) {
	fmt.Println(file.Filename)
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuadongZheJiang2
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}

	f, _ := file.Open()
	data, _ := io.ReadAll(f)

	err := formUploader.Put(context.Background(), &ret, upToken, file.Filename, bytes.NewReader(data), file.Size, &putExtra)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	path := "http://" + address + "/" + ret.Key
	return path, nil
}
