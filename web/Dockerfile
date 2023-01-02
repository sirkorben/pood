# FROM node:17-alpine as builder

# WORKDIR /app

# COPY package.json .
# COPY yarn.lock .
# RUN yarn install
# COPY . .
# RUN yarn build

# FROM nginx:1.19.0
# WORKDIR /usr/share/nginx/html
# RUN rm-rf ./*
# COPY --from=builder /app/build .

# CMD npm run build && npm run start

# EXPOSE 3000

FROM node:slim

WORKDIR /app

COPY . .

RUN yarn install

CMD ["npm", "start"]

EXPOSE 3000