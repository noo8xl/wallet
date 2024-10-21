
-----------------------  wallets tables <- 


CREATE TABLE IF NOT EXISTS btcWallets (
  id INT UNSIGNED auto_increment, 
  address VARCHAR(100) NOT NULL, 
  privateKey VARCHAR(500) NOT NULL, 
  publicKey VARCHAR(500) NOT NULL, 
  wif VARCHAR(500) NOT NULL, 
  scriptType VARCHAR(500), 
  originalAddress VARCHAR(500), 
  oapAddress VARCHAR(500), 
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
  updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
  userId VARCHAR(200) NOT NULL,

   PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS ethWallets(
  id INT UNSIGNED auto_increment PRIMARY KEY, 
  address VARCHAR(100) NOT NULL, 
  privateKey VARCHAR (500) NOT NULL, 
  publicKey VARCHAR (500) NOT NULL, 
  wif VARCHAR (500) NOT NULL, 
  scriptType VARCHAR (500), 
  originalAddress VARCHAR(500), 
  oapAddress VARCHAR(500), 
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
  updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
  userId VARCHAR(200) NOT NULL,
  
   PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS tonWallets(
  id INT UNSIGNED auto_increment, 
  address VARCHAR(100) NOT NULL, 
  addrType INT, 
  privateKey BLOB NOT NULL, 
  bitsLen INT UNSIGNED, 
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
  updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
  userId VARCHAR(200) NOT NULL,

   PRIMARY KEY(id)
);
