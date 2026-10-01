---
page_title: "xcsh_forward_proxy_policy landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_forward_proxy_policy landing."
---

# xcsh_forward_proxy_policy landing

<a id="canonical-1321032010100101-1332210300230231-2303222012301133-2310222233132131-3320301230022230-1212303322110131-3223203122203201-0023111200022203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303330033233230-1210213301202303-1112323000030121-3332301322123301-1110330211021003-2113230222002212-1232212113201211-0131302223001000"></a>

## xcsh_forward_proxy_policy — xcsh_forward_proxy_policy / 223010320201 / 2

Breadcrumbs:

- xcsh_forward_proxy_policy

Manages a Forward Proxy Policy resource in F5 Distributed Cloud for forward proxy policy
specification. configuration.

<a id="canonical-3020123223303230-1123100132320001-0230011223331213-0321333010022230-0231230303130022-1111321110003011-2231230021232020-0021310033132203"></a>

## Prerequisites — xcsh_forward_proxy_policy / 223010320201 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

<a id="canonical-1301103103321031-3032003032121131-2010310122303333-0212012121103032-2003131113232023-2131132100000313-2210030321203000-1132100001233033"></a>

## Minimal configuration — xcsh_forward_proxy_policy / 223010320201 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ForwardProxyPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ForwardProxyPolicy by name
data "xcsh_forward_proxy_policy" "example" {
  name      = "example-forward-proxy-policy"
  namespace = "staging"
}

output "forward_proxy_policy_id" {
  value = data.xcsh_forward_proxy_policy.example.id
}
```

<a id="canonical-2123000230110231-3002113113331233-1331221023323323-0013113132012112-0312021202123022-2023010122130220-0333010332001332-1032312000212220"></a>

## Root configuration — xcsh_forward_proxy_policy / 223010320201 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0013312021200233-2320131020110312-1312030220123221-2012200102211322-2112001020231113-2223002313031003-3221023001232120-2112202102310000"></a>

## Next pages — xcsh_forward_proxy_policy / 223010320201 / 6

- [Property reference](../guides/data-sources--forward_proxy_policy--reference--group-001.md#canonical-0201131330302000-1332220320301222-0231230133130020-1023003321203020-2130222120313300-0113211322120200-1311332023100113-2210312021113312)
- [Examples](../guides/data-sources--forward_proxy_policy--examples--group-001.md#canonical-3301002023232221-2123112022013300-2222032223232202-2230303310220002-1112212003003010-2213333013122132-2120321333320331-3201301212021132)
