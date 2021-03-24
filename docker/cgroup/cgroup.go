/*just for memory*/

package cgroup

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	CgroupForMemoryDir = "/sys/fs/cgroup/memory/limit"
	ProcFile           = "cgroup.procs"
	MemoryLimitFile    = "memory.limit_in_bytes"
	MemoryMaxUsageFile = "memory.max_usage_in_bytes"
)

func MkCgroupDir() error {
	err := os.Mkdir(CgroupForMemoryDir, 0700)
	if err != nil {
		return err
	}
	return nil
}

func writeFile(path string, value int) error {
	err := ioutil.WriteFile(path, []byte(fmt.Sprintf("%d", value)), 0700)
	if err != nil {
		return err
	}
	return nil
}

func SetProc(pid int) error {
	path := filepath.Join(CgroupForMemoryDir, ProcFile)
	err := writeFile(path, pid)
	if err != nil {
		return err
	}
	return nil
}

func LimitMemory(limitMemory int) error {
	path := filepath.Join(CgroupForMemoryDir, MemoryLimitFile)
	err := writeFile(path, limitMemory)
	if err != nil {
		return err
	}
	return nil
}

func ReadMemory() string {
	b, _ := ioutil.ReadFile(filepath.Join(CgroupForMemoryDir, MemoryMaxUsageFile))
	return strings.Trim(string(b), "\n")
}
