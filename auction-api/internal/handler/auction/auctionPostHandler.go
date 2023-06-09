package auction

import (
	"net/http"

	"auction_server/auction-api/internal/logic/auction"
	"auction_server/auction-api/internal/svc"
	"auction_server/auction-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AuctionPostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AuctionPostReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auction.NewAuctionPostLogic(r.Context(), svcCtx)
		resp, err := l.AuctionPost(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
