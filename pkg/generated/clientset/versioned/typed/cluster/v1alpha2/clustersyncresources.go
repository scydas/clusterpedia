// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	"context"

	v1alpha2 "github.com/clusterpedia-io/api/cluster/v1alpha2"
	scheme "github.com/clusterpedia-io/clusterpedia/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ClusterSyncResourcesGetter has a method to return a ClusterSyncResourcesInterface.
// A group's client should implement this interface.
type ClusterSyncResourcesGetter interface {
	ClusterSyncResources() ClusterSyncResourcesInterface
}

// ClusterSyncResourcesInterface has methods to work with ClusterSyncResources resources.
type ClusterSyncResourcesInterface interface {
	Create(ctx context.Context, clusterSyncResources *v1alpha2.ClusterSyncResources, opts v1.CreateOptions) (*v1alpha2.ClusterSyncResources, error)
	Update(ctx context.Context, clusterSyncResources *v1alpha2.ClusterSyncResources, opts v1.UpdateOptions) (*v1alpha2.ClusterSyncResources, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.ClusterSyncResources, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.ClusterSyncResourcesList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.ClusterSyncResources, err error)
	ClusterSyncResourcesExpansion
}

// clusterSyncResources implements ClusterSyncResourcesInterface
type clusterSyncResources struct {
	*gentype.ClientWithList[*v1alpha2.ClusterSyncResources, *v1alpha2.ClusterSyncResourcesList]
}

// newClusterSyncResources returns a ClusterSyncResources
func newClusterSyncResources(c *ClusterV1alpha2Client) *clusterSyncResources {
	return &clusterSyncResources{
		gentype.NewClientWithList[*v1alpha2.ClusterSyncResources, *v1alpha2.ClusterSyncResourcesList](
			"clustersyncresources",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1alpha2.ClusterSyncResources { return &v1alpha2.ClusterSyncResources{} },
			func() *v1alpha2.ClusterSyncResourcesList { return &v1alpha2.ClusterSyncResourcesList{} }),
	}
}
