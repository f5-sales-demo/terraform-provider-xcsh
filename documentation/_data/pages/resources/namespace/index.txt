---
page_title: "xcsh_namespace"
subcategory: ""
description: "Manages new namespace. Name of the object is name of the name space in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["namespace"], "body_bytes": 1485, "body_sha256": "sha256:099de36bbfe09bd6b9f96ad5d2ef829a169c968aeeadad9573458c8f9e66b877", "capabilities": ["administration"], "category": "administration", "child_ids": ["xcsh-docs:resources:namespace:reference", "xcsh-docs:resources:namespace:examples", "xcsh-docs:resources:namespace:import", "xcsh-docs:resources:namespace:timeouts"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:namespace:collection", "completeness": "complete", "id": "xcsh-docs:resources:namespace:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/namespace/index.md", "product": "distributed-cloud", "provider_name": "namespace", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1131103113212120-1020211231300120-1030023112132023-3013023011120122-0313023322131212-0001202033301111-3031320130100032-1113113021202310", "registry_path": "docs/resources/namespace.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/namespace/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages new namespace. Name of the object is name of the name space in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["namespaceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_namespace

Breadcrumbs:

- xcsh_namespace

Manages new namespace. Name of the object is name of the name space in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Namespace Resource Example
# Manages new namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Credentials are supplied externally.
provider "xcsh" {}

# Basic Namespace configuration
resource "xcsh_namespace" "this" {
  name = "example-namespace"
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/namespace/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/namespace/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/namespace/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/namespace/lifecycle/timeouts/)
