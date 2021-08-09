package transaction

import "embed"

// Resources has the entire ./resources directory embed in read-only mode
//go:embed resources/application.yml
var Resources embed.FS
