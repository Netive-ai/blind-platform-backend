FROM mhart/alpine-node:12
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci --prod


FROM mhart/alpine-node:12
WORKDIR /app
COPY --from=0 /app .
COPY ./src .

EXPOSE 8080
CMD ["node", "server.js"]