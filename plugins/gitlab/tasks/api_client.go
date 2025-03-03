/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tasks

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/apache/incubator-devlake/errors"
	"github.com/apache/incubator-devlake/plugins/core"
	"github.com/apache/incubator-devlake/plugins/gitlab/models"
	"github.com/apache/incubator-devlake/plugins/helper"
)

func NewGitlabApiClient(taskCtx core.TaskContext, connection *models.GitlabConnection) (*helper.ApiAsyncClient, errors.Error) {
	// create synchronize api client so we can calculate api rate limit dynamically
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %v", connection.Token),
	}
	apiClient, err := helper.NewApiClient(taskCtx.GetContext(), connection.Endpoint, headers, 0, connection.Proxy, taskCtx)
	if err != nil {
		return nil, err
	}

	// create rate limit calculator
	rateLimiter := &helper.ApiRateLimitCalculator{
		UserRateLimitPerHour: connection.RateLimitPerHour,
		DynamicRateLimit: func(res *http.Response) (int, time.Duration, errors.Error) {
			rateLimitHeader := res.Header.Get("RateLimit-Limit")
			if rateLimitHeader == "" {
				// use default
				return 0, 0, nil
			}
			rateLimit, err := strconv.Atoi(rateLimitHeader)
			if err != nil {
				return 0, 0, errors.Default.Wrap(err, "failed to parse RateLimit-Limit header")
			}
			// seems like gitlab rate limit is on minute basis
			if rateLimit > 200 {
				return 200, 1 * time.Minute, nil
			} else {
				return rateLimit, 1 * time.Minute, nil
			}
		},
	}
	asyncApiClient, err := helper.CreateAsyncApiClient(
		taskCtx,
		apiClient,
		rateLimiter,
	)
	if err != nil {
		return nil, err
	}
	return asyncApiClient, nil
}

func ignoreHTTPStatus403(res *http.Response) errors.Error {
	if res.StatusCode == http.StatusForbidden {
		return helper.ErrIgnoreAndContinue
	}
	return nil
}
