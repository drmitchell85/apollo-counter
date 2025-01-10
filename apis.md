# APIs

## Current APIs
GET /health
    - basic ping to the sever

### /users
GET /users/email/{email}
    - get a user by their email

POST /users/newUser
    - create a new user

### /events
POST /events/newEvent
    - create a new event
    - add to PostgreSQL
    - cache in Redis

## Future APIs
### For Redis Cachine Practice
GET /events/{id}
    - Perfect for caching as event details rarely change
    - High read-to-write ratio makes it ideal for learning cache invalidation

GET /events/popular
    - List of most viewed/purchased events
    - Good for learning sorted sets in Redis

### For Message Queue Practice
POST /orders
    - Create a new order (async processing)
    - Perfect for learning how to handle distributed transactions


POST /events/{id}/waitlist
    - Add user to waitlist
    - Good for learning pub/sub patterns with RabbitMQ

### For Rate Limiting Practice
GET /events/search?query={query}
    Search endpoint that you can rate limit


POST /orders/bulk
    - Bulk operation that should be rate limited
    - Good for testing token bucket algorithm

### For Load Balancer Practice
GET /events/stats
    - Returns different stats from different server instances
    - Good for learning load distribution


GET /system/status
    - Returns server instance details
    - Useful for health checks