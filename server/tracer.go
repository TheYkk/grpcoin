// Copyright 2021 Ahmet Alp Balkan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"

	"go.opentelemetry.io/otel/sdk/trace"
)

var _ trace.SpanExporter = dummyTraceExporter{}

type dummyTraceExporter struct{}

func (d dummyTraceExporter) ExportSpans(ctx context.Context, ss []*trace.SpanSnapshot) error {
	return nil
}

func (d dummyTraceExporter) Shutdown(ctx context.Context) error {
	return nil
}
