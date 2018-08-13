package stub

import (
	"context"
	"fmt"

	"github.com/surajssd/operators/using-operator-sdk/yoyo-operator/pkg/apis/yoyo/v1alpha1"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewHandler() sdk.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1alpha1.Yoyo:
		fmt.Printf("CR: [%s]\n", o.Name)
		if o.Spec.YoyoDemo {
			o.Status.Status = "We have a TRUE configuration"
			fmt.Println("Hey we have our boolean set to TRUE!")
		} else {
			o.Status.Status = "We have a FALSE configuration"
			fmt.Println("The boolean is set to FALSE")
		}
		if err := sdk.Update(o); err != nil {
			fmt.Printf("[error]: %v\n", err)
		}
		fmt.Printf("We received a Yoyo: %#v\n\n", o)
	case *v1alpha1.YoyoList:
		fmt.Printf("We received a YoyoList: %#v\n\n", o)
	}
	return nil
}

// newbusyBoxPod demonstrates how to create a busybox pod
func newbusyBoxPod(cr *v1alpha1.Yoyo) *corev1.Pod {
	labels := map[string]string{
		"app": "busy-box",
	}
	return &corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "busy-box",
			Namespace: cr.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cr, schema.GroupVersionKind{
					Group:   v1alpha1.SchemeGroupVersion.Group,
					Version: v1alpha1.SchemeGroupVersion.Version,
					Kind:    "Yoyo",
				}),
			},
			Labels: labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}
}
