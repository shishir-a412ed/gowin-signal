# gowin-signal

Windows is not POSIX compliant and doesn't support signals.
If your usecase is just to send a SIGTERM or gracefully shutdown a process,
utilities like taskkill are sufficient.

$$ taskkill /? (for help)
$$ taskkill /pid <pid> (to send SIGTERM or gracefully shutdown a process)

For custom or other signals, gowin-signal could be used.

# Usecase

I have a golang application (Let's call it myApp) running in windows.
 - myApp has a signal handler which can reload myApp config files upon receiving a SIGHUP. 
 - myApp has a signal handler which can take a core dump upon receiving a SIGQUIT (SIGBREAK in windows).

How do I test my signal handler?

# Setup

1. git clone git@github.com:shishir-a412ed/gowin-signal.git
2. make build
3. make clean (to cleanup the executable)

# How to send a signal to a windows process


