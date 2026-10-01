---
page_title: "xcsh_external_connector"
subcategory: ""
description: "xcsh_external_connector for xcsh_external_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1498, "body_sha256": "sha256:d5971fb8e6bcc78092e2eecfb9869502bedfd87fb8fd954275b71d3fedf710b4", "canonical_id": "xcsh-docs:resources:external_connector:fundamentals", "child_ids": ["xcsh-docs:resources:external_connector:reference", "xcsh-docs:resources:external_connector:examples", "xcsh-docs:resources:external_connector:import", "xcsh-docs:resources:external_connector:timeouts"], "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:fundamentals", "parent_id": null, "path": "docs/resources/external_connector.md", "provider_name": "external_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_external_connector for xcsh_external_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_external_connector

Breadcrumbs:

- xcsh_external_connector

Manages a External Connector resource in F5 Distributed Cloud for external\_connector configuration
specification. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ExternalConnector Resource Example
# Manages a External Connector resource in F5 Distributed Cloud for external_connector configuration specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ExternalConnector configuration
resource "xcsh_external_connector" "example" {
  name      = "example-external-connector"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--external_connector--reference.md)
- [Examples](../guides/resources--external_connector--examples.md)
- [Import](../guides/resources--external_connector--import.md)
- [Timeouts](../guides/resources--external_connector--timeouts.md)
