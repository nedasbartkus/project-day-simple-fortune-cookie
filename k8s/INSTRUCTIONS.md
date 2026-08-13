# Running the Application with Kubernetes

This guide explains how to deploy and run the Simple Fortune Cookie application using Kubernetes.

## Prerequisites

Make sure you have the following installed and running:

- Docker Desktop
- Kubernetes
- `kubectl`

Check that Kubernetes is running:

```powershell
kubectl get nodes
```

At least one node should have the status `Ready`.

---

## 1. Apply Redis Persistent Storage

Redis uses persistent storage so that fortune cookies are not lost when the Redis pod is recreated.

Apply the Redis PersistentVolumeClaim:

```powershell
kubectl apply -f k8s/redis-pvc.yaml
```

Check that the PVC was created:

```powershell
kubectl get pvc
```

The `redis-data` PVC should have the status:

```text
Bound
```

---

## 2. Deploy the Application

Apply all Kubernetes manifests:

```powershell
kubectl apply -f k8s/
```

This creates:

- Frontend Deployment
- Frontend Service
- Backend Deployment
- Backend Service
- Redis Deployment
- Redis Service
- Redis PersistentVolumeClaim

---

## 3. Check the Pods

Check that all pods are running:

```powershell
kubectl get pods
```

All application pods should eventually show:

```text
Running
```

For example:

```text
NAME                               READY   STATUS    RESTARTS   AGE
cookie-backend-xxxxxxxxx-xxxxx    1/1     Running   0          1m
cookie-frontend-xxxxxxxxx-xxxxx   1/1     Running   0          1m
redis-xxxxxxxxx-xxxxx              1/1     Running   0          1m
```

If a pod is not running, check its logs:

```powershell
kubectl logs <pod-name>
```

---

## 4. Check the Deployments

```powershell
kubectl get deployments
```

You should see all deployments with their replicas ready:

```text
NAME              READY   UP-TO-DATE   AVAILABLE
cookie-backend    1/1     1            1
cookie-frontend   1/1     1            1
redis             1/1     1            1
```

---

## 5. Check the Services

```powershell
kubectl get svc
```

You should see something similar to:

```text
NAME              TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)
cookie-backend    ClusterIP   10.x.x.x        <none>        9000/TCP
cookie-frontend   NodePort    10.x.x.x        <none>        8080:32153/TCP
redis             ClusterIP   10.x.x.x        <none>        6379/TCP
```

The frontend uses a `NodePort`, which makes it accessible from your local machine.

---

## 6. Access the Frontend

Find the frontend NodePort:

```powershell
kubectl get svc cookie-frontend
```

Look at the `PORT(S)` column.

For example:

```text
8080:32153/TCP
```

Here:

- `8080` is the application port.
- `32153` is the NodePort exposed on the machine.

Open the following address in your browser:

```text
http://127.0.0.1:32153
```

Replace `32153` with the NodePort shown by your own `kubectl get svc` command.

> The NodePort may be different if the service is recreated.

---

## 7. Test the Application

Once the frontend is open:

1. View the available fortune cookies.
2. Add a new fortune cookie.
3. Refresh the page.
4. Verify that the new fortune is still available.

The application communicates internally through Kubernetes services:

```text
Frontend
   |
   | cookie-backend:9000
   v
Backend
   |
   | redis:6379
   v
Redis
```

---

## 8. Test Redis Persistence

Redis is configured with a PersistentVolumeClaim so that its data survives pod recreation.

First, add a new fortune cookie through the frontend.

Then find the Redis pod:

```powershell
kubectl get pods
```

Delete the Redis pod:

```powershell
kubectl delete pod <redis-pod-name>
```

For example:

```powershell
kubectl delete pod redis-xxxxxxxxx-xxxxx
```

Because Redis is managed by a Deployment, Kubernetes will automatically create a replacement pod.

Watch the pods:

```powershell
kubectl get pods
```

Wait until the new Redis pod shows:

```text
1/1   Running
```

Refresh the frontend.

The fortune cookie you added should still be available.

This confirms that Redis is using persistent storage.

---

## 9. Check Persistent Storage

Check the PersistentVolumeClaim:

```powershell
kubectl get pvc
```

You should see something similar to:

```text
NAME         STATUS   VOLUME   CAPACITY   ACCESS MODES
redis-data   Bound    ...      1Gi        RWO
```

You can also check PersistentVolumes:

```powershell
kubectl get pv
```

The Redis PVC should be connected to a PersistentVolume.

---

## 10. View Application Logs

### Backend

```powershell
kubectl logs deployment/cookie-backend
```

### Frontend

```powershell
kubectl logs deployment/cookie-frontend
```

### Redis

```powershell
kubectl logs deployment/redis
```

To view logs from a specific pod:

```powershell
kubectl logs <pod-name>
```

---

## 11. Troubleshooting

### Check all resources

```powershell
kubectl get all
```

### Check pods

```powershell
kubectl get pods
```

### Check services

```powershell
kubectl get svc
```

### Check deployments

```powershell
kubectl get deployments
```

### Check persistent storage

```powershell
kubectl get pvc
kubectl get pv
```

### Get detailed information about a pod

```powershell
kubectl describe pod <pod-name>
```

### Get detailed information about the Redis PVC

```powershell
kubectl describe pvc redis-data
```

### View recent Kubernetes events

```powershell
kubectl get events --sort-by=.lastTimestamp
```

---

## 12. Remove the Application

To remove the Kubernetes resources defined in the `k8s` folder:

```powershell
kubectl delete -f k8s/
```

### Important: Redis Data

The Redis data is stored using the `redis-data` PersistentVolumeClaim.

Check whether the PVC still exists:

```powershell
kubectl get pvc
```

If you want to keep the Redis data, **do not delete the PVC**.

If you intentionally want to delete the persistent storage and its data:

```powershell
kubectl delete pvc redis-data
```

> Deleting the PVC may permanently remove the stored Redis data, depending on the storage configuration.

---

## Useful Commands

### View pods

```powershell
kubectl get pods
```

### View services

```powershell
kubectl get svc
```

### View deployments

```powershell
kubectl get deployments
```

### View persistent storage

```powershell
kubectl get pvc
kubectl get pv
```

### View everything

```powershell
kubectl get all
```

### Watch pods

```powershell
kubectl get pods -w
```

Press `Ctrl+C` to stop watching.

### Restart a deployment

```powershell
kubectl rollout restart deployment/<deployment-name>
```

For example:

```powershell
kubectl rollout restart deployment/redis
```

---

## Kubernetes Architecture

The application consists of three main components:

```text
                    Kubernetes
                         |
             +-----------+-----------+
             |                       |
        Frontend Service         Redis Service
          (NodePort)               (ClusterIP)
             |                       |
             v                       v
      Frontend Pod              Redis Pod
             |                       |
             v                       v
      Backend Service         Persistent Storage
        (ClusterIP)                 ^
             |                      |
             v                      |
       Backend Pod           PersistentVolumeClaim
```

The frontend is exposed through a `NodePort`.

The backend and Redis are only accessible inside the Kubernetes cluster through their respective `ClusterIP` services.

Redis uses persistent storage so that fortune cookies are not lost when the Redis pod is recreated.