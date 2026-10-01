---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 1241, "body_sha256": "sha256:e5dbfb1d536da3cc906df42e9f72974f1b4eff08fc8180e13bbb8d4c804f4296", "child_ids": [], "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:7b15b9a747e6ef39c623ab924043149ec709c96c07747307cc5c511433ad09f9", "source_path": "examples/resources/xcsh_nfv_service/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:nfv_service:example:resource", "parent_id": "xcsh-docs:resources:nfv_service:examples", "path": "documentation/resources/nfv_service/examples/resource/index.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_nfv_service/resource.tf`; digest `sha256:7b15b9a747e6ef39c623ab924043149ec709c96c07747307cc5c511433ad09f9`.

```terraform
# NfvService Resource Example
# Manages new NFV service with configured parameters in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NfvService configuration
resource "xcsh_nfv_service" "example" {
  name      = "example-nfv-service"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/examples/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
