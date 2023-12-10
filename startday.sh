usage() { echo "Usage: $0 <yyyy> <d> [-l <go|ipynb>]" 1>&2; exit 1; }

year=$1
[[ "$year" =~ ^[0-9]{4}$ && $year -ge 2015 && $year -le $(date +%Y) ]] || { echo "'$year' is not a valid year"; usage; }
shift
day=$1
[[ "$day" =~ ^[0-9]{1,2}$ && $day -ge 1 && $day -le 31 ]] || { echo "'$day' is not a valid day"; usage; }
shift

# language

language=""

while getopts "l:" opt; do
  case "$opt" in
    "l" ) language=${OPTARG} ;;
    * ) usage
  esac
done

case $language in
  "go" ) echo "Setting up day for Go" ;;
  "ipynb" ) echo "Setting up day for Jupyter Notebook" ;;
  "" ) { echo "No language chosen, proceeding with Go"; language="go"; } ;;
  * ) { echo "'$language' is not a valid programming language. If you have a template for it, add the language to the case in this script."; usage; }
esac

# Session cookie

if [ -f .env ]; then
  export "$(xargs < .env)"
fi

if [ "$AOC_SESSION" = "" ]; then
  echo "Set the 'AOC_SESSION' environment variable to your Advent of Code session cookie"
  exit 1
fi

# creating directory

aoc="https://adventofcode.com/$year/day/$day"
echo "Preparing Advent of Code $year day $day: $aoc"

dir="$year/day$day"

if test -d "$dir"; then
  echo "Directory ./$dir already exists"
  exit 0
fi
mkdir -p "$dir"

# input.txt

echo "Fetching $aoc/input"

curl -fsS --request GET "$aoc/input" -H "cookie: session=$AOC_SESSION" -o "$dir/input.txt" || {
  echo "It's possibly too soon access this day";
  rmdir "$dir";
  exit 1;
}

echo "Created $dir/index.txt"

# main script

if ! test -f "template/$language"; then
  echo "Template template/$language does not exist"
  exit 1
fi

cp "template/$language" "$dir/main.$language"
echo "Created $dir/main.$language"
