FROM scratch

COPY app-static .
COPY gophers.png .
COPY index.html .
COPY version .
COPY cacert.pem /etc/ssl/certs/ca-certificates.crt

ENV APP_PORT=80
ENTRYPOINT ["./app-static"]
EXPOSE ${APP_PORT}
