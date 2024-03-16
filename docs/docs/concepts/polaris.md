---
sidebar_position: 1
title: Polaris
---
# What is Polaris?


Polaris helps you create, store and run workflows. 

You don't need to worry about:
- Feeding the sequence of steps (it will figure out which the sequence along with the ones that can run concurrently)
- Explicitly pausing workflows (when it runs out of new data to move the workflow ahead, it pauses)

A workflow is a series of multiple steps, and can often be long running. 

Workflow orchestrators (like Polaris) help break workflows down into chunks, so they can be processed asynchronously. If a workflow is waiting on an event, state is stored and the logical execution of the workflow is paused (meaning, the CPU is free to move on with other tasks). When the event is received, state is recovered and execution is resumed.