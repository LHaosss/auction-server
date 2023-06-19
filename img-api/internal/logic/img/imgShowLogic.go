package img

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"auction_server/img-api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImgShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImgShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImgShowLogic {
	return &ImgShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImgShowLogic) ImgShow(r *http.Request) (buff []byte, err error) {
	query := r.URL.Query()
	imageName := query["img_name"]
	fmt.Println(imageName[0])
	file, err := os.Open("./images/" + imageName[0])
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("未找到图片")
	}
	defer file.Close()

	buff, err = ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New("读取图片失败")
	}

	return buff, nil

}
