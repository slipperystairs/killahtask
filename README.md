# killahtask
A simple TODO CLI tool for managing your task. The cobra package was used to create the CLI interface. Task are written to a CSV file in following header structure `task_id, description, created, completed`. The CSV file that `killahtask` reads/writes is locked to prevent concurrent access.

This CLI tool performs simple CRUD operations:
```
dilly@dilly:~$ killahtask add "my really cool task"
dill@dilly:~$ killahtask list
dilly@dilly:~$ killahtask complete <task_id>
dilly@dilly:~$ killahtask delete <task_id>
```

## Add
Running `killahtask add "my task"` will add a task to your to a CSV file in your home directory (not OS restricted). This command can get executed with an alias (i.e., `killahtask a "my task"`). Both will patterns will achieve the same thing.

```
dilly@dilly:~$ killahtask add "my really cool task"
Task "my really cool task" added successfully!
dilly@dilly:~$ killahtask a "some other really cool task"
Task "some other really cool task" added successfully!
```

The `add` command does require the task to be wrapped in double quotes. Otherwise the CLI will throw an error.

```
dilly@dilly:~$ killahtask a task without double quotes
Too many arguments passed to the "add" command
Usage: killahtask add "my description"
```

Task descriptions are also required to be unique.

```
dilly@dilly:~$ killahtask a "some other really cool task"
Task description isn't unique! "some other really cool task" already exist.
```

## List
By default the `list` command will display task that are not completed (i.e., "Completed" is set to `false`). This command has a sub flag called `--all` (short  hand: `-a`) that prints all task. 

**Default output:**
```
dilly@dilly:~$ killahtask list
ID     Description                    Created
0      my thing                       2 hours ago
2      Something else really cool     2 hours ago
3      My other thing                 an hour ago
```

**Sub flag output:**
```
dilly@dilly:~$ killahtask list --all
ID     Description                    Created         Completed
0      my thing                       2 hours ago     false
2      Something else really cool     2 hours ago     false
3      My other thing                 2 hours ago     false
4      My other thing 2               an hour ago     true
```

## Complete
Marking a task complete will set the "Completed" value to true, indicating that you finished your task. Passing the `task_id` as an argument is required.
```
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
This command will remove the task from your CSV file. A `task_id` must be passed as an argument is required for the command to succeed.
```
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
Cobra can generate completion scripts for most shells.
To enable completion:
1. Generate a completion script for your shell:
2. Source the file (or move it to your shell's completion directory).
3. Restart your shell.

Supported shells:
- bash
- zsh
- fish
- PowerShell

**Linux example:**
```bash
# Install bash completion (Linux)
sudo apt-get update
sudo apt-get install bash-completion

# Enable bash completion for this session
source /etc/profile.d/bash_completion.sh

# Generate completion script
killahtask completion bash > ~/.killahtask_completion

# Load completion immediately
source ~/.killahtask_completion

# (Optional) Load automatically in every new terminal
echo "source ~/.killahtask_completion" >> ~/.bashrc

# Example usage:
# Type "killahtask c" then press TAB to complete to "complete"
```
