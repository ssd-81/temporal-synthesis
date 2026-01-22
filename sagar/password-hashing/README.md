# password hashing 

## the core idea  
- a password, a salt and some algorithmic specific parameters will be provided 
- calculate the sha256, hmac-sha256 , pbkdf-sha256, scrypt
- the core idea to understand is understand how each step is progressively used to build the next one. 
- the problem statement also says something about "secret step"; but not sure what it is ; just keep in mind 


## learnings
- sha256 is the core pillar everything is standing on (at least in PS context)
- hmac prevents attacks which can use existing hash value and append new data it; it does this by surrounding the actual hash value by other data (data-hash-hash)
- pbkdf2 is used for deriving a key ; it uses hmac internally; it processes the password multiple times using a digest and finally providing a fixed length output 
