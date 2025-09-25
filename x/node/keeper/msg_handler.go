package keeper

import (
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/sentinelhub/v12/types"
	v1base "github.com/sentinel-official/sentinelhub/v12/types/v1"
	"github.com/sentinel-official/sentinelhub/v12/x/node/types"
	"github.com/sentinel-official/sentinelhub/v12/x/node/types/v3"
	sessiontypes "github.com/sentinel-official/sentinelhub/v12/x/session/types/v3"
)

// HandleMsgRegisterNode handles a request to register a new node.
// It validates the pricing fields, checks for duplicates, collects deposit, and stores the new node in an inactive state.
func (k *Keeper) HandleMsgRegisterNode(ctx sdk.Context, msg *v3.MsgRegisterNodeRequest) (*v3.MsgRegisterNodeResponse, error) {
	// Validate submitted gigabyte prices
	if !k.IsValidGigabytePrices(ctx, msg.GigabytePrices) {
		return nil, types.NewErrorInvalidPrices(msg.GigabytePrices)
	}

	// Validate submitted hourly prices
	if !k.IsValidHourlyPrices(ctx, msg.HourlyPrices) {
		return nil, types.NewErrorInvalidPrices(msg.HourlyPrices)
	}

	// Parse the account address from the sender string
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Convert account address to node address
	nodeAddr := base.NodeAddress(accAddr.Bytes())

	// Reject registration if a node with the same address already exists
	if k.HasNode(ctx, nodeAddr) {
		return nil, types.NewErrorDuplicateNode(nodeAddr)
	}

	// Deduct deposit from sender and send to the community pool
	deposit := k.Deposit(ctx)
	if err := k.FundCommunityPool(ctx, accAddr, deposit); err != nil {
		return nil, err
	}

	// Construct the new node object with default inactive state
	node := v3.Node{
		Address:        nodeAddr.String(),
		GigabytePrices: msg.GigabytePrices,
		HourlyPrices:   msg.HourlyPrices,
		RemoteAddrs:    msg.RemoteAddrs,
		Status:         v1base.StatusInactive,
		InactiveAt:     time.Time{},
		StatusAt:       ctx.BlockTime(),
	}

	// Save node and register inactive-at indexing
	k.SetNode(ctx, node)
	k.SetNodeForInactiveAt(ctx, node.InactiveAt, nodeAddr)

	// Emit creation event with pricing and metadata
	ctx.EventManager().EmitTypedEvent(
		&v3.EventCreate{
			NodeAddress:    node.Address,
			GigabytePrices: node.GetGigabytePrices().String(),
			HourlyPrices:   node.GetHourlyPrices().String(),
			RemoteAddrs:    node.RemoteAddrs,
		},
	)

	return &v3.MsgRegisterNodeResponse{}, nil
}

// HandleMsgUpdateNodeDetails handles a request to update a node's pricing and remote URL.
// It verifies node existence and applies new pricing and metadata, emitting an update event.
func (k *Keeper) HandleMsgUpdateNodeDetails(ctx sdk.Context, msg *v3.MsgUpdateNodeDetailsRequest) (*v3.MsgUpdateNodeDetailsResponse, error) {
	// Validate new gigabyte prices
	if !k.IsValidGigabytePrices(ctx, msg.GigabytePrices) {
		return nil, types.NewErrorInvalidPrices(msg.GigabytePrices)
	}

	// Validate new hourly prices
	if !k.IsValidHourlyPrices(ctx, msg.HourlyPrices) {
		return nil, types.NewErrorInvalidPrices(msg.HourlyPrices)
	}

	// Parse sender's node address
	nodeAddr, err := base.NodeAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Retrieve the existing node; fail if not found
	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	// Apply updated prices and optional remote addrs
	node.GigabytePrices = msg.GigabytePrices
	node.HourlyPrices = msg.HourlyPrices

	if len(msg.RemoteAddrs) > 0 {
		node.RemoteAddrs = msg.RemoteAddrs
	}

	// Save the updated node to state
	k.SetNode(ctx, node)

	// Emit details update event
	ctx.EventManager().EmitTypedEvent(
		&v3.EventUpdateDetails{
			NodeAddress:    node.Address,
			GigabytePrices: node.GetGigabytePrices().String(),
			HourlyPrices:   node.GetHourlyPrices().String(),
			RemoteAddrs:    node.RemoteAddrs,
		},
	)

	return &v3.MsgUpdateNodeDetailsResponse{}, nil
}

