package client

import (
	"context"
	"google.golang.org/grpc"
	"miaosha-jjl/common/proto/product"
)

type ProductHandler func(ctx context.Context, client product.ProductClient) (interface{}, error)

func ProductClient(ctx context.Context, handler ProductHandler) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	client := product.NewProductClient(dial)
	return handler(ctx, client)
}
func ProductDetail(ctx context.Context, in *product.ProductDetailRequest) (*product.ProductDetailResponse, error) {
	client, err := ProductClient(ctx, func(ctx context.Context, client product.ProductClient) (interface{}, error) {
		detail, err := client.ProductDetail(ctx, in)
		if err != nil {
			return nil, err
		}
		return detail, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*product.ProductDetailResponse), nil
}
func EsAddProduct(ctx context.Context, in *product.EsAddProductRequest) (*product.EsAddProductResponse, error) {
	client, err := ProductClient(ctx, func(ctx context.Context, client product.ProductClient) (interface{}, error) {
		detail, err := client.EsAddProduct(ctx, in)
		if err != nil {
			return nil, err
		}
		return detail, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*product.EsAddProductResponse), nil
}
func GetAllProduct(ctx context.Context, in *product.GetAllProductRequest) (*product.GetAllProductResponse, error) {
	client, err := ProductClient(ctx, func(ctx context.Context, client product.ProductClient) (interface{}, error) {
		detail, err := client.GetAllProduct(ctx, in)
		if err != nil {
			return nil, err
		}
		return detail, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*product.GetAllProductResponse), nil
}

func EsSearchByKeyWord(ctx context.Context, in *product.EsSearchByKeyWordRequest) (*product.EsSearchByKeyWordResponse, error) {
	client, err := ProductClient(ctx, func(ctx context.Context, client product.ProductClient) (interface{}, error) {
		detail, err := client.EsSearchByKeyWord(ctx, in)
		if err != nil {
			return nil, err
		}
		return detail, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*product.EsSearchByKeyWordResponse), nil
}
func PriceFind(ctx context.Context, i *product.PriceFindRequest) (*product.PriceFindResponse, error) {
	client, err := ProductClient(ctx, func(ctx context.Context, client product.ProductClient) (interface{}, error) {
		find, err := client.PriceFind(ctx, i)
		if err != nil {
			return nil, err
		}
		return find, err
	})
	if err != nil {
		return nil, err
	}
	return client.(*product.PriceFindResponse), nil
}

func GetCollectProduct(ctx context.Context, in *product.GetCollectProductRequest) (*product.GetCollectProductResponse, error) {
	client, err := ProductClient(ctx, func(ctx context.Context, client product.ProductClient) (interface{}, error) {
		detail, err := client.GetCollectProduct(ctx, in)
		if err != nil {
			return nil, err
		}
		return detail, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*product.GetCollectProductResponse), err
}
func GetProductCategory(ctx context.Context, i *product.ProductCategoryRequest) (*product.ProductCategoryResponse, error) {
	client, err := ProductClient(ctx, func(ctx context.Context, client product.ProductClient) (interface{}, error) {
		find, err := client.ProductCategory(ctx, i)
		if err != nil {
			return nil, err
		}
		return find, err
	})
	if err != nil {
		return nil, err
	}
	return client.(*product.ProductCategoryResponse), err
}
func WebsiteProductList(ctx context.Context, in *product.WebsiteProductListRequest) (*product.WebsiteProductListResponse, error) {
	client, err := ProductClient(ctx, func(ctx context.Context, client product.ProductClient) (interface{}, error) {
		find, err := client.WebsiteProductList(ctx, in)
		if err != nil {
			return nil, err
		}
		return find, err
	})
	if err != nil {
		return nil, err
	}
	return client.(*product.WebsiteProductListResponse), err
}

func GroupByProductList(ctx context.Context, i *product.GroupByProductListRequest) (*product.GroupByProductListResponse, error) {
	client, err := ProductClient(ctx, func(ctx context.Context, client product.ProductClient) (interface{}, error) {
		find, err := client.GroupByProductList(ctx, i)

		if err != nil {
			return nil, err
		}
		return find, err
	})
	return client.(*product.GroupByProductListResponse), err
}
func ProductSort(ctx context.Context, in *product.ProductSortRequest) (*product.ProductSortResponse, error) {

	client, err := ProductClient(ctx, func(ctx context.Context, client product.ProductClient) (interface{}, error) {
		find, err := client.ProductSort(ctx, in)
		if err != nil {
			return nil, err
		}
		return find, err
	})
	if err != nil {
		return nil, err
	}
	return client.(*product.ProductSortResponse), err
}
func ProductRanking(ctx context.Context, in *product.ProductRankingRequest) (*product.ProductRankingResponse, error) {

	client, err := ProductClient(ctx, func(ctx context.Context, client product.ProductClient) (interface{}, error) {
		find, err := client.ProductRanking(ctx, in)
		if err != nil {
			return nil, err
		}
		return find, err
	})
	if err != nil {
		return nil, err
	}
	return client.(*product.ProductRankingResponse), err
}
func ProductFilter(ctx context.Context, in *product.ProductFilterRequest) (*product.ProductFilterResponse, error) {

	client, err := ProductClient(ctx, func(ctx context.Context, client product.ProductClient) (interface{}, error) {
		find, err := client.ProductFilter(ctx, in)
		if err != nil {
			return nil, err
		}
		return find, err
	})
	if err != nil {
		return nil, err
	}
	return client.(*product.ProductFilterResponse), err
}
func AddSpikeProduct(ctx context.Context, in *product.AddSpikeProductReq) (*product.AddSpikeProductResp, error) {

	client, err := ProductClient(ctx, func(ctx context.Context, client product.ProductClient) (interface{}, error) {
		find, err := client.AddSpikeProduct(ctx, in)
		if err != nil {
			return nil, err
		}
		return find, err
	})
	if err != nil {
		return nil, err
	}
	return client.(*product.AddSpikeProductResp), err
}
