FROM debian:jessie 
 
COPY bookingService /bookingService 
RUN  useradd bookingService 
USER bookingService 
 
ENV LISTEN_URL=0.0.0.0:8181 
EXPOSE 8181 
CMD ["/bookingService"]
