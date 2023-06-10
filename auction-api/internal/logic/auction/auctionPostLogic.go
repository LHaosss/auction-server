package auction

import (
	"context"
	"errors"
	"fmt"

	"auction_server/auction-api/internal/svc"
	"auction_server/auction-api/internal/types"
	"auction_server/auction-rpc/auctioncenter"
	"auction_server/auction-rpc/pb"
	"auction_server/user-rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuctionPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuctionPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuctionPostLogic {
	return &AuctionPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuctionPostLogic) AuctionPost(req *types.AuctionPostReq) (resp *types.AuctionPostResp, err error) {
	// 提交竞拍逻辑

	// 1、验证参数是否规范
	err = validAuctionPostReqParamers(req)
	if err != nil {
		return nil, err
	}

	// 2、检验post_username是否合法 调用user-rpc的getUserInfo信息，判断是否存在该用户
	_, err = l.svcCtx.UserRpcClient.UserGetInfo(l.ctx, &usercenter.UserGetInfoReq{
		Xid: req.PostUserXid,
	})
	if err != nil {
		fmt.Println("##########", err, "#########")
		return nil, errors.New("提交竞拍的用户不合规")
	}

	// 3、添加auction信息 调用auction-rpc后端服务
	auctionInfo, err := l.svcCtx.AuctionRpcClient.AuctionPost(l.ctx, &auctioncenter.AuctionPostReq{
		AuctionName: req.AuctionName,
		PostUserXid: req.PostUserXid,
		BasePrice:   int32(req.BasePrice),
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
		}(req.Type),
		AuctionImgUrl: req.AuctionImgUrl,
	})
	if err != nil {
		return nil, errors.New("提交竞拍信息失败")
	}

	// 4、添加user_auction信息（向用户提交的竞拍信息中添加信息）

	return &types.AuctionPostResp{
		Id:          int(auctionInfo.GetId()),
		Xid:         auctionInfo.GetXid(),
		AuctionName: auctionInfo.GetAuctionName(),
		PostUserXid: auctionInfo.GetPostUserXid(),
		BasePrice:   int(auctionInfo.GetBasePrice()),
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
		}(auctionInfo.GetType()),
		AuctionImgUrl: auctionInfo.GetAuctionImgUrl(),
	}, nil
}

func validAuctionPostReqParamers(req *types.AuctionPostReq) error {
	const ErrWithParamers = "请求参数出错"
	if req.AuctionName == "" || req.BasePrice == 0 || req.PostUserXid == "" || req.Type == "" || req.AuctionImgUrl == "" {
		return errors.New(ErrWithParamers)
	}
	return nil
}
