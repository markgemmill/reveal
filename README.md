# Reveal 

Show thy daily password!

Reveal is a convenience app to help store and read a single 
password that you need repeated access throughout the day. 

### How It Works

When you open Reveal you are presented with a START button. 
Clicking the START button will copy the last clipboard value
into Reveal. 

Reveal shows the password in large characters, split across
multiple lines for easy of reading. (The number of character
per line can be configured with a value between 3 and 5.)

The password will remain visible for the number of settings
you have configured (default is 30).

The copy button will copy the password to the clipboard.

The reveal button will display the password again.

### Configuration

The configuration files can be found at C:\Users\<username>\AppData\Local\reveal\config.toml.

Configuration consists of four entries:

```
timeout = 30
segment_size = 4
copy_timeout = 10
dock_position = "ur"
```

`timeout` is a value in seconds.
`segement_size` can be a value between 3 and 5.
`copy_timeout` is a value in seconds and controls how long the value will be available on the clipboard. 
`dock_position` can be one of "center", "ul", "ur", "ll", "lr" (upper/lower/left/right)


### Install

Download the latest release package, unzip and place the exe wherever you 
choose.