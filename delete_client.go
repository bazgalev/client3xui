/* Copyright 2024 Digilol OÜ
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package client3xui

import (
	"context"
	"net/http"
)

// Add client to an inbound.
func (c *Client) DeleteClient(ctx context.Context, inboundId uint, clientUuid string) (*ApiResponse, error) {
	resp := &ApiResponse{}
	err := c.Do(ctx, http.MethodPost, "/panel/api/inbounds/"+string(inboundId)+"/delClient/"+clientUuid, nil, resp)
	return resp, err
}