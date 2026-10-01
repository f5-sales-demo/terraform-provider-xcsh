---
page_title: "xcsh_cloud_connect"
subcategory: ""
description: "xcsh_cloud_connect for xcsh_cloud_connect."
xcsh_docs: {"aliases": [], "body_bytes": 1434, "body_sha256": "sha256:d7682d16ec1dc5095ededb61a183658dbdd5b1e17b027499c9cff8e2a8ccb5dd", "canonical_id": "xcsh-docs:resources:cloud_connect:fundamentals", "child_ids": ["xcsh-docs:resources:cloud_connect:reference", "xcsh-docs:resources:cloud_connect:examples", "xcsh-docs:resources:cloud_connect:import", "xcsh-docs:resources:cloud_connect:timeouts"], "collection_id": "xcsh-docs:resources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_connect:fundamentals", "parent_id": null, "path": "docs/resources/cloud_connect.md", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_connect/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_cloud_connect for xcsh_cloud_connect.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_cloud_connect

Breadcrumbs:

- xcsh_cloud_connect

Manages a Cloud Connect resource in F5 Distributed Cloud for establishing connectivity to cloud
provider networks.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudConnect Resource Example
# Manages a Cloud Connect resource in F5 Distributed Cloud for establishing connectivity to cloud provider networks.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudConnect configuration
resource "xcsh_cloud_connect" "example" {
  name      = "example-cloud-connect"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--cloud_connect--reference.md)
- [Examples](../guides/resources--cloud_connect--examples.md)
- [Import](../guides/resources--cloud_connect--import.md)
- [Timeouts](../guides/resources--cloud_connect--timeouts.md)
