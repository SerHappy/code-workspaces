BINARY := cw
CMD_DIR := ./cmd/cw

BINDIR ?= $(HOME)/go/bin

build:
	go build -o $(BINARY) $(CMD_DIR)

install:
	mkdir -p "$(BINDIR)"
	go build -o "$(BINDIR)/$(BINARY)" "$(CMD_DIR)"

uninstall:
	rm -f "$(BINDIR)/$(BINARY)"