FROM golang:1.21.0

WORKDIR /project/api-regapp/

# COPY go.mod, go.sum and download the dependencies
COPY go.* ./
RUN go mod download

# COPY All things inside the project and build
COPY . .
RUN go build -o /project/auth-regapp/build/api .


EXPOSE 3000
ENTRYPOINT [ "/project/auth-regapp/build/api" ]