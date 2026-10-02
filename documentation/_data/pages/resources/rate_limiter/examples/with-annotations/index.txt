---
page_title: "With annotations"
subcategory: "Security"
description: "With annotations for xcsh_rate_limiter."
xcsh_docs: {"aliases": ["with-annotations"], "body_bytes": 1280, "body_sha256": "sha256:699c4d9382abdb6c2fa581fb0d52f9e42ceebffe29219fd78781fa305882f0aa", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:rate_limiter:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:b278928b01f70454355f22fa0b9f5a1fb258167336f683c4e01b990a20527c97", "source_path": "examples/resources/xcsh_rate_limiter/with-annotations.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:rate_limiter:example:with-annotations", "parent_id": "xcsh-docs:resources:rate_limiter:examples", "path": "documentation/resources/rate_limiter/examples/with-annotations/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3021303320102011-3302301221013102-3031033332031213-2012311302022123-1211002100113233-3213212230301310-1322023210310203-1103303111322103", "registry_path": "docs/guides/resources--rate_limiter--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["with-annotations"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter/examples/with-annotations/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "With annotations for xcsh_rate_limiter.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# With annotations

Breadcrumbs:

- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/examples/)
- With annotations

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_rate_limiter/with-annotations.tf`; digest `sha256:b278928b01f70454355f22fa0b9f5a1fb258167336f683c4e01b990a20527c97`.

```terraform
# WithAnnotations — Acceptance-test-derived Configuration
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
  name      = "example"
  namespace = "system"

  annotations = {
    example-key = "example-value"
  }
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/examples/)
- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/)
