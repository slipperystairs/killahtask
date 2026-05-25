# killahtask
A simple TODO CLI tool for managing your tasks. The Cobra package was used to create the CLI interface. Tasks are written to a CSV file in the following header structure `task_id, description, created, completed`. The CSV file that `killahtask` reads/writes is locked to prevent concurrent access. The file gets created in your home directory: `~/killahtask_<username>.csv`

This CLI tool performs simple CRUD operations:
```
dilly@dilly:~$ killahtask add "my really cool task"
dill@dilly:~$ killahtask list
dilly@dilly:~$ killahtask complete <task_id>
dilly@dilly:~$ killahtask delete <task_id>
```

After cloning the repo, you can build the tool directly into /usr/local/bin (Note: `go build` needs to get executed in the project directory):
```bash
dilly@dilly:~/projects/killahtask$ sudo go build -o /usr/local/bin/killahtask
```
Then run:
```
killahtask
```
Output:
```
No commands were passed to the killah...see below
A todo task tool that performs simple CRUD operations.

Usage:
  killahtask [flags]
  killahtask [command]

Available Commands:
  add         Adds a new item
  complete    Completes an item on the list
  completion  Generate completion script
  delete      Delete an item in your list
  help        Help about any command
  list        List the items in your list

Flags:
  -h, --help   help for killahtask

Use "killahtask [command] --help" for more information about a command.
```

## Add
Add a task (aliases `add` or `a`):

```
dilly@dilly:~$ killahtask add my really cool task
Task "my really cool task" added successfully!
dilly@dilly:~$ killahtask a some other really cool task
Task "some other really cool task" added successfully!
```

Task descriptions must be unique.

```
dilly@dilly:~$ killahtask a some other really cool task
Task description isn't unique! "some other really cool task" already exists.
```

## List
By default, list shows incomplete tasks. Use `--all (-a)` to show all tasks.

**Default output:**
```
dilly@dilly:~$ killahtask list
ID     Description                    Created
0      my thing                       2 hours ago
2      Something else really cool     2 hours ago
3      My other thing                 an hour ago
```

**With `--all` output:**
```
dilly@dilly:~$ killahtask list --all
ID     Description                    Created         Completed
0      my thing                       2 hours ago     false
2      Something else really cool     2 hours ago     false
3      My other thing                 2 hours ago     false
4      My other thing 2               an hour ago     true
```

## Complete
Mark a task complete:
```
dilly@dilly:~$ killahtask complete
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
Delete a task:
```
dilly@dilly:~$ killahtask delete 0
Task removed successfully!
dilly@dilly:~$ killahtask list -a
ID     Description                    Created         Completed
2      Something else really cool     2 hours ago     false
3      My other thing                 2 hours ago     true
4      My other thing 2               2 hours ago     true
``` 

## Cowsay
The `add`, `list`, `complete`, and `delete` commands have a `--cowsay` subcommand that will display the output using `cowsay`.
```
dilly@dilly:~/projects/killahtask$ killahtask add "Do nothing" --cowsay
 ----------------------------------
| "Do nothing" added successfully! |
 ----------------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||

dilly@dilly:~/projects/killahtask$ killahtask list --all --cowsay
 ---------------------------------------------------------------
| ID     Description            Created               Completed |
| 0      Take out the trash     2 minutes ago         false     |
| 1      Go to sleep            2 minutes ago         false     |
| 2      Write some code        a minute ago          false     |
| 3      Walk dog               a minute ago          false     |
| 4      Do nothing             a few seconds ago     false     |
 ---------------------------------------------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||

dilly@dilly:~/projects/killahtask$ killahtask complete 0 --cowsay
 ------------------------------
| ID 0 was marked as complete! |
 ------------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||

dilly@dilly:~/projects/killahtask$ killahtask delete  0 --cowsay
 ------------------------------
| Task 0 removed successfully! |
 ------------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

## Tab Completion
Cobra can generate completion scripts for several shells.
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
