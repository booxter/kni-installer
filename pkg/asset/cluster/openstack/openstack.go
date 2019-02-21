// Package openstack extracts OpenStack metadata from install
// configurations.
package openstack

import (
	"github.com/metalkube/kni-installer/pkg/types"
	"github.com/metalkube/kni-installer/pkg/types/openstack"
)

// Metadata converts an install configuration to OpenStack metadata.
func Metadata(clusterID string, config *types.InstallConfig) *openstack.Metadata {
	return &openstack.Metadata{
		Region: config.Platform.OpenStack.Region,
		Cloud:  config.Platform.OpenStack.Cloud,
		Identifier: map[string]string{
			"openshiftClusterID": clusterID,
		},
	}
}
