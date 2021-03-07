package flag

import (
	"flag"
	"fmt"
	"os"
)

type Default struct {
	flagSet  *flag.FlagSet
	FilePath string
	Language string
	Time     int
	Memory   int
}

func New() *Default {
	return &Default{
		flagSet: flag.NewFlagSet("", flag.ExitOnError),
	}
}

func (d *Default) Parse(args []string) {
	d.flagSet.StringVar(&d.FilePath, "file", "", "")
	d.flagSet.StringVar(&d.FilePath, "f", "", "")
	d.flagSet.StringVar(&d.Language, "language", "", "")
	d.flagSet.StringVar(&d.Language, "l", "", "")
	d.flagSet.IntVar(&d.Time, "Time", 0, "")
	d.flagSet.IntVar(&d.Time, "t", 0, "")
	d.flagSet.IntVar(&d.Memory, "m", 0, "")
	d.flagSet.IntVar(&d.Memory, "Memory", 0, "")
	d.flagSet.Usage = d.helpCallback
	err := d.flagSet.Parse(args)
	if err != nil {
		os.Exit(0)
	}
}

func (d *Default) helpCallback() {
	fmt.Printf(
		"Usage: %s [options]\n"+
			"Options:\n"+
			"    -f, --file <name>               FilePath to be executed. ex: xxx.cpp\n"+
			"    -l, --language <language>       Code language. ex: C++\n"+
			"    -t, --Time <Time>               Limit Time. unit: ms\n"+
			"    -m, --momery <Memory>           Limit Memory. unit: kb\n"+
			"Common Options:\n"+
			"	 -h, --help                      Show this message\n",
		os.Args[0],
	)
	os.Exit(0)
}
