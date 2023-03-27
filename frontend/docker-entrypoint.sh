#!/bin/sh

# Replace placeholders in Nginx configuration file with environment variables
envsubst '$${PROXY_IP}' < /etc/nginx/default.conf.template > /etc/nginx/conf.d/default.conf

# Start Nginx
nginx -g "daemon off;"