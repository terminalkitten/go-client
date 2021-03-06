// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package one

import (
	"context"
	"net/http"
)

// AccountsService handles communication with the accounts related
// methods of the Ark Core API - Version 1.
type AccountsService Service

// Get all accounts.
func (s *AccountsService) List(ctx context.Context) (*Accounts, *http.Response, error) {
	var responseStruct *Accounts
	resp, err := s.client.SendRequest(ctx, "GET", "accounts/getAllAccounts", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a account by the given address.
func (s *AccountsService) Get(ctx context.Context, address string) (*AccountSingle, *http.Response, error) {
	query := &AddressQuery{Address: address}

	var responseStruct *AccountSingle
	resp, err := s.client.SendRequest(ctx, "GET", "accounts", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Count all accounts.
func (s *AccountsService) Count(ctx context.Context) (*AccountsCount, *http.Response, error) {
	var responseStruct *AccountsCount
	resp, err := s.client.SendRequest(ctx, "GET", "accounts/count", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get a delegate by the given address.
func (s *AccountsService) Delegate(ctx context.Context, address string) (*AccountDelegates, *http.Response, error) {
	query := &AddressQuery{Address: address}

	var responseStruct *AccountDelegates
	resp, err := s.client.SendRequest(ctx, "GET", "accounts/delegates", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the delegate registration fee.
func (s *AccountsService) DelegateFee(ctx context.Context) (*DelegateFee, *http.Response, error) {
	var responseStruct *DelegateFee
	resp, err := s.client.SendRequest(ctx, "GET", "accounts/delegates/fee", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the balance for an account by the given address.
func (s *AccountsService) Balance(ctx context.Context, address string) (*AccountBalance, *http.Response, error) {
	query := &AddressQuery{Address: address}

	var responseStruct *AccountBalance
	resp, err := s.client.SendRequest(ctx, "GET", "accounts/getBalance", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the public key for an account by the given address.
func (s *AccountsService) PublicKey(ctx context.Context, address string) (*PublicKey, *http.Response, error) {
	query := &AddressQuery{Address: address}

	var responseStruct *PublicKey
	resp, err := s.client.SendRequest(ctx, "GET", "accounts/getPublicKey", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get all wallets sorted by balance in descending order.
func (s *AccountsService) Top(ctx context.Context, query *TopQuery) (*AccountsTop, *http.Response, error) {
	var responseStruct *AccountsTop
	resp, err := s.client.SendRequest(ctx, "GET", "accounts/top", query, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}
