package handlers

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/coredns/coredns/plugin/metrics"
	"github.com/miekg/dns"
	"golang.org/x/net/context"
)

type CorednsMetricsWrapper struct {
	metrics *metrics.Metrics
	logger  boshlog.Logger
	logTag  string
}

type metricsContext struct {
	withMetrics bool
}

func NewCorednsMetricsWrapper(next dns.Handler, logger boshlog.Logger) CorednsMetricsWrapper {
	metrics := metrics.New("127.0.0.1:53088")
	metrics.Next = corednsMetricsHandlerWrapper{Next: next}
	return CorednsMetricsWrapper{
		metrics: metrics,
		logTag:  "DNSMetricsHandler",
		logger:  logger,
	}
}

func (m CorednsMetricsWrapper) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	indicator := &metricsContext{
		withMetrics: true,
	}
	requestContext := context.WithValue(context.Background(), "indicator", indicator)
	_, err := m.metrics.ServeDNS(requestContext, w, r)
	if err != nil {
		m.logger.Error(m.logTag, "Error getting dns metrics:", err.Error())
	}
}

func (m CorednsMetricsWrapper) Run() error {
	return m.metrics.OnStartup()
}

type corednsMetricsHandlerWrapper struct {
	Next dns.Handler
}

func (m corednsMetricsHandlerWrapper) ServeDNS(ctx context.Context, writer dns.ResponseWriter, msg *dns.Msg) (int, error) {
	requestContext := ctx.Value("indicator").(*metricsContext)
	requestContext.withMetrics = false
	m.Next.ServeDNS(writer, msg)
	return 0, nil
}

func (m corednsMetricsHandlerWrapper) Name() string {
	return "CorednsMetricsHandlerWrapper"
}
