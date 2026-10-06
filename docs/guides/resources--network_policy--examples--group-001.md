---
page_title: "xcsh_network_policy examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_policy examples."
---

# xcsh_network_policy examples

<a id="canonical-3231201002231112-2221112210320001-1011322112323203-3100031231131033-3123320131213212-3333312212303233-3020021213001320-2300231102202221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- Examples

<a id="canonical-1231202001021020-3312131233322000-0010111213022011-0113212200010011-1323333031310323-2101021130333032-2312133300030011-2323122030322212"></a>

### Complete configurations for `xcsh_network_policy`

- [Resource](resources--network_policy--examples--group-001.md#canonical-0213011323220312-3331100322223123-1210233332003233-0302322131301111-1210212201300100-3023300103002122-2311220212102332-2313013332003121): valid configuration.

<a id="canonical-0213011323220312-3331100322223123-1210233332003233-0302322131301111-1210212201300100-3023300103002122-2311220212102332-2313013332003121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md#canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112)
- [Examples](resources--network_policy--examples--group-001.md#canonical-3231201002231112-2221112210320001-1011322112323203-3100031231131033-3123320131213212-3333312212303233-3020021213001320-2300231102202221)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_network_policy/resource.tf`; digest `sha256:4d106a33f6bcd90b712f1a25c8242900f48156f66650342f586b89690227f887`.

```terraform
# NetworkPolicy Resource Example
# Manages new network policy with configured parameters in specified namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkPolicy configuration
resource "xcsh_network_policy" "example" {
  name      = "example-network-policy"
  namespace = "staging"
}
```
