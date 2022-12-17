#!/bin/bash

helm upgrade "server" \
  --install \
  --wait \
  --debug \
  --create-namespace \
  --namespace "test" \
  --set image.repository="uturkarslan/test-go-server-arm" \
  --set image.tag="1.0.0" \
  "../charts/server"
