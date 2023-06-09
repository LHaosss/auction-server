package logic

import (
	"context"

	"auction_server/auction-rpc/internal/svc"
	"auction_server/auction-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuctionPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuctionPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuctionPostLogic {
	return &AuctionPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuctionPostLogic) AuctionPost(in *pb.AuctionPostReq) (*pb.AuctionPostResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AuctionPostResp{}, nil
}
