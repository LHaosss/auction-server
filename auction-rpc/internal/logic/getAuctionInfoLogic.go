package logic

import (
	"context"
	"database/sql"

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
	user, err := l.svcCtx.AuctionModel.FindOneByXid(l.ctx, sql.NullString{
		String: in.GetXid(),
		Valid:  true,
	})
	if err != nil {
		return nil, err
	}

	return &pb.GetAuctionInfoResp{
		AuctionInfo: &pb.AuctionInfo{
			AuctionXid:    user.Xid.String,
			AuctionName:   user.AuctionName,
			AuctionImgUrl: user.AuctionimgUrl,
			BasePrice:     int32(user.BasePrice.Int64),
			CurrentPrice:  int32(user.CurrentPrice.Int64),
			PostUserXid:   user.PostUserXid,
			OfferUserXid:  user.OfferUserXid.String,
			Type: func(typeZH string) pb.AuctionType { // 将对请求中的拍品类别进行转换
				switch typeZH {
				case "艺术品":
					return pb.AuctionType_ArtWork
				case "文玩":
					return pb.AuctionType_ArtsCrafts
				case "珠宝":
					return pb.AuctionType_Jewelry
				case "房产":
					return pb.AuctionType_Home
				case "汽车":
					return pb.AuctionType_Car
				case "烟酒":
					return pb.AuctionType_SmokeDrink
				case "服饰":
					return pb.AuctionType_Clothes
				default:
					return pb.AuctionType_Others
				}
			}(user.Type),
		},
	}, nil
}
