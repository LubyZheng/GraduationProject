/*just for memory*/

package cgroup

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	CgroupForMemoryDir = "/sys/fs/cgroup/memory/limit"
	ProcFile           = "cgroup.procs"
	MemoryLimitFile    = "memory.limit_in_bytes"
	SwapLimitFile      = "memory.swappiness"
	MMUIB              = "memory.max_usage_in_bytes"
)

func MkCgroupDir() error {
	err := os.Mkdir(CgroupForMemoryDir, 0777)
	if err != nil {
		return err
	}
	return nil
}

func writeFile(path string, value int) error {
	err := ioutil.WriteFile(path, []byte(fmt.Sprintf("%d", value)), 0777)
	if err != nil {
		return err
	}
	return nil
}

func LimitMemory(limitMemory int) {
	path := filepath.Join(CgroupForMemoryDir, MemoryLimitFile)
	writeFile(path, limitMemory)
}

func SetProc(pid int) error {
	path := filepath.Join(CgroupForMemoryDir, ProcFile)
	err := writeFile(path, pid)
	if err != nil {
		return err
	}
	return nil
}

func Read() string {
	b, _ := ioutil.ReadFile(filepath.Join(CgroupForMemoryDir, MMUIB))
	return string(b)
}
