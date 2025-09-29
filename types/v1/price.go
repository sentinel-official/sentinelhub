package v1

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// QuotePriceFunc defines a function signature for converting a base price to a quote price.
type QuotePriceFunc func(ctx sdk.Context, basePrice sdk.DecCoin) (sdk.Coin, error)

// NewPriceFromString parses a string like "denom:base,quote" into a Price.
func NewPriceFromString(s string) (Price, error) {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return Price{}, errors.New("empty string")
	}

	// Split denom and values
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		return Price{}, errors.New("invalid format")
	}

	denom := parts[0]
	values := parts[1]

	// Split base and quote values
	parts = strings.SplitN(values, ",", 2)
	if len(parts) != 2 {
		return Price{}, errors.New("invalid format")
	}

	baseValue, err := sdkmath.LegacyNewDecFromStr(parts[0])
	if err != nil {
		return Price{}, fmt.Errorf("invalid base value: %w", err)
	}

	quoteValue, ok := sdkmath.NewIntFromString(parts[1])
	if !ok {
		return Price{}, errors.New("invalid quote value")
	}

	price := Price{
		Denom:      denom,
		BaseValue:  baseValue,
		QuoteValue: quoteValue,
	}

	// Ensure the price is valid before returning
	if err := price.Validate(); err != nil {
		return Price{}, fmt.Errorf("invalid price: %w", err)
	}

	return price, nil
}

// NewPriceFromCoin constructs a Price with only a quote value.
func NewPriceFromCoin(coin sdk.Coin) Price {
	return Price{
		Denom:      coin.Denom,
		BaseValue:  sdkmath.LegacyZeroDec(),
		QuoteValue: coin.Amount,
	}
}

// ZeroPrice returns a Price with zero base and quote values.
func ZeroPrice(denom string) Price {
	return Price{
		Denom:      denom,
		BaseValue:  sdkmath.LegacyZeroDec(),
		QuoteValue: sdkmath.ZeroInt(),
	}
}

// BasePrice converts Price to sdk.DecCoin.
func (p Price) BasePrice() sdk.DecCoin {
	return sdk.DecCoin{
		Denom:  p.Denom,
		Amount: p.BaseValue,
	}
}

// QuotePrice converts Price to sdk.Coin.
func (p Price) QuotePrice() sdk.Coin {
	return sdk.Coin{
		Denom:  p.Denom,
		Amount: p.QuoteValue,
	}
}

// Copy returns a deep copy of Price.
func (p Price) Copy() Price {
	return Price{
		Denom:      p.Denom,
		BaseValue:  p.BaseValue,
		QuoteValue: p.QuoteValue,
	}
}

// String converts Price to string format "denom:base,quote".
func (p Price) String() string {
	return fmt.Sprintf("%s:%s,%s", p.Denom, p.BaseValue, p.QuoteValue)
}

// IsEqual checks if two Price values are equal.
func (p Price) IsEqual(v Price) bool {
	if p.Denom != v.Denom {
		return false
	}

	if !p.BaseValue.IsZero() || !v.BaseValue.IsZero() {
		return p.BaseValue.Equal(v.BaseValue)
	}

	return p.QuoteValue.Equal(v.QuoteValue)
}

// IsGT checks if this Price is greater than another.
func (p Price) IsGT(v Price) bool {
	if p.Denom != v.Denom {
		return false
	}

	if !p.BaseValue.IsZero() || !v.BaseValue.IsZero() {
		return p.BaseValue.GT(v.BaseValue)
	}

	return p.QuoteValue.GT(v.QuoteValue)
}

// IsGTE checks if this Price is greater than or equal to another.
func (p Price) IsGTE(v Price) bool {
	if p.Denom != v.Denom {
		return false
	}

	return !p.IsLT(v)
}

// IsLT checks if this Price is less than another.
func (p Price) IsLT(v Price) bool {
	if p.Denom != v.Denom {
		return false
	}

	if !p.BaseValue.IsZero() || !v.BaseValue.IsZero() {
		return p.BaseValue.LT(v.BaseValue)
	}

	return p.QuoteValue.LT(v.QuoteValue)
}

