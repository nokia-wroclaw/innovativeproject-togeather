## STAGE 1: Build Angular app ##
FROM node:latest as builder

COPY togeather /togeather

WORKDIR /togeather

RUN npm install
RUN $(npm bin)/ng build


## STAGE 2: Run nginx to serve the app ##
FROM nginx

COPY --from=builder /togeather/dist/* /usr/share/nginx/html/

EXPOSE 80
