---
page_title: "xcsh_service_policy_rule"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_service_policy_rule."
---

# xcsh_service_policy_rule

<a id="canonical-0131222010231231-3010331102212121-0000330200210011-0233010022022120-1322113333232331-1300100010221333-1012020132010301-1111121233113032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_service_policy_rule

Manages service\_policy\_rule creates a new object in the storage backend for metadata.namespace in
F5 Distributed Cloud.

<a id="canonical-0132013033221331-2230332102321221-0333032131233202-1112022332101002-0313203031333122-0000330110102122-0231101331223301-3112303221223123"></a>

### Prerequisites for `xcsh_service_policy_rule`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2011000002201100-0011321213212003-0210233300231221-3322223122111231-0132002123120122-0332311123013200-0033102223110123-1113110230232011"></a>

### Minimal configuration for `xcsh_service_policy_rule`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ServicePolicyRule Resource Example
# Manages service_policy_rule creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ServicePolicyRule configuration
resource "xcsh_service_policy_rule" "example" {
  name      = "example-service-policy-rule"
  namespace = "staging"

  action = "DENY"
}
```

<a id="canonical-0200001000231120-2101023011321102-0232121120031332-0132202230300110-1233300201121202-0332220101102110-0213023211301233-1221100201012313"></a>

### Root configuration for `xcsh_service_policy_rule`

Required root properties: `action`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2122233102113000-3013203120230203-2322322321031223-0132312010020111-1020321133202021-1333023020310213-3013032230231233-0313120222020111"></a>

### Explore this collection for `xcsh_service_policy_rule`

- [Property reference](../guides/resources--service_policy_rule--reference--group-001.md#canonical-3123121120223132-2100202032330223-1100112013210300-3233003003201021-3231131300131211-0202303330311130-2112220030322031-1130322203311231)
- [Examples](../guides/resources--service_policy_rule--examples--group-001.md#canonical-0023101333202233-2312130003023013-2031120302010112-1223202002333100-1103023303201230-3132030320300230-2102312130202330-0111300211010120)
- [Import](../guides/resources--service_policy_rule--lifecycle--group-001.md#canonical-3202023003231322-3231211231032003-2210332220021222-2031331210113132-1021002212333332-3100312112210200-2302121123012212-2010212022122122)
- [Timeouts](../guides/resources--service_policy_rule--lifecycle--group-001.md#canonical-0332313202002331-3330123122023333-2130323213003013-0002031221321303-1202003213001223-1201222102311111-1322132311221020-0031031110230120)
