# envconfig
Package for defining configuration based on environment variables. Sort of a cross between `kelseyhightower/envconfig` and the `flag` package.

Unlike `flag` envconfig does not need to parse os.Args. This means that envconfig reads the environment and sets values immediately.

Note: All variable names are normalized to uppercase when reading the environment.

## Example Usage

```
import "github.com/mattaitchison/envconfig"

// Basic types
var hostIP = envconfig.String("registrator_ip", "", "IP for ports mapped to the host")
var internal = envconfig.Bool("registrator_internal", false, "Use internal ports instead of published ones")
var ttlRefresh = envconfig.Int("registrator_ttl_refresh", 0,
  "Frequency with which service TTLs are refreshed")
var velocity = envconfig.Float64("velocity", 1.23,
  "Floating point value with no prefix")

// Convenience + validation
var deregister = envconfig.StringOption("registrator_deregister", "always",
  []string{"always", "never", "on-success"},
  "Deregister mode")
```
