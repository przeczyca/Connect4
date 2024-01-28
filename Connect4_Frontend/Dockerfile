FROM node:20-alpine3.19

COPY . /app

WORKDIR /app

EXPOSE 3000

CMD npm i && npm run build && npx serve dist