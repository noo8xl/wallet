
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
  
  balance DOUBLE UNSIGNED NOT NULL DEFAULT 0.0,
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  userId INT UNSIGNED NOT NULL,

  PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS oneTimeBtcWallets (
  id INT UNSIGNED auto_increment,
  address VARCHAR(100) NOT NULL,
  privateKey VARCHAR(500) NOT NULL,
  publicKey VARCHAR(500) NOT NULL,
  wif VARCHAR(500) NOT NULL,
  isUsed BOOL NOT NULL DEFAULT 0,
  balance DOUBLE UNSIGNED NOT NULL DEFAULT 0.0,

  
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  userId INT UNSIGNED NOT NULL,

  PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS ethWallets(
  id INT UNSIGNED auto_increment,
  address VARCHAR(100) NOT NULL,
  privateKey VARCHAR (500) NOT NULL,
  publicKey VARCHAR (500) NOT NULL,
  wif VARCHAR (500) NOT NULL,
  scriptType VARCHAR (500),
  originalAddress VARCHAR(500),
  oapAddress VARCHAR(500),

  balance DOUBLE UNSIGNED NOT NULL DEFAULT 0.0,
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  userId INT UNSIGNED NOT NULL,

  PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS oneTimeEthWallets(
  id INT UNSIGNED auto_increment,
  address VARCHAR(100) NOT NULL,
  privateKey VARCHAR (500) NOT NULL,
  publicKey VARCHAR (500) NOT NULL,
  wif VARCHAR (500) NOT NULL,
  isUsed BOOL NOT NULL DEFAULT 0,
  balance DOUBLE UNSIGNED NOT NULL DEFAULT 0.0,

  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  userId INT UNSIGNED NOT NULL,

  PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS tonWallets(
  id INT UNSIGNED auto_increment,
  address VARCHAR(100) NOT NULL,
  addrType INT,
  privateKey BLOB NOT NULL,
  bitsLen INT UNSIGNED,

  balance DOUBLE UNSIGNED NOT NULL DEFAULT 0.0,
  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  userId INT UNSIGNED NOT NULL,

  PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS oneTimeTonWallets(
  id INT UNSIGNED auto_increment,
  address VARCHAR(100) NOT NULL,
  addrType INT,
  privateKey BLOB NOT NULL,
  bitsLen INT UNSIGNED,
  isUsed BOOL NOT NULL DEFAULT 0,
  balance DOUBLE UNSIGNED NOT NULL DEFAULT 0.0,

  createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  userId INT UNSIGNED NOT NULL,

  PRIMARY KEY(id)
);


# common user rights only to insert data and use a few procedures* 
# admin user -> have rights to insert, select
# root -> have all privileges


DELIMITER $$
CREATE PROCEDURE IF NOT EXISTS update_wallet_balance(
  IN tableName VARCHAR(50),
  IN addr VARCHAR(100),
  IN newBalance FLOAT
)
BEGIN

  BEGIN TRY

    UPDATE tableName
    SET balance=newBalance, 
        updated_at=CURRENT_TIMESTAMP()
    WHERE address=addr;

	COMMIT ;

  END TRY

  BEGIN CATCH

    ROLLBACK ;

  END CATCH

END $$
DELIMITER ;

CALL update_wallet_balance(1,0.332);



