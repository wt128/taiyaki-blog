FROM node:18.14.0

USER node

WORKDIR /usr/src/app
RUN mkdir node_modules

CMD ["/bin/bash", "-c", "yarn install --frozen-lockfile && yarn dev"]
EXPOSE 5173
