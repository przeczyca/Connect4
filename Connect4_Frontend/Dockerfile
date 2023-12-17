FROM node:20-alpine3.19
WORKDIR /app

# Run "npm run build" first so the "dist" folder exists
COPY dist dist/

EXPOSE 3000

CMD ["npx", "serve", "dist"]