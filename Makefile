run.app:
	go build
	mv ./aicup2020 ./bin
	export LOCAL_DEBUG=true && \
	./bin/aicup2020