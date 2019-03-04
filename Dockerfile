# Compiler image
FROM didstopia/base:go-alpine-3.5 AS go-builder

# Copy the project 
COPY . /tmp/remailer/
WORKDIR /tmp/remailer/

# Install dependencies
RUN make deps

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/remailer



# Runtime image
FROM scratch

# Copy certificates
COPY --from=go-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Copy the built binary
COPY --from=go-builder /go/bin/remailer /go/bin/remailer

# Expose environment variables


# Expose ports
EXPOSE 2525

# Run the binary
ENTRYPOINT ["/go/bin/remailer"]
