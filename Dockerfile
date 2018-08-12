FROM jgsqware/go-glide:1.8 as builder

ARG SERVICE

ARG PROJECT

RUN mkdir -p ${PROJECT}

WORKDIR ${PROJECT}

COPY . .

RUN glide install

RUN make ${SERVICE}

FROM alpine:latest

ARG SERVICE

ARG PROJECT

WORKDIR /root

COPY --from=builder ${PROJECT}/${SERVICE} .

# TODO: create user for running application in non-root way

CMD ./${SERVICE}
