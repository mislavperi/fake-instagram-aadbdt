package metrics

import "github.com/prometheus/client_golang/prometheus"

type UserMetrics struct {
	userCreationDuration prometheus.SummaryVec
	userLoginDuration    prometheus.SummaryVec

	userLoginCounter prometheus.CounterVec
}

func NewUserMetrics(namespace string) *UserMetrics {
	userCreationDuration := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:      "user_creation_duration",
			Namespace: namespace,
		},
		[]string{"creation_type"},
	)

	userLoginDuration := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:      "user_login_duration",
			Namespace: namespace,
		},
		[]string{"login_type"},
	)

	userLoginCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "user_login_counter",
		},
		[]string{"login_type"},
	)

	return &UserMetrics{
		userLoginDuration:    *userLoginDuration,
		userCreationDuration: *userCreationDuration,
		userLoginCounter:     *userLoginCounter,
	}
}

func (m *UserMetrics) OnLoginStart(label string) *prometheus.Timer {
	m.userLoginCounter.WithLabelValues(label).Inc()
	return prometheus.NewTimer(m.userLoginDuration.WithLabelValues(label))
}

func (m *UserMetrics) OnLoginFinish(timer *prometheus.Timer) {
	timer.ObserveDuration()
}

func (m *UserMetrics) OnCreationStart(label string) *prometheus.Timer {
	return prometheus.NewTimer(m.userCreationDuration.WithLabelValues(label))
}

func (m *UserMetrics) OnCreationFinish(timer *prometheus.Timer) {
	timer.ObserveDuration()
}
