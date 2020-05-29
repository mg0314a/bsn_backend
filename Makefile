.PHONY:

proto:
	cd api && go generate

deps:proto
	cd internal/dao && go generate
	cd internal/di && go generate
