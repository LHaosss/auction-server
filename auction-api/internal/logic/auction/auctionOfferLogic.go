package auction

import (
	"context"

	"auction_server/auction-api/internal/svc"
	"auction_server/auction-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuctionOfferLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuctionOfferLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuctionOfferLogic {
	return &AuctionOfferLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuctionOfferLogic) AuctionOffer(req *types.AuctionOfferReq) (resp *types.AuctionOfferResp, err error) {
	// todo: add your logic here and delete this line

	return
}
