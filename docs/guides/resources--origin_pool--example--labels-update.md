---
page_title: "Labels update"
subcategory: "Load Balancing"
description: "Labels update for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1201, "body_sha256": "sha256:0141768d4f9b35ff82c8a8a2c37a474325b28b8839d215d67f2c446858cad9b7", "canonical_id": "xcsh-docs:resources:origin_pool:example:labels-update", "child_ids": [], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:d2a6e4ca39384f01c4d80efec2f06fdf2e82a0293580f75d7c78c1f61e96c022", "source_path": "examples/resources/xcsh_origin_pool/labels-update.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:origin_pool:example:labels-update", "parent_id": "xcsh-docs:resources:origin_pool:examples", "path": "docs/guides/resources--origin_pool--example--labels-update.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["labels-update"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/examples/labels-update/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Labels update for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Labels update

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Examples](resources--origin_pool--examples.md)
- Labels update

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_origin_pool/labels-update.tf`; digest `sha256:d2a6e4ca39384f01c4d80efec2f06fdf2e82a0293580f75d7c78c1f61e96c022`.

```terraform
# LabelsUpdate — Acceptance-test-derived Configuration
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

  port = 443

  labels = {
    environment = "example-value"
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
