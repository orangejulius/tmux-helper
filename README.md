# Julian's tmux helper

A very simple Tmux helper command.

If called with no arguments, in a directory following the format `~/repos/dir1/dir2`, creates a new tmux session named `dir1/dir2`
If the directory does not match that format, runs `tmux attach`

If an argument is passed, it runs `tmux new-session -s $arg1`
