---
page_title: "xcsh_secret_management_access landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_secret_management_access landing."
---

# xcsh_secret_management_access landing

<a id="canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103013333023101-1001333333231231-3312023303122200-1031330001033233-0320121120332302-2333212033033112-0030212122321230-0103233031003031"></a>

## xcsh_secret_management_access — xcsh_secret_management_access / 122033110210 / 2

Breadcrumbs:

- xcsh_secret_management_access

Manages secret\_management\_access creates a new object in storage backend for metadata.namespace in
F5 Distributed Cloud.

<a id="canonical-3202302213302220-2132113220333101-1001031131001103-0022211223011110-2320211303202301-1331300131223031-0120010302131331-0223332201121122"></a>

## Prerequisites — xcsh_secret_management_access / 122033110210 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1321213121233221-1210012212023010-0212202120000011-3203121113013032-1031033312133232-3130023003323210-2210130103223021-0120313101302332"></a>

## Minimal configuration — xcsh_secret_management_access / 122033110210 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SecretManagementAccess Resource Example
# Manages secret_management_access creates a new object in storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SecretManagementAccess configuration
resource "xcsh_secret_management_access" "example" {
  name      = "example-secret-management-access"
  namespace = "staging"

  provider_name = "example-value"
}
```

<a id="canonical-1313023303123002-0231311201221210-2133101202103321-1123210103132202-0212213203033310-0330320221110031-3312321030222210-2202232030223322"></a>

## Root configuration — xcsh_secret_management_access / 122033110210 / 5

Required root properties: `name`, `namespace`, `provider_name`. Full root flags and choices appear in the property reference.

<a id="canonical-3022003300022321-2321022133301123-1230030232200231-1103112133010323-0331220300201202-0101020123013122-0100032210020132-0331330333202020"></a>

## Next pages — xcsh_secret_management_access / 122033110210 / 6

- [Property reference](../guides/resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [Examples](../guides/resources--secret_management_access--examples--group-001.md#canonical-2111202301110310-0313132001022331-3123020103123130-2301210312131331-1013011323103223-2100201310303012-1333233323020312-3022111021121303)
- [Import](../guides/resources--secret_management_access--lifecycle--group-001.md#canonical-1311023322202000-3031202311101120-0012032203312013-3302002203322322-0331232113331102-1310131321133112-2030200011222103-3101131100333230)
- [Timeouts](../guides/resources--secret_management_access--lifecycle--group-001.md#canonical-3020232303212322-0302201011002102-3303322122303013-3130001212102213-0013233023020113-3002213100321102-0320001201233103-3213233111103320)
