FROM golang:1.18

ARG docker_id
ARG docker_user

RUN useradd -G sudo,root -u ${docker_id} -d /home/${docker_user} ${docker_user}
RUN mkdir /home/${docker_user}
RUN chown ${docker_user}:${docker_user} /home/${docker_user}

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Build Delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

#Google-chrome repo
RUN wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add - \
&& echo "deb http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google.list

RUN apt-get update
RUN apt-get -y install net-tools lsof google-chrome-stable

WORKDIR /usr/src/app

CMD ["air"]