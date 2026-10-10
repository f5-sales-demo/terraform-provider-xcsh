---
page_title: "With labels"
subcategory: "Load Balancing"
description: "With labels for xcsh_origin_pool."
xcsh_docs: {"aliases": ["with-labels"], "body_bytes": 1239, "body_sha256": "sha256:d2a58fbf3be367dc436b8ce1a4ee5e60121efa6adcf34e2cfe861f5218f7d866", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:492fffafa075a1f28ad253873263057c634ae541eb762e94f4e227737c234dcd", "source_path": "examples/resources/xcsh_origin_pool/with-labels.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:origin_pool:example:with-labels", "parent_id": "xcsh-docs:resources:origin_pool:examples", "path": "documentation/resources/origin_pool/examples/with-labels/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2311121333023321-0221132131023221-1102302002223221-2320012302100233-0020022101223001-1023222202000033-2322203032102330-3002213122132033", "registry_path": "docs/guides/resources--origin_pool--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["with-labels"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/examples/with-labels/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "With labels for xcsh_origin_pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["origin_poolCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# With labels

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/examples/)
- With labels

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_origin_pool/with-labels.tf`; digest `sha256:492fffafa075a1f28ad253873263057c634ae541eb762e94f4e227737c234dcd`.

```terraform
# WithLabels — Acceptance-test-derived Configuration
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

resource "xcsh_origin_pool" "test" {
  name        = "example"
  namespace   = "system"
  description = "Test origin pool"

  port = 443

  labels = {
    environment = "test"
    team        = "platform"
  }

  origin_servers {
    public_name {
      dns_name = "example.com"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}
```
