#!/bin/bash

export REVISION=$(git rev-parse --short HEAD)
skaffold build -p ci
