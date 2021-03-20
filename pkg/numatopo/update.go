package numatopo

import (
	"context"
	"os"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"

	"volcano.sh/noderestopo/pkg/apis/nodeinfo/v1alpha1"
	"volcano.sh/noderestopo/pkg/args"
	"volcano.sh/noderestopo/pkg/client/clientset/versioned"
)

func NodeInfoRefresh(opt *args.Argument) bool {
	isChange := false

	config := GetkubeletConfig(opt.KubeletConf)
	if config != nil {
		isChange = true
	}

	if TopoInfoUpdate(opt) {
		isChange = true
	}

	return isChange
}

func CreateOrUpdateNumatopo(client *versioned.Clientset) {
	hostname := os.Getenv("MY_NODE_NAME")
	if hostname == "" {
		klog.Errorf("get Hostname failed.")
		return
	}

	numaInfo, err := client.NodeinfoV1alpha1().Numatopos("default").Get(context.TODO(), hostname, metav1.GetOptions{})
	if err != nil {
		if !apierrors.IsNotFound(err) {
			klog.Errorf("get Numatopo for node %s failed, err=%v", hostname, err)
			return
		}

		numaInfo = &v1alpha1.Numatopo{
			ObjectMeta: metav1.ObjectMeta{
				Name: hostname,
			},
			Spec: v1alpha1.NumatopoSpec{
				Policies:   GetPolicy(),
				NumaResMap: GetAllResTopoInfo(),
			},
		}

		_, err = client.NodeinfoV1alpha1().Numatopos("default").Create(context.TODO(), numaInfo, metav1.CreateOptions{})
		if err != nil {
			klog.Errorf("create Numatopo for node %s failed, err=%v", hostname, err)
		}
	} else {
		numaInfo.Spec = v1alpha1.NumatopoSpec{
			Policies:   GetPolicy(),
			NumaResMap: GetAllResTopoInfo(),
		}
		_, err = client.NodeinfoV1alpha1().Numatopos("default").Update(context.TODO(), numaInfo, metav1.UpdateOptions{})
		if err != nil {
			klog.Errorf("update Numatopo for node %s failed, err=%v", hostname, err)
		}
	}
}
