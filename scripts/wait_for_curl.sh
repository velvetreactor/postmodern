#!/usr/bin/env sh

until curl $1; do
  echo "Server not ready, repeating..."
  sleep 2
done
eval $2
