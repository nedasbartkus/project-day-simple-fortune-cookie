# Docker

## Backend

The backend can be built into a Docker image using:

```powershell
docker build -t cookie-backend ./backend
```

To run the backend container:

```powershell
docker run -p 9000:9000 cookie-backend
```

To stop the container:~

```powershell
docker ps
docker stop <container-id>
```