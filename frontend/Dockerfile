#FROM node:latest
#
#WORKDIR /usr/src/app
#
#COPY package*.json ./
#
#RUN npm install
#
#COPY . .
#
#EXPOSE 5173
#CMD [ "npm", "run", "dev"]
FROM node:lts-alpine as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# production stage
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]