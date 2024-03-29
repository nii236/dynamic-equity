[![](https://godoc.org/github.com/nii236/dynamic-equity?status.svg)](https://godoc.org/github.com/nii236/dynamic-equity) [![Go Report Card](https://goreportcard.com/badge/github.com/nii236/dynamic-equity)](https://goreportcard.com/report/github.com/nii236/dynamic-equity)

# Slicing Pie

> Slicing Pie is a universal, one-size-fits all model that creates a perfectly fair equity split in an early-stage, bootstrapped start-up company.

More information [here](https://slicingpie.com)

# Database

```
docker run -d -p 5432:5432 \
--name pie-db \
-e POSTGRES_USER=pie \
-e POSTGRES_PASSWORD=dev \
-e POSTGRES_DB=pie \
postgres:11-alpine
```
