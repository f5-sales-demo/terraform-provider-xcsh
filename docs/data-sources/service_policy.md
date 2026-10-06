---
page_title: "xcsh_service_policy"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_service_policy."
---

# xcsh_service_policy

<a id="canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_service_policy

Reads Service Policy information from F5 Distributed Cloud.

<a id="canonical-3300211201003011-1330120233232002-3203032332111311-0210112313303312-1112313003010123-2321220230031331-1333011111021032-2310313131132232"></a>

### Prerequisites for `xcsh_service_policy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

<a id="canonical-0213313033031100-3220333301202313-2100223122301022-2102121322301221-2213003022302221-3031001111202231-0033312233230211-0332302013100312"></a>

### Minimal configuration for `xcsh_service_policy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ServicePolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ServicePolicy by name
data "xcsh_service_policy" "example" {
  name      = "example-service-policy"
  namespace = "staging"
}

output "service_policy_id" {
  value = data.xcsh_service_policy.example.id
}
```

<a id="canonical-1013210103100321-2230330333022332-2210200211113023-0122211002013212-0212001202213113-1211121020303221-2003220033002300-2203130302310202"></a>

### Root configuration for `xcsh_service_policy`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2302112330132031-0132231020011232-2003310100330103-3000110001231321-3030230112130100-0220002222103321-1120321323233300-0321123323232211"></a>

### Explore this collection for `xcsh_service_policy`

- [Property reference](../guides/data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [Examples](../guides/data-sources--service_policy--examples--group-001.md#canonical-2310032000213002-0013023100320201-1122323331333212-3213302213103111-0003331221123320-2033213121331022-2322303123023213-0001233132023000)
