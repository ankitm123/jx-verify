package pods_test

import (
	"context"
	"github.com/jenkins-x/jx-verify/pkg/cmd/pods"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"testing"
)

func TestPods(t *testing.T) {
	ns := "jx"
	podName := "my-pod"

	_, o := pods.NewCmdVerifyPods()
	kubeClient := fake.NewSimpleClientset(&v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName,
			Namespace: ns,
		},
	})
	o.KubeClient = kubeClient
	o.Namespace = ns

	podInterface := kubeClient.CoreV1().Pods(ns)

	ctx := context.TODO()

	RequirePodCount(t, ctx, podInterface, 1)

	o.OnEvent(&v1.Event{
		InvolvedObject: v1.ObjectReference{
			Kind: "ConfigMap",
			Name: "cheese",
		},
		Message: "",
	}, ns)

	RequirePodCount(t, ctx, podInterface, 1)

	o.OnEvent(&v1.Event{
		InvolvedObject: v1.ObjectReference{
			Kind:      "Pod",
			Name:      podName,
			Namespace: ns,
		},
		Message: pods.ErrImagePullMessage,
	}, ns)

	RequirePodCount(t, ctx, podInterface, 0)
}

// RequirePodCount requires the given number of pods to exist
func RequirePodCount(t *testing.T, ctx context.Context, podInterface corev1.PodInterface, expectedLen int) {
	podList, err := podInterface.List(ctx, metav1.ListOptions{})
	require.NoError(t, err, "failed to list pods")
	require.NotNil(t, podList, "no PodList returned")

	require.Len(t, podList.Items, expectedLen, "expected PodList.Items count")
	//t.Logf("now has %d pods\n", len(podList.Items))
}
