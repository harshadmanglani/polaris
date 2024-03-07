---
sidebar_position: 1
---

# Introduction

## What is a workflow?
A workflow is a series of multiple steps, and can often be long running. 

Workflow orchestrators (like Polaris) help break workflows down into chunks, so they can be processed asynchronously. If a workflow is waiting on an event, state is stored and the logical execution of the workflow is paused (meaning, the CPU is free to move on with other tasks). When the event is received, state is recovered and execution is resumed.

## Use cases
1. You have multi-step workflow executions where each step is dependent on data generated from previous steps.
2. Executions can span one request scope or multiple scopes.
3. Your systems works with reusable components that can be combined in different ways to generate different end-results.
4. Your workflows can pause, resume or even restart from the beginning.

## Limitations
1. Workflow versioning is tricky to implement:
   1. Unless you can afford a 100% downtime ensuring all active workflows move into a terminal state, deploying new code requires ensuring backward compatibility.
   2. What this means is - you'll need to a deploy a version of code that is backward compatible for older non terminal workflows while newer ones will execute on the new code.
   3. Once the older workflows have completed, a deployment to clean up stale code will be required.
2. The level of abstraction is lower in this framework compared to Cadence, Conductor:
   1. Workflows can be made fault oblivious if there is an external (reliable) service giving callbacks per workflow id.
   2. Instrumentation can be set up by adding your custom code to push events via listeners.

## How does the framework perform at scale?
The framework itself has extremely low overhead. Since execution graphs are generated pre-runtime, all the orchestrator will do at runtime is use the graph and available data to run whichever builders can be run. 