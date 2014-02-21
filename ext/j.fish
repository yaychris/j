function j
  if test 0 -eq (count $argv)
    echo "usage:"
    echo "  j <regex>    Jump to the directory matching <regex> (matched by frecency)"
    echo "  j -l <regex> Show all directories matching <regex>"
    echo "  j --add      Add the current directory to the database"

    return
  end

  switch $argv[1]
  case "--add"
    # add the current directory
    eval $J_COMMAND --add=(pwd)
  case "-l"
    # list all matches
    eval $J_COMMAND --match=$argv[2..-1]
  case '*'
    # perform a search and cd to the first matching directory
    set dir (eval $J_COMMAND --match="$argv" --limit=1)

    if test -d "$dir"
      cd $dir
    end
  end
end
