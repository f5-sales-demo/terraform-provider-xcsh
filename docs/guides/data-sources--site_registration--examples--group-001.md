---
page_title: "xcsh_site_registration examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_registration examples."
---

# xcsh_site_registration examples

<a id="canonical-2ffc7a37abcd0056a9f2d90bd2820c24ce6c8a60078d36dfa1a9cc17da2afecf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34306646f9c8c799172a8e7e09c1620082bafdecf61a7d0f952c64997a339193"></a>

## Examples — Examples / 21f367980fc2 / 2

Breadcrumbs:

- [xcsh_site_registration](../data-sources/site_registration.md#canonical-29e210d357a5bdcac6bb8895520d01f565fe4c298fee3fb79b2fce630978c912)
- Examples

<a id="canonical-bfc28d7a8e10293ed984a10229c3e6934beb5cd6ff59143e59d6ab2d0c8e12d7"></a>

## Complete configurations — Examples / 21f367980fc2 / 3

- [Data source](data-sources--site_registration--examples--group-001.md#canonical-e56616d0b49148a7c04af59867369b0b7a8b797f03311fd2692eb9d3e8d342b4): valid configuration.

<a id="canonical-2bae1fe5d96510fbd6ce5812fc692e590c26e1c1b561e2822c31cf52b13858c8"></a>

## Next pages — Examples / 21f367980fc2 / 4

- [Data source](data-sources--site_registration--examples--group-001.md#canonical-e56616d0b49148a7c04af59867369b0b7a8b797f03311fd2692eb9d3e8d342b4)
- [xcsh_site_registration](../data-sources/site_registration.md#canonical-29e210d357a5bdcac6bb8895520d01f565fe4c298fee3fb79b2fce630978c912)

<a id="canonical-e56616d0b49148a7c04af59867369b0b7a8b797f03311fd2692eb9d3e8d342b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a79d62ae3e03b189867e430d7d49adcd49068aafb12e89aac18f8dd3778dbe33"></a>

## Data source — Data source / 37e96d568835 / 2

Breadcrumbs:

- [xcsh_site_registration](../data-sources/site_registration.md#canonical-29e210d357a5bdcac6bb8895520d01f565fe4c298fee3fb79b2fce630978c912)
- [Examples](data-sources--site_registration--examples--group-001.md#canonical-2ffc7a37abcd0056a9f2d90bd2820c24ce6c8a60078d36dfa1a9cc17da2afecf)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_registration/data-source.tf`; digest `sha256:61863eead29a93d9046672296900d790c36ec0a2b1c8200672a0a876134858bb`.

```terraform
# Example: Resolve a Customer Edge registration so it can be approved
#
# A registration is named "r-<uuid>", NOT after the site, so it cannot be read
# by site name. This data source finds the registration that belongs to a site.
#
# It returns found = false — with no error — until the CE has booted and
# registered, so an approval can safely be gated on it.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_registration" "ce" {
  site_name = "my-ce-site"
  namespace = "system"
}

# A multi-node site returns one registration per node; pick one by hostname.
data "xcsh_site_registration" "ha_ce_node_0" {
  site_name = "my-ha-ce-site"
  hostname  = "master-0"
}

output "registration_name" {
  description = "Registration name (r-<uuid>) to approve, null until the CE registers"
  value       = data.xcsh_site_registration.ce.name
}

output "registration_state" {
  description = "Current registration state (PENDING, ONLINE, ...)"
  value       = data.xcsh_site_registration.ce.state
}

output "ha_node_0_registration_name" {
  description = "Registration name of the master-0 node of the three-node site"
  value       = data.xcsh_site_registration.ha_ce_node_0.name
}

# Approve the registration only once it exists — on the first apply the CE has
# not registered yet, so nothing is planned; re-apply after the node boots.
resource "xcsh_registration_approval" "ce" {
  count = data.xcsh_site_registration.ce.found ? 1 : 0

  name         = data.xcsh_site_registration.ce.name
  namespace    = data.xcsh_site_registration.ce.namespace
  cluster_size = 1
}
```

<a id="canonical-4818d6c771d6a1a1a9fccf2e19081c3414dfa4a75e089185848609d5c6c34b2e"></a>

## Next pages — Data source / 37e96d568835 / 3

- [Examples](data-sources--site_registration--examples--group-001.md#canonical-2ffc7a37abcd0056a9f2d90bd2820c24ce6c8a60078d36dfa1a9cc17da2afecf)
- [xcsh_site_registration](../data-sources/site_registration.md#canonical-29e210d357a5bdcac6bb8895520d01f565fe4c298fee3fb79b2fce630978c912)
