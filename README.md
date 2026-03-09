# UberEats practice

I've used Grab in SouthEastAsia and heard of HungerStation and Foodora in the MiddleEast. Will use Golang and Postgres because that's my current preferred stack.

I'm splitting the operations by writes and reads because I want to experiment with non-REST APIs. I've done RoR and REST APIs for 13 years now and want to challenge myself with something new.

## Writes
Will probably put all of this in one endpoint, and then switch on opcodes or sth. If that pattern is good enough for low-level C programming, it's good enough for the web.

- Restaurant new (cuisines, locations, name, desc)
  - Maybe this should just be business to allow for groceries, pharmacies, etc? Use the encoding approach to have subtypes, or just have them all separate?
- Dish new (pic, name, desc, cuisine tags, restaurant)
- Rating new (delivery or product)
- Order new (dishes, time, cost)
- Delivery
  - ACK'd by driver
  - Pickup done
  - Update location during delivery
  - Delivery done (by driver)
  - Delivery done (by client)

## Reads
This should be more traditional URLs, to allow for better caching at scale at all levels of the stack.
- Redis/Memcache at server side
- CDNs in the middle
- Browser caches on the client

Operations:
- List dishes
- List cuisines of dishes
- List restaurants by cuisine (tag)
- List restaurants by location
