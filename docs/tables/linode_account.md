# Table: linode_account

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| address_2 | string | X | √ | Second line of this Account’s billing address. | 
| city | string | X | √ | The city for this Account’s billing address. | 
| company | string | X | √ | The company name associated with this Account. | 
| last_name | string | X | √ | The last name of the person associated with this Account. | 
| phone | string | X | √ | The phone number associated with this Account. | 
| email | string | X | √ | The email address of the person associated with this Account. | 
| address_1 | string | X | √ | First line of this Account’s billing address. | 
| balance_uninvoiced | string | X | √ | This Account’s current estimated invoice in US dollars. This is not your final invoice balance. Transfer charges are not included in the estimate. | 
| country | string | X | √ | The two-letter country code of this Account’s billing address. | 
| state | string | X | √ | The state for this Account’s billing address. | 
| tax_id | string | X | √ | The tax identification number associated with this Account, for tax calculations in some countries. If you do not live in a country that collects tax, this should be null. | 
| balance | string | X | √ | This Account’s balance, in US dollars. | 
| credit_card | json | X | √ | Credit Card information associated with this Account. | 
| first_name | string | X | √ | The first name of the person associated with this Account. | 
| zip | string | X | √ | The zip code of this Account’s billing address. | 


