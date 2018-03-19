minkube: eval $(minikube docker-env)

building: habitus  --host=https://192.168.99.100:2376 --use-tls
