---
page_title: "xcsh_shape_bot_defense_instance"
subcategory: ""
description: "Manages a Shape Bot Defense Instance resource in F5 Distributed Cloud for get virtual host from a given namespace. configuration. (read-only data source)"
xcsh_docs: {"aliases": ["shape bot defense instance"], "body_bytes": 1536, "body_sha256": "sha256:b9a431e6ad220c49f1b29ebba7a651de45cb111b130298ac3cf33923bb9a566c", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:shape_bot_defense_instance:reference", "xcsh-docs:data-sources:shape_bot_defense_instance:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:shape_bot_defense_instance:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:shape_bot_defense_instance:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/shape_bot_defense_instance/index.md", "product": "distributed-cloud", "provider_name": "shape_bot_defense_instance", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3000123112021030-0302103133011033-2210121132011122-2132220211231033-3233312310101331-1222032010313211-1333211301300023-0231320200212323", "registry_path": "docs/data-sources/shape_bot_defense_instance.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/shape_bot_defense_instance/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages a Shape Bot Defense Instance resource in F5 Distributed Cloud for get virtual host from a given namespace. configuration. (read-only data source)", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
