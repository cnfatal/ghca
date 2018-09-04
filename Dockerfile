FROM scratch
COPY ghca /
EXPOSE 8080
ENTRYPOINT ["/ghca"]
CMD ["--server"]