### Customers service

![Customers](/github/logo.png)

## Description

The service collects information about buyers from the CSV file and sends it to the test server (by default `localhost:9090/messages`), the test service is called commservice.

Sends this data to it according to the schedule.

### Example

InvoicesSchedule service:
```
go run main.go
2018/01/29 04:03:58 [INFO] All invoices for customer: 'bskentelberyl@mozilla.org' have been sent
2018/01/29 04:04:01 [INFO] All invoices for customer: 'charrimanr@ucla.edu' have been sent
2018/01/29 04:04:13 [INFO] All invoices for customer: 'esealove2@go.com' have been sent
2018/01/29 04:04:14 [INFO] All invoices for customer: 'rforrester11@indiegogo.com' have been sent
2018/01/29 04:04:15 [INFO] All invoices for customer: 'vdaybell0@seattletimes.com' have been sent
2018/01/29 04:04:15 [INFO] All invoices for customers have been sent!
```

Commservice service:
```
2018/01/29 04:03:51 listening on port 9090
2018/01/29 04:03:55 POST /messages 169.723µs
2018/01/29 04:03:57 POST /messages 72.154µs
2018/01/29 04:03:58 POST /messages 127.391µs
2018/01/29 04:04:01 POST /messages 70.104µs
2018/01/29 04:04:03 POST /messages 129.003µs
2018/01/29 04:04:06 POST /messages 66.062µs
2018/01/29 04:04:09 POST /messages 80.26µs
2018/01/29 04:04:13 POST /messages 74.289µs
2018/01/29 04:04:14 POST /messages 113.012µs
2018/01/29 04:04:15 POST /messages 47.797µs
^C

Message Report
==================================================

Message Count
 ✓  3 messages for "vdaybell0@seattletimes.com" received
 ✓  2 messages for "esealove2@go.com" received
 ✓  1 messages for "bskentelberyl@mozilla.org" received
 ✓  1 messages for "charrimanr@ucla.edu" received
 ✓  3 messages for "rforrester11@indiegogo.com" received

Message Timings
 ✓  1. message for "vdaybell0@seattletimes.com" received after 8s
 ✓  2. message for "vdaybell0@seattletimes.com" received after 14s
 ✓  3. message for "vdaybell0@seattletimes.com" received after 20s
 ✓  1. message for "esealove2@go.com" received after 0s
 ✓  2. message for "esealove2@go.com" received after 18s
 ✓  1. message for "bskentelberyl@mozilla.org" received after 3s
 ✓  1. message for "charrimanr@ucla.edu" received after 6s
 ✓  1. message for "rforrester11@indiegogo.com" received after 2s
 ✓  2. message for "rforrester11@indiegogo.com" received after 11s
 ✓  3. message for "rforrester11@indiegogo.com" received after 19s

==================================================
```

### How to run?

```
go get -u github.com/pcherednichenko/invoicesSchedule
```
Then from folder
```
./commservice
go run main.go
```

Created by [Pavel Cherednichenko](https://www.linkedin.com/in/pavel-cherednichenko-0a2a0b118/)