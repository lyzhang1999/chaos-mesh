// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package chaosdaemon

import (
	pb "github.com/chaos-mesh/chaos-mesh/pkg/chaosdaemon/pb"
	"github.com/chaos-mesh/chaos-mesh/pkg/mock"
)

func applyTbf(tbf *pb.Tbf, pid uint32) error {
	// Mock point to return error in unit test
	if err := mock.On("TbfApplyError"); err != nil {
		if e, ok := err.(error); ok {
			return e
		}
		if ignore, ok := err.(bool); ok && ignore {
			return nil
		}
	}

	panic("unimplemented")
}

func deleteTbf(tbf *pb.Tbf, pid uint32) error {
	// Mock point to return error in unit test
	if err := mock.On("TbfDeleteError"); err != nil {
		if e, ok := err.(error); ok {
			return e
		}
		if ignore, ok := err.(bool); ok && ignore {
			return nil
		}
	}

	panic("unimplemented")
}
