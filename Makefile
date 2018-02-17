RACEPATH = "log_path=./report strip_path_prefix=pkg/"

all: scene 

scene:
	cd bin && GOPATH=$(PWD) go build -gcflags "-N -l" scene 
