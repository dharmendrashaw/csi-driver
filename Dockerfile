FROM alpine

#TODO: Change to buiding project here
COPY csi-driver /usr/local/bin/csi-driver 

ENTRYPOINT [ "/usr/local/bin/csi-driver" ]