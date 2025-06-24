
# WIP :wrench:

## OminiDB - SQL Client

### Credentials
> [!NOTE]
> - User: admin
> - Password: admin

### Connection string for connect to OminiDB
<div align="center">

| Field      | Value      |
| :--------: | :-------:  |
| Host       | postgres   |
| Port       | 5432       |
| User       | icash      |
| Password   | 123456     |

</div>

![connection_string_example](assets/db_connection_string_example.png)

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
    DEBTS {
        string id
        string card_id
        string ticket_id
        string name
        string description
    }
    INSTALLMENTS {
        string id
        string debt_id
        string due_date
        bool paid
        int number
    }
    TICKETS {
        string id
        string debt_id
        string code_bars
    }

    CARDS ||--o{ BANKS: bank_id
    DEBTS ||--o{ CARDS: card_id
    DEBTS ||--o{ TICKETS: ticket_id
    INSTALLMENTS ||--o{ DEBTS: debt_id
    TICKETS ||--o{ DEBTS: debt_id 
```

## Packages

    Netflix/go-env  - Enviroments variables
    Fiber           - REST API
    Sonic           - JSON Marshal/Unmarshal 
    Snowflake       - Generate id
    Testify         - Tests

