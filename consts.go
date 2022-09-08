package main

const ALPHA = "a"
const ALPHANUMERIC = "c"
const NUMERIC = "n"

var BBAN_TO_REGEX = map[string]string{ALPHA: "[A-Z]", ALPHANUMERIC: "[A-Za-z0-9]", NUMERIC: "[0-9]"}
