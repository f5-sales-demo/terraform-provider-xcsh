---
page_title: "xcsh_forwarding_class"
subcategory: ""
description: "Manages a Forwarding Class resource in F5 Distributed Cloud for forwarding class is created by users in system namespace. configuration."
xcsh_docs: {"aliases": ["forwarding class"], "body_bytes": 1423, "body_sha256": "sha256:40470a43698f6fb4687cf5cedac7135ddbf3d93c72a4926df58dfd440710e45e", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:forwarding_class:reference", "xcsh-docs:data-sources:forwarding_class:examples"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:forwarding_class:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:forwarding_class:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/forwarding_class/index.md", "product": "distributed-cloud", "provider_name": "forwarding_class", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0323010303031321-2222112303012320-2022001320033212-3211310231002211-0310131010121230-1201102203203131-0230233130001111-3201321012320203", "registry_path": "docs/data-sources/forwarding_class.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forwarding_class/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a Forwarding Class resource in F5 Distributed Cloud for forwarding class is created by users in system namespace. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["forwarding_classCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_forwarding_class

Breadcrumbs:

- xcsh_forwarding_class

Manages a Forwarding Class resource in F5 Distributed Cloud for forwarding class is created by users
in system namespace. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ForwardingClass Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ForwardingClass by name
data "xcsh_forwarding_class" "example" {
  name      = "example-forwarding-class"
  namespace = "staging"
}

output "forwarding_class_id" {
  value = data.xcsh_forwarding_class.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forwarding_class/examples/)
