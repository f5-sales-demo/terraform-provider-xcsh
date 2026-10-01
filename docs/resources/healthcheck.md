---
page_title: "xcsh_healthcheck"
subcategory: "Monitoring"
description: "xcsh_healthcheck for xcsh_healthcheck."
xcsh_docs: {"aliases": [], "body_bytes": 1764, "body_sha256": "sha256:a3d2082f3effa6102e81d633a34561f171a77ece5e12179af24188d32c7a5229", "canonical_id": "xcsh-docs:resources:healthcheck:fundamentals", "child_ids": ["xcsh-docs:resources:healthcheck:reference", "xcsh-docs:resources:healthcheck:examples", "xcsh-docs:resources:healthcheck:import", "xcsh-docs:resources:healthcheck:timeouts"], "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "id": "xcsh-docs:resources:healthcheck:fundamentals", "parent_id": null, "path": "docs/resources/healthcheck.md", "provider_name": "healthcheck", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_healthcheck for xcsh_healthcheck.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["healthcheckCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_healthcheck

Breadcrumbs:

- xcsh_healthcheck

Manages a Healthcheck resource in F5 Distributed Cloud for healthcheck object defines method to
determine if the given endpoint is healthy. single healthcheck object can be referred to by one or
many cluster objects. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `healthy_threshold`, `interval`, `name`, `namespace`, `timeout`, `unhealthy_threshold`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--healthcheck--reference.md)
- [Examples](../guides/resources--healthcheck--examples.md)
- [Import](../guides/resources--healthcheck--import.md)
- [Timeouts](../guides/resources--healthcheck--timeouts.md)
