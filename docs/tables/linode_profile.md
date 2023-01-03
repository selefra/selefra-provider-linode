# Table: linode_profile

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| username | string | X | √ | Your username, used for logging in to our system. | 
| email | string | X | √ | Your email address. This address will be used for communication with Linode as necessary. | 
| email_notifications | bool | X | √ | If true, you will receive email notifications about account activity. If false, you may still receive business-critical communications through email. | 
| lish_auth_method | string | X | √ | The authentication methods that are allowed when connecting to the Linode Shell (Lish): password_keys, keys_only, disabled. | 
| referrals | json | X | √ | Information about your status in our referral program. This information becomes accessible after this Profile’s Account has established at least $25.00 USD of total payments. | 
| restricted | bool | X | √ | If true, your User has restrictions on what can be accessed on your Account. | 
| two_factor_auth | bool | X | √ | If true, logins from untrusted computers will require Two Factor Authentication. | 
| uid | string | X | √ | Your unique ID in our system. This value will never change, and can safely be used to identify your User. | 
| authorized_keys | json | X | √ | The list of SSH Keys authorized to use Lish for your User. | 
| ip_whitelist_enabled | bool | X | √ | If true, logins for your User will only be allowed from whitelisted IPs. This setting is currently deprecated, and cannot be enabled. | 
| timezone | string | X | √ | The timezone you prefer to see times in. | 


