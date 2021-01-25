Shows the time offset to the given NTP server in various formats. Breaks if the time offset reaches a given limit.

Useful for initContainers when accurate time is required (e.g. at minikube, when the host has  not yet synchronized the time after wakeup).


## Usage

    -format string
          output formats:
            s    seconds
            ms   miliseconds
            us   microseconds
            h    human readable
           (default "ms")
    -limit int
          exit(2) if offset diff is greather than <n> ms
    -ntpserver string
          ntp server hostname
    -quiet
          suppress output to standard output

## Docker example

    docker run --rm jb5r/clockoffset \
      -ntpserver time.google.com \
      -format h \
      -limit 20000
