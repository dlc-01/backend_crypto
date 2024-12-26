package controller

import (
	"net/http"

	"github.com/dlc-01/BackendCrypto/internal/model"
	"github.com/dlc-01/BackendCrypto/internal/service"
	"github.com/gin-gonic/gin"
)

type CryptoController struct {
	service service.ICryptoService
}

func NewCryptoController(svc service.ICryptoService) *CryptoController {
	return &CryptoController{service: svc}
}

func (c *CryptoController) GetAllCryptos(ctx *gin.Context) {
	cryptos, err := c.service.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, cryptos)
}

func (c *CryptoController) GetCryptoBySymbol(ctx *gin.Context) {
	symbol := ctx.Param("symbol")
	crypto, err := c.service.GetBySymbol(ctx, symbol)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, crypto)
}

func (c *CryptoController) CreateCrypto(ctx *gin.Context) {
	var request struct {
		Symbol      string `json:"symbol" binding:"required"`
		Description string `json:"description" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	crypto, err := c.service.Create(ctx, request.Symbol, request.Description)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, crypto)
}

func (c *CryptoController) UpdateCrypto(ctx *gin.Context) {
	var crypto model.CryptoCurrency
	if err := ctx.ShouldBindJSON(&crypto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad json format"})
		return
	}

	if err := c.service.Update(ctx, &crypto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "crypto updated"})
}

func (c *CryptoController) DeleteCrypto(ctx *gin.Context) {
	symbol := ctx.Param("symbol")
	crypto := &model.CryptoCurrency{Symbol: symbol}
	if err := c.service.Delete(ctx, crypto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "crypto deleted"})
}
