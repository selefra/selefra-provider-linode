# Table: linode_domain

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| soa_email | string | X | √ | Start of Authority email address. This is required for master Domains. | 
| status | string | X | √ | Used to control whether this Domain is currently being rendered: disabled, active, edit_mode, has_errors. | 
| domain | string | X | √ | The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain. | 
| master_ips | json | X | √ | The IP addresses representing the master DNS for this Domain. | 
| description | string | X | √ | A description for this Domain. This is for display purposes only. | 
| ttl_sec | int | X | √ | Time to Live - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. | 
| domain_type | string | X | √ | If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave). | 
| retry_sec | int | X | √ | The interval, in seconds, at which a failed refresh should be retried. | 
| tags_src | json | X | √ | List of Tags applied to this domain. | 
| id | int | X | √ | The unique ID of this Domain. | 
| axfr_ips | json | X | √ | The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it. | 
| tags | json | X | √ | Tags applied to this domain as a map. | 
| expire_sec | int | X | √ | The amount of time in seconds that may pass before this Domain is no longer authoritative. | 
| refresh_sec | int | X | √ | The amount of time in seconds before this Domain should be refreshed. | 


