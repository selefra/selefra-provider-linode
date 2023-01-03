# Table: linode_volume

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| size | int | X | √ | The Volume’s size, in GiB. | 
| created | timestamp | X | √ | When this Volume was created. | 
| filesystem_path | string | X | √ | The full filesystem path for the Volume based on the Volume’s label. | 
| region | string | X | √ | Region where the volume resides. | 
| updated | timestamp | X | √ | When this Volume was last updated. | 
| id | int | X | √ | The unique ID of this Volume. | 
| label | string | X | √ | The Volume’s label is for display purposes only. | 
| linode_id | int | X | √ | If a Volume is attached to a specific Linode, the ID of that Linode will be displayed here. | 
| status | string | X | √ | The current status of the volume: creating, active, resizing, contact_support. | 
| tags | json | X | √ | Tags applied to this volume as a map. | 
| tags_src | json | X | √ | List of Tags applied to this volume. | 


