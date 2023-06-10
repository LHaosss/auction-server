package logic

import (
	"context"

	"auction_server/auctionInfoManager-rpc/internal/svc"
	"auction_server/auctionInfoManager-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserOfferAuctionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserOfferAuctionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserOfferAuctionLogic {
	return &UserOfferAuctionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserOfferAuctionLogic) UserOfferAuction(in *pb.UserOfferAuctionReq) (*pb.UserOfferAuctionResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UserOfferAuctionResp{}, nil
}
