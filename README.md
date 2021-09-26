# Corndogs
Cloud Native Background Task state manager for kubernetes and scaling out

Inspired by Celery and Sidekick, but meant to be language and UI agnostic.

## Status
Building a workflow. We are pre-alpha and not to be used for anything. Don't look at me, I'm hideous!

We are a deploy-first, CI/CD oriented project, and thus are working on deployment before feature work.

# Intended Design
## Data Structures

A "task" is just a row in a db with the following fields:
 * Task UUID
 * Queue string
 * Current State string
 * Auto Target State (the state to move it to when picked up or timed out)
 * Submit time
 * Update time
 * Timeout (null if waiting for pulling, otherwise number of seconds until it "times out")
 * Payload (a bytestring, package it however you want. JSON, msgpack, whatever)


## Flow

Corndogs doesn't work on tasks. It is a task state manager.

Submitters submit to a queue. If a state string isn't given, it will be "submitted" and the auto target state will be "submitted-working" which for simple workflows should be fine. Any time something is submitted with a state but not an auto target state "-working" will be added to the state for the auto target state.

Workers can pull a new task from a queue and state. State defaults to "submitted" and it will get the next task based on Submit Time, not Update Time. This means a task that is failing will keep getting picked up. Error handling and submitting to "dead letter" queues is a responsibility of clients.

When workers "complete" a task, they submit it with an optional next task. This means they can submit the next state of the workflow in an atomic way. This optional "next_task" is the same as the submit requirements.

This allows simple and complicated workflows, alongside workers dedicated to each phase of a workflow. This should also allow highly horizontally scalable workloads using an appropriate datastore.

## Supported datastores

Current targets are Postgres and CockroachDB. The design is such that Corndogs should not know where it's storing things, it just gets pointed to a URL and picks up where it needs to. This is Cloud Native.

## Metrics and such

Aside from logs and Prometheus metrics, a number of endpoints should be provided that allow intelligence around operation of or working against tasks. Some examples of statistical metrics that should be available in some way through normal request methods:

 * How many and what queues exist?
 * How many tasks are in flight in each queue? For each state?
 * How many tasks are waiting in each queue? For each state?
 * How many timeouts for each queue in each state?
 * What workers exist for each queue in the last N minutes?
 * What workers have what timeouts for a given window?
 * How many tasks completed in a time window? (This may allow historic views)
 * Length of time for tasks to be worked.

