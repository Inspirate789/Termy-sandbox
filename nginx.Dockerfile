FROM byjg/nginx-extras:1.21

RUN mkdir -p /static/img
RUN mkdir -p /etc/nginx/html/documentation/docs
COPY frontend_old/index.html /static
COPY frontend_old/authorization-window.png /static/img
COPY README.md /etc/nginx/html/documentation/readme.md
COPY docs/img /etc/nginx/html/documentation/docs/img
COPY scripts/md-renderer.html /md/md-renderer.html

# Copying nginx.conf
COPY nginx/test.nginx.conf /etc/nginx/conf.d/default.conf
