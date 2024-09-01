
# Bash Logger

Bash Logger is a lightweight logging tool designed to capture and log every command executed in Bash, sending the log information to a specified Telegram chat. The tool is composed of two main parts: a shell script `history.sh` and a Go program `history_log`.

## Features

- Automatically captures each command entered in Bash.
- Logs the command execution time, user, terminal information, and working directory.
- Sends these log entries to a pre-configured Telegram chat for real-time monitoring.

## File Structure

- `history.sh`: Captures user-entered Bash commands and passes them to the `history_log` program.
- `history_log`: A Go program that sends the log information to the specified Telegram chat using the Telegram API.

## Installation Steps

1. **Install Go Environment**  
   If you haven't installed Go yet, follow the [official installation guide](https://golang.org/doc/install).

2. **Compile the `history_log` Program**  
   Navigate to the project directory and run the following command to compile the `history_log` program:

   ```bash
   go build -o /usr/local/bin/history_log history_log.go
   ```

   This will generate the executable `history_log` in the `/usr/local/bin/` directory.

3. **Configure `history.sh`**  
   Copy the `history.sh` script to your system and ensure it loads automatically when a user logs in. You can add the following line to `/etc/profile` or the user's `~/.bash_profile`:

   ```bash
   source /path/to/history.sh
   ```

   Replace `/path/to/history.sh` with the actual script path.

## Usage

Once configured, `history.sh` will automatically capture commands executed by users and invoke the `history_log` program to send the log information to the configured Telegram chat.

The log format is as follows:

```
YYYY-MM-DD HH:MM:SS ## HOSTNAME ## LOGIN_USER ## LOGIN_TTY ## LOGIN_IP ## USER=current_user PWD=current_directory CMD=executed_command
```

Example:

```
2024-09-01 14:32:10 ## myserver ## root ## pts/0 ## 192.168.1.100 ## USER=root PWD=/root CMD=ls -la
```

## Custom Configuration

- **Telegram API Token**: Defined as `tgAPIToken` in `history_log.go`.
- **Telegram User ID**: Defined as `tgUserID` in `history_log.go`.

You can modify these constants to configure the target chat.

## Security

As log entries may contain sensitive data, ensure that the information is appropriately protected during transmission and storage.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.
