year=$1
day=$2
language=$3
[[ "$year" =~ ^[0-9]{4}$ && $year -ge 2015 && $year -le $(date +%Y) ]] || { echo "'$year' is not a valid year"; exit 1; }
[[ "$day" =~ ^[0-9]{1,2}$ && $day -ge 1 && $day -le 31 ]] || { echo "'$day' is not a valid day"; exit 1; }


case $language in
  "go" ) echo "Setting up day for Go" ;;
  "ipynb" ) echo "Setting up day for Jupyter Notebook" ;;
  "" ) { echo "No language chosen, proceeding with Go"; language="go"; } ;;
esac

dir="$year/day$day"
aoc="https://adventofcode.com/$year/day/$day/input"

if [ ! -f .env ]; then
  export "$(xargs < .env)"
fi

if [ "$AOC_SESSION" = "" ]; then
  echo "Set the 'AOC_SESSION' environment variable to your Advent of Code session cookie"
fi

if test -d "$dir"; then
  echo "Directory ./$dir already exists"
  exit 0
fi
mkdir -p "$dir"

echo "Fetching $aoc"

curl -fsS --request GET "$aoc" -H "cookie: session=$AOC_SESSION" -o "$dir/input.txt" || {
  echo "It's possibly too soon access this day";
  rmdir "$dir";
  exit 1;
}

echo "Created $dir/index.txt"

case $language in
  "go" ) { cp template/go "$dir/main.go"; echo "Created $dir/main.go"; } ;;
  "ipynb" ) { cp template/ipynb "$dir/$day.ipynb"; echo "Created $dir/$day.ipynb"; } ;;
esac
