package sidecar

import (
	"github.com/scylladb/scylla-operator/pkg/naming"
	v1 "k8s.io/api/core/v1"
)

type Metadata struct {
	Members []*MemberMetadata `json:"members"`
}

type MemberMetadata struct {
	datacenter string   `json:"datacenter"`
	rack       string   `json:"rack"`
	clusterIP  string   `json:"clusterIP"`
	podIP      []string `json:"podIP"`
}

func NewMemberMetadata(service *v1.Service, pods []*v1.Pod) *MemberMetadata {
	ips := make([]string, len(pods))
	for i, pod := range pods {
		ips[i] = pod.Status.PodIP
	}
	return &MemberMetadata{
		datacenter: service.Labels[naming.DatacenterNameLabel],
		rack:       service.Labels[naming.RackNameLabel],
		clusterIP:  service.Spec.ClusterIP,
		podIP:      ips,
	}
}
