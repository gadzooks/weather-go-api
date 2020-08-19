FROM node:10.11.0-alpine

RUN npm install -g newman newman-reporter-html

WORKDIR /etc/newman

COPY places-API.postman_collection.json .

#ENTRYPOINT ["newman"]
CMD ["newman", "run", "places-API.postman_collection.jso"]
