# Aplikasi Manajemen Data Mahasiswa

Aplikasi Full-Stack dengan arsitektur Microservices, dirancang untuk berjalan di lingkungan Kubernetes.

## Teknologi

- **Frontend:** React.js (Vite) + Nginx
- **Backend:** Golang (Gin Framework)
- **Database:** PostgreSQL
- **Infrastruktur:** Kubernetes (Minikube)

## Cara Menjalankan

### 1. Persiapan Cluster

Jalankan di PowerShell:

```bash
minikube start
minikube -p minikube docker-env | Invoke-Expression
```

### 2. Build Image

```bash
# Backend
cd backend
docker build -t backend-api:1.0 .
cd ..

# Frontend
cd frontend
docker build -t frontend-web:1.0 .
cd ..

```

### 3. Deploy

```bash
kubectl apply -f k8s/database.yaml
kubectl apply -f k8s/backend.yaml
kubectl apply -f k8s/frontend.yaml

```

_(Cek dengan `kubectl get pods -w` hingga semua berstatus `Running`)_

### 4. Akses Aplikasi

Buka terminal baru, jalankan port-forward untuk web:

```bash
kubectl port-forward svc/frontend-service 3000:8080

```

Buka browser: **`http://localhost:3000`**

_(Opsional) Untuk dokumentasi API Swagger, buka terminal baru:_

```bash
kubectl port-forward svc/backend 8080:8080

```

Buka browser: **`http://localhost:8080/swagger/index.html`**
