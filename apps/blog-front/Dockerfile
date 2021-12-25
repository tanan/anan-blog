FROM node:14-alpine as build-stage

ADD . /usr/src/blog-front
WORKDIR /usr/src/blog-front
RUN yarn
RUN yarn build

FROM nginx:latest as production-stage
COPY --from=build-stage /usr/src/blog-front/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf
COPY default.conf /etc/nginx/conf.d/default.conf