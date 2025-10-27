# Unpacking

## core idea
The core is pretty simple; access the problem endpoint to grab a base64 encoded data; now we need to extract some particular values from that. 
- Seems pretty simple, let's see how it goes. 
- Unpacking would be a better word than extracting; though I am not sure of the difference. 
- the pack is always going to contain the following in the same order: 
    - a regular int (signed), to start off
    - an unsigned int
    - a short (signed) to make things interesting
    - a float because floating point is important
    - a double as well
    - another double but this time in big endian (network byte order)
- focus on 
    - base64 pack of **bytes**

- a few fundamental questions 
    - what is encoding and why do we need it ? 
    - can't we simply operate without an encoding? 
    - what is the difference between an encoding and a type? 
    