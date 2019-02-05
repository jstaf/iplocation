# iplocation

This is a small script designed to make looking up the physical location of an 
IP address or hostname just a little bit easier. This is a very thin wrapper 
around the free ip-api.com JSON endpoint.

To build this tool, run `go build` (you'll need to have Go installed).

To use, invoke `iplocation` with either an IP address or DNS address:

```bash
./iplocation 1.2.3.4
./iplocation google.ca
```

The output will be human-readable JSON describing the location and organization
behind a given address. The only caveat is that if multiple addresses are found 
for a hostname, the result may vary from query to query (this tool can only 
retrieve 1 address at a time).

```
$ ./iplocation 1.2.3.4
{
    "as": "",
    "city": "South Brisbane",
    "country": "Australia",
    "countryCode": "AU",
    "isp": "Debogon",
    "lat": -27.4748,
    "lon": 153.017,
    "org": "",
    "query": "1.2.3.4",
    "region": "QLD",
    "regionName": "Queensland",
    "status": "success",
    "timezone": "Australia/Brisbane",
    "zip": "4101"
}
```
