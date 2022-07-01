module github.com/chanwit/spike-bucket

go 1.16

require (
	github.com/johannesboyne/gofakes3 v0.0.0-20220517215058-83a58ec253b6
	github.com/spf13/afero v1.2.1
)

replace github.com/johannesboyne/gofakes3 => ./gofakes3