// IsLTE checks if this Price is less than or equal to another.
func (p Price) IsLTE(v Price) bool {
	if p.Denom != v.Denom {
		return false
	}

	return !p.IsGT(v)
}

// Validate checks if the Price has a valid denom and non-negative values.
func (p Price) Validate() error {
	if err := sdk.ValidateDenom(p.Denom); err != nil {
		return fmt.Errorf("invalid denom: %w", err)
	}

	if p.BaseValue.IsNegative() {
		return errors.New("base value cannot be negative")
	}

	if p.QuoteValue.IsNegative() {
		return errors.New("quote value cannot be negative")
	}

	return nil
}

// negative returns the negated Price (both base and quote).
func (p Price) negative() Price {
	return Price{
		Denom:      p.Denom,
		BaseValue:  p.BaseValue.Neg(),
		QuoteValue: p.QuoteValue.Neg(),
	}
}

// Add combines two Price values with the same denom.
func (p Price) Add(v Price) Price {
	if p.Denom != v.Denom {
		panic(errors.New("denoms do not match"))
	}

	return Price{
		Denom:      p.Denom,
		BaseValue:  p.BaseValue.Add(v.BaseValue),
		QuoteValue: p.QuoteValue.Add(v.QuoteValue),
	}
}

// Sub subtracts one Price from another with the same denom.
func (p Price) Sub(v Price) Price {
	if p.Denom != v.Denom {
		panic(errors.New("denoms do not match"))
	}

	return Price{
		Denom:      p.Denom,
		BaseValue:  p.BaseValue.Sub(v.BaseValue),
		QuoteValue: p.QuoteValue.Sub(v.QuoteValue),
	}
}

// UpdateQuoteValue applies a pricing function to compute a new quote value from the base.
func (p Price) UpdateQuoteValue(ctx sdk.Context, fn QuotePriceFunc) (Price, error) {
	// If BaseValue is zero, return the original Price
	if p.BaseValue.IsZero() {
		return p, nil
	}

	// Get the base price using BasePrice()
	basePrice := p.BasePrice()

	// Compute the new quote value using the provided function
	newQuote, err := fn(ctx, basePrice)
	if err != nil {
		return Price{}, err
	}

	// If newQuote is zero, return the original Price
	if newQuote.IsZero() {
		return p, nil
	}

	// Return a new Price instance with the updated QuoteValue
	return Price{
		Denom:      p.Denom,
		BaseValue:  p.BaseValue,
		QuoteValue: newQuote.Amount,
	}, nil
}

// Prices is a collection of Price values.
type Prices []Price

// NewPricesFromString parses a semicolon-separated string of prices.
func NewPricesFromString(s string) (Prices, error) {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return Prices{}, nil
	}

	parts := strings.Split(s, ";")

	prices := make(Prices, len(parts))
	for i, part := range parts {
		price, err := NewPriceFromString(part)
		if err != nil {
			return nil, fmt.Errorf("failed to parse price at index %d: %w", i, err)
		}

		prices[i] = price
	}

	prices = prices.Sort()
	if err := prices.Validate(); err != nil {
		return nil, fmt.Errorf("invalid prices: %w", err)
	}

	return prices, nil
}

// NewPricesFromCoins constructs a Prices list from a slice of sdk.Coin.
func NewPricesFromCoins(coins ...sdk.Coin) (Prices, error) {
	prices := make(Prices, len(coins))
	for i, coin := range coins {
		prices[i] = NewPriceFromCoin(coin)
	}

	prices = prices.Sort()
	if err := prices.Validate(); err != nil {
		return nil, fmt.Errorf("invalid prices: %w", err)
	}

	return prices, nil
}

// Copy returns a deep copy of the Prices slice.
func (p Prices) Copy() Prices {
	v := make(Prices, len(p))
	for i := range p {
		v[i] = p[i].Copy()
	}

	return v
}

