# First stage: build the executable.
FROM golang:1.17.0-alpine AS builder

# git is required to fetch go dependencies
RUN apk add --no-cache git

# Create the user and group files that will be used in the running
# container to run the process as an unprivileged user.
# RUN mkdir /user && \
#     cat /etc/passwd > /user/passwd && \
#     echo 'nobody:x:65534:65534:nobody:/:' >> /user/passwd && \
#     cat /etc/group > /user/group && \
#     echo 'nobody:x:65534:' >> /user/group && \
#     echo 'nobody ALL=(ALL) NOPASSWD: /bin/cp /context/.env /app' > /user/nobody

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
COPY ./go.mod ./go.sum ./

ARG GITLAB_USER_NAME
ARG GITLAB_USER_PAT
RUN git config --global url."https://$GITLAB_USER_NAME:$GITLAB_USER_PAT@gitlab.com/".insteadOf https://gitlab.com/
RUN go mod download

# Import the code from the context.
COPY . .

# Build the executable. Mark the build as statically linked.
ENV GO111MODULE "on"
ENV CGO_ENABLED 0
RUN go build \
    -installsuffix 'static' \
    -o main .

# Final stage: the running container.
FROM alpine AS final

# RUN apk --no-cache update \
#     && apk --no-cache add sudo

WORKDIR /app

# Import the user and group files from the first stage.
# RUN adduser -D somebody somegroup \
#     && echo "somebody ALL=(ALL) NOPASSWD: /bin/cp /context/.env /app" > /etc/sudoers.d/somebody \
#     && chmod 0440 /etc/sudoers.d/somebody \
#     && echo "somebody:$GITLAB_USER_PAT" | chpasswd
# COPY --from=builder /user/group /user/passwd /etc/
# COPY --from=builder /user/nobody /etc/sudoers.d/
# RUN chmod 0440 /etc/sudoers.d/nobody
# RUN echo "nobody:$GITLAB_USER_PAT" | chpasswd

# Import the Certificate-Authority certificates for enabling HTTPS.
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Import required artifacts first stage.
COPY --from=builder /src/main .

# Give nobody access to working folder
# RUN chown -vR somebody:somegroup /app

# Perform any further action as an unprivileged user.
# USER somebody

# Run the compiled binary.
CMD [ "/bin/sh", "-c", "cp -v /context/.env /app && sleep 3 && /app/main"]


