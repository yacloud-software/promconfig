.PHONY:	client server
all:	client server
client:
	cd client && go install ${LDFLAGS} promconfig-client.go `ls -1 *.go|grep -v promconfig-client.go`
server:
	cd server && go install ${LDFLAGS} promconfig-server.go `ls -1 *.go|grep -v promconfig-server.go`
