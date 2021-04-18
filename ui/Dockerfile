FROM node

RUN mkdir /application
ADD . /application
WORKDIR /application
RUN npm install
EXPOSE 3000
CMD ["npm", "start"]