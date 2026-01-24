# killahtask
A simple TODO CLI tool for managing your task. This The cobra package was used to create the CLI interface. Task are written to a CSV file in following header structure `task_id, description, created, completed`. The CSV file that `killahtask` reads/writes to gets locked to prevent concurrent access.

This CLI tool performs simple CRUD operations:
```bash
dilly@dilly:$ killahtask add "my really cool task"
dill@dilly:$ killahtask list
dilly@dilly:$ killahtask complete <task_id>
```

## Add
Running `killahtask add "my task"` will add a task to your to a CSV file in your home directory (not OS restricted). This command can also be used with an alias (i.e., `killahtask a "my task"`). Both will patterns will achieve the same thing.

```bash
dilly@dilly:$ killahtask add "my really cool task"
Task "my really cool task" added successfully!
dilly@dilly:$ killahtask a "some other really cool task"
Task "some other really cool task" added successfully!
```

The `add` command does require the task to be wrapped in double quotes. Otherwise the CLI will throw an error.

```bash
dilly@dilly:~$ killahtask a task without double quotes
Too many arguments passed to the "add" command
Usage: killahtask add "my description"
```

Task descriptions are also required to be unique.

```bash
dilly@dilly:~$ killahtask a "some other really cool task"
Task description isn't unique! "some other really cool task" already exist.
```

## List

## Completed

## Delete

## Tab Completion
