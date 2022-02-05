package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/fundraising module sentinel errors
var (
	ErrInvalidAuctionType          = sdkerrors.Register(ModuleName, 2, "invalid auction type")
	ErrInvalidStartPrice           = sdkerrors.Register(ModuleName, 3, "invalid start price")
	ErrInvalidVestingSchedules     = sdkerrors.Register(ModuleName, 4, "invalid vesting schedules")
	ErrInvalidAuctionStatus        = sdkerrors.Register(ModuleName, 5, "invalid auction status")
	ErrInsufficientRemainingAmount = sdkerrors.Register(ModuleName, 6, "insufficient remaining amount")
)
