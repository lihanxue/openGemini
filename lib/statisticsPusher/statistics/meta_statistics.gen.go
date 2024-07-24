// Code generated by tmpl; DO NOT EDIT.
// https://github.com/benbjohnson/tmpl
//
// Source: statistics.tmpl

/*
Copyright 2024 Huawei Cloud Computing Technologies Co., Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package statistics

import (
	"sync"
	"sync/atomic"

	"github.com/openGemini/openGemini/lib/statisticsPusher/statistics/opsStat"
)

type MetaStatistics struct {
	itemSnapshotTotal              int64
	itemSnapshotDataSize           int64
	itemSnapshotUnmarshalDuration  int64
	itemLeaderSwitchTotal          int64
	itemStoreApplyTotal            int64
	itemGetFromOpsMapTotal         int64
	itemGetFromOpsMapLenTotal      int64
	itemGetFromDataMarshalTotal    int64
	itemGetFromDataMarshalLenTotal int64

	mu  sync.RWMutex
	buf []byte

	tags map[string]string
}

var instanceMetaStatistics = &MetaStatistics{}

func NewMetaStatistics() *MetaStatistics {
	return instanceMetaStatistics
}

func (s *MetaStatistics) Init(tags map[string]string) {
	s.tags = make(map[string]string)
	for k, v := range tags {
		s.tags[k] = v
	}
}

func (s *MetaStatistics) Collect(buffer []byte) ([]byte, error) {
	data := map[string]interface{}{
		"SnapshotTotal":              s.itemSnapshotTotal,
		"SnapshotDataSize":           s.itemSnapshotDataSize,
		"SnapshotUnmarshalDuration":  s.itemSnapshotUnmarshalDuration,
		"LeaderSwitchTotal":          s.itemLeaderSwitchTotal,
		"StoreApplyTotal":            s.itemStoreApplyTotal,
		"GetFromOpsMapTotal":         s.itemGetFromOpsMapTotal,
		"GetFromOpsMapLenTotal":      s.itemGetFromOpsMapLenTotal,
		"GetFromDataMarshalTotal":    s.itemGetFromDataMarshalTotal,
		"GetFromDataMarshalLenTotal": s.itemGetFromDataMarshalLenTotal,
	}

	buffer = AddPointToBuffer("meta", s.tags, data, buffer)
	if len(s.buf) > 0 {
		s.mu.Lock()
		buffer = append(buffer, s.buf...)
		s.buf = s.buf[:0]
		s.mu.Unlock()
	}

	return buffer, nil
}

func (s *MetaStatistics) CollectOps() []opsStat.OpsStatistic {
	data := map[string]interface{}{
		"SnapshotTotal":              s.itemSnapshotTotal,
		"SnapshotDataSize":           s.itemSnapshotDataSize,
		"SnapshotUnmarshalDuration":  s.itemSnapshotUnmarshalDuration,
		"LeaderSwitchTotal":          s.itemLeaderSwitchTotal,
		"StoreApplyTotal":            s.itemStoreApplyTotal,
		"GetFromOpsMapTotal":         s.itemGetFromOpsMapTotal,
		"GetFromOpsMapLenTotal":      s.itemGetFromOpsMapLenTotal,
		"GetFromDataMarshalTotal":    s.itemGetFromDataMarshalTotal,
		"GetFromDataMarshalLenTotal": s.itemGetFromDataMarshalLenTotal,
	}

	return []opsStat.OpsStatistic{
		{
			Name:   "meta",
			Tags:   s.tags,
			Values: data,
		},
	}
}

func (s *MetaStatistics) AddSnapshotTotal(i int64) {
	atomic.AddInt64(&s.itemSnapshotTotal, i)
}

func (s *MetaStatistics) AddSnapshotDataSize(i int64) {
	atomic.AddInt64(&s.itemSnapshotDataSize, i)
}

func (s *MetaStatistics) AddSnapshotUnmarshalDuration(i int64) {
	atomic.AddInt64(&s.itemSnapshotUnmarshalDuration, i)
}

func (s *MetaStatistics) AddLeaderSwitchTotal(i int64) {
	atomic.AddInt64(&s.itemLeaderSwitchTotal, i)
}

func (s *MetaStatistics) AddStoreApplyTotal(i int64) {
	atomic.AddInt64(&s.itemStoreApplyTotal, i)
}

func (s *MetaStatistics) AddGetFromOpsMapTotal(i int64) {
	atomic.AddInt64(&s.itemGetFromOpsMapTotal, i)
}

func (s *MetaStatistics) AddGetFromOpsMapLenTotal(i int64) {
	atomic.AddInt64(&s.itemGetFromOpsMapLenTotal, i)
}

func (s *MetaStatistics) AddGetFromDataMarshalTotal(i int64) {
	atomic.AddInt64(&s.itemGetFromDataMarshalTotal, i)
}

func (s *MetaStatistics) AddGetFromDataMarshalLenTotal(i int64) {
	atomic.AddInt64(&s.itemGetFromDataMarshalLenTotal, i)
}

func (s *MetaStatistics) Push(item *MetaStatItem) {
	if !item.Validate() {
		return
	}

	data := item.Values()
	tags := item.Tags()
	AllocTagMap(tags, s.tags)

	s.mu.Lock()
	s.buf = AddPointToBuffer("meta", tags, data, s.buf)
	s.mu.Unlock()
}

type MetaStatItem struct {
	validateHandle func(item *MetaStatItem) bool

	Status int64
	LTime  int64

	NodeID string
	Host   string
}

func (s *MetaStatItem) Push() {
	NewMetaStatistics().Push(s)
}

func (s *MetaStatItem) Validate() bool {
	if s.validateHandle == nil {
		return true
	}
	return s.validateHandle(s)
}

func (s *MetaStatItem) Values() map[string]interface{} {
	return map[string]interface{}{
		"Status": s.Status,
		"LTime":  s.LTime,
	}
}

func (s *MetaStatItem) Tags() map[string]string {
	return map[string]string{
		"NodeID": s.NodeID,
		"Host":   s.Host,
	}
}
