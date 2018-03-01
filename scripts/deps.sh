#!/usr/bin/env bash

a="""
github.com/spf13/cobra
github.com/spf13/viper
github.com/fsnotify/fsnotify
github.com/gin-contrib/sse
github.com/go-kit/kit/log
github.com/go-logfmt/logfmt
github.com/spf13/afero
github.com/spf13/cast
github.com/spf13/jwalterweatherman
github.com/spf13/pflag
"""

echo $a | xargs gopm get -l