package kubernetes

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
	"strconv"
)

var (
	activeDeployment = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "startupheroes_deployment_active",
			Help: "Deployment info for project and environment.",
		},
		[]string{"app_family", "namespace", "user", "image", "timestamp"},
	)
)

func init() {
	prometheus.MustRegister(activeDeployment)
}
func addPrometheusEvent(appFamily string, namespace string, user string, image string) {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	timestampString := strconv.Itoa(int(timestamp))
	activeDeployment.WithLabelValues(appFamily, namespace, user, image, timestampString).Set(1)
	go time.AfterFunc(45 * time.Second, func() {
		activeDeployment.WithLabelValues(appFamily, namespace, user, image, timestampString).Set(0)
	})
}
