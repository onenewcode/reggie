package obs

import (
	"github.com/google/uuid"
	"github.com/minio/minio-go"
	"log"
	"mime/multipart"
	"path"
	"strings"
	"time"
)

type OBSClient interface {
	UploadImg(fh *multipart.FileHeader) *string
}

/*
实现类
*/
type MyMinio struct {
}

func (*MyMinio) UploadImg(fh *multipart.FileHeader) *string {
	var str strings.Builder
	str.WriteString(time.Now().Format("2006/01/02/"))
	// 生成一个新的UUIDv4
	id := uuid.New()
	str.WriteString(id.String())
	str.WriteString(path.Ext(fh.Filename))
	// 构建文件在Minio的存储路径
	filepath := str.String()
	// 获取文件的读取流
	file_body, _ := fh.Open()
	_, err := minioClient.PutObject(bucketName, filepath, file_body, fh.Size, minio.PutObjectOptions{
		ContentType: fh.Header.Get("Content-Type"),
	})
	// 拼接返回路径
	filepath = "http://" + path.Join(endpoint, bucketName, filepath)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return &filepath

}
