# Table: linode_instance

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| backups | json | X | √ | Information about this Linode’s backups status. | 
| image | string | X | √ | An Image ID to deploy the Disk from. | 
| ipv6 | cidr | X | √ | This Linode’s IPv6 SLAAC address. | 
| status | string | X | √ | The current status of the instance: creating, active, resizing, contact_support. | 
| label | string | X | √ | The Instance’s label is for display purposes only. | 
| tags_src | json | X | √ | List of Tags applied to this instance. | 
| watchdog_enabled | bool | X | √ | The watchdog, named Lassie, is a Shutdown Watchdog that monitors your Linode and will reboot it if it powers off unexpectedly. | 
| ipv4 | json | X | √ | Array of this Linode’s IPv4 Addresses. | 
| region | string | X | √ | Region where the instance resides. | 
| tags | json | X | √ | Tags applied to this instance as a map. | 
| specs | json | X | √ | Information about the resources available to this Linode, e.g. disk space. | 
| updated | timestamp | X | √ | When this Instance was last updated. | 
| id | int | X | √ | The unique ID of this Instance. | 
| alerts | json | X | √ | Alerts are triggered if CPU, IO, etc exceed these limits. | 
| created | timestamp | X | √ | When this Instance was created. | 
| hypervisor | string | X | √ | The virtualization software powering this Linode, e.g. kvm. | 
| instance_type | string | X | √ | This is the Linode Type that this Linode was deployed with. | 


