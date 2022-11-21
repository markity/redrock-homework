package tools

// database connection info, you need to create database project6 first
var dns = "root:mark2004119@/project6"

// if there is something wrong with query, retry most 3 times
var queryRetryAttempt = 3

func GetDNS() string {
	return dns
}

func GetQueryRetryAttempt() int {
	return queryRetryAttempt
}
