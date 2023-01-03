# Table: linode_kubernetes_cluster

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| status | string | X | √ |  | 
| tags | json | X | √ | Tags applied to the Kubernetes cluster as a map. | 
| tags_src | json | X | √ | List of Tags applied to the Kubernetes cluster. | 
| label | string | X | √ | This Kubernetes cluster’s unique label for display purposes only. | 
| api_endpoints | json | X | √ | API endpoints for the cluster. | 
| region | string | X | √ | This Kubernetes cluster’s location. | 
| kubeconfig | string | X | √ | Kube config for the cluster. | 
| pools | json | X | √ | Pools for the cluster. | 
| updated | timestamp | X | √ | When this Kubernetes cluster was updated. | 
| id | int | X | √ | This Kubernetes cluster’s unique ID. | 
| created | timestamp | X | √ | When this Kubernetes cluster was created. | 
| k8s_version | string | X | √ | The desired Kubernetes version for this Kubernetes cluster in the format of <major>.<minor>, and the latest supported patch version will be deployed. | 


