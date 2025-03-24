FROM golang:1.23

# Set the working directory
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Install any needed packages specified in requirements.txt
RUN go mod tidy
RUN apt-get update && apt-get install make && apt-get install curl

RUN go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
ENV PATH=$PATH:/go/bin

# Make port 8080 available to the world outside this container
EXPOSE 8080

# RUN
CMD make migrate-up && make run