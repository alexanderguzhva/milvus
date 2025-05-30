// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package datacoord

import (
	"context"
	"time"

	"github.com/milvus-io/milvus/internal/types"
	"github.com/milvus-io/milvus/pkg/v2/proto/indexpb"
)

type Task interface {
	GetTaskID() int64
	GetNodeID() int64
	ResetTask(mt *meta)
	GetTaskSlot() int64
	PreCheck(ctx context.Context, dependency *taskScheduler) bool
	CheckTaskHealthy(mt *meta) bool
	SetState(state indexpb.JobState, failReason string)
	GetState() indexpb.JobState
	GetFailReason() string
	UpdateVersion(ctx context.Context, nodeID int64, meta *meta, compactionHandler compactionPlanContext) error
	UpdateMetaBuildingState(meta *meta) error
	AssignTask(ctx context.Context, client types.DataNodeClient, meta *meta) bool
	QueryResult(ctx context.Context, client types.DataNodeClient)
	DropTaskOnWorker(ctx context.Context, client types.DataNodeClient) bool
	SetJobInfo(meta *meta) error
	SetQueueTime(time.Time)
	GetQueueTime() time.Time
	SetStartTime(time.Time)
	GetStartTime() time.Time
	SetEndTime(time.Time)
	GetEndTime() time.Time
	GetTaskType() string
	DropTaskMeta(ctx context.Context, meta *meta) error
}
