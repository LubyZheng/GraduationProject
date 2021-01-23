package flags

import (
	"flag"
	"fmt"
	"os"
)

type Default struct {
	flagSet  *flag.FlagSet
	FilePath string
}

func New() *Default {
	return &Default{
		flagSet: flag.NewFlagSet("", flag.ExitOnError),
	}
}

func (d *Default) Parse(args []string) {
	d.flagSet.StringVar(&d.FilePath, "file", "", "")
	d.flagSet.StringVar(&d.FilePath, "f", "", "")
	d.flagSet.Usage = d.helpCallback
	err := d.flagSet.Parse(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func (d *Default) helpCallback() {
	fmt.Println("-f file")
	os.Exit(0)
}
