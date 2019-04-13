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

#
# Build poller object
build:
  $(GOBUILD) -o $(BINARY_NAME) -v *.go
