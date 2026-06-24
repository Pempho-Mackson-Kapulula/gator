# Gator

Gator is a high-performance, concurrent CLI RSS feed aggregator built in Go. It allows users to manage, track, and read summaries of their favorite blogs, podcasts, and news sites directly from the terminal.  This project was built to master multi-layered backend development, database schema management, and concurrent data fetching in Go.


## Prerequisites

To run gator, you need the following installed:

- Go
- PostgreSQL

## Installation

Use `go install github.com/Pempho-Mackson-Kapulula/gator@latest` to install the `gator` CLI.

## Configuration

Create a file named `.gatorconfig.json` in your home directory.

Add the following JSON:

```json
{
  "db_url": "connection_string_goes_here",
  "current_user_name": "username_goes_here"
}
```



## Usage

After installing, run the CLI with:

```bash
gator <command>
```

**Example commands:**

```bash
gator register <name>
gator login <name>
gator addfeed <name> <url>
gator feeds
gator follow <url>
gator following
gator unfollow <url>
gator agg 30s
gator browse 10
```
