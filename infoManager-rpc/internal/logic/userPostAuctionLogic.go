package logic

import (
	"context"

	"auction_server/infoManager-rpc/internal/svc"
	"auction_server/infoManager-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPostAuctionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserPostAuctionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPostAuctionLogic {
	return &UserPostAuctionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserPostAuctionLogic) UserPostAuction(in *pb.UserPostAuctionReq) (*pb.UserPostAuctionResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UserPostAuctionResp{}, nil
}
