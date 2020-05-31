apiVersion: apps/v1
kind: Deployment
metadata:
  name: jaeger
spec:
  selector:
    matchLabels:
      run: jaeger
  replicas: 1
  template:
    metadata:
      labels:
        run: jaeger
    spec:
      containers:
        - name: jaeger
          image: jaegertracing/all-in-one:latest
          imagePullPolicy: {{ .Values.pullPolicy }}
          ports:
            - name: http
              containerPort: 16686
            - name: jaeger
              containerPort: 6831
              protocol: UDP
---
apiVersion: v1
kind: Service
metadata:
  name: jaeger
  labels:
    run: jaeger
spec:
  ports:
    - name: http
      port: 16686
      protocol: TCP
    - name: jaeger
      port: 6831
      protocol: UDP
  selector:
    run: jaeger
