/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v3

import (
	envoy_core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_api "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
)

type PipeListenerConfigurer struct {
	SocketPath string
}

func (c *PipeListenerConfigurer) Configure(l *envoy_api.Listener) error {
	l.Address = &envoy_core.Address{
		Address: &envoy_core.Address_Pipe{
			Pipe: &envoy_core.Pipe{
				Path: c.SocketPath,
			},
		},
	}

	return nil
}
