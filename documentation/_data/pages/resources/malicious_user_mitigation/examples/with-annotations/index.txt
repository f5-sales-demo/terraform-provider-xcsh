---
page_title: "With annotations"
subcategory: ""
description: "With annotations for xcsh_malicious_user_mitigation."
xcsh_docs: {"aliases": ["with-annotations"], "body_bytes": 1706, "body_sha256": "sha256:4a67c50c4d5c3c2b6ebce54db34c526148e2058a9a40d9801c08359ae8545982", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:malicious_user_mitigation:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:10af11e1fcedf0c1a38131c01416f578defc093012b6a6669770dd539168e15b", "source_path": "examples/resources/xcsh_malicious_user_mitigation/with-annotations.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:malicious_user_mitigation:example:with-annotations", "parent_id": "xcsh-docs:resources:malicious_user_mitigation:examples", "path": "documentation/resources/malicious_user_mitigation/examples/with-annotations/index.md", "product": "distributed-cloud", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0332011323203221-2212021231330001-3022103131310300-1030022012220312-0103300311201021-3223103001302132-0300002300321220-3012322030320300", "registry_path": "docs/guides/resources--malicious_user_mitigation--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["with-annotations"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/malicious_user_mitigation/examples/with-annotations/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "With annotations for xcsh_malicious_user_mitigation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# With annotations

Breadcrumbs:

- [xcsh_malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/examples/)
- With annotations

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_malicious_user_mitigation/with-annotations.tf`; digest `sha256:10af11e1fcedf0c1a38131c01416f578defc093012b6a6669770dd539168e15b`.

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
    time = {
      source  = "hashicorp/time"
      version = "= 0.13.1"
    }
  }
}

resource "xcsh_namespace" "test" {
  name = "example"
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [xcsh_namespace.test]
  create_duration = "5s"
}

resource "xcsh_malicious_user_mitigation" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = "example-value"
  namespace  = xcsh_namespace.test.name

  annotations = {
    example-key = "example-value"
  }
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/examples/)
- [xcsh_malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/)
