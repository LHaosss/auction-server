package img

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"auction_server/img-api/internal/svc"
	"auction_server/img-api/internal/types"

	"github.com/rs/xid"
	"github.com/zeromicro/go-zero/core/logx"
)

type ImgUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImgUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImgUploadLogic {
	return &ImgUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

const maxFileSize = 10 << 20 // 10MB

func (l *ImgUploadLogic) ImgUpload(r *http.Request) (resp *types.ImgUploadResp, err error) {
	_ = r.ParseMultipartForm(maxFileSize)
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("图片参数出错")
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	names := strings.Split(handler.Filename, ".")
	imgName := xid.New().String() + "." + names[len(names)-1]
	imgUrl := "http://127.0.0.1:7771/img/v1/img/show?img_name=" + imgName
	tempFile, err := os.Create(path.Join(l.svcCtx.Config.Path, imgName))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer tempFile.Close()
	io.Copy(tempFile, file)
	return &types.ImgUploadResp{
		OK:     0,
		ImgUrl: imgUrl,
	}, nil
}
