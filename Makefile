SCOW_REF = branch=master
# SCOW_REF = tag=v0.4.0

protos: 
	buf generate --template buf.gen.yaml https://github.com/PKUHPC/SCOW.git#subdir=protos,${SCOW_REF}

run-api: FORCE
	go run *.go api

run-hook: FORCE
	go run *.go hook

FORCE: ;