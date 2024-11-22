# `change-go-pkgs`

Get Go packages that changed between two Git versions, usage:

    NAME:
       changed-go-packages - Get the changed Go packages between two commits
    
    USAGE:
       changed-go-packages [global options] command [command options]
    
    COMMANDS:
       help, h  Shows a list of commands or help for one command
    
    GLOBAL OPTIONS:
       --from-ref value
       --to-ref value
       --repo-dir value   The Git repo to inspect (default: ".")
       --mod-dir value    Path to the directory containing go.mod. Used to find local packages (default: ".")
       --log-level value  The level to log at. Valid values are: debug, info, warn, error (default: WARN)
       --help, -h         show help
