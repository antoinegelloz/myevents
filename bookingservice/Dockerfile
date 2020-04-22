FROM scratch
 
COPY bookingservice /
ADD certificate /certificate
 
ENV LISTEN_URL=0.0.0.0:8282
EXPOSE 8282

CMD ["./bookingservice"]