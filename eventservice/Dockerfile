FROM scratch
 
COPY eventservice /
ADD certificate /certificate
 
ENV LISTEN_URL=0.0.0.0:8181
EXPOSE 8181 

CMD ["./eventservice"]
