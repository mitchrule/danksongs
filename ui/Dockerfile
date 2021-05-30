FROM node

RUN mkdir /application
ADD . /application
WORKDIR /application
RUN npm install
RUN npx next build
EXPOSE 3000
CMD ["npx", "next", "start"]