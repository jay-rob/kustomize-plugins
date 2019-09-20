// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"errors"
	"fmt"

	"sigs.k8s.io/kustomize/v3/pkg/ifc"
	"sigs.k8s.io/kustomize/v3/pkg/resmap"
	"sigs.k8s.io/kustomize/v3/pkg/types"
	"sigs.k8s.io/yaml"

	"golang.org/x/oauth2/google"
	iam "google.golang.org/api/iam/v1"
)

type plugin struct {
	rf               *resmap.Factory
	ldr              ifc.Loader
	types.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	GcpProjectId     string `json:"gcpProjectId,omitempty" file:"gcpProjectId,omitempty"`
	ServiceAccount   string `json:"serviceAccount,omitempty" file:"serviceAccount,omitempty"`
}

//noinspection GoUnusedGlobalVariable
//nolint: golint
var KustomizePlugin plugin

func (p *plugin) Config(ldr ifc.Loader, rf *resmap.Factory, c []byte) error {
	p.rf = rf
	p.ldr = ldr
	return yaml.Unmarshal(c, p)
}

// The plan here is to convert the plugin's input
// into the format used by the builtin secret generator plugin.
func (p *plugin) Generate() (resmap.ResMap, error) {
	args := types.SecretArgs{}
	args.Name = p.Name
	args.Namespace = p.Namespace

	keys, err := listKeys(p.GcpProjectId, p.ServiceAccount)
	if err != nil {
		return nil, err
	}

	if len(keys) < 1 {
		return nil, errors.New("No key provisioned for account")
	}

	key, err := keys[0].MarshalJSON()
	if err != nil {
		return nil, err
	}

	args.LiteralSources = []string{
		fmt.Sprintf("credentials.json=%s", key),
	}

	return p.rf.FromSecretArgs(p.ldr, nil, args)
}

// listKey lists a service account's keys.
func listKeys(project, serviceAccountEmail string) ([]*iam.ServiceAccountKey, error) {
	client, err := google.DefaultClient(context.Background(), iam.CloudPlatformScope)
	if err != nil {
		return nil, fmt.Errorf("google.DefaultClient: %v", err)
	}
	service, err := iam.New(client)
	if err != nil {
		return nil, fmt.Errorf("iam.New: %v", err)
	}

	resource := fmt.Sprintf("projects/%v/serviceAccounts/%v", project, serviceAccountEmail)
	response, err := service.Projects.ServiceAccounts.Keys.List(resource).Do()
	if err != nil {
		return nil, fmt.Errorf("Projects.ServiceAccounts.Keys.List: %v", err)
	}

	return response.Keys, nil
}
