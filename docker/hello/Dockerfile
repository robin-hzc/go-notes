FROM debian:latest

ADD  hello /var/www/
COPY start.sh /var/www/
RUN  chmod +x /var/www/start.sh
EXPOSE 8081-8090
CMD /var/www/start.sh