package initialition

import (
	"context"
	"fmt"
	"log"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/initialition"
	"time"
)

type Integration struct {
	initialition.UnimplementedIntegrateServer
}

func (i *Integration) Integration(ctx context.Context, in *initialition.IntegrationRequest) (*initialition.IntegrationResponse, error) {
	initialitions := model.UserIntegralog{
		IntegralType:  in.IntegralType,
		Integral:      in.Integral,
		Bak:           in.Bak,
		OperationTime: time.Time{},
		CreateTime:    time.Time{},
	}
	err := initialitions.UserIntegralog()
	if err != nil {
		log.Fatalf("添加积分失败%v", err)
	}
	fmt.Println("成功添加200积分")
	err = initialitions.UserIntegralogCx(in.Integral)
	if err != nil {
		log.Fatalf("查询积分失败:%v", err)
	}
	fmt.Printf("当前积分:%d\n", err)

	return &initialition.IntegrationResponse{Bigint: int64(initialitions.UserId)}, nil
}
