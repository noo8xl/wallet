# wallet

-> Documentation:

Commands:

--- gwlt: the 'gwlt' command implements get wallet function.

- flags: blockchain(btc) OR (create) userId(stringVal);
  --> first flag can be (create) instead of currency name. It provides create wallet function in each available blockchain <--
- example: $ wallet gwlt btc your-user-id _OR_ $ wallet gwlt create your-user-id;
- result: gwlt returns a new btc address OR wallet []list if don't catch an error.

--- gb: the 'gb' command implements get balance function.

- flags: coinName(btc) address(btc address) currencyType(usd);
- example: $ wallet gb btc bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh usd;
- result: gb returns a btc value and value in fiat by chosen currencyType if don't catch an error.

--- tsx: the 'tsx' command implements send transaction function.

- flags: blockchain(btc) addressFrom(your-btc-address) addressTo(beneficiar-btc-address) amount(amount-in-crypto);
  -> transaction should be signed ONLY by your private key. The tsx method available ONLY if you send coins from wallet,
  which created by wallet app! <-
- example: $ wallet tsx btc bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhxIIos 0.2
- result: tsx returns a link to scanner or transaction hash.

--- psr: the 'psr' command implements a balance parser function.

- flags: coinName(btc) from(int64 value stamp) to(int64 value stamp);
- example: $ wallet prs btc 790819802479 192846192841;
- result: prs returns a []list of addresses in chosen blockchain with NOT-EMPTY(0.001) balance.

Additional usage:

--- The app can be used via cli or could be connected with the main api service via redis.

- In redis you could find a requirement data as userId, coinName, address to use it outside of cli app.
- If you get a data from redis store or not - data would be deleted in 5 minutes.

Available settings:

--- blockchains: btc(bitcoin), ton(the-open-network), eth(ethereum);
--- currency types: usd, eur, aud, aed, rub;
--- fees: to each transaction in any blockchain the fees will be setted automatically.
--- balance parser: work ONLY amoung wallets, created by wallet application.
