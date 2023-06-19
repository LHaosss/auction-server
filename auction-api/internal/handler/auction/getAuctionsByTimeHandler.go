package auction

import (
	"net/http"

	"auction_server/auction-api/internal/logic/auction"
	"auction_server/auction-api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAuctionsByTimeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := auction.NewGetAuctionsByTimeLogic(r.Context(), svcCtx)
		resp, err := l.GetAuctionsByTime()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
