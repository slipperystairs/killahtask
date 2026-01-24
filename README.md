# killahtask
A simple TODO CLI tool for managing your task. The cobra package was used to create the CLI interface. Task are written to a CSV file in following header structure `task_id, description, created, completed`. The CSV file that `killahtask` reads/writes is locked to prevent concurrent access.

This CLI tool performs simple CRUD operations:
```bash
dilly@dilly:~$ killahtask add "my really cool task"
dill@dilly:~$ killahtask list
dilly@dilly:~$ killahtask complete <task_id>
dilly@dilly:~$ killahtask delete <task_id>
```

## Add
Running `killahtask add "my task"` will add a task to your to a CSV file in your home directory (not OS restricted). This command can get executed with an alias (i.e., `killahtask a "my task"`). Both will patterns will achieve the same thing.

```bash
dilly@dilly:~$ killahtask add "my really cool task"
Task "my really cool task" added successfully!
dilly@dilly:~$ killahtask a "some other really cool task"
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
By default the `list` command will display task that are not completed (i.e., "Completed" is set to `false`). This command has a sub flag called `--all` (short  hand: `-a`) that prints all task. 

**Default output:**
```bash
dilly@dilly:~$ killahtask list
ID     Description                    Created
0      my thing                       2 hours ago
2      Something else really cool     2 hours ago
3      My other thing                 an hour ago
```

**Sub flag output:**
```bash
dilly@dilly:~$ killahtask list --all
ID     Description                    Created         Completed
0      my thing                       2 hours ago     false
2      Something else really cool     2 hours ago     false
3      My other thing                 2 hours ago     false
4      My other thing 2               an hour ago     true
```

## Complete
Marking a task complete will set the "Completed" value to true, indicating that you finished your task. Passing the `task_id` as an argument is required.
```bash
dilly@dilly:~$ killahtask complete
Task ID is missing!
Usage: killahtask complete <task_id>
dilly@dilly:~$ killahtask complete 123 123
Too many arguments passed to the "complete" command.
Usage: killahtask complete <task_id>
dilly@dilly:~$ killahtask complete 4
ID 4 was marked as complete!
dilly@dilly:~$ killahtask list -a
ID     Description                    Created         Completed
0      my thing                       3 hours ago     false
2      Something else really cool     2 hours ago     false
3      My other thing                 2 hours ago     true
4      My other thing 2               2 hours ago     true
```

## Delete
This command will remove the task from your CSV file. A `task_id` must be passed as an argument is required for this command to succeed.
```bash
dilly@dilly:~$ killahtask delete
Task ID is missing!
Usage: killahtask delete <task_id>
dilly@dilly:~$ killahtask delete 12903 123
Too many arguments passed to the "delete" command.
dilly@dilly:~$ killahtask delete 0
Task removed successfully!
dilly@dilly:~$ killahtask list -a
ID     Description                    Created         Completed
2      Something else really cool     2 hours ago     false
3      My other thing                 2 hours ago     true
4      My other thing 2               2 hours ago     true
``` 
## Tab Completion
