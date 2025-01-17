// Copyright 2018 Kaleido, a ConsenSys business

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	kld "github.com/kaleido-io/kaleido-sdk-go/kaleido"
	"github.com/spf13/cobra"
)

var serviceListCmd = &cobra.Command{
	Use:   "service",
	Short: "List deployed services in an environment",
	Run: func(cmd *cobra.Command, args []string) {
		if consortiumId == "" {
			fmt.Println("Missing required parameter: --consortiumId for the consortium to list services of")
			os.Exit(1)
		}

		if environmentId == "" {
			fmt.Println("Missing required parameter: --environmentId for the environment to list services of")
			os.Exit(1)
		}

		client := getNewClient()
		var services []kld.Service
		_, err := client.ListServices(consortiumId, environmentId, &services)

		if err != nil {
			fmt.Printf("Failed to list services. %v\n", err)
			os.Exit(1)
		}

		encoded, _ := json.Marshal(services)
		fmt.Printf("\n%+v\n", string(encoded))
	},
}

var serviceGetCmd = &cobra.Command{
	Use:   "service",
	Short: "Retrieves service details",
	Run: func(cmd *cobra.Command, args []string) {
		if consortiumId == "" {
			fmt.Println("Missing required parameter: --consortiumId for the consortium that the service belongs to")
			os.Exit(1)
		}

		if environmentId == "" {
			fmt.Println("Missing required parameter: --environmentId for the environment that the service belongs to")
			os.Exit(1)
		}

		if serviceId == "" {
			fmt.Println("Missing required parameter: --id for the service to retrieve")
			os.Exit(1)
		}

		client := getNewClient()
		var service kld.Service
		res, err := client.GetService(consortiumId, environmentId, serviceId, &service)

		validateGetResponse(res, err, "service")
	},
}

var serviceCreateCmd = &cobra.Command{
	Use:   "service",
	Short: "Deploy a service",
	Run: func(cmd *cobra.Command, args []string) {
		validateName()
		validateServiceType()
		validateConsortiumId("service")
		validateEnvironmentId("service")
		validateMembershipId("service")

		client := getNewClient()
		service := kld.NewService(name, serviceType, membershipId)
		res, err := client.CreateService(consortiumId, environmentId, &service)

		validateCreationResponse(res, err, "service")
	},
}

func newServiceListCmd() *cobra.Command {
	flags := serviceListCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to retrieve the services from")
	flags.StringVarP(&environmentId, "environment", "e", "", "Id of the environment to retrieve the services from")

	return serviceListCmd
}

func newServiceGetCmd() *cobra.Command {
	flags := serviceGetCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to retrieve the service from")
	flags.StringVarP(&environmentId, "environment", "e", "", "Id of the environment to retrieve the service from")
	flags.StringVarP(&serviceId, "id", "i", "", "Id of the service to retrieve")

	return serviceGetCmd
}

func newServiceCreateCmd() *cobra.Command {
	flags := serviceCreateCmd.Flags()
	flags.StringVarP(&name, "name", "n", "", "Name of the service")
	flags.StringVarP(&serviceType, "service", "s", "", "Type of the service")
	flags.StringVarP(&membershipId, "membership", "m", "", "Id of the membership this service belongs to")
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium this service is created under")
	flags.StringVarP(&environmentId, "environment", "e", "", "Id of the environment this service is created for")

	return serviceCreateCmd
}
