---
page_title: "With labels"
subcategory: "Load Balancing"
description: "With labels for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1252, "body_sha256": "sha256:6d0c00d7f89cdaa4cd6d9f141cdb1256f2d453182a2a33950c3bb8ae01d0de1f", "canonical_id": "xcsh-docs:resources:origin_pool:example:with-labels", "child_ids": [], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:492fffafa075a1f28ad253873263057c634ae541eb762e94f4e227737c234dcd", "source_path": "examples/resources/xcsh_origin_pool/with-labels.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:origin_pool:example:with-labels", "parent_id": "xcsh-docs:resources:origin_pool:examples", "path": "docs/guides/resources--origin_pool--example--with-labels.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["with-labels"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/examples/with-labels/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "With labels for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# With labels

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Examples](resources--origin_pool--examples.md)
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

## Next pages

- [Examples](resources--origin_pool--examples.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
