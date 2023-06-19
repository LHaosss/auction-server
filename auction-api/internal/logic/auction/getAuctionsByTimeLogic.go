package auction

import (
	"context"
	"errors"

	"auction_server/auction-api/internal/svc"
	"auction_server/auction-api/internal/types"
	"auction_server/auction-rpc/auctioncenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuctionsByTimeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAuctionsByTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuctionsByTimeLogic {
	return &GetAuctionsByTimeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAuctionsByTimeLogic) GetAuctionsByTime() (resp *types.GetAuctionsByTimeResp, err error) {
	auctionsInfo, err := l.svcCtx.AuctionRpcClient.GetAuctionsByTime(l.ctx, &auctioncenter.GetAuctionsByTimeReq{})
	if err != nil {
		return nil, errors.New("获取拍卖信息表失败")
	}

	return &types.GetAuctionsByTimeResp{
		AuctionsInfo: func() []*types.AuctionInfo {
			results := make([]*types.AuctionInfo, len(auctionsInfo.AuctionInfos))
			for i, v := range auctionsInfo.AuctionInfos {
				results[i] = &types.AuctionInfo{
					AuctionName:   v.AuctionName,
					AuctionXid:    v.AuctionXid,
					AuctionImgUrl: v.AuctionImgUrl,
					PostUserXid:   v.PostUserXid,
					OfferUserXid:  v.OfferUserXid,
					BasePrice:     int(v.BasePrice),
					CurrentPrice:  int(v.CurrentPrice),
					Type:          v.Type.String(),
				}
			}
			return results
		}(),
	}, nil
}
