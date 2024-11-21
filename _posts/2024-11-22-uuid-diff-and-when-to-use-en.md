---
layout: post
title: Choosing UUIDs for Use
tags: uuid rfc
keywords: uuid, rfc
description: UUIDs are reference codes used to represent resources in a system. Designed
  to eliminate the need for a central authority when generating them, they reduce
  management costs while maintaining uniqueness and traceability. Over time, they
  have gained popularity as primary keys (PK) in databases, catering to the modern
  need for distributed systems.
date: 2024-11-22 04:01 +0700
---

## What is a UUID?

A UUID serves as a reference code for resources in a system. Its design principle
emphasizes decentralization in code generation to reduce management overhead while
ensuring unique, reusable identifiers.
Over time, UUIDs have become a popular choice as primary keys (PK) in databases,
supporting the shift towards distributed systems.

```sh
# Example:
e5c01d00-0cc6-41a9-86ef-91dab8b92219

# Format: 8-4-4-4-12 characters
# Total: 32 characters = 128 bits
```

## The Different Versions

The UUID specification has 7 well-defined versions, with version 4 and 7 as
the most commonly used in recent software development.

```sh
# Example: UUID 4
e07d07cb-afe4-4402-be35-dc431fc02764

# Example: UUID 7
0193503d-d8d0-7ae0-beca-c005478f5650

# The 13th digit indicates the version.
```

## Choosing Between V4 and V7

UUID 4 was not originally designed with databases in mind,
which resulted in difficulties with sorting and indexing. Because
UUID 4 doesn't have a natural ordering, making them unsuitable for use with
tree-based data structures like B-trees or binary search trees,
which are used in most opensource databases. Since all UUID 4 are randomly
generated, their insertion into such trees can result in inefficient storage
and slower query performance.

To address these issues, UUID 6 was introduced, and later UUID 7,
which relies on Unix Epoch Time for sequential ordering.
This makes UUID 7 ideal for databases.

> Use UUID 7 as the default. If you're starting a new system.

## V8

UUID 8 is an open-ended version designed for customization and flexibility
while remaining compatible with other UUIDs.

For example, the  [Go library](https://pkg.go.dev/github.com/samborkent/uuidv8#section-readme)
introduces UUID 8 with more time granularity compared to UUID 7. While UUID 7
use 1 millisecond precision, UUID 8 from this lib increases it to 20 microseconds,
enabling more IDs to be generated in the same timeframe.

## Null and The Opposite

If you need to represent an empty UUID in your system, use all zeros.

```sh
# Commonly used by databases for empty UUIDs:
00000000-0000-0000-0000-000000000000
```

Conversely, you can use the maximum value as a placeholder for fields
requiring a UUID.

```sh
# Opposite or maximum value:
FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF
```
