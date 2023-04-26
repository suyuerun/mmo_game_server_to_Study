#!/bin/bash
protoc --go_out=. *.proto;
protoc --csharp_out=./c# *.proto;