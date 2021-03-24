package judge

import (
	"flag"
	"fmt"
	"os"
)

type Arguments struct {
	flagSet    *flag.FlagSet
	FilePath   string
	Language   string
	Time       int
	Memory     int
	StudentID  string
	QuestionID string
}

func NewFlag() *Arguments {
	return &Arguments{
		flagSet: flag.NewFlagSet("", flag.ExitOnError),
	}
}

func (a *Arguments) Parse(args []string) {
	if len(os.Args) == 1 {
		a.helpCallback()
	}

	a.flagSet.StringVar(&a.FilePath, "file", "", "")
	a.flagSet.StringVar(&a.FilePath, "f", "", "")
	a.flagSet.StringVar(&a.Language, "language", "", "")
	a.flagSet.StringVar(&a.Language, "l", "", "")
	a.flagSet.IntVar(&a.Time, "time", 0, "")
	a.flagSet.IntVar(&a.Time, "t", 0, "")
	a.flagSet.IntVar(&a.Memory, "m", 0, "")
	a.flagSet.IntVar(&a.Memory, "memory", 0, "")
	a.flagSet.StringVar(&a.StudentID, "sid", "nobody", "")
	a.flagSet.StringVar(&a.QuestionID, "qid", "", "")

	a.flagSet.Usage = a.helpCallback

	err := a.flagSet.Parse(args)
	if err != nil {
		os.Exit(1)
	}

	if a.QuestionID == "" {
		fmt.Println("-qid <Question> is necessary")
		a.helpCallback()
	}
}

func (a *Arguments) helpCallback() {
	fmt.Printf(
		"\nUsage: %s [options]\n"+
			"\nOptions:\n"+
			"    -f,   --file <Name>               FileName. ex: .\\xxx.cpp\n"+
			"    -l,   --language <Language>       Code language. ex: C++\n"+
			"    -t,   --time <Time>               Limit Time. Unit: ms. Default:10000ms\n"+
			"    -m,   --momery <Memory>           Limit Memory. Unit: kb. Default:65536kb\n"+
			"    -sid, <Memory>                    Student's ID\n"+
			"    -qid, <Question>                  Question's ID\n"+
			"    -h,   --help                      Show this message\n\n",
		os.Args[0],
	)
	os.Exit(0)
}
