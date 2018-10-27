update-vendor:
	dep ensure
	dep ensure -update

release:
	goreleaser --rm-dist
