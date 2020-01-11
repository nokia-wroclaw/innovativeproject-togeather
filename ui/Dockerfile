## STAGE 1: Build Angular app ##
FROM node:10.18-jessie-slim as builder

WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install and cache app dependencies
COPY package.json /app/package.json
COPY package-lock.json /app/package-lock.json
RUN npm ci

COPY . /app

# generate build
RUN ng build --output-path=dist --aot=true --prod


## STAGE 2: Run nginx to serve the app ##
FROM nginx

COPY --from=builder /app/dist /usr/share/nginx/html

# expose app container on port 80
EXPOSE 80

# run nginx
 CMD ["nginx", "-g", "daemon off;"]
