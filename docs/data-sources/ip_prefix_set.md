---
page_title: "xcsh_ip_prefix_set landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ip_prefix_set landing."
---

# xcsh_ip_prefix_set landing

<a id="canonical-abae8fdf0da4e4dbdf40819b5ddfd7d2d0567431640b74f7240d07839e34bedc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae0b05a00e24896a81f1b4214410225971d21eb7641709d6bfbc8890516c8e33"></a>

## xcsh_ip_prefix_set — xcsh_ip_prefix_set / 1954bccaecde / 2

Breadcrumbs:

- xcsh_ip_prefix_set

Manages ip\_prefix\_set creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

<a id="canonical-11920be82115a2f510812abc07e366b24d07ef432377e5b50b9fbb6488112af3"></a>

## Prerequisites — xcsh_ip_prefix_set / 1954bccaecde / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-a3d80f78eb26a514ba9aa287d0701db0b2629f24f44248ef42ec5538a927989d"></a>

## Minimal configuration — xcsh_ip_prefix_set / 1954bccaecde / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# IPPrefixSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing IPPrefixSet by name
data "xcsh_ip_prefix_set" "example" {
  name      = "example-ip-prefix-set"
  namespace = "staging"
}

output "ip_prefix_set_id" {
  value = data.xcsh_ip_prefix_set.example.id
}
```

<a id="canonical-0fea1510f367646b24f3b1059092ad92eeca4fc714386c9f70ae90793ea2df6b"></a>

## Root configuration — xcsh_ip_prefix_set / 1954bccaecde / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-524f4f219521e5f4c8622a113293b2fd0d3fc1331eb36f6015d83cf98dc85276"></a>

## Next pages — xcsh_ip_prefix_set / 1954bccaecde / 6

- [Property reference](../guides/data-sources--ip_prefix_set--reference--group-001.md#canonical-24ad13b3f3a6891fd4f419a489a2eede859b177520a1a420539bf19661644552)
- [Examples](../guides/data-sources--ip_prefix_set--examples--group-001.md#canonical-5100a4a30085b44b789ae0073db4bba5a4ec81d101479bf4216798464fee5841)
