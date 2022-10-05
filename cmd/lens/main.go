/*
 * Copyright 2022 CoreLayer BV
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package main

import (
	"fmt"
	"github.com/corelayer/go-corelogic-nitro/pkg/lego/providers/http/corelogic"
	"github.com/corelayer/go-netscaler-nitro/pkg/client"
	"log"
)

func main() {
	// Banner source: https://patorjk.com/software/taag/#p=display&v=3&f=Ivrit&t=LENS
	var banner = "  _     _____ _   _ ____  \n | |   | ____| \\ | / ___| \n | |   |  _| |  \\| \\___ \\ \n | |___| |___| |\\  |___) |\n |_____|_____|_| \\_|____/ \n                          "
	fmt.Println(banner)

	env := client.Environment{
		Name: "Test",
		Type: client.Standalone,
		SNIP: client.Node{
			Name:    "Test-VPX",
			Address: "127.0.0.1",
		},
		Nodes: nil,
		Settings: client.Settings{
			Scheme:                   client.HTTPS,
			InsecureSkipVerify:       true,
			Timeout:                  60,
			LogTlsSecrets:            false,
			LogTlsSecretsDestination: "",
		},
	}

	_, err := corelogic.NewCoreLogicHttpProvider(env)
	if err != nil {
		log.Fatal(err)
	}

}
