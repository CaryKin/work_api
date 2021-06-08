package request

import "gin-vue-admin/model"

type ShopProductCateSearch struct {
	model.ShopProductCate
	PageInfo
}

type ShopProductCateList struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Cover string `json:"cover"`
}

type ShopProductCateTree struct {
	ID       uint                  `json:"id"`
	Title    string                `json:"title"`
	Cover    string                `json:"cover"`
	Pid      int                   `json:"pid"`
	Children []ShopProductCateTree `json:"children"`
}

func ModelByShopProductCateList(model []model.ShopProductCate) (ShopProductCateLists []ShopProductCateList) {
	for _, cate := range model {
		ShopProductCateLists = append(ShopProductCateLists, ShopProductCateList{
			ID:    cate.ID,
			Title: cate.Title,
			Cover: cate.Cover,
		})
	}

	return ShopProductCateLists
}

func ModelByShopProductCateTree(model []model.ShopProductCate) (ShopProductCateTrees []ShopProductCateTree) {
	for _, cate := range model {
		ShopProductCateTrees = append(ShopProductCateTrees, ShopProductCateTree{
			ID:    cate.ID,
			Title: cate.Title,
			Cover: cate.Cover,
			Pid:   cate.Pid,
		})
	}
	return ShopProductCateTrees
}
