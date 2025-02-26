# Course Backend

Backend API for Course Management, built with Go and Fiber, and deployed using Docker on Koyeb.

## 🚀 Deployment on Koyeb

### 1️⃣ Build and Push Docker Image
Ensure you have Docker installed and configured before running the following commands:

```sh
# Build Docker image for Linux/AMD64 platform
docker buildx build --platform linux/amd64 -t lshinkuro/coursebackend:latest .

# Push the image to Docker Hub
docker push lshinkuro/coursebackend:latest
```

### 2️⃣ Deploy to Koyeb
1. Go to [Koyeb Dashboard](https://app.koyeb.com/).
2. Click on **Create Service**.
3. Select **Docker Image** as the deployment source.
4. Enter your Docker image name:
   ```
   lshinkuro/coursebackend:latest
   ```
5. Set the **port** to `3000` (or match your application’s configured port).
6. Configure **environment variables** (e.g., `JWT_SECRET`, `DATABASE_URL`).
7. Click **Deploy**.

### 3️⃣ Verify Deployment
Once deployed, check the logs and access the API at:
```sh
https://your-koyeb-app.koyeb.app/
```

## 🛠 Environment Variables
Ensure these variables are set in Koyeb:

| Variable       | Description |
|---------------|-------------|
| `PORT`        | Set to `3000` |
| `JWT_SECRET`  | Secret key for JWT authentication |
| `DATABASE_URL`| Database connection URL |

## 📜 API Endpoints
| Method | Endpoint     | Description |
|--------|-------------|-------------|
| GET    | `/health`   | Check API health |
| POST   | `/login`    | User authentication |
| GET    | `/courses`  | Fetch all courses |

## 🛠 Troubleshooting
### ❌ **TCP Health Check Failed on Port 8000**
Ensure that your service is listening on **port 3000** in your application and Koyeb configuration.

### 🔄 **Updates Not Reflecting?**
1. Rebuild and push the Docker image:
   ```sh
   docker buildx build --platform linux/amd64 -t lshinkuro/coursebackend:latest .
   docker push lshinkuro/coursebackend:latest
   ```
2. Redeploy in Koyeb.

## 📝 License
MIT License. Feel free to use and contribute!

---

Happy Coding! 🚀
