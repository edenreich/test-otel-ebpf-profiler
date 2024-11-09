## Running OpenTelemetry eBPF Profiler in Kubernetes

1. task create-cluster
2. task deploy-otel-collector
3. task build-ebpf-profiler (necessary to build it ourselves, see: https://github.com/open-telemetry/opentelemetry-ebpf-profiler/issues/43)
4. task deploy-ebpf-profiler
