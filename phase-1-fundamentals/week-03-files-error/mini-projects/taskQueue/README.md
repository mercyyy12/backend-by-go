# Producer-Consumer Task Queue

## Overview
This is a **simple Producer-Consumer Task Queue**.
**Objective:**

- Add different types of tasks: **Email**, **Image**, **Report**
- List all tasks in the queue
- Process tasks concurrently using worker goroutines

## Features

- Users can enter tasks of different types with payloads
- **Concurrent processing:** Uses a buffered channel and multiple goroutines (workers)
- **Task tracking:** Each task has a unique ID and timestamp
