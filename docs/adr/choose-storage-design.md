# Choose storage design

[OBSERVATORY LANDING PAGE](../../README.md) | [TO ADRs List](./index.md)

## Status

**proposed**, accepted, rejected, deprecated, superseded

## Context

We need to have storage to collect and keep information about targets (IP addresses or hosts). User will be able to send information (IP or host name, port, etc.) through our endpoint and web layer should be able to retain it somewhere. And this information should be able on demand.

Requirements:

1. it is easy to implement;
2. data from user is passed to this storage
3. data in storage should be changed with the next user's request with new data
4. storage collects only valid information
5. storage has multiple buckets:

    - information about target;
    - information about results;

## In consideration

- in application (runtime) memory;
- database - declined;
- browser local storage declined, but we can return to this in future;

## Decision :star:

For now it is **in application memory**

## Positive Consequences

- we don't need any additional application;
- the simplest way for implementation;
- we can use in different browsers;

## Negative Consequences

- all data is lost on reload application

## Links

Links to related ADR(s)
