# `kube-2fa`: apply files with DUO mfa controls

This repository provides `kube-2fa` tool
[Install &rarr;](#installation)


**`kube-2fa`** is simple wrapper of kubectl which send push or mfa code into 2fa in every apply
Since it send push notification async, new resources are applied into the cluster. There should be control mechanism which check push-id in label



```
Usage:
  kube-2fa apply [flags]

Flags:
  -c, --code string       mfa code
      --config string     config file (default is $HOME/.kube_mfa.yaml)
  -f, --fileName string   fileName
  -h, --help              help for apply

```

### Usage

```

$ kube-2fa apply -f test.yaml
send push notification and add push id as a label.


$ kube-2fa apply -f test.yaml --code=123456
send push notification with code and add push id as a label.

```

-----
## Example 
Convert following deployment
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
    push-id: %AUTH_ID%
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80
```
into 
 ```

apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
    push-id: 1234-1213-12314-12313
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80
```
## Installation


It is so simple to build and use it.

```
go build .
cp kube-2fa /usr/local/bin
```
and save binary file somewhere in under PATH

##Configuration
```
    mfa_config:
       current_api: "duo"
       duo:
        username: "duo_name"
        ikey: "duo_ikey"
        skey: "duo_skey"
        host: "duo_host"
        userAgent: "duo_user_agent"
```

## Contribution

Adding new mfa service is so simple.

Implementing mfaApi and adding new service into tools.go is simple enough to use it.
Write configuration under your configuration yaml and select your mfa service as a current_api
```
    mfa_config:
       current_api: "custom-mfa"
       custom-mfa:
           custom-mfa-conf: "123"
```

