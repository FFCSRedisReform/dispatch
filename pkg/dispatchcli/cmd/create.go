///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	apiModels "github.com/vmware/dispatch/pkg/api-manager/gen/models"
	"github.com/vmware/dispatch/pkg/dispatchcli/i18n"
	eventModels "github.com/vmware/dispatch/pkg/event-manager/gen/models"
	functionModels "github.com/vmware/dispatch/pkg/function-manager/gen/models"
	iamModels "github.com/vmware/dispatch/pkg/identity-manager/gen/models"
	imageModels "github.com/vmware/dispatch/pkg/image-manager/gen/models"
	secretModels "github.com/vmware/dispatch/pkg/secret-store/gen/models"
	serviceModels "github.com/vmware/dispatch/pkg/service-manager/gen/models"
	"github.com/vmware/dispatch/pkg/utils"
)

var (
	createLong = i18n.T(`Create a resource. See subcommands for resources that can be created.`)

	// TODO: Add examples
	createExample = i18n.T(``)
	file          = i18n.T(``)
	workDir       = i18n.T(``)
)

type modelAction func(interface{}) error

type importFunction struct {
	functionModels.Function
}

func importFile(out io.Writer, errOut io.Writer, cmd *cobra.Command, args []string, actionMap map[string]modelAction) error {
	fullPath := path.Join(workDir, file)
	b, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return errors.Wrapf(err, "Error reading file %s", fullPath)
	}

	// Manually split up the yaml doc.  This is NOT a streaming parser.
	docs := bytes.Split(b, []byte("---"))

	type kind struct {
		Kind string `json:"kind"`
	}

	type output struct {
		APIs             []*apiModels.API                 `json:"api"`
		BaseImages       []*imageModels.BaseImage         `json:"baseImages"`
		Images           []*imageModels.Image             `json:"images"`
		DriverTypes      []*eventModels.DriverType        `json:"driverTypes"`
		Drivers          []*eventModels.Driver            `json:"drivers"`
		Subscriptions    []*eventModels.Subscription      `json:"subscriptions"`
		Functions        []*functionModels.Function       `json:"functions"`
		Secrets          []*secretModels.Secret           `json:"secrets"`
		Policies         []*iamModels.Policy              `json:"policies"`
		ServiceInstances []*serviceModels.ServiceInstance `json:"serviceInstances"`
		ServiceAccounts  []*iamModels.ServiceAccount      `json:"serviceaccounts"`
	}

	o := output{}

	for _, doc := range docs {
		k := kind{}
		err = yaml.Unmarshal(doc, &k)
		if err != nil {
			return errors.Wrapf(err, "Error decoding document %s", string(doc))
		}
		switch docKind := k.Kind; docKind {
		case utils.APIKind:
			m := &apiModels.API{}
			err = yaml.Unmarshal(doc, &m)
			if err != nil {
				return errors.Wrapf(err, "Error decoding api document %s", string(doc))
			}
			err = actionMap[docKind](m)
			if err != nil {
				return err
			}
			o.APIs = append(o.APIs, m)
			fmt.Fprintf(out, "Created %s: %s\n", docKind, *m.Name)
		case utils.BaseImageKind:
			m := &imageModels.BaseImage{}
			err = yaml.Unmarshal(doc, &m)
			if err != nil {
				return errors.Wrapf(err, "Error decoding base image document %s", string(doc))
			}
			err = actionMap[docKind](m)
			if err != nil {
				return err
			}
			o.BaseImages = append(o.BaseImages, m)
			fmt.Fprintf(out, "Created %s: %s\n", docKind, *m.Name)
		case utils.ImageKind:
			m := &imageModels.Image{}
			err = yaml.Unmarshal(doc, &m)
			if err != nil {
				return errors.Wrapf(err, "Error decoding image document %s", string(doc))
			}
			err = actionMap[docKind](m)
			if err != nil {
				return err
			}
			o.Images = append(o.Images, m)
			fmt.Fprintf(out, "Created %s: %s\n", docKind, *m.Name)
		case utils.FunctionKind:
			m := &functionModels.Function{}
			err = yaml.Unmarshal(doc, &m)
			if err != nil {
				return errors.Wrapf(err, "Error decoding function document %s", string(doc))
			}
			if strings.HasPrefix(*m.Code, "@") {
				functionPath := path.Join(workDir, (*m.Code)[1:])
				codeFileContent, err := ioutil.ReadFile(functionPath)
				if err != nil {
					return errors.Wrapf(err, "Error when reading content of %s", functionPath)
				}
				codeEncoded := string(codeFileContent)
				m.Code = &codeEncoded
			}
			err = actionMap[docKind](m)
			if err != nil {
				return err
			}
			o.Functions = append(o.Functions, m)
			fmt.Fprintf(out, "Created %s: %s\n", docKind, *m.Name)
		case utils.DriverTypeKind:
			m := &eventModels.DriverType{}
			err = yaml.Unmarshal(doc, &m)
			if err != nil {
				return errors.Wrapf(err, "Error when decoding driver type document of %s", string(doc))
			}
			err = actionMap[docKind](m)
			if err != nil {
				return err
			}
			o.DriverTypes = append(o.DriverTypes, m)
			fmt.Fprintf(out, "Created %s: %s\n", docKind, *m.Name)
		case utils.DriverKind:
			m := &eventModels.Driver{}
			err = yaml.Unmarshal(doc, &m)
			if err != nil {
				return errors.Wrapf(err, "Error decoding driver document %s", string(doc))
			}
			err = actionMap[docKind](m)
			if err != nil {
				return err
			}
			o.Drivers = append(o.Drivers, m)
			fmt.Fprintf(out, "Created %s: %s\n", docKind, *m.Name)
		case utils.SubscriptionKind:
			m := &eventModels.Subscription{}
			err = yaml.Unmarshal(doc, &m)
			if err != nil {
				return errors.Wrapf(err, "Error decoding subscription document %s", string(doc))
			}
			err = actionMap[docKind](m)
			if err != nil {
				return err
			}
			o.Subscriptions = append(o.Subscriptions, m)
			fmt.Fprintf(out, "Created %s: %s\n", docKind, *m.Name)
		case utils.SecretKind:
			m := &secretModels.Secret{}
			err = yaml.Unmarshal(doc, &m)
			if err != nil {
				return errors.Wrapf(err, "Error decoding secret document %s", string(doc))
			}
			err = actionMap[docKind](m)
			if err != nil {
				return err
			}
			o.Secrets = append(o.Secrets, m)
			fmt.Fprintf(out, "Created %s: %s\n", docKind, *m.Name)
		case utils.PolicyKind:
			m := &iamModels.Policy{}
			err = yaml.Unmarshal(doc, &m)
			if err != nil {
				return errors.Wrapf(err, "Error decoding policy document &s", string(doc))
			}
			err = actionMap[docKind](m)
			if err != nil {
				return err
			}
			o.Policies = append(o.Policies, m)
			fmt.Fprintf(out, "Created %s: %s\n", docKind, *m.Name)
		case utils.ServiceInstanceKind:
			m := &serviceModels.ServiceInstance{}
			err := yaml.Unmarshal(doc, &m)
			if err != nil {
				return errors.Wrapf(err, "Error decoding service instance document &s", string(doc))
			}
			err = actionMap[docKind](m)
			if err != nil {
				return err
			}
			o.ServiceInstances = append(o.ServiceInstances, m)
			fmt.Fprintf(out, "Created %s: %s\n", docKind, *m.Name)
		case utils.ServiceAccountKind:
			m := &iamModels.ServiceAccount{}
			err = yaml.Unmarshal(doc, &m)
			if err != nil {
				return errors.Wrapf(err, "Error decoding service account document &s", string(doc))
			}
			err = actionMap[docKind](m)
			if err != nil {
				return err
			}
			o.ServiceAccounts = append(o.ServiceAccounts, m)
			fmt.Fprintf(out, "Created %s: %s\n", docKind, *m.Name)
		default:
			continue
		}
	}
	if dispatchConfig.JSON {
		encoder := json.NewEncoder(out)
		encoder.SetIndent("", "    ")
		return encoder.Encode(o)
	}
	return nil
}

