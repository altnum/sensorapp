FROM grafana/grafana-enterprise:8.2.1

VOLUME  ["/var/lib/grafana" "$PWD/data:/var/lib/grafana"]

ENV GF_SECURITY_ADMIN_USER=admin
ENV GF_AUTH_GOOGLE_CLIENT_SECRET=admin

EXPOSE 3000