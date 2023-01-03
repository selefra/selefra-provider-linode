# Table: linode_image

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| image_type | string | X | √ | How the Image was created: manual, automatic. | 
| is_public | bool | X | √ | True if the Image is public. | 
| size | int | X | √ | The minimum size this Image needs to deploy. Size is in MB. | 
| id | string | X | √ | The unique ID of this Image. | 
| label | string | X | √ | A short description of the Image. | 
| created | timestamp | X | √ | When this Image was created. | 
| deprecated | bool | X | √ | Whether or not this Image is deprecated. Will only be true for deprecated public Images. | 
| created_by | string | X | √ | The name of the User who created this Image, or 'linode' for official Images. | 
| description | string | X | √ | A detailed description of this Image. | 
| expiry | timestamp | X | √ | Only Images created automatically (from a deleted Linode; type=automatic) will expire. | 
| vendor | string | X | √ | The upstream distribution vendor. None for private Images. | 


