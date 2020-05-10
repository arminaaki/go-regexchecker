# go-regexchecker

go-regexchecker is a simple tool for validating strings given a Regex pattern.

## Installation

Download and extract the [latest release](https://github.com/arminaaki/go-regexchecker/releases/latest) to your user directory.

## Usage
```sh
$ ./go-regexchecker -rule "/(?<month>\d{1,2})\/(?<day>\d{1,2})\/(?<year>\d{4})/"

_______      ________     ______    ________   ____  ____
|_   __ \    |_   __  |  .' ___  |  |_   __  | |_  _||_  _|
  | |__) |     | |_ \_| / .'   \_|    | |_ \_|   \ \  / /
  |  __ /      |  _| _  | |   ____    |  _| _     > '' <
 _| |  \ \_   _| |__/ | \ '.___]  |  _| |__/ |  _/ /''\ \_
|____| |___| |________|  '._____.'  |________| |____||____|
   ______   ____  ____   ________     ______   ___  ____    ________   _______      _   _
 .' ___  | |_   ||   _| |_   __  |  .' ___  | |_  ||_  _|  |_   __  | |_   __ \    | | | |
/ .'   \_|   | |__| |     | |_ \_| / .'   \_|   | |_/ /      | |_ \_|   | |__) |   | | | |
| |          |  __  |     |  _| _  | |          |  __'.      |  _| _    |  __ /    | | | |
\ '.___.'\  _| |  | |_   _| |__/ | \  .___.'\  _| |  \ \_   _| |__/ |  _| |  \ \_  |_| |_|
 '.____ .' |____||____| |________|  '.____ .' |____||____| |________| |____| |___| (_) (_)
==========================================================================================

Today's date is: 5/9/2020.
✔ Today's date is: 5/9/2020.
year  => 2020
month => 5
day   => 9

The invalid date is: 5-8-2020
✘ The invalid date is: 5-8-2020
```
