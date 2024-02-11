# Choose Backend Program Language (Jan 26, 2024)

[OBSERVATORY LANDING PAGE](../../README.md) | [TO ADRs List](./index.md)

## Status

proposed, **accepted**, rejected, deprecated, superseded

## Context

We have to choose the program language for futher developing. Program language should meet requirements:

- popular;
- well known or at least it is not hard to figure out in the code;
- better to have executable file for whole application;

## In consideration

Due to I'm only one who is working on this progect, I'm considering the next program languages:

1. Ruby (pure Ruby or Sinatra framework) - rejected:

    \+ easy to learn and understand;

    \- slow;

    \- not compilable, so it is easy to brake code or add there Malware.

2. Golang (Fiber or Echo frameworks):

    \+ easy to learn and understand;

    \+ fast;

    \+ compilable;

    \- requires developers with specific knowlage in this language;

3. Javascript (NextJS & ExpressJS) - rejected:

    \+ easy to learn and understand;

    \- not compilable, so it is easy to brake code or add there Malware.

    \- requires developers with specific knowlage in these frameworks;

## Decision :star:

Golang on Fiber framework

## Positive Consequences

We can take into process many endpoints, it is not a big dial for Golang to process and group information for dozens processes.

It is compiled, so it's more safer to use it on local PC

## Negative Consequences

Golang requires specific knowlage i this language and in this framework. So it could take a while for code amendment.

## Links

- no links yet
