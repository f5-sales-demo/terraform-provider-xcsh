---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_crl."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1263, "body_sha256": "sha256:de6778cbadf10257b5e4306988e3b88113686a922568d87a0d078ac8e954d43c", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:crl:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:787cb67d26f54e510168be44dfb1793c35d97676385e49506a4c31f4d8662bf3", "source_path": "examples/resources/xcsh_crl/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:crl:example:resource", "parent_id": "xcsh-docs:resources:crl:examples", "path": "documentation/resources/crl/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "crl", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1222233233230320-2223120312320112-2110210320101320-1002302222132120-0230233333021011-3312133333020310-0132322131300302-2012300022303313", "registry_path": "docs/guides/resources--crl--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/crl/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_crl.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["crlCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/crl/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/crl/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_crl/resource.tf`; digest `sha256:787cb67d26f54e510168be44dfb1793c35d97676385e49506a4c31f4d8662bf3`.

```terraform
# CRL Resource Example
# Manages a CRL resource in F5 Distributed Cloud for api to create crl object.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CRL configuration
resource "xcsh_crl" "example" {
  name      = "example-crl"
  namespace = "staging"

  refresh_interval = 6
  server_address   = "example-value"
  server_port      = 1
  timeout          = 1
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/crl/examples/)
- [xcsh_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/crl/)
