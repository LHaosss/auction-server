package logic

import (
	"context"

	"auction_server/infoManager-rpc/internal/svc"
	"auction_server/infoManager-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserOfferInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserOfferInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserOfferInfoLogic {
	return &GetUserOfferInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserOfferInfoLogic) GetUserOfferInfo(in *pb.GetUserOfferInfoReq) (*pb.GetUserOfferInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserOfferInfoResp{}, nil
}
