// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package two

import (
	"context"
	"net/http"
)

// NodeService handles communication with the node related
// methods of the Ark Core API - Version 2.
type NodeService Service

// Get the node status.
func (s *NodeService) Status(ctx context.Context) (*GetNodeStatus, *http.Response, error) {
	var responseStruct *GetNodeStatus
	resp, err := s.client.SendRequest(ctx, "GET", "node/status", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the node syncing status.
func (s *NodeService) Syncing(ctx context.Context) (*GetNodeSyncing, *http.Response, error) {
	var responseStruct *GetNodeSyncing
	resp, err := s.client.SendRequest(ctx, "GET", "node/syncing", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}

// Get the node configuration.
func (s *NodeService) Configuration(ctx context.Context) (*GetNodeConfiguration, *http.Response, error) {
	var responseStruct *GetNodeConfiguration
	resp, err := s.client.SendRequest(ctx, "GET", "node/configuration", nil, &responseStruct)

	if err != nil {
		return nil, resp, err
	}

	return responseStruct, resp, err
}
