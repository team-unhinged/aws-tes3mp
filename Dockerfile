FROM alpine

RUN apk update
RUN apk add python py2-pip
RUN pip install awscli