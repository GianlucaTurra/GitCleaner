# GitCleaner

A simple TUI tool, developed with golang's BubbleTea library, to manage local
git branches not present in the remote repository.  
Performs `git fetch -p` at startup to ensure the refs are up to date.
**Breware that no confirmation is asked to perform the git commands**.

## Honest disclaimer

If you want a real tool to manage git repository look out to
[LazyGit](https://github.com/jesseduffield/lazygit), this one is just a fancy
laerning project. :)
