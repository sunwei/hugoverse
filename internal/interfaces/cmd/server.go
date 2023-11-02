package cmd

import (
	"errors"
	"flag"
	"fmt"
	"github.com/sunwei/hugoverse/internal/interfaces/api"
	"github.com/sunwei/hugoverse/pkg/log"
	"net/http"
	"os"
)

type serverCmd struct {
	parent       *flag.FlagSet
	cmd          *flag.FlagSet
	hugoProjPath *string
}

func NewServerCmd(parent *flag.FlagSet) (*serverCmd, error) {
	nCmd := &serverCmd{
		parent: parent,
	}

	nCmd.cmd = flag.NewFlagSet("normal", flag.ExitOnError)
	nCmd.hugoProjPath = nCmd.cmd.String("p", "", fmt.Sprintf(
		"[required] target hugo project path \n(e.g. %s)", "path/to/your/hugo/project"))

	err := nCmd.cmd.Parse(parent.Args()[1:])
	if err != nil {
		return nil, err
	}

	return nCmd, nil
}

func (oc *serverCmd) Usage() {
	oc.cmd.Usage()
}

func (oc *serverCmd) Run() error {
	if *oc.hugoProjPath == "" {
		oc.cmd.Usage()
		return errors.New("please specify a target hugo project path")
	}

	_, err := os.Stat(*oc.hugoProjPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("directory %s does not exist", *oc.hugoProjPath)
	}

	if err != nil {
		return err
	}

	l := log.NewStdLogger()
	s, err := api.NewServer(func(s *api.Server) error {
		s.Log = l
		s.ProjPath = *oc.hugoProjPath

		return nil
	})
	if err != nil {
		l.Fatalf("Error creating server: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "1314"
	}

	l.Printf("Listening on :%v ...", port)
	l.Fatalf("Error listening on :%v: %v", port, http.ListenAndServe(":"+port, s))

	return nil
}
