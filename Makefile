BIN=bin/sidupaste

SRC = main.go \
      sidupaste/sidupaste.go \
      go.sum \
      go.mod

all: build man

build: $(BIN)

man:
	mkdir -p man
	pandoc --standalone --to man doc/sidupaste.1.md -o man/sidupaste.1
	pandoc --standalone --to man doc/sidupaste.conf.5.md -o man/sidupaste.conf.5

install: build man
	install -m 555 bin/sidupaste /usr/local/bin/sidupaste
	cp man/sidupaste.1 /usr/local/man/man1/sidupaste.1
	cp man/sidupaste.conf.5 /usr/local/man/man5/sidupaste.conf.5

deinstall:
	rm -f /usr/local/bin/sidupaste
	rm -f /usr/local/man/man1/sidupaste.1
	rm -f /usr/local/man/man5/sidupaste.conf.5

clean:
	rm -rf bin
	rm -rf man

$(BIN): $(SRC)
	go build -o bin/sidupaste

.PHONY: all man build clean
