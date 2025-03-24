FROM golang:1.23-alpine

# Set the working directory
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Install any needed packages specified in requirements.txt
RUN go mod tidy
RUN apk update && apk add make

# Make port 8080 available to the world outside this container
EXPOSE 8080

# RUN
CMD ["make", "run"]