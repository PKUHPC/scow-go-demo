SCOW_REF = \#branch=master
# SCOW_REF = \#tag=v0.4.0

protos: 
	buf generate --template buf.gen.yaml https://github.com/PKUHPC/SCOW.git${SCOW_REF}

FORCE: ;