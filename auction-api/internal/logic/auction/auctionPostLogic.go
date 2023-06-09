package auction

import (
	"context"

	"auction_server/auction-api/internal/svc"
	"auction_server/auction-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuctionPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuctionPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuctionPostLogic {
	return &AuctionPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuctionPostLogic) AuctionPost(req *types.AuctionPostReq) (resp *types.AuctionPostResp, err error) {
	// todo: add your logic here and delete this line

	return
}
