/*
 * Minio Client (C) 2018 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"github.com/fatih/color"
	"github.com/minio/cli"
	"github.com/minio/mc/pkg/console"
	"github.com/minio/mc/pkg/probe"
	"github.com/minio/minio/pkg/madmin"
)

var adminUsersEnableCmd = cli.Command{
	Name:   "enable",
	Usage:  "Enable users",
	Action: mainAdminUsersEnable,
	Before: setGlobalsFromContext,
	Flags:  globalFlags,
	CustomHelpTemplate: `NAME:
  {{.HelpName}} - {{.Usage}}

USAGE:
  {{.HelpName}} TARGET USERNAME

FLAGS:
  {{range .VisibleFlags}}{{.}}
  {{end}}
EXAMPLES:
  1. Enable a disabled user 'newuser' for Minio server.
     $ {{.HelpName}} myminio newuser
`,
}

// checkAdminUsersEnableSyntax - validate all the passed arguments
func checkAdminUsersEnableSyntax(ctx *cli.Context) {
	if len(ctx.Args()) != 2 {
		cli.ShowCommandHelpAndExit(ctx, "enable", 1) // last argument is exit code
	}
}

// mainAdminUsersEnable is the handle for "mc admin users enable" command.
func mainAdminUsersEnable(ctx *cli.Context) error {
	checkAdminUsersEnableSyntax(ctx)

	console.SetColor("UserMessage", color.New(color.FgGreen))

	// Get the alias parameter from cli
	args := ctx.Args()
	aliasedURL := args.Get(0)

	// Create a new Minio Admin Client
	client, err := newAdminClient(aliasedURL)
	fatalIf(err, "Cannot get a configured admin connection.")

	e := client.SetUserStatus(args.Get(1), madmin.AccountEnabled)
	fatalIf(probe.NewError(e).Trace(args...), "Cannot enable user")

	printMsg(userMessage{
		op:        "enable",
		AccessKey: args.Get(1),
	})

	return nil
}
