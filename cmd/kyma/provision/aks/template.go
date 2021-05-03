package aks

import (
	"errors"
	"fmt"
	"strings"

	"github.com/kyma-incubator/hydroform/provision/types"
	"github.com/kyma-project/cli/internal/cli"
)

type AksCmd struct {
	opts *Options
	cli.Command
}

func (p *AksCmd) NewCluster() *types.Cluster {
	return &types.Cluster{
		Name:              p.opts.Name,
		KubernetesVersion: p.opts.KubernetesVersion,
		DiskSizeGB:        p.opts.DiskSizeGB,
		NodeCount:         p.opts.NodeCount,
		Location:          p.opts.Location,
		MachineType:       p.opts.MachineType,
	}
}

func (prov *AksCmd) NewProvider() (*types.Provider, error) {
	p := &types.Provider{
		Type:                types.Azure,
		ProjectName:         prov.opts.Project,
		CredentialsFilePath: prov.opts.CredentialsFile,
	}

	p.CustomConfigurations = make(map[string]interface{})
	for _, e := range prov.opts.Extra {
		v := strings.Split(e, "=")

		if len(v) != 2 {
			return p, fmt.Errorf("wrong format for extra configuration %s, please provide NAME=VALUE pairs", e)
		}
		p.CustomConfigurations[v[0]] = v[1]
	}
	return p, nil
}

func (p *AksCmd) ProviderName() string { return "AKS" }

func (p *AksCmd) Attempts() uint { return p.opts.Attempts }

func (p *AksCmd) KubeconfigPath() string { return p.opts.KubeconfigPath }

func (p *AksCmd) ValidateFlags() error {
	var errMessage strings.Builder
	// mandatory flags]
	if p.opts.Name == "" {
		errMessage.WriteString("\nRequired flag `name` has not been set.")
	}
	if p.opts.Project == "" {
		errMessage.WriteString("\nRequired flag `project` has not been set.")
	}
	if p.opts.CredentialsFile == "" {
		errMessage.WriteString("\nRequired flag `credentials` has not been set.")
	}

	if errMessage.Len() != 0 {
		return errors.New(errMessage.String())
	}
	return nil
}

func (p *AksCmd) IsVerbose() bool { return p.opts.Verbose }
