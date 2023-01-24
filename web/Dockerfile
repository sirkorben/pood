FROM node:latest as builder

COPY package.json package-lock.json ./

RUN npm install && mkdir /react-ui && mv ./node_modules ./react-ui

WORKDIR /react-ui

COPY . .

RUN npm run build

#!/bin/sh instead of /bin/bash
FROM nginx:alpine
COPY nginx.conf /etc/nginx/nginx.conf

RUN rm -rf /usr/share/nginx/html/*

COPY --from=builder /react-ui/build /usr/share/nginx/html

EXPOSE 3000 80

ENTRYPOINT ["nginx", "-g", "daemon off;"]


# THIS PART WORKS with `npm run build` done before running up nginx container and placing build folder inside
# FROM nginx:alpine
# COPY nginx.conf /etc/nginx/nginx.conf
# RUN rm -rf /usr/share/nginx/html/*
# COPY /build /usr/share/nginx/html
# EXPOSE 3000 80
# ENTRYPOINT ["nginx", "-g", "daemon off;"]