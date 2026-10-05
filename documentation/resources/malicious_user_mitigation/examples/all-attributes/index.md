---
page_title: "All attributes"
subcategory: ""
description: "All attributes for xcsh_malicious_user_mitigation."
xcsh_docs: {"aliases": ["all-attributes"], "body_bytes": 1854, "body_sha256": "sha256:805f59a4d828de05f1cca9b561a4541d7dc2873147d6d913a68c6d2a9dd26aab", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:malicious_user_mitigation:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:e1a870d0ac8f57cd049b19713024efa0b33bb0d6f7d3e2cc75464e43a5026739", "source_path": "examples/resources/xcsh_malicious_user_mitigation/all-attributes.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:malicious_user_mitigation:example:all-attributes", "parent_id": "xcsh-docs:resources:malicious_user_mitigation:examples", "path": "documentation/resources/malicious_user_mitigation/examples/all-attributes/index.md", "product": "distributed-cloud", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1213311311203230-3220000000201113-3120133333033102-3320211002020021-0013323130120321-1002021332120123-0223222331233022-3133103001133121", "registry_path": "docs/guides/resources--malicious_user_mitigation--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["all-attributes"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/malicious_user_mitigation/examples/all-attributes/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "All attributes for xcsh_malicious_user_mitigation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# All attributes

Breadcrumbs:

- [xcsh_malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/examples/)
- All attributes

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_malicious_user_mitigation/all-attributes.tf`; digest `sha256:e1a870d0ac8f57cd049b19713024efa0b33bb0d6f7d3e2cc75464e43a5026739`.

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
  depends_on  = [time_sleep.wait_for_namespace]
  name        = "example-value"
  namespace   = xcsh_namespace.test.name
  description = "Test malicious user mitigation with all attributes"
  disable     = false

  labels = {
    environment = "test"
    team        = "security"
  }

  annotations = {
    purpose = "testing"
  }
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/examples/)
- [xcsh_malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/)
