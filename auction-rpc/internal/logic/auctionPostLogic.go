package logic

import (
	"context"
	"database/sql"

	"auction_server/auction-rpc/internal/svc"
	"auction_server/auction-rpc/model/auction_model"
	"auction_server/auction-rpc/pb"

	"github.com/rs/xid"
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
	data := &auction_model.Auction{
		Xid: sql.NullString{
			String: xid.New().String(),
			Valid:  true,
		},
		AuctionName: in.GetAuctionName(),
		PostUserXid: in.GetPostUserXid(),
		Type: func(t pb.AuctionType) string {
			switch t {
			case pb.AuctionType_ArtWork:
				return "艺术品"
			case pb.AuctionType_ArtsCrafts:
				return "文玩"
			case pb.AuctionType_Jewelry:
				return "珠宝"
			case pb.AuctionType_Home:
				return "房产"
			case pb.AuctionType_Car:
				return "汽车"
			case pb.AuctionType_SmokeDrink:
				return "烟酒"
			case pb.AuctionType_Clothes:
				return "服饰"
			default:
				return "其他"
			}
		}(in.GetType()),
		BasePrice: sql.NullInt64{
			Int64: int64(in.GetBasePrice()),
			Valid: true,
		},
		AuctionimgUrl: in.GetAuctionImgUrl(),
	}
	_, err := l.svcCtx.AuctionModel.Insert(l.ctx, data)
	if err != nil {
		return nil, err
	}

	return &pb.AuctionPostResp{
		Id:          int32(data.Id),
		Xid:         data.Xid.String,
		AuctionName: data.AuctionName,
		PostUserXid: data.PostUserXid,
		BasePrice:   int32(data.BasePrice.Int64),
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
		}(data.Type),
		AuctionImgUrl: data.AuctionimgUrl,
	}, nil
}
