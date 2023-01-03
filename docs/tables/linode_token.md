# Table: linode_token

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| label | string | X | √ | This token's label. This is for display purposes only, but can be used to more easily track what you're using each token for. | 
| token | string | X | √ | First 16 characters of the token. | 
| scopes | json | X | √ | Array of scopes for the token, e.g. *, account:read_write, domains:read_only. | 
| created | timestamp | X | √ | The date and time this token was created. | 
| expiry | timestamp | X | √ | When this token will expire. | 
| id | int | X | √ | This token's unique ID, which can be used to revoke it. | 


