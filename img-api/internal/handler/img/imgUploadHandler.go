package img

import (
	"net/http"

	"auction_server/img-api/internal/logic/img"
	"auction_server/img-api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ImgUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := img.NewImgUploadLogic(r.Context(), svcCtx)
		resp, err := l.ImgUpload(r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
