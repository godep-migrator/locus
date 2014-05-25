# locus

This is a tiny little Go web app using Martini and gogeos that purely exists to scratch an itch I had. As such, a lot of the niceties aren't present. I might add them later.

![Locus](https://raw.github.com/anachronistic/locus/master/locus-home.png)

## Running

```
$ PORT=5000 go run *.go
```

## Accessing Programmatically

```
$ curl -sd "point-x=1.0&point-y=1.0&polygon-wkt=MULTIPOLYGON(((0.0 0.0, 2.0 0.0, 2.0 2.0,0.0 2.0, 0.0 0.0)))" -X POST http://localhost:5000/ | jq '.'
{
  "timestamp": 1401046824,
  "contained": true
}
```
