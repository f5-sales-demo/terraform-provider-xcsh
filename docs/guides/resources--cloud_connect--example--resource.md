---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_cloud_connect."
xcsh_docs: {"aliases": [], "body_bytes": 1096, "body_sha256": "sha256:e94c3a420c4fb47cb898e1987833c1e160d2d8805ddd5f8e6b44270965280de2", "canonical_id": "xcsh-docs:resources:cloud_connect:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:cloud_connect:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1c68a2114f267127cdbf735730210ab219b41bb6ec8f0bf42263953d188b9f46", "source_path": "examples/resources/xcsh_cloud_connect/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:cloud_connect:example:resource", "parent_id": "xcsh-docs:resources:cloud_connect:examples", "path": "docs/guides/resources--cloud_connect--example--resource.md", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_connect/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_cloud_connect.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md)
- [Examples](resources--cloud_connect--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cloud_connect/resource.tf`; digest `sha256:1c68a2114f267127cdbf735730210ab219b41bb6ec8f0bf42263953d188b9f46`.

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

## Next pages

- [Examples](resources--cloud_connect--examples.md)
- [xcsh_cloud_connect](../resources/cloud_connect.md)
