package info

import (
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/duration"

	"github.com/codefresh-io/argo-rollouts/pkg/apis/rollouts/v1alpha1"
	"github.com/codefresh-io/argo-rollouts/utils/annotations"
)

const (
	IconWaiting     = "◷"
	IconProgressing = "◌"
	IconWarning     = "⚠"
	IconUnknown     = "?"
	IconOK          = "✔"
	IconBad         = "✖"
	IconPaused      = "॥"
	IconNeutral     = "•"
)

const (
	InfoTagCanary  = "canary"
	InfoTagStable  = "stable"
	InfoTagActive  = "active"
	InfoTagPreview = "preview"
)

type Metadata v1.ObjectMeta

type ImageInfo struct {
	Image string
	Tags  []string
}

func Age(m v1.ObjectMeta) string {
	return duration.HumanDuration(metav1.Now().Sub(m.CreationTimestamp.Time))
}

func ownerRef(ownerRefs []metav1.OwnerReference, uids []types.UID) *metav1.OwnerReference {
	for _, ownerRef := range ownerRefs {
		for _, uid := range uids {
			if ownerRef.UID == uid {
				return &ownerRef
			}
		}
	}
	return nil
}

func parseRevision(annotations_ map[string]string) int {
	if annotations_ != nil {
		if revision, err := strconv.Atoi(annotations_[annotations.RevisionAnnotation]); err == nil {
			return revision
		}
	}
	return 0
}

func parseExperimentTemplateName(annotations_ map[string]string) string {
	if annotations_ != nil {
		return annotations_[v1alpha1.ExperimentTemplateNameAnnotationKey]
	}
	return ""
}

func parseScaleDownDeadline(annotations_ map[string]string) string {
	if annotations_ != nil {
		return annotations_[v1alpha1.DefaultReplicaSetScaleDownDeadlineAnnotationKey]
	}
	return ""
}
