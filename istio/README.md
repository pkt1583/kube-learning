istioctl install --set meshConfig.accessLogFile=/dev/stdout --set profile=demo

kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.16/samples/addons/kiali.yaml

kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.16/samples/addons/grafana.yaml

kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.16/samples/addons/prometheus.yaml

kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.16/samples/addons/jaeger.yaml
