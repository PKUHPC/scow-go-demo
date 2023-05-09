protos: 
	buf generate --template buf.gen.yaml https://github.com/PKUHPC/SCOW.git#subdir=protos,tag=api-${SCOW_API_VERSION}

run-api: FORCE
	go run *.go api

run-hook: FORCE
	go run *.go hook

FORCE: ;