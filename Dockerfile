FROM golang:1.7
ADD javbus /
RUN chmod +x /javbus
EXPOSE 8080
CMD ["/javbus"]
