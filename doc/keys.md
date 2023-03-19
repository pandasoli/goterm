## Keys <img width='32' src='https://raw.githubusercontent.com/tstamborski/pixelart-icons/main/png/macos32.png' alt='Cmd16 Icon'/>


### `Getch() (string, error)` - Get a char from the input

Here I decided to return a `string` because if you press an accent and then a letter that accepts an accent (e.g.: `á`) it's alright, but if it doesn't (e.g.: `´b`) it would return only the accent, and also for special keys like `LeftArrow` (`\033[D`).

So instead of you have to do this:
```go
key, err := term.Getch()
if err != nil { panic(err) }

if key == '\033' {
  key, err := goterm.Getch()
  if err != nil { panic(err) }

  if key == '[' {
    key, err := goterm.Getch()
    if err != nil { panic(err) }

    if key == 'D' {
      fmt.Println("Pressed LeftKey")
    }
  }
}
```

Now, you can just do this:
```go
str, err := goterm.Getch()
if err != nil { panic(err) }

if str == "\033[D" {
  fmt.Println("Pressed LeftKey")
}
```

<br/>

### `KBHit() bool` - Returns if a key was pressed

Thus you can use the `Getch` to get it or not.  
So you don't have to pause the program (get `Getch` does it until a key get pressed).
