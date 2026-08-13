# Docker

## Frontend

The frontend can be built into a Docker image using:

```powershell
docker build -t cookie-frontend ./frontend
```

To run the frontend container:

```powershell
docker run -p 8080:8080 cookie-frontend
```

To stop the container:

```powershell
docker ps
docker stop <container-id>
```