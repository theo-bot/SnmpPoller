#
# Makefile for SnmpPoller project
#

# Go related variables.
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=SnmpPoller

all: build

build:
  $(GOBUILD) -o $(BINARY_NAME) -v Poller.go
