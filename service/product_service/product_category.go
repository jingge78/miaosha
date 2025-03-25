package product_service

import (
	"context"
	"fmt"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/product"
)

func (p *ServerProduct) ProductCategory(ctx context.Context, in *product.ProductCategoryRequest) (*product.ProductCategoryResponse, error) {
	var class model.StoreCategory

	// 获取父类 ID 下的分类和其子分类
	cate, err := class.GetCateWithChild(int(in.ParentId))
	if err != nil {
		return nil, fmt.Errorf("分类列表展示失败: %s", err.Error())
	}

	// 将查询到的分类和子分类转换为 gRPC 响应格式
	var slice []*product.ProductCategoryResponse_Category // 注意这里是指针类型
	for _, category := range cate.Children {
		slice = append(slice, &product.ProductCategoryResponse_Category{ // 使用指针
			Id:            category.ID,
			ParentId:      category.Pid,
			CateName:      category.CateName,
			Pic:           category.Pic,
			Sort:          category.Sort,
			IsShow:        int32(category.IsShow),
			AddTime:       category.AddTime,
			SubCategories: convertToGrpcCategories(category.Children), // 递归处理子分类
		})
	}

	return &product.ProductCategoryResponse{
		Categories: slice,
	}, nil
}

// 辅助函数：将 StoreCategory 转换成 gRPC 分类格式
func convertToGrpcCategories(categories []model.StoreCategory) []*product.ProductCategoryResponse_Category {
	var grpcCategories []*product.ProductCategoryResponse_Category // 注意这里是指针类型的切片
	for _, category := range categories {
		grpcCategories = append(grpcCategories, &product.ProductCategoryResponse_Category{ // 使用指针
			Id:            category.ID,
			ParentId:      category.Pid,
			CateName:      category.CateName,
			Pic:           category.Pic,
			Sort:          category.Sort,
			IsShow:        int32(category.IsShow),
			AddTime:       category.AddTime,
			SubCategories: convertToGrpcCategories(category.Children), // 递归处理子分类
		})
	}
	return grpcCategories
}
