# Build stage
FROM node:22-alpine3.22 AS builder

WORKDIR /app

COPY package.json .
COPY package-lock.json .
COPY tsconfig.json .
COPY ecosystem.config.js .
COPY .env .
COPY ./src ./src
COPY ./prisma ./prisma

RUN npm install
RUN npx prisma generate
RUN npm run build

# Production stage
FROM node:23-alpine3.22 as production

WORKDIR /app

COPY package.json .
COPY package-lock.json .
COPY ecosystem.config.js .
COPY .env .
COPY ./prisma ./prisma
COPY --from=builder /app/dist ./dist
COPY ./src/app/domain/proto /app/src/app/domain/proto
COPY ./src/public /app/src/public
RUN npm install --production
RUN npm install pm2 -g

EXPOSE 8080

CMD ["pm2-runtime", "start", "ecosystem.config.js", "--env", "production"]
