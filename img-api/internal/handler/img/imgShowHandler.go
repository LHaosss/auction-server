package img

import (
	"net/http"

	"auction_server/img-api/internal/logic/img"
	"auction_server/img-api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ImgShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := img.NewImgShowLogic(r.Context(), svcCtx)
		resp, err := l.ImgShow(r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			w.Write(resp)
			httpx.OkJsonCtx(r.Context(), w, nil)
		}
	}
}
