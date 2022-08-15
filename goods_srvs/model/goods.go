package model

// Category 商品类目
type Category struct {
	BaseModel
	Name             string      `gorm:"type:varchar(20);not null  comment '类目名'" json:"name"`
	ParentCategoryID int32       `json:"parent"`
	ParentCategory   *Category   `json:"-"`
	SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID;references:ID" json:"sub_category"`
	Level            int32       `gorm:"type:int;not null;default:1" json:"level"`
	IsTab            bool        `gorm:"default:false;not null  comment '是否展示在tab栏'" json:"is_tab"`
}

// Brands 品牌
type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null  comment '品牌名称'"`
	Logo string `gorm:"type:varchar(200);default:'';not null comment 'LOGO'"`
}

// GoodsCategoryBrand 商品分类与品牌的关系
type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands     Brands
}

func (GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}

// Banner 轮播图
type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200);not null comment '图片链接'"`
	Url   string `gorm:"type:varchar(200);not null comment '商品详情页'"`
	Index int32  `gorm:"type:int;default:1;not null"`
}

type Goods struct {
	BaseModel
	// 外键
	CategoryID int32 `gorm:"type:int;not null"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;not null"`
	Brands     Brands
	// 状态信息
	OnSale   bool `gorm:"default:false;not null comment '是否上架'"`
	ShipFree bool `gorm:"default:false;not null comment '是否免运费'"`
	IsNew    bool `gorm:"default:false;not null comment '是否新品'"`
	IsHot    bool `gorm:"default:false;not null comment '是否热卖商品'"`

	Name            string   `gorm:"type:varchar(50);not null comment '商品名称'"`
	GoodsSn         string   `gorm:"type:varchar(50);not null comment '商品编号'"`
	ClickNum        int32    `gorm:"type:int;default:0;not null comment '点击量'"`
	SoldNum         int32    `gorm:"type:int;default:0;not null comment '销量'"`
	FavNum          int32    `gorm:"type:int;default:0;not null comment '收藏量'"`
	MarketPrice     float32  `gorm:"not null comment '促销价'"`
	ShopPrice       float32  `gorm:"not null comment '售价'"`
	GoodsBrief      string   `gorm:"type:varchar(100);not null comment '商品简介'"`
	Images          GormList `gorm:"type:varchar(1000);not null comment '商品图'"`
	DescImages      GormList `gorm:"type:varchar(1000);not null comment '商品详情图'"`
	GoodsFrontImage string   `gorm:"type:varchar(200);not null comment '商品封面图'"`
}
