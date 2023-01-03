# Table: linode_event

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| action | string | X | √ | The action that caused this Event. New actions may be added in the future. | 
| status | string | X | √ | Current status of the Event, Enum: 'failed' 'finished' 'notification' 'scheduled' 'started'. | 
| entity | json | X | √ | Detailed information about the Event's entity, including ID, type, label, and URL used to access it. | 
| read | bool | X | √ | If this Event has been read. | 
| created | timestamp | X | √ | The date and time this event was created. | 
| username | string | X | √ | The username of the User who caused the Event. | 
| percent_complete | int | X | √ | A percentage estimating the amount of time remaining for an Event. Returns null for notification events. | 
| rate | string | X | √ | The rate of completion of the Event. Only some Events will return rate; for example, migration and resize Events. | 
| seen | bool | X | √ | If this Event has been seen. | 
| secondary_entity | json | X | √ | Detailed information about the Event's secondary or related entity, including ID, type, label, and URL used to access it. | 
| time_remaining | int | X | √ | The estimated time remaining until the completion of this Event. This value is only returned for in-progress events. | 
| id | int | X | √ | The unique ID of this Event. | 


