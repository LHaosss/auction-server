package logic

import (
	"context"

	"auction_server/auction-rpc/internal/svc"
	"auction_server/auction-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuctionInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAuctionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuctionInfoLogic {
	return &GetAuctionInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAuctionInfoLogic) GetAuctionInfo(in *pb.GetAuctionInfoReq) (*pb.GetAuctionInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetAuctionInfoResp{}, nil
}
