[OBSERVATORY LANDING PAGE](../../README.md) | [TO ADRs List](./index.md)

### Title
# Choose Worker Protocol (Jan 29, 2024)

## Status

What is the status, such as proposed, **accepted**, rejected, deprecated, superseded, etc.?

## Context

We need to collect information about availability from remote server/endpoint/IP.

For now wervice should work with the next prerequisites:
- request sends different size of information, depends on user's needs;
- we should anderstand whether endpoint alive;
- on response we know or can calculate latency;

## In consideration

Under consideration:
1. ICMP request (standart `ping` command) [MVP](https://github.com/olegsobchuk/go-observatory/commit/e332febb4bef9bbdb2afdb99072024cd3f7f3bf5#diff-f3745aa1c297d2f52cadee411b82233718223881de1b71771ec3bfbe37275e4a) - declined as it requires admin/root privileges for building and sending ICMP reqiest.
2. TCP request [MVP](https://github.com/olegsobchuk/go-observatory/commit/e332febb4bef9bbdb2afdb99072024cd3f7f3bf5#diff-e8a1113f516e1e586a914c2d9bb1794b12f3d21949f82bb34d184cde208de00c)

## Decision :star:

TCP request

## Positive Consequences

TCP protocol could be used without admin or root privileges.

## Negative Consequences

\- (i'm not sure)

## Links

\- no
