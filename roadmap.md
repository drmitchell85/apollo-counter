# Roadmap

## Plan
1. API Service (current)

2. Cache Service (before CLI)

3. CLI Tool

4. Message Queue

5. Load Balancer

6. Rate Limiter

## Reasoning
Implementing caching before CLI tool will give us more meaningful metrics to monitor with our CLI

Message queue implementation will help us understand distributed systems better before adding a load balancer

Rate limiting works best when we have multiple instances behind a load balancer,
so it makes sense to implement this last.