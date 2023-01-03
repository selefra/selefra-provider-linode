# Table: linode_bucket

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| label | string | X | √ | The name of this bucket. | 
| cluster | string | X | √ | The ID of the Object Storage Cluster this bucket is in. | 
| created | timestamp | X | √ | When this bucket was created. | 
| hostname | string | X | √ | The hostname where this bucket can be accessed. This hostname can be accessed through a browser if the bucket is made public. | 


