---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_crl."
xcsh_docs: {"aliases": [], "body_bytes": 958, "body_sha256": "sha256:7738774a2c3b1840c75f7247e17e8dd4c8bb3aeff90abe340e34ba948054ad75", "canonical_id": "xcsh-docs:resources:crl:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:crl:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:787cb67d26f54e510168be44dfb1793c35d97676385e49506a4c31f4d8662bf3", "source_path": "examples/resources/xcsh_crl/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:crl:example:resource", "parent_id": "xcsh-docs:resources:crl:examples", "path": "docs/guides/resources--crl--example--resource.md", "provider_name": "crl", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/crl/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_crl.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["crlCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_crl](../resources/crl.md)
- [Examples](resources--crl--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_crl/resource.tf`; digest `sha256:787cb67d26f54e510168be44dfb1793c35d97676385e49506a4c31f4d8662bf3`.

```terraform
# CRL Resource Example
# Manages a CRL resource in F5 Distributed Cloud for api to create crl object.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CRL configuration
resource "xcsh_crl" "example" {
  name      = "example-crl"
  namespace = "staging"

  refresh_interval = 6
  server_address   = "example-value"
  server_port      = 1
  timeout          = 1
}
```

## Next pages

- [Examples](resources--crl--examples.md)
- [xcsh_crl](../resources/crl.md)
