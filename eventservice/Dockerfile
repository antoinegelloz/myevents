FROM debian:jessie 
 
COPY eventsService /eventsService 
RUN  useradd eventsService 
USER eventsService 
 
ENV LISTEN_URL=0.0.0.0:8181 
EXPOSE 8181 
CMD ["/eventsService"]
