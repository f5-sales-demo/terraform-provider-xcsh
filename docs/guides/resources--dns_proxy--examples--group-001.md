---
page_title: "xcsh_dns_proxy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_proxy examples."
---

# xcsh_dns_proxy examples

<a id="canonical-9fa90dae695ab8998220060b8a3ee60cd4dfcfd96133704fb891d295262d40e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1877cd4b243fa62e392d4281ab669bff67ecb547f215bae765a846226a3c1409"></a>

## Examples — Examples / 6f63b30ecca9 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- Examples

<a id="canonical-8cdc53b19e21fa6845d2991fe8322dfe6fe3131818a59c864ce57bda0d956a71"></a>

## Complete configurations — Examples / 6f63b30ecca9 / 3

- [Resource](resources--dns_proxy--examples--group-001.md#canonical-4b2e7cb147f935486b2c8f4f6a5a353e59e3a7106e046e0182fa111b7be50b3c): valid configuration.

<a id="canonical-652a755fd7e5e13e7adc347e3daf211653f5331a9372ced300ca7abaf8520ed7"></a>

## Next pages — Examples / 6f63b30ecca9 / 4

- [Resource](resources--dns_proxy--examples--group-001.md#canonical-4b2e7cb147f935486b2c8f4f6a5a353e59e3a7106e046e0182fa111b7be50b3c)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)

<a id="canonical-4b2e7cb147f935486b2c8f4f6a5a353e59e3a7106e046e0182fa111b7be50b3c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4ae78d693e007a88dbfb60b0bed970a860194da572ab2500c388155ecd74d59"></a>

## Resource — Resource / f175d213f562 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
- [Examples](resources--dns_proxy--examples--group-001.md#canonical-9fa90dae695ab8998220060b8a3ee60cd4dfcfd96133704fb891d295262d40e5)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dns_proxy/resource.tf`; digest `sha256:943728fe8d11ffeb7bc560a6a0e3b2ff777ab7fb03003f4fd0ef80073c303500`.

```terraform
# DNSProxy Resource Example
# Manages DNS Proxy in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSProxy configuration
resource "xcsh_dns_proxy" "example" {
  name      = "example-dns-proxy"
  namespace = "system"
}
```

<a id="canonical-33a5f829df180a3456bf92e0b8d952fd75ab33e2cff2561b29ccf8aa52d98d1c"></a>

## Next pages — Resource / f175d213f562 / 3

- [Examples](resources--dns_proxy--examples--group-001.md#canonical-9fa90dae695ab8998220060b8a3ee60cd4dfcfd96133704fb891d295262d40e5)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-6f40f0710353e90bdbf38f43f5d0424c972f0377902408914de7f101e829d41a)
