/**
 * Copyright (c) F5, Inc.
 *
 * This source code is licensed under the Apache License, Version 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package os

type Process struct {
	Pid  int32
	Ppid int32
	Name string
	Cmd  string
	Exe  string
}