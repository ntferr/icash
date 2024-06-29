# WIP :wrench:

## OminiDB - SQL Client

**Connection string for connect to OminiDB**

| Field    | Value    |
| -------- | -------  |
| Host     | postgres |
| Port     | 5432     |
| User     | icash    |
| Password | 123456   |


**Credentials**
- User: admin
- Password: admin

![connection_string_example](https://github.com/ntferr/icash/blob/main/assets/db_connection_string_example.png)


```mermaid
    erDiagram

    BANKS {
        string id
        string name
        string code
    }
    CARDS {
        string id
        string bank_id
        string number
        string expire_at
    }
```

## DER
![arquitetura-banco](https://github.com/ntferr/icash/blob/main/assets/der.png)

---
## Packages

    Netflix/go-env  - Enviroments variables
    Fiber           - REST API
    Sonic           - JSON Marshal/Unmarshal 
    Snowflake       - Generate id
    Testify         - Tests
    