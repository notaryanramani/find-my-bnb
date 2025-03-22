# Find My BnB

Find My BnB is a **vector search application** designed to help users find the best BnB listings using vector embeddings for search optimization. The project is still under development.

## Project Structure

```
find-my-bnb/
├── api/        # Go backend
├── embedder/   # Python vector embedding service
├── web/        # Frontend (JS, HTML, Tailwind CSS, Express for routing)
└── README.md   # Documentation
```

## Prerequisites

Ensure you have the following installed:

- **Go** (latest version)
- **Python** (3.8+)
- **Node.js** (latest LTS version)
- **Docker & Docker Compose**

## Running the Application

The application will be containerized in the future, ensuring ease of deployment and scalability.

Follow these steps to set up and run the application in the correct order:

### 1️⃣ Start PostgreSQL Database

```sh
cd api
docker-compose up -d
```

This will start the PostgreSQL database in a detached mode.

---

### 2️⃣ Run Python Embedder Service

```sh
cd embedder
source venv/bin/activate  # Activate virtual environment
python main.py            # Start the embedding service
```

Ensure that `venv` is properly set up before running the script.

---

### 3️⃣ Run Go Backend

```sh
cd api
go mod tidy  # Ensure all dependencies are installed
make run      # Start the backend server
```

The backend will now be running and listening for requests.

---

### 4️⃣ Run Frontend

```sh
cd web
npm install   # Install frontend dependencies
npm start     # Start the frontend server
```

This will serve the frontend application, making the full system functional.

## Contributing

If you want to contribute, ensure you follow best practices:

- Format code properly before committing.
- Use meaningful commit messages.
- Follow the directory structure.

## License

This project is licensed under the MIT License.

---

For any issues or feature requests, feel free to open a GitHub issue.

