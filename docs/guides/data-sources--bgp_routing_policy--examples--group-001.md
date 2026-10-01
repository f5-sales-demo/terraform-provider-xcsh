---
page_title: "xcsh_bgp_routing_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_routing_policy examples."
---

# xcsh_bgp_routing_policy examples

<a id="canonical-081e2a3754d1a66d63c5257d16e6901391a00cf262454a13ddc7c0ead7261e26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-27287554ad1abb4e3b5cc8b8e5ec8e13db52a41db1f701727385865a9df4e7c4"></a>

## Examples — Examples / b6f29e9b1dbe / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)
- Examples

<a id="canonical-59afd2d57345e3c6c5931ecd561d7e03e25373ab2ed274aaeee63fafd9aba6f3"></a>

## Complete configurations — Examples / b6f29e9b1dbe / 3

- [Data source](data-sources--bgp_routing_policy--examples--group-001.md#canonical-5e1dd66f63ac703c4c7e0b7ab1263e87011972ee7ec64ad2ed4530e3aa444996): valid configuration.

<a id="canonical-d86d51d6d15fa2840761df8212ddaf07732939c781cdbb7110df261a141e643f"></a>

## Next pages — Examples / b6f29e9b1dbe / 4

- [Data source](data-sources--bgp_routing_policy--examples--group-001.md#canonical-5e1dd66f63ac703c4c7e0b7ab1263e87011972ee7ec64ad2ed4530e3aa444996)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)

<a id="canonical-5e1dd66f63ac703c4c7e0b7ab1263e87011972ee7ec64ad2ed4530e3aa444996"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c103fe347961165c6c877a429c207b3727b199ce80e22041fe8be5b6cec55c47"></a>

## Data source — Data source / d918f2105b5f / 2

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)
- [Examples](data-sources--bgp_routing_policy--examples--group-001.md#canonical-081e2a3754d1a66d63c5257d16e6901391a00cf262454a13ddc7c0ead7261e26)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bgp_routing_policy/data-source.tf`; digest `sha256:b3d6f35be703edce6abe0703f58285587174e0deae9414b75725a102a2aa63cb`.

```terraform
# BGPRoutingPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BGPRoutingPolicy by name
data "xcsh_bgp_routing_policy" "example" {
  name      = "example-bgp-routing-policy"
  namespace = "staging"
}

output "bgp_routing_policy_id" {
  value = data.xcsh_bgp_routing_policy.example.id
}
```

<a id="canonical-75e634e6971235002f8e69c218dea1aaa5e3d16856c5149634aad0d3469bcacf"></a>

## Next pages — Data source / d918f2105b5f / 3

- [Examples](data-sources--bgp_routing_policy--examples--group-001.md#canonical-081e2a3754d1a66d63c5257d16e6901391a00cf262454a13ddc7c0ead7261e26)
- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-d283a4dc63d88327e17c750f2480ee46fe9394b55d272ace958f6f84d1ad4985)
