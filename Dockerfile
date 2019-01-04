FROM alpine

RUN apk update
RUN apk add python py2-pip less groff
RUN apk add go make git
RUN apk add musl-dev
RUN pip install awscli

# We can use this to package our go lambdas.
COPY src/ /src/
