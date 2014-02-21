              .__.
             /  /
      .__.  /  /
     /  /  /  /
    /  /__/  /
    \_______/


Port of rupa/j to Go.


Installation
------------

    go build
    cd ~/.config/fish/functions
    ln -s /path/to/repo/ext/j.fish

Add this inside your `fish_prompt` function:

    function fish_prompt
      j --add

      # other fish prompt stuffs
    end

Add the environment variables to your fish config:

    set -x J_COMMAND /path/to/repo/jgo
    set -x J_DATA "$HOME/.j"

`cd` around to build up the database.

Commands
--------

    j thing      cd to the highest ranking directory matching "thing"
    j -l thing   list all matches for "thing"

