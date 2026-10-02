---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_namespace."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1195, "body_sha256": "sha256:556c0cc1aebac49e96abfb00d739864b7f0df6dfa6f362f9031f0624905099c6", "capabilities": ["administration"], "category": "administration", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:namespace:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:95871f5bf448147bdd129adaa3690bae9af907a5a988941728cc5c87a1b9e514", "source_path": "examples/resources/xcsh_namespace/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:namespace:example:resource", "parent_id": "xcsh-docs:resources:namespace:examples", "path": "documentation/resources/namespace/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "namespace", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1000320232300203-3003333110003100-1310301310102123-3031221121110122-1122112210013322-2320202220323123-0003330023101033-2000233332213300", "registry_path": "docs/guides/resources--namespace--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/namespace/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_namespace.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["namespaceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/namespace/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/namespace/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_namespace/resource.tf`; digest `sha256:95871f5bf448147bdd129adaa3690bae9af907a5a988941728cc5c87a1b9e514`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/namespace/examples/)
- [xcsh_namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/namespace/)
