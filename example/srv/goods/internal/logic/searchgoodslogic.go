package logic

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/ryon-wen/own-utils/own"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"2106A-zg6/srv/goods/internal/svc"
	"2106A-zg6/srv/goods/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchGoodsLogic {
	return &SearchGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func checkSearchGoods(in *pb.SearchGoodsReq) string {
	if in == nil {
		return "invalid search goods request"
	}
	if in.Content == "" {
		return "search content is required"
	}
	if in.Offset < 0 {
		return "search offset is negative"
	}
	if in.Limit <= 0 {
		return "search limit not be negative"
	}
	return ""
}

func (l *SearchGoodsLogic) SearchGoods(in *pb.SearchGoodsReq) (*pb.SearchGoodsResp, error) {
	if errStr := checkSearchGoods(in); errStr != "" {
		return nil, status.Error(codes.InvalidArgument, errStr)
	}
	var sReq SearchReq
	sReq.From = int(in.Offset)
	sReq.Size = int(in.Limit)
	sReq.Query.Match.Name = in.Content
	sReq.Highlight.Fields.Name.PreTags = "<span style='color:red;'>"
	sReq.Highlight.Fields.Name.PostTags = "</span>"
	str, err := own.EsSearch(sReq)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if strings.Contains(str, "error") {
		return nil, status.Error(codes.Internal, str)
	}
	var sResp SearchResp
	_ = json.Unmarshal([]byte(str), &sResp)
	if sResp.Hits.Total.Value == 0 {
		return &pb.SearchGoodsResp{Goods: nil}, nil
	}
	var goods []*pb.Good
	for _, v := range sResp.Hits.Hits {
		goods = append(goods, &pb.Good{
			ID:    int64(v.Source.Id),
			Name:  v.Source.Name,
			Price: v.Source.Price,
			Stock: int64(v.Source.Stock),
		})
	}
	return &pb.SearchGoodsResp{
		Goods: goods,
	}, nil
}

type SearchReq struct {
	Query struct {
		Match struct {
			Name string `json:"name"`
		} `json:"match"`
	} `json:"query"`
	From      int `json:"from"`
	Size      int `json:"size"`
	Highlight struct {
		Fields struct {
			Name struct {
				PreTags  string `json:"pre_tags"`
				PostTags string `json:"post_tags"`
			} `json:"name"`
		} `json:"fields"`
	} `json:"highlight"`
}

type SearchResp struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string  `json:"_index"`
			Type   string  `json:"_type"`
			Id     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				Stock     int         `json:"stock"`
				Price     float64     `json:"price"`
				CreatedAt time.Time   `json:"created_at"`
				DeletedAt interface{} `json:"deleted_at"`
				Name      string      `json:"name"`
				Version   string      `json:"@version"`
				Timestamp time.Time   `json:"@timestamp"`
				Id        int         `json:"id"`
				UpdatedAt time.Time   `json:"updated_at"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
