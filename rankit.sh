#!/bin/bash
docker run -it --rm -v $(pwd):/app -w /app golang go run .