// Teleport
// Copyright (C) 2024 Gravitational, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package decisionv1

import (
	"github.com/gravitational/trace"

	decisionpb "github.com/gravitational/teleport/api/gen/proto/go/teleport/decision/v1alpha1"
	"github.com/gravitational/teleport/lib/authz"
)

// ServiceConfig holds creation parameters for [Service].
type ServiceConfig struct {
	// Authorizer used by the service.
	Authorizer authz.Authorizer
}

// Service implements the teleport.decision.v1alpha1.DecisionService gRPC API.
type Service struct {
	decisionpb.UnimplementedDecisionServiceServer

	authorizer authz.Authorizer
}

// NewService creates a new [Service] instance.
func NewService(cfg ServiceConfig) (*Service, error) {
	if cfg.Authorizer == nil {
		return nil, trace.BadParameter("param Authorizer required")
	}

	return &Service{
		authorizer: cfg.Authorizer,
	}, nil
}
