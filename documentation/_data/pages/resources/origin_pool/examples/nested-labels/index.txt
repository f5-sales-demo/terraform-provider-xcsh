---
page_title: "Nested labels"
subcategory: "Load Balancing"
description: "Nested labels for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1409, "body_sha256": "sha256:48fe204f3b975061add86e3fc59c4b8545b081a2e921af40b404d7bbbf65940b", "child_ids": [], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:2a2c2b7a5d041bafee98f9e4a0464b2e4b5c967ac5f5b30fa2f474ff4f6ef535", "source_path": "examples/resources/xcsh_origin_pool/nested-labels.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:origin_pool:example:nested-labels", "parent_id": "xcsh-docs:resources:origin_pool:examples", "path": "documentation/resources/origin_pool/examples/nested-labels/index.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["nested-labels"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/examples/nested-labels/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Nested labels for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Nested labels

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/examples/)
- Nested labels

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_origin_pool/nested-labels.tf`; digest `sha256:2a2c2b7a5d041bafee98f9e4a0464b2e4b5c967ac5f5b30fa2f474ff4f6ef535`.

```terraform
# NestedLabels — Acceptance-test-derived Configuration
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
  name      = "example"
  namespace = "system"

  port = 8080

  origin_servers {
    public_ip {
      ip = "192.0.2.1"
    }
    labels = {
      "env" = "test"
      "app" = "demo"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/examples/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