// HandleMsgUpdateNodeStatus handles a request to update the node's operational status.
// It updates internal indexing and lifecycle state based on the transition between active/inactive.
func (k *Keeper) HandleMsgUpdateNodeStatus(ctx sdk.Context, msg *v3.MsgUpdateNodeStatusRequest) (*v3.MsgUpdateNodeStatusResponse, error) {
	// Parse node address from sender
	nodeAddr, err := base.NodeAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Fetch the node; fail if it doesn't exist
	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	// Perform pre-hook actions if transitioning to inactive
	if msg.Status.Equal(v1base.StatusInactive) {
		if err := k.NodeInactivePreHook(ctx, nodeAddr); err != nil {
			return nil, err
		}
	}

	// Remove node from active index if transitioning from inactive to active
	if msg.Status.Equal(v1base.StatusActive) {
		if node.Status.Equal(v1base.StatusInactive) {
			k.DeleteInactiveNode(ctx, nodeAddr)
		}
	}

	// Remove node from active list if moving to inactive
	if msg.Status.Equal(v1base.StatusInactive) {
		if node.Status.Equal(v1base.StatusActive) {
			k.DeleteActiveNode(ctx, nodeAddr)
		}
	}

	// Clear existing inactive index before updating node
	k.DeleteNodeForInactiveAt(ctx, node.InactiveAt, nodeAddr)

	// Update status fields
	node.Status = msg.Status
	node.InactiveAt = time.Time{}
	node.StatusAt = ctx.BlockTime()

	// If becoming active, set inactiveAt for future expiry tracking
	if node.Status.Equal(v1base.StatusActive) {
		node.InactiveAt = k.GetInactiveAt(ctx)
	}

	// Save updated node and index inactiveAt
	k.SetNode(ctx, node)
	k.SetNodeForInactiveAt(ctx, node.InactiveAt, nodeAddr)

	// Emit event signaling node status change
	ctx.EventManager().EmitTypedEvent(
		&v3.EventUpdateStatus{
			NodeAddress: node.Address,
			Status:      node.Status,
		},
	)

	return &v3.MsgUpdateNodeStatusResponse{}, nil
}

// HandleMsgStartSession handles a request to initiate a session between a user and a node.
// It validates session parameters, calculates pricing, deducts deposit, and creates session state.
func (k *Keeper) HandleMsgStartSession(ctx sdk.Context, msg *v3.MsgStartSessionRequest) (*v3.MsgStartSessionResponse, error) {
	// Validate requested bandwidth and time allocations
	if msg.Gigabytes != 0 {
		if ok := k.IsValidSessionGigabytes(ctx, msg.Gigabytes); !ok {
			return nil, types.NewErrorInvalidGigabytes(msg.Gigabytes)
		}
	}

	if msg.Hours != 0 {
		if ok := k.IsValidSessionHours(ctx, msg.Hours); !ok {
			return nil, types.NewErrorInvalidHours(msg.Hours)
		}
	}

	// Convert and validate node address
	nodeAddr, err := base.NodeAddressFromBech32(msg.NodeAddress)
	if err != nil {
		return nil, err
	}

	// Ensure node exists and is active
	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	if !node.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidNodeStatus(nodeAddr, node.Status)
	}

	// Determine appropriate price based on requested denomination and service type
	var price v1base.Price
	if msg.Gigabytes != 0 {
		price, found = node.GigabytePrice(msg.MaxPrice.Denom)
		if !found {
			return nil, types.NewErrorPriceNotFound(msg.MaxPrice.Denom)
		}
	}

	if msg.Hours != 0 {
		price, found = node.HourlyPrice(msg.MaxPrice.Denom)
		if !found {
			return nil, types.NewErrorPriceNotFound(msg.MaxPrice.Denom)
		}
	}

	// Adjust price using current quote mechanism
	price, err = price.UpdateQuoteValue(ctx, k.QuotePriceFunc)
	if err != nil {
		return nil, err
	}

	// Reject if quoted price exceeds user's maximum offer
	if price.IsGT(msg.MaxPrice) {
		return nil, types.NewErrorInvalidPrice(price)
	}

	// Parse user's account address
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Generate new session ID and expiration time
	count := k.GetSessionCount(ctx)
	inactiveAt := k.GetSessionInactiveAt(ctx)

	// Construct the session object with bandwidth/time limits and metadata
	session := &v3.Session{
		BaseSession: &sessiontypes.BaseSession{
			ID:            count + 1,
			AccAddress:    accAddr.String(),
			NodeAddress:   nodeAddr.String(),
			DownloadBytes: sdkmath.ZeroInt(),
			UploadBytes:   sdkmath.ZeroInt(),
			MaxBytes:      msg.GetBytes(),
			Duration:      0,
			MaxDuration:   msg.GetDuration(),
			Status:        v1base.StatusActive,
			InactiveAt:    inactiveAt,
			StartAt:       ctx.BlockTime(),
			StatusAt:      ctx.BlockTime(),
		},
		Price: price,
	}

	// Deduct deposit from user to fund the session
	deposit := session.DepositAmount()
	if err := k.AddDeposit(ctx, accAddr, deposit); err != nil {
		return nil, err
	}

	// Persist session and update relevant indexes
	k.SetSessionCount(ctx, count+1)
	k.SetSession(ctx, session)
	k.SetSessionForAccount(ctx, accAddr, session.ID)
	k.SetSessionForNode(ctx, nodeAddr, session.ID)
	k.SetSessionForInactiveAt(ctx, session.InactiveAt, session.ID)

	// Emit event indicating session creation
	ctx.EventManager().EmitTypedEvent(
		&v3.EventCreateSession{
			ID:          session.ID,
			AccAddress:  session.AccAddress,
			NodeAddress: session.NodeAddress,
			Price:       session.Price.String(),
			MaxBytes:    session.MaxBytes.String(),
			MaxDuration: session.MaxDuration.String(),
		},
	)

	return &v3.MsgStartSessionResponse{
		ID: session.ID,
	}, nil
}

// HandleMsgUpdateParams handles a governance-authorized update to node module parameters.
// Only the module's authority account can invoke this change.
func (k *Keeper) HandleMsgUpdateParams(ctx sdk.Context, msg *v3.MsgUpdateParamsRequest) (*v3.MsgUpdateParamsResponse, error) {
	// Reject if the caller is not the module's authority
	if msg.From != k.authority {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Save updated parameters to state
	k.SetParams(ctx, msg.Params)

	return &v3.MsgUpdateParamsResponse{}, nil
}
