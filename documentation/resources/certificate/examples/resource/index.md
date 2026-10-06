---
page_title: "Resource"
subcategory: "Security"
description: "Resource for xcsh_certificate."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1057, "body_sha256": "sha256:78caad79c98f3f39b1c5abcd03453c7bc61a2e96997fd7dc47ad7d52cdc62f77", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:certificate:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a89bc2bb9f8d9310b1230ada6b9c044b0eeaf77e938fa8139bb314ba9412a092", "source_path": "examples/resources/xcsh_certificate/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:certificate:example:resource", "parent_id": "xcsh-docs:resources:certificate:examples", "path": "documentation/resources/certificate/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "certificate", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1112221302211031-2031103122010300-1001020001123022-3203200030003230-3021230330303200-2232230301223022-2311312301202222-1003332233301021", "registry_path": "docs/guides/resources--certificate--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/certificate/examples/resource/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Resource for xcsh_certificate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["certificateCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_certificate/resource.tf`; digest `sha256:a89bc2bb9f8d9310b1230ada6b9c044b0eeaf77e938fa8139bb314ba9412a092`.

```terraform
# Certificate Resource Example
# Manages a Certificate resource in F5 Distributed Cloud for certificate.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Certificate configuration
resource "xcsh_certificate" "example" {
  name      = "example-certificate"
  namespace = "staging"

  certificate_url = "example-value"
}
```
