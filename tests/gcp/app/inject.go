// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//+build gooseinject

package main

import (
	"context"

	"github.com/google/go-cloud/gcp"
	"github.com/google/go-cloud/gcp/gcpcloud"
	"github.com/google/go-cloud/goose"
	"github.com/google/go-cloud/server"
)

func initialize(ctx context.Context) (*server.Server, error) {
	panic(goose.Build(
		appSet,
		gcpcloud.Services,
		gcp.DefaultIdentity,
	))
}