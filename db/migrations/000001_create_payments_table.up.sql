BEGIN;
CREATE TABLE IF NOT EXISTS payments.transactions
(
   user_id INT PRIMARY KEY,
   client_id VARCHAR (50) NOT NULL,
   from_acct VARCHAR (50) NOT NULL,
   to_acct VARCHAR (50) NOT NULL,
   amount INT NOT NULL
);
COMMIT;