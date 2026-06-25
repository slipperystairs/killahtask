.PHONY: killahtask
killahtask:
	podman run --rm -ti -v .:/app -w /app docker.io/golang go build .

.PHONY: install
install:
	install killahtask /usr/local/bin

.PHONY: clean
clean:
	rm -f killahtask

.PHONY: dist-clean
dist-clean:
	rm -f /usr/local/bin/killahtask
