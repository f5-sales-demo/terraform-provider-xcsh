---
page_title: "Resource"
subcategory: "Monitoring"
description: "Resource for xcsh_healthcheck."
xcsh_docs: {"aliases": [], "body_bytes": 1206, "body_sha256": "sha256:5358ffe20de1730ff8d2e5bb01af978a4c54227a6b35bbbec266f647c1ad4fad", "canonical_id": "xcsh-docs:resources:healthcheck:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3cd7562f625a6457911202f74161bf5f6a0eec910fe3f71c8a9f0c0e1ea3d639", "source_path": "examples/resources/xcsh_healthcheck/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:healthcheck:example:resource", "parent_id": "xcsh-docs:resources:healthcheck:examples", "path": "docs/guides/resources--healthcheck--example--resource.md", "provider_name": "healthcheck", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_healthcheck.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["healthcheckCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md)
- [Examples](resources--healthcheck--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_healthcheck/resource.tf`; digest `sha256:3cd7562f625a6457911202f74161bf5f6a0eec910fe3f71c8a9f0c0e1ea3d639`.

```terraform
# Healthcheck Resource Example
# Manages a Healthcheck resource in F5 Distributed Cloud for healthcheck object defines method to determine if the given endpoint is healthy.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Healthcheck configuration
resource "xcsh_healthcheck" "example" {
  name      = "example-healthcheck"
  namespace = "staging"

  healthy_threshold   = 1
  interval            = 1
  timeout             = 1
  unhealthy_threshold = 1
}
```

## Next pages

- [Examples](resources--healthcheck--examples.md)
- [xcsh_healthcheck](../resources/healthcheck.md)
