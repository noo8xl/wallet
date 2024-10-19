

-- 0
-- -> false 
-- 1 -> true

-- create table users(
--   ID int PRIMARY KEY AUTO_INCREMENT, 
--   permission BOOLEAN NOT NULL,
--   user_error varchar
-- (1200),
--   user_id varchar
-- (200) NOT NULL,
--   created_at int NOT NULL, 
--   updated_at int NOT NULL, 
-- );

-----------------------  wallets tables <- 


CREATE TABLE IF NOT EXISTS btcWallets (
  id INT
  auto_increment PRIMARY KEY, 
  address VARCHAR(100) NOT NULL, 
  privateKey VARCHAR(500) NOT NULL, 
  publicKey VARCHAR(500) NOT NULL, 
  wif VARCHAR(500) NOT NULL, 
  scriptType VARCHAR(500), 
  originalAddress VARCHAR(500), 
  oapAddress VARCHAR(500), 
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
  updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
  userId VARCHAR(200) NOT NULL 
);

CREATE TABLE IF NOT EXISTS ethWallets(
  id int auto_increment PRIMARY KEY, 
  address VARCHAR(100) NOT NULL, 
  privateKey VARCHAR (500) NOT NULL, 
  publicKey VARCHAR (500) NOT NULL, 
  wif VARCHAR (500) NOT NULL, 
  scriptType VARCHAR (500), 
  originalAddress VARCHAR(500), 
  oapAddress VARCHAR(500), 
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
  updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
  userId VARCHAR(200) NOT NULL 
);

CREATE TABLE IF NOT EXISTS tonWallets(
  id INT auto_increment PRIMARY KEY, 
  address VARCHAR(100) NOT NULL, 
  addrType INT, 
  privateKey BLOB NOT NULL, 
  bitsLen INT UNSIGNED, 
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
  updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
  userId VARCHAR(200) NOT NULL
);
