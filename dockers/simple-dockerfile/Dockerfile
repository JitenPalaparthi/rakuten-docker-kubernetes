FROM ubuntu:latest

LABEL "author"="rakuten"

LABEL "purpose"="demo"

ENV APP_PORT 9999

WORKDIR /data

COPY file1.txt .

COPY file2.txt .

#ENTRYPOINT ["cat","file1.txt"]

RUN  apt-get update -y && \
     apt-get install vim -y
      
CMD [ "/bin/ls", "-l"]

#CMD ["/bin/sh" ,"while true;do sleep 3600; done"]



