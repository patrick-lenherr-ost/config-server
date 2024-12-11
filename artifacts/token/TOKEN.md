kubectl port-forward service/config-server-metrics 9443:8443 -n network-system

TOKEN=$(kubectl get secret config-server-token -n network-system -o jsonpath='{.data.token}' | base64 --decode)

curl -k -H "Authorization: Bearer $TOKEN" https://localhost:9443/metrics