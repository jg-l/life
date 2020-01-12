# life

Life is a simple command-line application to keep track of life's events.

![life in action](screeenshot.png)

# Features
- Keep track of your life with micro bloggin
- Read and Append only. Mistakes are a part of life.
- Cross-platform.

# Usage

`$ life`

Displays all of the _life_ entries in reverse chronological order (latest at the bottom)

`$ life "My businesss was sold for $100,000 USD!"`

Will add the entry to the life log.

# Your life log

Your _.life_ database can be found in `$HOME/.life`

# Install

`go get github.com/jg-l/life`

# Tricks

## Tag and Filter entries using grep

- Add a hashtag or any symbol
`$ life "Meditated for 50 minutes! #spiritual"`

- Filter life log using grep
`$ life | grep "#spiritual"`

## Delete Entries

If you REALLY want to delete an entry, you can get [boltbrowser](https://github.com/br0xen/boltbrowser) and open the the life log at `$HOME/.life`



# Future features (probably not...)
- Filter entries based on date ranges
