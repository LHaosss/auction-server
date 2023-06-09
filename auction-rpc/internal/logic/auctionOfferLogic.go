package logic

import (
	"context"

	"auction_server/auction-rpc/internal/svc"
	"auction_server/auction-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuctionOfferLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuctionOfferLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuctionOfferLogic {
	return &AuctionOfferLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuctionOfferLogic) AuctionOffer(in *pb.AuctionOfferReq) (*pb.AuctionOfferResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AuctionOfferResp{}, nil
}
