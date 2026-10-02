---
page_title: "xcsh_shape_bot_defense_instance"
subcategory: ""
description: "Manages a Shape Bot Defense Instance resource in F5 Distributed Cloud for get virtual host from a given namespace. configuration. (read-only data source)"
xcsh_docs: {"aliases": ["shape bot defense instance"], "body_bytes": 1536, "body_sha256": "sha256:b9a431e6ad220c49f1b29ebba7a651de45cb111b130298ac3cf33923bb9a566c", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:shape_bot_defense_instance:reference", "xcsh-docs:data-sources:shape_bot_defense_instance:examples"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:shape_bot_defense_instance:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:shape_bot_defense_instance:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/shape_bot_defense_instance/index.md", "product": "distributed-cloud", "provider_name": "shape_bot_defense_instance", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3000123112021030-0302103133011033-2210121132011122-2132220211231033-3233312310101331-1222032010313211-1333211301300023-0231320200212323", "registry_path": "docs/data-sources/shape_bot_defense_instance.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/shape_bot_defense_instance/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a Shape Bot Defense Instance resource in F5 Distributed Cloud for get virtual host from a given namespace. configuration. (read-only data source)", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
