.PHONY: build check gen-authors

build:
	go build github.com/mrostecki/gomountinfo

check:
	go test -cover -v ./...

gen-authors:
	out="`git log --pretty=format:'%aN <%aE>' | sort -u`" && \
	perl -p -e "s/#authorslist#// and print '$$out'" \
	< AUTHORS.in > AUTHORS-tmp && mv -f AUTHORS-tmp AUTHORS