// String converts the Prices slice to a semicolon-separated string.
func (p Prices) String() string {
	parts := make([]string, 0, len(p))
	for _, price := range p {
		parts = append(parts, price.String())
	}

	return strings.Join(parts, ";")
}

// IsEqual checks if two Prices slices are equal.
func (p Prices) IsEqual(v Prices) bool {
	if len(p) != len(v) {
		return false
	}

	for i := range p {
		if !p[i].IsEqual(v[i]) {
			return false
		}
	}

	return true
}

// IsSorted returns true if the Prices slice is sorted by denom.
func (p Prices) IsSorted() bool {
	for i := 1; i < len(p); i++ {
		if p[i].Denom < p[i-1].Denom {
			return false
		}
	}

	return true
}

// Validate checks that the Prices slice is sorted and all prices are valid.
func (p Prices) Validate() error {
	for i := range p {
		if i > 0 && p[i].Denom <= p[i-1].Denom {
			return errors.New("denoms must be sorted")
		}

		if err := p[i].Validate(); err != nil {
			return fmt.Errorf("invalid price: %w", err)
		}
	}

	return nil
}

// Len returns the number of Prices.
func (p Prices) Len() int {
	return len(p)
}

// Swap swaps two Prices elements.
func (p Prices) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Less compares Prices by denom for sorting.
func (p Prices) Less(i, j int) bool {
	return p[i].Denom < p[j].Denom
}

// Sort sorts the Prices by denom.
func (p Prices) Sort() Prices {
	sort.Sort(p)

	return p
}

// AmountOf returns the base and quote amount for a specific denom.
func (p Prices) AmountOf(denom string) (sdkmath.LegacyDec, sdkmath.Int) {
	index := p.IndexOf(denom)
	if index != -1 {
		return p[index].BaseValue, p[index].QuoteValue
	}

	return sdkmath.LegacyZeroDec(), sdkmath.ZeroInt()
}

// IndexOf returns the index of a Price with the given denom.
func (p Prices) IndexOf(denom string) int {
	index := sort.Search(len(p), func(i int) bool {
		return p[i].Denom >= denom
	})
	if index < len(p) && p[index].Denom == denom {
		return index
	}

	return -1
}

// Find retrieves the Price for a given denom, if it exists.
func (p Prices) Find(denom string) (Price, bool) {
	index := p.IndexOf(denom)
	if index != -1 {
		return p[index], true
	}

	return Price{}, false
}

// Add adds multiple Price values to the Prices slice.
func (p Prices) Add(items ...Price) Prices {
	return p.add(items)
}

// Sub subtracts multiple Price values from the Prices slice.
func (p Prices) Sub(items ...Price) Prices {
	return p.sub(items)
}

// Map converts the Prices slice into a map[denom]Price.
// Panics if a duplicate denom exists in the result.
func (p Prices) Map() map[string]Price {
	m := make(map[string]Price, len(p))
	for _, price := range p {
		if _, exists := m[price.Denom]; exists {
			panic(fmt.Errorf("duplicate denom %s", price.Denom))
		}

		m[price.Denom] = price
	}

	return m
}

// add adds values of another Prices slice to this one.
func (p Prices) add(v Prices) Prices {
	return p.merge(v)
}

// merge merges and sums two Prices slices by denom.
func (p Prices) merge(v Prices) Prices {
	m := make(map[string]Price)
	for _, price := range p {
		m[price.Denom] = price
	}

	for _, price := range v {
		curr, ok := m[price.Denom]
		if !ok {
			curr = ZeroPrice(price.Denom)
		}

		m[price.Denom] = curr.Add(price)
	}

	res := make(Prices, 0, len(m))
	for _, price := range m {
		res = append(res, price)
	}

	return res.Sort()
}

// negative returns Prices with all values negated.
func (p Prices) negative() Prices {
	v := make(Prices, len(p))
	for i, price := range p {
		v[i] = price.negative()
	}

	return v
}

// sub subtracts values of another Prices slice from this one.
func (p Prices) sub(v Prices) Prices {
	return p.merge(v.negative())
}
