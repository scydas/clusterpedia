package scheme

import (
	// These imports are the API groups the API server will support.
	apiextensionsinstall "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/install"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	apiregistrationinstall "k8s.io/kube-aggregator/pkg/apis/apiregistration/install"
	_ "k8s.io/kubernetes/pkg/apis/admission/install"
	_ "k8s.io/kubernetes/pkg/apis/admissionregistration/install"
	_ "k8s.io/kubernetes/pkg/apis/apiserverinternal/install"
	_ "k8s.io/kubernetes/pkg/apis/apps/install"
	_ "k8s.io/kubernetes/pkg/apis/authentication/install"
	_ "k8s.io/kubernetes/pkg/apis/authorization/install"
	_ "k8s.io/kubernetes/pkg/apis/autoscaling/install"
	_ "k8s.io/kubernetes/pkg/apis/batch/install"
	_ "k8s.io/kubernetes/pkg/apis/certificates/install"
	_ "k8s.io/kubernetes/pkg/apis/coordination/install"
	_ "k8s.io/kubernetes/pkg/apis/core/install"
	_ "k8s.io/kubernetes/pkg/apis/discovery/install"
	v1 "k8s.io/kubernetes/pkg/apis/discovery/v1"
	"k8s.io/kubernetes/pkg/apis/discovery/v1beta1"
	_ "k8s.io/kubernetes/pkg/apis/events/install"
	_ "k8s.io/kubernetes/pkg/apis/extensions/install"
	_ "k8s.io/kubernetes/pkg/apis/flowcontrol/install"
	_ "k8s.io/kubernetes/pkg/apis/imagepolicy/install"
	_ "k8s.io/kubernetes/pkg/apis/networking/install"
	_ "k8s.io/kubernetes/pkg/apis/node/install"
	_ "k8s.io/kubernetes/pkg/apis/policy/install"
	_ "k8s.io/kubernetes/pkg/apis/rbac/install"
	_ "k8s.io/kubernetes/pkg/apis/resource/install"
	_ "k8s.io/kubernetes/pkg/apis/scheduling/install"
	_ "k8s.io/kubernetes/pkg/apis/storage/install"
	_ "k8s.io/kubernetes/pkg/apis/storagemigration/install"
)

func init() {
	apiextensionsinstall.Install(LegacyResourceScheme)
	apiregistrationinstall.Install(LegacyResourceScheme)

	// Fix the version order of 'discovery.k8s.io' resources
	utilruntime.Must(LegacyResourceScheme.SetVersionPriority(v1.SchemeGroupVersion, v1beta1.SchemeGroupVersion))
}
