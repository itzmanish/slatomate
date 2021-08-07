FROM alpine
ADD slatomate-service /slatomate-service
ENTRYPOINT [ "/slatomate-service" ]
