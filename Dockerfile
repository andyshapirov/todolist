FROM ubuntu:latest

WORKDIR /app

COPY todolist ./  
COPY web ./web/

EXPOSE 7540

CMD ["./todolist"]