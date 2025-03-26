package product_service

import (
	"context"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/proto/product"
)

// 网站商品展示
// 可以展示一级分类 二级分类
// 按照商品分类进行商品展示，会显示一级分类的商品，并显示二级分类，点击分类可以进入分类查看对应的商品
func (p *ServerProduct) WebsiteProductList(ctx context.Context, in *product.WebsiteProductListRequest) (*product.WebsiteProductListResponse, error) {

	var cate model.StoreCategory
	category, err := cate.PageSizeListByCategory(int64(in.CategoryId), int(in.Page), int(in.PageSize))
	if err != nil {
		return nil, err
	}
	var Price []*product.ProductInfo
	for _, v := range category {
		Price = append(Price, &product.ProductInfo{
			Name:       v.CateName,
			ImageUrl:   v.Pic,
			CategoryId: v.Pid,
		})
	}
	return &product.WebsiteProductListResponse{Products: Price}, nil
}