// NewCmdCreate creates a command object for the "create" action.
// Currently, one must use subcommands for specific resources to be created.
// In future create should accept file or stdin with multiple resources specifications.
// TODO: add create command implementation
func NewCmdCreate(out io.Writer, errOut io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Short:   i18n.T("Create resources."),
		Long:    createLong,
		Example: createExample,
		Run: func(cmd *cobra.Command, args []string) {
			if file == "" {
				runHelp(cmd, args)
				return
			}

			createMap := map[string]modelAction{
				utils.ImageKind:           CallCreateImage,
				utils.BaseImageKind:       CallCreateBaseImage,
				utils.FunctionKind:        CallCreateFunction,
				utils.SecretKind:          CallCreateSecret,
				utils.ServiceInstanceKind: CallCreateServiceInstance,
			}

			err := importFile(out, errOut, cmd, args, createMap)
			CheckErr(err)
		},
	}

	cmd.Flags().StringVarP(&cmdFlagApplication, "application", "a", "", "associate with an application")
	cmd.Flags().StringVarP(&file, "file", "f", "", "Path to YAML file")
	cmd.Flags().StringVarP(&workDir, "work-dir", "w", "", "Working directory relative paths are based on")

	cmd.AddCommand(NewCmdCreateBaseImage(out, errOut))
	cmd.AddCommand(NewCmdCreateImage(out, errOut))
	cmd.AddCommand(NewCmdCreateFunction(out, errOut))
	cmd.AddCommand(NewCmdCreateSecret(out, errOut))
	cmd.AddCommand(NewCmdCreateAPI(out, errOut))
	cmd.AddCommand(NewCmdCreateSubscription(out, errOut))
	cmd.AddCommand(NewCmdCreateEventDriver(out, errOut))
	cmd.AddCommand(NewCmdCreateEventDriverType(out, errOut))
	cmd.AddCommand(NewCmdCreateApplication(out, errOut))
	cmd.AddCommand(NewCmdCreateServiceInstance(out, errOut))
	return cmd
}
