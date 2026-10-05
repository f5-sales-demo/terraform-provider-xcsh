---
page_title: "xcsh_shape_bot_defense_instance"
subcategory: ""
description: "Manages a Shape Bot Defense Instance resource in F5 Distributed Cloud for get virtual host from a given namespace. configuration. (read-only data source)"
xcsh_docs: {"aliases": ["shape bot defense instance"], "body_bytes": 1536, "body_sha256": "sha256:b9a431e6ad220c49f1b29ebba7a651de45cb111b130298ac3cf33923bb9a566c", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:shape_bot_defense_instance:reference", "xcsh-docs:data-sources:shape_bot_defense_instance:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:shape_bot_defense_instance:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:shape_bot_defense_instance:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/shape_bot_defense_instance/index.md", "product": "distributed-cloud", "provider_name": "shape_bot_defense_instance", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3000123112021030-0302103133011033-2210121132011122-2132220211231033-3233312310101331-1222032010313211-1333211301300023-0231320200212323", "registry_path": "docs/data-sources/shape_bot_defense_instance.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/shape_bot_defense_instance/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Manages a Shape Bot Defense Instance resource in F5 Distributed Cloud for get virtual host from a given namespace. configuration. (read-only data source)", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_shape_bot_defense_instance

Breadcrumbs:

- xcsh_shape_bot_defense_instance

Manages a Shape Bot Defense Instance resource in F5 Distributed Cloud for get virtual host from a
given namespace. configuration. (read-only data source)

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ShapeBotDefenseInstance Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ShapeBotDefenseInstance by name
data "xcsh_shape_bot_defense_instance" "example" {
  name      = "example-shape-bot-defense-instance"
  namespace = "staging"
}

output "shape_bot_defense_instance_id" {
  value = data.xcsh_shape_bot_defense_instance.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/shape_bot_defense_instance/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/shape_bot_defense_instance/examples/)
