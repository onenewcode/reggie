package obs

import (
	"github.com/minio/minio-go"
	"log"
)

var (
	minioClient *minio.Client
	OBS         OBSSave
)

const (
	endpoint        = "121.37.143.160:9000" //兼容对象存储服务endpoint,也可以设置自己的服务器地址
	accessKeyID     = "minioadmin"          // 对象存储的Access key
	secretAccessKey = "minioadmin"          /// 对象存储的Secret key
	ssl             = false                 //true代表使用HTTPS
	bucketName      = "sky-take-out"        // 设置同名称
)

func init() {
	// 初使化minio client对象。
	mc, err := minio.New(endpoint, accessKeyID, secretAccessKey, ssl)
	if err != nil {
		log.Println(err)
	} else {
		minioClient = mc
	}
	OBS = &MyMinio{}
}
