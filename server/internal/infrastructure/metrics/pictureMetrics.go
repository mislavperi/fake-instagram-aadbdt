package metrics

import "github.com/prometheus/client_golang/prometheus"

type PictureMetrics struct {
	pictureUploadDuration prometheus.Summary

	pictureUploadCounter prometheus.Counter
}

func NewPictureMetrics(namespace string) *PictureMetrics {
	pictureUploadDuration := prometheus.NewSummary(
		prometheus.SummaryOpts{
			Name: "picture_upload_duration",
		},
	)

	pictureUploadCounter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "picture_upload_counter",
		},
	)

	return &PictureMetrics{
		pictureUploadDuration: pictureUploadDuration,
		pictureUploadCounter:  pictureUploadCounter,
	}
}

func (m *PictureMetrics) OnUploadStart() *prometheus.Timer {
	m.pictureUploadCounter.Inc()
	return prometheus.NewTimer(m.pictureUploadDuration)
}

func (m *PictureMetrics) OnUploadFinish(timer *prometheus.Timer) {
	timer.ObserveDuration()
}
