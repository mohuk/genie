# Stage 0, "build-stage", based on Node.js, to build and compile Angular
FROM mohuk/angular-cli:7.3.3 as build-stage

WORKDIR /app

COPY package*.json /app/

RUN npm install

COPY ./ /app/

RUN npm run build -- --output-path=./dist/out


# Stage 1, based on Nginx, to have only the compiled app, ready for production with Nginx
FROM nginx:1.15

COPY --from=build-stage /app/dist/out/ /usr/share/nginx/html

# COPY ./nginx-custom.conf /etc/nginx/conf.d/default.conf