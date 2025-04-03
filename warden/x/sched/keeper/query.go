package keeper

import (
	"github.com/warden-protocol/wardenprotocol/x/sched/types"
)

var _ types.QueryServer = Keeper{}
