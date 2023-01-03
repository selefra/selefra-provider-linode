# Table: linode_type

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| price_hourly | int | X | √ | Cost (in US dollars) per hour. | 
| label | string | X | √ | The Linode Type’s label is for display purposes only. | 
| addons | json | X | √ | A list of optional add-on services for Linodes and their associated costs. | 
| memory | int | X | √ | Amount of RAM included in this Linode Type. | 
| transfer | int | X | √ | The monthly outbound transfer amount, in MB. | 
| vcpus | int | X | √ | The number of VCPU cores this Linode Type offers. | 
| id | string | X | √ | The ID representing the Linode Type. | 
| disk | int | X | √ | The Disk size, in MB, of the Linode Type. | 
| network_out | int | X | √ | The Mbits outbound bandwidth allocation. | 
| class | string | X | √ | The class of the Linode Type: nanode, standard, dedicated, gpu, highmem. | 
| price_monthly | int | X | √ | Cost (in US dollars) per month. | 


