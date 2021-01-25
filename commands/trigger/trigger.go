//
// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
package trigger

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/apache/skywalking-infra-e2e/internal/constant"
	"github.com/apache/skywalking-infra-e2e/internal/flags"
)

func init() {
	Trigger.Flags().StringVar(&flags.Interval, "interval", "3s", "trigger the action every N seconds")
	Trigger.Flags().IntVar(&flags.Times, "times", 0, "how many times to trigger the action, 0=infinite")
	Trigger.Flags().StringVar(&flags.Action, "action", "", "the action of the trigger")
	Trigger.Flags().StringVar(&flags.HttpUrl, "url", "", "the url of the http action")
	Trigger.Flags().StringVar(&flags.HttpMethod, "method", "get", "the method of the http action")
}

var Trigger = &cobra.Command{
	Use:   "trigger",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		var action Action
		if strings.EqualFold(flags.Action, constant.ActionHTTP) {
			action = NewHTTPAction()
		}
		if action == nil {
			return fmt.Errorf("no such action for args")
		}
		return action.Do()
	},
}

type Action interface {
	Do() error
}
