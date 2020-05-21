// Copyright 2020 The PipeCD Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stackdriver

import "time"

const ProviderType = "StackdriverLogging"

// Provider is a client for stackdriver.
type Provider struct {
	serviceAccount []byte

	timeout time.Duration
}

func NewProvider(serviceAccount []byte) (*Provider, error) {
	return &Provider{
		serviceAccount: serviceAccount,
	}, nil
}

func (p *Provider) Type() string {
	return ProviderType
}

func (p *Provider) RunQuery(query string, threshold int) (bool, error) {
	return false, nil
}