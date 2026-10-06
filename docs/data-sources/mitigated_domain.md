---
page_title: "xcsh_mitigated_domain"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_mitigated_domain."
---

# xcsh_mitigated_domain

<a id="canonical-3233303323320020-3201003222322030-3103103212101100-1321230323023131-3310220303032323-0211212333301333-1330120302113020-2012202211332310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_mitigated_domain

Reads Mitigated Domain information from F5 Distributed Cloud.

<a id="canonical-0110211322303221-0232020022022030-2122301020033230-2023220210133312-3230003302022111-1131103113212111-1120113023312002-1331101133200331"></a>

### Prerequisites for `xcsh_mitigated_domain`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2122111212012210-2211101130301131-1032121001113013-2033023130313331-2313013012103131-2311012211021321-3101102313130312-2130103222011100"></a>

### Minimal configuration for `xcsh_mitigated_domain`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# MitigatedDomain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing MitigatedDomain by name
data "xcsh_mitigated_domain" "example" {
  name      = "example-mitigated-domain"
  namespace = "staging"
}

output "mitigated_domain_id" {
  value = data.xcsh_mitigated_domain.example.id
}
```

<a id="canonical-0122203213203320-0120001200200121-2002000001302202-1023000213300211-1020203030131202-1120221310131220-2120103020002312-0030121032003011"></a>

### Root configuration for `xcsh_mitigated_domain`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2133132031200332-2232033122212231-2230033021230323-1332113233231130-0031033320031333-2222130012223100-3233313031302103-2232331311231332"></a>

### Explore this collection for `xcsh_mitigated_domain`

- [Property reference](../guides/data-sources--mitigated_domain--reference--group-001.md#canonical-0201213223313332-2132200331130222-1202220001302010-1022310120223013-3020032223131202-2123112211203031-2213120030001221-2330123113023202)
- [Examples](../guides/data-sources--mitigated_domain--examples--group-001.md#canonical-1230120121021001-2100330101330313-0000131010233210-3233021320010013-2322121011013130-2330100102101222-0132023121130001-2233232331220210)
