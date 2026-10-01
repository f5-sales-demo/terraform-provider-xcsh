---
page_title: "xcsh_public_ip landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_public_ip landing."
---

# xcsh_public_ip landing

<a id="canonical-1031112030301010-2311310020031031-2211013213012022-1133000123132320-1321011231003203-2132100011003132-1013002332020322-0201311232021222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123232321013230-3100311223322100-2331021210020303-1301310201102233-1110113211321333-1121013320121323-1231330203313112-2030033303012001"></a>

## xcsh_public_ip — xcsh_public_ip / 023112110012 / 2

Breadcrumbs:

- xcsh_public_ip

Manages a Public IP resource in F5 Distributed Cloud for get public\_ip will get the object from the
storage backend for namespace metadata.namespace. configuration. (read-only data source)

<a id="canonical-1322001102010210-1303300310131001-3303013330200312-2330030122310320-2200113002213011-1101030310313322-3313221101331332-0312302320231231"></a>

## Prerequisites — xcsh_public_ip / 023112110012 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2222133120133033-0213332302130113-3102123022322212-2131102213331301-1333132320030301-1101003112021300-3210313123312211-0012131312103003"></a>

## Minimal configuration — xcsh_public_ip / 023112110012 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-1312301133233303-3133210230112323-2023113011100323-0203021002232331-3321302120320033-3030020113211002-1013320323230201-0330223002220013"></a>

## Root configuration — xcsh_public_ip / 023112110012 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1310032010303013-0103313222332030-0022312002213002-1301122202110301-1323310211232202-3332102332012211-3130111100331011-1220022121120233"></a>

## Next pages — xcsh_public_ip / 023112110012 / 6

- [Property reference](../guides/data-sources--public_ip--reference--group-001.md#canonical-1320331200230230-1110321111312310-1320320001202310-3322211120133022-3333120310113022-0300120302121212-2210030111311112-1221332130113230)
- [Examples](../guides/data-sources--public_ip--examples--group-001.md#canonical-1212333320001102-0203132121132100-2013320123101220-2103330233200323-3331232102031102-2300100221033032-3000232033221101-3010310211300333)
