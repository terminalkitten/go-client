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

// SignaturesService handles communication with the signatures
// methods of the Ark Core API - Version 1.
type SignaturesService Service

// Get the second signature registration fee.
func (s *SignaturesService) Fee(ctx context.Context) (*SignaturesFee, *http.Response, error) {
	var responseStruct *SignaturesFee
	resp, err := s.client.SendRequest(ctx, "GET", "signatures/fee", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}
