# Table: linode_user

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| username | string | X | √ | This User’s username. This is used for logging in, and may also be displayed alongside actions the User performs (for example, in Events or public StackScripts). | 
| email | string | X | √ | The email address for this User, for account management communications, and may be used for other communications as configured. | 
| restricted | bool | X | √ | If true, this User must be granted access to perform actions or access entities on this Account. | 
| ssh_keys | json | X | √ | A list of SSH Key labels added by this User. These are the keys that will be deployed if this User is included in the authorized_users field of a create Linode, rebuild Linode, or create Disk request. | 


