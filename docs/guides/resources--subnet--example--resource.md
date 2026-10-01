---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_subnet."
xcsh_docs: {"aliases": [], "body_bytes": 1027, "body_sha256": "sha256:3c950f5b01dfae198ea8dd8bb1c9a27970ef90705389f565e778db44cd3596b8", "canonical_id": "xcsh-docs:resources:subnet:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:subnet:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:4b17c21430cf56ad4295d650386f0dba8c0f28367f60769889a0c3628d334ae0", "source_path": "examples/resources/xcsh_subnet/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:subnet:example:resource", "parent_id": "xcsh-docs:resources:subnet:examples", "path": "docs/guides/resources--subnet--example--resource.md", "provider_name": "subnet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/subnet/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_subnet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["subnetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md)
- [Examples](resources--subnet--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_subnet/resource.tf`; digest `sha256:4b17c21430cf56ad4295d650386f0dba8c0f28367f60769889a0c3628d334ae0`.

```terraform
# Subnet Resource Example
# Manages a Subnet resource in F5 Distributed Cloud for subnet object contains configuration for an interface of a vm/pod.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Subnet configuration
resource "xcsh_subnet" "example" {
  name      = "example-subnet"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--subnet--examples.md)
- [xcsh_subnet](../resources/subnet.md)
