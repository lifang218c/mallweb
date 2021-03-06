package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/king-tu/mallweb/app/common/errors"
	"github.com/king-tu/mallweb/app/common/middlewares"
	"github.com/king-tu/mallweb/app/common/utils"
	"github.com/king-tu/mallweb/app/global"
	"github.com/king-tu/mallweb/app/services/goods/proto/goods"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// 获取商品首页商品分类列表、轮播图、促销商品
func GetFreshGoodsIndex(c *gin.Context) {
	ctx, ok := middlewares.ContextWithSpan(c)
	if ok == false {
		global.Logger.Bg().Debug("get context err")
	}

	var req goods.FreshGoodsIndexRequest
	resp, err := goodsServiceClient.FreshGoodsIndex(ctx, &req)
	if err != nil {
		e := errors.ConvertFrom(err)
		global.Logger.Bg().Error(false, "FreshGoodsIndex 服务调用出错", zap.String("error", e.Detail))
		utils.AbortWithError(c, e)
		return
	}

	// Return the normal response
	c.JSON(http.StatusOK, resp)
}

// 获取商品详情
func GetGoodsDetail(c *gin.Context) {
	ctx, ok := middlewares.ContextWithSpan(c)
	if ok == false {
		global.Logger.Bg().Debug("get context err")
	}

	id := c.Param("id")
	global.Logger.Bg().Debug("GetGoodsDetail Param", zap.String("id", id))
	goodsId, _ := strconv.Atoi(id)
	req := goods.GoodsDetailRequest{GoodsSkuId: int32(goodsId)}

	resp, err := goodsServiceClient.GetGoodsDetail(ctx, &req)
	if err != nil {
		e := errors.ConvertFrom(err)
		global.Logger.Bg().Error(false, "GetGoodsDetail 服务调用出错", zap.String("error", e.Detail))
		utils.AbortWithError(c, e)
		return
	}

	// Return the normal response
	c.JSON(http.StatusOK, resp)
}

// 搜索商品
func SearchGoods(c *gin.Context) {
	goodsName := c.Query("goodsName")
	sort := c.Query("sort")

	req := goods.SearchGoodsRequest{
		GoodsName: goodsName,
		Sort:      sort,
	}

	ctx, ok := middlewares.ContextWithSpan(c)
	if ok == false {
		global.Logger.Bg().Debug("get context err")
	}

	resp, err := goodsServiceClient.SearchGoods(ctx, &req)
	if err != nil {
		e := errors.ConvertFrom(err)
		global.Logger.Bg().Error(false, "SearchGoods 服务调用出错", zap.String("error", e.Detail))
		utils.AbortWithError(c, e)
		return
	}

	// Return the normal response
	c.JSON(http.StatusOK, resp)
}
