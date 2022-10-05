#/*
# * Copyright 2022 CoreLayer BV
# *
# *    Licensed under the Apache License, Version 2.0 (the "License");
# *    you may not use this file except in compliance with the License.
# *    You may obtain a copy of the License at
# *
# *        http://www.apache.org/licenses/LICENSE-2.0
# *
# *    Unless required by applicable law or agreed to in writing, software
# *    distributed under the License is distributed on an "AS IS" BASIS,
# *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# *    See the License for the specific language governing permissions and
# *    limitations under the License.
# */

FROM golang:1.17-alpine AS BUILD

WORKDIR $GOPATH/src/github.com/corelayer/lens
COPY . .

RUN go get -d -v
RUN go install -v


RUN GOOS=linux GOARCH=amd64 go build -o /tmp/lens main.go


FROM gcr.io/distroless/base:latest

COPY --from=BUILD /tmp/lens /usr/local/bin/lens

CMD ["/usr/local/bin/lens"]