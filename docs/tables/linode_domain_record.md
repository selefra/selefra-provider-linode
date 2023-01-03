# Table: linode_domain_record

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | int | X | √ | This Record’s unique ID. | 
| priority | int | X | √ | The priority of the target host for this Record. Lower values are preferred. Only valid for MX and SRV record requests. Required for SRV record requests. | 
| protocol | string | X | √ | The protocol this Record’s service communicates with. An underscore (_) is prepended automatically to the submitted value for this property. Only valid for SRV record requests. | 
| service | string | X | √ | The name of the service. Only valid and required for SRV record requests. | 
| tag | string | X | √ | The tag portion of a CAA record. Only valid and required for CAA record requests. | 
| weight | int | X | √ | The relative weight of this Record used in the case of identical priority. Higher values are preferred. Only valid and required for SRV record requests. | 
| record_type | string | X | √ | The type of Record this is in the DNS system: A, AAAA, NS, MX, CNAME, TXT, SRV, PTR, CAA. | 
| name | string | X | √ | The name of this Record. This property’s actual usage and whether it is required depends on the type of record it represents. For example, for CNAME, it is the hostname. | 
| port | int | X | √ | The port this Record points to. Only valid and required for SRV record requests. | 
| target | string | X | √ | The target for this Record. For requests, this property’s actual usage and whether it is required depends on the type of record this represents. For example, for CNAME it is the domain target. | 
| ttl_sec | int | X | √ | Time to Live - the amount of time in seconds that the domain record may be cached by resolvers or other domain servers. | 


