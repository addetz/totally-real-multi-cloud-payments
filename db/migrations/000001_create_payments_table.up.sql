BEGIN;
CREATE TABLE IF NOT EXISTS payments.transactions
(
   tx_id INT PRIMARY KEY,
   user_id VARCHAR (50) NOT NULL,
   client_id VARCHAR (50) NOT NULL,
   from_acct VARCHAR (50) NOT NULL,
   to_acct VARCHAR (50) NOT NULL,
   amount INT NOT NULL
);
COMMIT;