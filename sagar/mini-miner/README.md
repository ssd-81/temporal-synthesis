# Mini-Miner

## Problem Statement 
This project is a simplified mining essentially. The core idea to get the fundamental process going underneath various blockchains to validate blocks (especially) proof of work blockchains like bitcoin and initially Ethereum. 
---
The problem to be solved is available at 
`GET https://hackattic.com/challenges/mini_miner/problem?access_token=aaa699dde38ea86a`

Submission is made to
`POST https://hackattic.com/challenges/mini_miner/solve?access_token=aaa699dde38ea86a`


## Core ideas
- The idea is mainly to brute force to find the nonce for the provided problem. 
- But, there a catches to this simple problem: 
    - utilize all system cores (or as much as possible)
    - reduce inefficient compute
    - avoid expensive operations (like JSON serialization and hex -> hex string)
    

## Status logs
- the current issue is that I have setup the basic mining mechanism, but it is not really working out because of computational inefficiency. I need to figure out a way to utilize all the cores for the algorithm to run effectively. 
- introducing concurrency 
    - need to understand how to initilize new threads with task by creating a thread pool (or something)
- putting concurrency on hold
    - reason: the internal packages are not even able to handle easier hashes; so likely there is a flaw in the hash package (there is just method ; CheckSolution)
    - did ask gemini to check for the hash function; it said it was extremely inefficient due to hex encoding; first I will try to fix it and then do the research why it is so inefficient. 
    - ideas to explore : masking 
    - gained a bit of clarity about how low level representation works ; uint8 is a 8 bit unsigned integer
    so it can store values from 0-255
    - managed to get it right; there were minor inconsistencies while json parsing and requests
    - no need for concurrency. but likely will try to introduce that eventually if I am interested. 