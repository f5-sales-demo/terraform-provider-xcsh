---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_authentication."
xcsh_docs: {"aliases": [], "body_bytes": 954, "body_sha256": "sha256:fd1822d0ebd412f1a70c20262d63f7b88c1c8a65607452e36384a4b912afff57", "canonical_id": "xcsh-docs:resources:authentication:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:authentication:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:43532e046dd0e813a92a49d472a5eaa915d3ad0e96dcd6b49097f16328170876", "source_path": "examples/resources/xcsh_authentication/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:authentication:example:resource", "parent_id": "xcsh-docs:resources:authentication:examples", "path": "docs/guides/resources--authentication--example--resource.md", "provider_name": "authentication", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authentication/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_authentication.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["authenticationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md)
- [Examples](resources--authentication--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_authentication/resource.tf`; digest `sha256:43532e046dd0e813a92a49d472a5eaa915d3ad0e96dcd6b49097f16328170876`.

```terraform
# Authentication Resource Example
# Manages a Authentication resource in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Authentication configuration
resource "xcsh_authentication" "example" {
  name      = "example-authentication"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--authentication--examples.md)
- [xcsh_authentication](../resources/authentication.md)
