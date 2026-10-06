---
page_title: "xcsh_public_ip examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_public_ip examples."
---

# xcsh_public_ip examples

<a id="canonical-1212333320001102-0203132121132100-2013320123101220-2103330233200323-3331232102031102-2300100221033032-3000232033221101-3010310211300333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_public_ip](../data-sources/public_ip.md#canonical-1031112030301010-2311310020031031-2211013213012022-1133000123132320-1321011231003203-2132100011003132-1013002332020322-0201311232021222)
- Examples

<a id="canonical-0030333323011303-0211320132222012-1311231222321322-2031122003230112-3112221120103203-3113020031212133-3213312102200210-2130000133321013"></a>

### Complete configurations for `xcsh_public_ip`

- [Data source](data-sources--public_ip--examples--group-001.md#canonical-2222001222003213-2122001223202112-0013211320323330-2122030312130011-1020202330211320-2122303320102033-2022122111200311-0013303122122100): valid configuration.

<a id="canonical-2222001222003213-2122001223202112-0013211320323330-2122030312130011-1020202330211320-2122303320102033-2022122111200311-0013303122122100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_public_ip](../data-sources/public_ip.md#canonical-1031112030301010-2311310020031031-2211013213012022-1133000123132320-1321011231003203-2132100011003132-1013002332020322-0201311232021222)
- [Examples](data-sources--public_ip--examples--group-001.md#canonical-1212333320001102-0203132121132100-2013320123101220-2103330233200323-3331232102031102-2300100221033032-3000232033221101-3010310211300333)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_public_ip/data-source.tf`; digest `sha256:3249669279b5070f3336fdc2ea70d7926ac4af4c76c377a1b4478e268c3c5d4d`.

```terraform
# PublicIP Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing PublicIP by name
data "xcsh_public_ip" "example" {
  name      = "example-public-ip"
  namespace = "staging"
}

output "public_ip_id" {
  value = data.xcsh_public_ip.example.id
}
```
