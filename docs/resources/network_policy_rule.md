---
page_title: "xcsh_network_policy_rule"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_rule."
---

# xcsh_network_policy_rule

<a id="canonical-2112020201001103-2201031331322001-0223011022132003-2100033230200220-3103222330021332-3321002202302211-2131122200310202-1220303211023221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_network_policy_rule

Manages network policy rule with configured parameters in specified namespace in F5 Distributed
Cloud.

<a id="canonical-3022102103333102-2033320230331030-2131203121121101-3123332233022033-2003222302331311-1021103301331321-1121031133032213-1000210001030021"></a>

### Prerequisites for `xcsh_network_policy_rule`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0022011003102021-3000103230132231-2211012023302100-0310110222113121-1312130000212103-0032120223330101-2332310200001022-1312333211003202"></a>

### Minimal configuration for `xcsh_network_policy_rule`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkPolicyRule Resource Example
# Manages network policy rule with configured parameters in specified namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkPolicyRule configuration
resource "xcsh_network_policy_rule" "example" {
  name      = "example-network-policy-rule"
  namespace = "staging"
}
```

<a id="canonical-1332032211303113-0300100131132303-2310030330110211-3023230033113300-2102121211112213-3212320310223113-3130100221002032-3313021210303101"></a>

### Root configuration for `xcsh_network_policy_rule`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1031103223330320-0103311100232212-0230212213021330-3011110110312323-1130213222221112-3300031133013230-0332032301323110-0322002213120013"></a>

### Explore this collection for `xcsh_network_policy_rule`

- [Property reference](../guides/resources--network_policy_rule--reference--group-001.md#canonical-3012303312123012-2212103332102002-1022222311101132-2300213221103332-2210103322300222-0100331323202201-2332222132333001-2303211101203300)
- [Examples](../guides/resources--network_policy_rule--examples--group-001.md#canonical-0332032102031030-3122310133212130-3301012133100130-0213200313233111-3110022130033331-2331310200120230-1230030121103030-3020102112122231)
- [Import](../guides/resources--network_policy_rule--lifecycle--group-001.md#canonical-0001132230022202-3122001332130110-1010122332130112-2231032100322333-1031010012332133-1231322322323201-0110103113230313-3120322203033011)
- [Timeouts](../guides/resources--network_policy_rule--lifecycle--group-001.md#canonical-0111030200102323-0210123122230303-2331001311100323-1330300322122201-2220122033232023-2033323223023202-1323301211222332-2320030303030022)
