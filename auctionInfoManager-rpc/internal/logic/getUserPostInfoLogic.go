package logic

import (
	"context"

	"auction_server/auctionInfoManager-rpc/internal/svc"
	"auction_server/auctionInfoManager-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPostInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserPostInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPostInfoLogic {
	return &GetUserPostInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserPostInfoLogic) GetUserPostInfo(in *pb.GetUserPostInfoReq) (*pb.GetUserPostInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserPostInfoResp{}, nil
}
