package traveler

import (
	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"k8s.io/klog"
)

func IsGPUNode() bool {
	nvmlReturn := nvml.Init()
	if nvmlReturn != nvml.SUCCESS {
		return false
	} else {
		return true
	}
}

func GetGPUs() int {
	nvml.Init()
	defer func() {
		nvmlReturn := nvml.Shutdown()
		if nvmlReturn != nvml.SUCCESS {
			//log.Fatalf("Unable to shutdown NVML: %v", ret)
			klog.Infof("Unable to shutdown NVML: %v\n", nvmlReturn)
		}
	}()
	count, nvmlReturn := nvml.DeviceGetCount()
	if nvmlReturn != nvml.SUCCESS {
		//log.Fatalf("Unable to get device count: %v", ret)
		klog.Infof("Unable to get device count: %v", nvmlReturn)
		count = 0
	}

	return count
}
