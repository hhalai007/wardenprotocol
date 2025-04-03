package sched_test

import (
	"testing"

	keepertest "github.com/warden-protocol/wardenprotocol/testutil/keeper"
	"github.com/warden-protocol/wardenprotocol/testutil/nullify"
	sched "github.com/warden-protocol/wardenprotocol/x/sched/module"
	"github.com/warden-protocol/wardenprotocol/x/sched/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SchedKeeper(t)
	sched.InitGenesis(ctx, k, genesisState)
	got := sched.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
