package numatopo

import (
	"volcano.sh/noderestopo/pkg/apis/nodeinfo/v1alpha1"
	"volcano.sh/noderestopo/pkg/args"
)

type NumaInfo interface {
	Name() v1alpha1.ResourceName
	Update(opt *args.Argument) NumaInfo
	GetResourceInfoMap() v1alpha1.ResourceInfoMap
}
