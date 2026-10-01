---
page_title: "xcsh_network_bot_defense landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_bot_defense landing."
---

# xcsh_network_bot_defense landing

<a id="canonical-0f1230987fc786b4cb7c4b344761f3dee154f819970a4e6a082efc66aa2b0e4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-76ce487c04679e7babd86460145e2308d94e120b987bbeb05e54be3d0db1c45f"></a>

## xcsh_network_bot_defense — xcsh_network_bot_defense / 873f87da7267 / 2

Breadcrumbs:

- xcsh_network_bot_defense

Bot Defense domains for an FQDN-aware firewall or proxy. Values are bundled from the pinned OpenAPI
release; this data source performs no network request. Ports and traffic direction are not encoded
in the manifest.

<a id="canonical-e2a2037b53398c1281173294e0300c33bdee6f20ea9c91856143151bbda7479a"></a>

## Prerequisites — xcsh_network_bot_defense / 873f87da7267 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-8d0acfe9e53bd7c030bcc4fd5ab7ad9c2e93a5ab3a585b55bf4e743fc6f2b721"></a>

## Minimal configuration — xcsh_network_bot_defense / 873f87da7267 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_bot_defense" "proxy" {}

# Configure these exact domains in a suffix-aware proxy or FQDN firewall.
output "bot_defense_https_proxy_rule" {
  value = {
    direction = "egress"
    protocol  = "tcp"
    port      = 443
    domains   = data.xcsh_network_bot_defense.proxy.domains
  }
}
```

<a id="canonical-3dba987a0f2a3d7e316b402462c3f7bcba1d2011accb117a9d31a34fd543f2d2"></a>

## Root configuration — xcsh_network_bot_defense / 873f87da7267 / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-0d975b7d0de11bb65cf3cc989a79bc2108e7e99e2a089fec4bc8bebbc20480d1"></a>

## Next pages — xcsh_network_bot_defense / 873f87da7267 / 6

- [Property reference](../guides/data-sources--network_bot_defense--reference--group-001.md#canonical-e9b989ac97501f94c8334812a9f1281e2f3df1d62b05ae468e99c21685977e41)
- [Examples](../guides/data-sources--network_bot_defense--examples--group-001.md#canonical-02700dbe54bec0d993bff3ebb664b29b8813b46344c33b64034f20f39624ad6c)
