---
page_title: "xcsh_bot_network_policy"
subcategory: ""
description: "Manages a Bot Network Policy resource in F5 Distributed Cloud for get bot network policy. configuration. (read-only data source)"
xcsh_docs: {"aliases": ["bot network policy"], "body_bytes": 1433, "body_sha256": "sha256:9125d14d799f045440b59255ce067b935bcad243ae4e0702b984f9fda4a12e55", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_network_policy:reference", "xcsh-docs:data-sources:bot_network_policy:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_network_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_network_policy:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/bot_network_policy/index.md", "product": "distributed-cloud", "provider_name": "bot_network_policy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0131231022333300-2321113231232321-2013223130013312-1311321020233021-2120000102000123-3300101133122113-2333131010121013-3020322100131202", "registry_path": "docs/data-sources/bot_network_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_network_policy/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Manages a Bot Network Policy resource in F5 Distributed Cloud for get bot network policy. configuration. (read-only data source)", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_bot_network_policy

Breadcrumbs:

- xcsh_bot_network_policy

Manages a Bot Network Policy resource in F5 Distributed Cloud for get bot network policy.
configuration. (read-only data source)

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotNetworkPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotNetworkPolicy by name
data "xcsh_bot_network_policy" "example" {
  name      = "example-bot-network-policy"
  namespace = "staging"
}

output "bot_network_policy_id" {
  value = data.xcsh_bot_network_policy.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_network_policy/examples/)
