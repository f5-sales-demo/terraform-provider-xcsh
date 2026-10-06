---
page_title: "xcsh_certificate examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate examples."
---

# xcsh_certificate examples

<a id="canonical-3322313303332131-0131020110320022-1221212301120013-2231031201130012-3302302113112112-3032012230000112-3133333123033310-0120311001333332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)
- Examples

<a id="canonical-0133132232113222-3022302220003022-0003030333102121-3103011201003122-2201000213323233-0202312212033103-0111020000312113-1330322033102201"></a>

### Complete configurations for `xcsh_certificate`

- [Resource](resources--certificate--examples--group-001.md#canonical-1112221302211031-2031103122010300-1001020001123022-3203200030003230-3021230330303200-2232230301223022-2311312301202222-1003332233301021): valid configuration.

<a id="canonical-1112221302211031-2031103122010300-1001020001123022-3203200030003230-3021230330303200-2232230301223022-2311312301202222-1003332233301021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_certificate](../resources/certificate.md#canonical-2010033013323302-3111123220113301-2322122303101210-2003020210333032-0232021333220031-3000020103303100-3202112103030121-1013321102200130)
- [Examples](resources--certificate--examples--group-001.md#canonical-3322313303332131-0131020110320022-1221212301120013-2231031201130012-3302302113112112-3032012230000112-3133333123033310-0120311001333332)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_certificate/resource.tf`; digest `sha256:a89bc2bb9f8d9310b1230ada6b9c044b0eeaf77e938fa8139bb314ba9412a092`.

```terraform
# Certificate Resource Example
# Manages a Certificate resource in F5 Distributed Cloud for certificate.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Certificate configuration
resource "xcsh_certificate" "example" {
  name      = "example-certificate"
  namespace = "staging"

  certificate_url = "example-value"
}
```
