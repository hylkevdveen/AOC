# AOC

Hylke's AOC solutions

## Set up a day

You can use [`startday.sh`](startday.sh) to set up the template script for your preferred programming language and fetch the `input.txt` file.

```
./startday.sh <yyyy> <d> [-l <go|ipynb>]
```

### Session cookie

Because you must be logged in to get your input, you need to get your Advent of Code session ID and set the `$AOC_SESSION` environment variable.
The script will try to read it from the `.env` file if it exists so you can store it there.

You can find your session ID by visiting the input page of any one challenge, opening the Network tab on your browsers inspection tool and looking for the Cookie field under Request Headers on the input request.

### Code templates

Add a folder [template](./template) in which you store the templates to your preferred language as plaintext files using just the scripts file extension as name (e.g. `template/ipynb`). 

You might have to add this language to the script if it's not already included.

_Happy coding!_
