#!/bin/bash

protoc service/steamgame/steamgamepb/steamgame.proto --go_out=plugins=grpc:.
protoc service/game/gamepb/game.proto --go_out=plugins=grpc:.

protoc service/game/gamepb/game.proto \
--js_out=import_style=commonjs:./service/game/static/generated \
--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./service/game/static/generated