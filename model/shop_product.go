// 自动生成模板ShopProduct
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type ShopProduct struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:商品标题;type:varchar(255);size:255;"`
      Picture  string `json:"picture" form:"picture" gorm:"column:picture;comment:商品主图;type:varchar(255);size:255;"`
      CateId  int `json:"cateId" form:"cateId" gorm:"column:cate_id;comment:商品分类编号;type:int;size:10;"`
      Sketch  string `json:"sketch" form:"sketch" gorm:"column:sketch;comment:简述;type:varchar(200);size:200;"`
      Intro  string `json:"intro" form:"intro" gorm:"column:intro;comment:商品描述;type:text;"`
      TotalSales  int `json:"totalSales" form:"totalSales" gorm:"column:total_sales;comment:总销量;type:int;size:10;"`
      Price  float64 `json:"price" form:"price" gorm:"column:price;comment:商品价格;type:decimal;size:8,2;"`
      MarketPrice  float64 `json:"marketPrice" form:"marketPrice" gorm:"column:market_price;comment:市场价格;type:decimal;size:8,2;"`
      CostPrice  float64 `json:"costPrice" form:"costPrice" gorm:"column:cost_price;comment:成本价;type:decimal;size:19,2;"`
      Stock  int `json:"stock" form:"stock" gorm:"column:stock;comment:库存量;type:int;size:10;"`
      WarningStock  int `json:"warningStock" form:"warningStock" gorm:"column:warning_stock;comment:库存警告;type:int;size:10;"`
      Covers  string `json:"covers" form:"covers" gorm:"column:covers;comment:幻灯片;type:text;"`
      ProductStatus  *bool `json:"productStatus" form:"productStatus" gorm:"column:product_status;comment:商品状态 0下架，1正常，10违规（禁售）;type:tinyint;"`
      ProductionDate  int `json:"productionDate" form:"productionDate" gorm:"column:production_date;comment:生产日期;type:int;size:10;"`
      ShelfLife  int `json:"shelfLife" form:"shelfLife" gorm:"column:shelf_life;comment:保质期;type:int;size:10;"`
}


func (ShopProduct) TableName() string {
  return "shop_product"
}

