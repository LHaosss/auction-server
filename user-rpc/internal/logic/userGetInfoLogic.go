package logic

import (
	"context"
	"fmt"
	"os"

	"auction_server/user-rpc/internal/svc"
	"auction_server/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserGetInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserGetInfoLogic {
	return &UserGetInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserGetInfoLogic) UserGetInfo(in *pb.UserGetInfoReq) (*pb.UserGetInfoResp, error) {
	fmt.Println("****************************")
	user, err := l.svcCtx.UserModel.FindOneByXid(l.ctx, in.GetXid())
	if err != nil {
		os.WriteFile("er.txt", []byte(err.Error()), os.ModeAppend)
		return nil, err
	}

	return &pb.UserGetInfoResp{
		Id:       int32(user.Id),
		Xid:      user.Xid,
		Username: user.Username.String,
	}, nil
}
