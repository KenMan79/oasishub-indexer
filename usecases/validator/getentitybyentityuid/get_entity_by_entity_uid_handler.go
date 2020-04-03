package getentitybyentityuid

import (
	"github.com/figment-networks/oasishub-indexer/domain/delegationdomain"
	"github.com/figment-networks/oasishub-indexer/domain/entitydomain"
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
	EntityUID types.PublicKey `form:"entity_uid" binding:"required"`
}

type Response struct {
	*entitydomain.EntityAgg

	LastHeight                 types.Height                               `json:"last_height"`
	TotalValidated             int64                                      `json:"total_validated"`
	TotalMissed                int64                                      `json:"total_missed"`
	TotalProposed              int64                                      `json:"total_proposed"`
	LastDelegations            []*delegationdomain.DelegationSeq          `json:"last_delegations"`
	RecentDebondingDelegations []*delegationdomain.DebondingDelegationSeq `json:"recent_debonding_delegations"`
}

func (h *httpHandler) Handle(c *gin.Context) {
	var req Request
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Error(err)
		err := errors.NewError("invalid height", errors.ServerInvalidParamsError, err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.useCase.Execute(req.EntityUID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
