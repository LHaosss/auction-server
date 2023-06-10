package logic

import (
	"context"
	"database/sql"

	"auction_server/auction-rpc/internal/svc"
	"auction_server/auction-rpc/model/auction_model"
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
	auctionInfo, err := l.svcCtx.AuctionModel.FindOneByXid(l.ctx, sql.NullString{
		String: in.GetAuctionXid(),
		Valid:  true,
	})
	if err != nil {
		return nil, err
	}

	// 更新Auction信息
	data := &auction_model.Auction{
		Id: auctionInfo.Id,
		Xid: sql.NullString{
			String: in.GetAuctionXid(),
			Valid:  true,
		},
		AuctionName: auctionInfo.AuctionName,
		OfferUserXid: sql.NullString{
			String: in.GetOfferUserXid(),
			Valid:  true,
		},
		CurrentPrice: sql.NullInt64{
			Int64: int64(in.GetOfferPrice()),
			Valid: true,
		},
		PostUserXid:   auctionInfo.PostUserXid,
		Type:          auctionInfo.Type,
		BasePrice:     auctionInfo.BasePrice,
		AuctionimgUrl: auctionInfo.AuctionimgUrl,
	}

	err = l.svcCtx.AuctionModel.Update(l.ctx, data)
	if err != nil {
		return nil, err
	}

	return &pb.AuctionOfferResp{
		Id:           int32(data.Id),
		Xid:          data.Xid.String,
		AuctionName:  data.AuctionName,
		PostUserXid:  data.PostUserXid,
		OfferUserXid: data.OfferUserXid.String,
		BasePrice:    int32(data.BasePrice.Int64),
		CurrentPrice: int32(data.CurrentPrice.Int64),
		Type: func(typeZH string) pb.AuctionType {
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
