SCOW_VERSION = v0.4.0

protos: 
	buf generate --template buf.gen.yaml https://github.com/PKUHPC/SCOW.git#tag=${SCOW_VERSION}

FORCE: ;