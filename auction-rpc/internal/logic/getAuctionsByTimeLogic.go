package logic

import (
	"context"
	"fmt"

	"auction_server/auction-rpc/internal/svc"
	"auction_server/auction-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuctionsByTimeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAuctionsByTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuctionsByTimeLogic {
	return &GetAuctionsByTimeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAuctionsByTimeLogic) GetAuctionsByTime(in *pb.GetAuctionsByTimeReq) (*pb.GetAuctionsByTimeResp, error) {
	auctionMajority, err := l.svcCtx.AuctionModel.FindMajority(l.ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &pb.GetAuctionsByTimeResp{
		AuctionInfos: func() []*pb.AuctionInfo {
			auctionInfos := make([]*pb.AuctionInfo, len(*auctionMajority))
			for i, v := range *auctionMajority {
				fmt.Println(v.Xid.String)
				auctionInfos[i] = &pb.AuctionInfo{
					AuctionXid:    v.Xid.String,
					AuctionName:   v.AuctionName,
					AuctionImgUrl: v.AuctionimgUrl,
					BasePrice:     int32(v.BasePrice.Int64),
					CurrentPrice:  int32(v.CurrentPrice.Int64),
					PostUserXid:   v.PostUserXid,
					OfferUserXid:  v.OfferUserXid.String,
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
					}(v.Type),
				}
			}
			return auctionInfos
		}(),
	}, nil
}
