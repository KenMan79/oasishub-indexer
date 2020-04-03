package getaccountbypublickey

import (
	"github.com/figment-networks/oasishub-indexer/domain/accountdomain"
	"github.com/figment-networks/oasishub-indexer/domain/delegationdomain"
	"github.com/figment-networks/oasishub-indexer/types"
	"github.com/figment-networks/oasishub-indexer/utils/errors"
	"github.com/figment-networks/oasishub-indexer/utils/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type httpHandler struct {
	useCase UseCase
}

func NewHttpHandler(useCase UseCase) types.HttpHandler {
	return &httpHandler{useCase: useCase}
}

type Request struct {
	PublicKey types.PublicKey `form:"public_key" binding:"required"`
}

type Response struct {
	*accountdomain.AccountAgg

	LastHeight                 types.Height                               `json:"last_height"`
	LastDelegations            []*delegationdomain.DelegationSeq          `json:"last_delegations"`
	RecentDebondingDelegations []*delegationdomain.DebondingDelegationSeq `json:"recent_debonding_delegations"`
}

func (h *httpHandler) Handle(c *gin.Context) {
	var req Request
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Error(err)
		err := errors.NewError("invalid public key", errors.ServerInvalidParamsError, err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res, err := h.useCase.Execute(req.PublicKey)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
