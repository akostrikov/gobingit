FROM ubuntu:16.04

RUN apt-get update; apt-get install -y git

#ADD . /opt/git

CMD ["/opt/git/test.sh"]