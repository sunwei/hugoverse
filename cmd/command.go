package cmd

import (
	"errors"
	"flag"
	"fmt"
	"github.com/sunwei/hugoverse/internal/interfaces/cmd"
	"os"
)

func New() error {
	topLevel := flag.NewFlagSet("hugov", flag.ExitOnError)
	topLevel.Usage = func() {
		fmt.Println("Usage:\n  hugov [command]")
		fmt.Println("\nCommands:")
		fmt.Println("   server:  start the headless CMS server")
		fmt.Println("  version:  show hugoverse command version")

		fmt.Println("\nExample:")
		fmt.Println("  hugov server")
	}

	err := topLevel.Parse(os.Args[1:])
	if err != nil {
		return err
	}

	if topLevel.Parsed() {
		if len(topLevel.Args()) == 0 {
			topLevel.Usage()
			return errors.New("please specify a sub-command")
		}

		// 获取子命令及参数
		subCommand := topLevel.Args()[0]

		switch subCommand {
		case "version":
			versionCmd, err := cmd.NewVersionCmd(topLevel)
			if err != nil {
				return err
			}
			if err := versionCmd.Run(); err != nil {
				return err
			}
		case "server":
			openCmd, err := cmd.NewServerCmd(topLevel)
			if err != nil {
				return err
			}
			if err := openCmd.Run(); err != nil {
				return err
			}

		default:
			topLevel.Usage()
			return errors.New("invalid sub-command")
		}
	}

	return nil
}
