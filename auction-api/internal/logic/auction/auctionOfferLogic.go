package auction

import (
	"context"
	"errors"

	"auction_server/auction-api/internal/svc"
	"auction_server/auction-api/internal/types"
	"auction_server/auction-rpc/auctioncenter"
	"auction_server/auction-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuctionOfferLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuctionOfferLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuctionOfferLogic {
	return &AuctionOfferLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuctionOfferLogic) AuctionOffer(req *types.AuctionOfferReq) (resp *types.AuctionOfferResp, err error) {
	// 验证参数是否合法
	err = validAuctionOfferReq(req)
	if err != nil {
		return nil, err
	}

	// 获取auction信息，判断出价是否大于当前价格
	auctionInfo, err := l.svcCtx.AuctionRpcClient.GetAuctionInfo(l.ctx, &auctioncenter.GetAuctionInfoReq{Xid: req.AuctionXid})
	if err != nil {
		return nil, errors.New("用户查询出错")
	}
	if int(auctionInfo.GetAuctionInfo().GetCurrentPrice()) >= req.OfferPrice {
		return nil, errors.New("出价低于拍品现价")
	}

	// 更新auction信息
	auctionOfferInfo, err := l.svcCtx.AuctionRpcClient.AuctionOffer(l.ctx, &auctioncenter.AuctionOfferReq{
		AuctionXid:   req.AuctionXid,
		OfferUserXid: req.OffeUserXid,
		OfferPrice:   int32(req.OfferPrice),
	})
	if err != nil {
		return nil, errors.New("更新用户信息失败")
	}

	// 更新offer_info信息

	return &types.AuctionOfferResp{
		Id:           int(auctionOfferInfo.GetId()),
		Xid:          auctionOfferInfo.GetXid(),
		AuctionName:  auctionOfferInfo.GetAuctionName(),
		PostUserXid:  auctionOfferInfo.GetPostUserXid(),
		OfferUserXid: auctionOfferInfo.GetOfferUserXid(),
		BasePrice:    int(auctionOfferInfo.GetBasePrice()),
		CurrentPrice: int(auctionOfferInfo.GetCurrentPrice()),
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
		}(auctionOfferInfo.GetType()),
		AuctionImgUrl: auctionOfferInfo.GetAuctionImgUrl(),
	}, nil
}

func validAuctionOfferReq(req *types.AuctionOfferReq) error {
	if req.AuctionXid == "" || req.OfferPrice <= 0 || req.OffeUserXid == "" {
		return errors.New("请求参数出错")
	}
	return nil
}
