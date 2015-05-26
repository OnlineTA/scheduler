package main

import (
  "log"
  "golang.org/x/exp/inotify"

  "github.com/onlineta/common"
)

func main() {
  cfg := common.Config{}
  if err := cfg.Parse("../onlineta.conf"); err != nil {
    log.Fatal(err)
    return
  }
  cfg.Serve()

  // Setup inotify listeners
  watcher, err := inotify.NewWatcher()
  if err != nil {
    log.Fatal(err)
    return
  }

  if err := watcher.AddWatch(common.ConfigValue("IncomingDir"), inotify.IN_CLOSE_WRITE); err != nil {
    log.Fatal(err)
    return
  }

  for {
    select {
    case ev := <- watcher.Event:
      log.Print(ev)
    case err := <- watcher.Error:
      log.Print(err)
    }
  }
}
