all:
	true

clean:
	rm -rf {models,cmd,vendor,backend}
	find . -type file -path './restapi/*' -not -path './restapi/configure_*.go' -delete
	find . -type d -empty -delete
