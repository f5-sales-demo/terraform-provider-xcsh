---
page_title: "All attributes"
subcategory: "Security"
description: "All attributes for xcsh_rate_limiter."
xcsh_docs: {"aliases": ["all-attributes"], "body_bytes": 1419, "body_sha256": "sha256:f6f750358adcfaf22acc67d7920a8cdd2933ce3d13c87919324536cbe4c58424", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:rate_limiter:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:0fa6c0ef5d6c3a83afac2353f4434ae5a9fdbf3ec95f2509bae99373fcdd2e92", "source_path": "examples/resources/xcsh_rate_limiter/all-attributes.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:rate_limiter:example:all-attributes", "parent_id": "xcsh-docs:resources:rate_limiter:examples", "path": "documentation/resources/rate_limiter/examples/all-attributes/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1000020310210121-2331323211133230-2333302022102102-3031311031001311-3021313023130333-1201112123300013-0300001323323023-2010102033120223", "registry_path": "docs/guides/resources--rate_limiter--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["all-attributes"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter/examples/all-attributes/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "All attributes for xcsh_rate_limiter.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# All attributes

Breadcrumbs:

- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/examples/)
- All attributes

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_rate_limiter/all-attributes.tf`; digest `sha256:0fa6c0ef5d6c3a83afac2353f4434ae5a9fdbf3ec95f2509bae99373fcdd2e92`.

```terraform
# AllAttributes — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_rate_limiter" "test" {
  name        = "example"
  namespace   = "system"
  description = "Test rate limiter with all attributes"
  disable     = false

  labels = {
    environment = "test"
    team        = "engineering"
  }

  annotations = {
    purpose = "testing"
  }
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/examples/)
- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/)
