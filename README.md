[![Build Status](https://travis-ci.org/0intro/ripple-keypairs.svg?branch=master)](https://travis-ci.org/0intro/ripple-keypairs)

Ripple-KeyPairs
===============

This tool generates Ripple key pairs.

Usage
-----

```
usage: ripple-keypairs [ -n nWorkers ] [ -p prefix | -s seed ]
```

Examples
--------

Generate a Ripple key pair from a random seed:

```
$ ripple-keypairs
Seed (secret key) shkSRjSLNX4FUmDKFyQ96VyHJHbqf
AccountID r1Lr1474jgGmtLr1UoUWfwFgeQvEZUfBd
NodePublicKey n9KxW47TUSEcUfh9Z7T5NuZtTnZfXdj834sCoKZyMFFh7GRFhhzg
NodePrivateKey pfaQU75qFzMv9L9NgXQfdL5My1qtDndFiTFH3a2RgELzN72EcNc
AccountPublicKey aB43KJd1aqcYs7ZfBkgq1s9w6HeCyhDRE5KudgzXK5J7NzMNvDBQ
AccountPrivateKey pwrjcMKBrvufQDHeUUkjeaduLK9mcorHRKUqT4AJDzVH7cBikwf
```

Generate a Ripple key pair from the specified seed:

```
$ ripple-keypairs -seed shkSRjSLNX4FUmDKFyQ96VyHJHbqf
Seed (secret key) shkSRjSLNX4FUmDKFyQ96VyHJHbqf
AccountID r1Lr1474jgGmtLr1UoUWfwFgeQvEZUfBd
NodePublicKey n9KxW47TUSEcUfh9Z7T5NuZtTnZfXdj834sCoKZyMFFh7GRFhhzg
NodePrivateKey pfaQU75qFzMv9L9NgXQfdL5My1qtDndFiTFH3a2RgELzN72EcNc
AccountPublicKey aB43KJd1aqcYs7ZfBkgq1s9w6HeCyhDRE5KudgzXK5J7NzMNvDBQ
AccountPrivateKey pwrjcMKBrvufQDHeUUkjeaduLK9mcorHRKUqT4AJDzVH7cBikwf
```

Generate a Ripple key pair with an account ID beginning by the specified prefix (with 4 workers):

```
$ ripple-keypairs -prefix rBob -n 4
Seed (secret key) sp8Qdg2FhSayDNAHgGtDzvJesEoCJ
AccountID rBobUXSghgBwSKHqWd5DgaveWFDdqA1b7Q
NodePublicKey n9L1zogoPuKHqDeRVbZHZvzGSgS3vZEf5gJo7G9CX1SbEUsQTfnK
NodePrivateKey pnjNqY81Caxug9mvPZ74iwFoZKAWAGU1QJq46y4rmjo667ezUaA
AccountPublicKey aBQ9dfB2ghZ1ayXseBUSZjYbkqemcHTCMTbLtoKFnSrns1QqnU2X
AccountPrivateKey p97DJ5QgvGzw6iMB6H7DgzRSmFCmrSVPf7iwGEj5wA63JFjinjA
```

Thanks
------

This tools relies on the very nice [github.com/rubblelabs/ripple](https://github.com/rubblelabs/ripple) package from [Donovan Hide](https://github.com/donovanhide).
