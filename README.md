# ClickbankGo
A Golang library to help parse the incomming json notification from the Clickbank Instant Payment Notification service.

Save yourself from having to loop through the elements or create the huge struct represenation of the notification.

## Getting Started
Start off by grabbing the library and putting it into your `$GOPATH`.

```
$ go get github.com/centricwebestate/clickbankgo
```

## Example webserver
An example web server is located under the cmd/ folder in the example.go file.

Now in your terminal run

```sh
$ go run webserver.go
```

Now in your browser navigate to `http://localhost:8080` and using a tool such as the [Advanced Rest Client](https://chrome.google.com/webstore/detail/advanced-rest-client/hgmloofddffdnphfgcellkdfbfbjeloo/) Chrome Plugin do a post to your file using this as an example:

```json
{"iv":"YVQXiZn5I+Nb4Oscqn8fWw==","notification":"I/6cBGpvkpQ+bgj42sP+YcQcTgquYuyUhtjxgGFPbS7JPpPyZwlifpSt2Zy1D1GF8Rq0EeZDrCIWB6onrcs+NJK6WhT3GryMuj77WvRDZi5HaEiZrbt9vp0u4klJwXTQMOrs+71dU+u8ecJdEtB4e5RkMf3JXxqEeA2Z3WdzHszIwz0P9o4D2LoGpmI5Vn2zUzfhRgil76YSwbTBWNbBVp/4ho5JSgcfS8wjFNTChOlcnN35MdhWZrJtqmCy9WnpQOCHpcU4UGedmh1nySIuF10c6E2cxpeI5ZbFX6ZsfwFRultTMMHX+uShkepgTUxjUjE2iujzvWPCEtQNWkj7SkAd4vErL3oBxRzg6oWaywZMtOSWexcRlYqY+Dlpd5y5Is6n6QpDfKE+7BnbQwsQNpwSH3YIvPiGBcrWX76QCUaLUWNpcH81WT1UAGry5R0LRjOZVHGz0cMBkQ9UwNXplem5Il78p3l2VP/iw7/i0RyAq7f7YxJAU4d9VtPD7tU3g5SP0K6KwkXJk2AGbLuI0La7SJGCsbLeuoKgSHKfecNcPK9oRjJTKgVeUJGSEf7J6SNj2b1StZfX6ySoc5n7OmZy/YJWszEeouf2d8JC7jpIiNhrQOhtlM4c8UPtit6QIpo8r2iBwr9Jz4Dggvgsgp2WIhjUrqL07m+5/o9v1duN9++liO7ndmTvjZEqZoi71X62QrOHk7z6iWDnHqPenrprC9iqWoZvL2aMPlt5lXfaNVh15maGuxAD80mYfbYKWOhtDjlj4oIcReE5S8sU9WVzUTALqxICOUmoukKAW9cqWVo8mWZL/cW+fBdrPJlGM8CaYJ3E8ftIbNp+WhVFRiNz05Bd/4D2tDgSrYzi6vXmajjkLZPuSNiiqFnygVO81bp8vC8SawEJzU6oc0Wqsg=="}
```

Now click send and you should have a win!

## Licence
This project is Copyright 2015 Centric Web Estate. All rights reserved. Licenced out under MIT.
