---
page_title: "Resource"
subcategory: "Infrastructure"
description: "Resource for xcsh_cloud_credentials."
xcsh_docs: {"aliases": [], "body_bytes": 1130, "body_sha256": "sha256:bc2d4ee2b728cfbfc4e90f60c449a41a2dd9f2cea203e5d7fc213e11b44becda", "canonical_id": "xcsh-docs:resources:cloud_credentials:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:cloud_credentials:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:59ca9345886180c4c5420989ec1dac0bd8bd93a36d37acaaca6a81918a8c91f1", "source_path": "examples/resources/xcsh_cloud_credentials/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:cloud_credentials:example:resource", "parent_id": "xcsh-docs:resources:cloud_credentials:examples", "path": "docs/guides/resources--cloud_credentials--example--resource.md", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_credentials/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_cloud_credentials.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_cloud_credentials](../resources/cloud_credentials.md)
- [Examples](resources--cloud_credentials--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cloud_credentials/resource.tf`; digest `sha256:59ca9345886180c4c5420989ec1dac0bd8bd93a36d37acaaca6a81918a8c91f1`.

```terraform
# CloudCredentials Resource Example
# Manages a Cloud Credentials resource in F5 Distributed Cloud for api to create cloud_credentials object.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudCredentials configuration
resource "xcsh_cloud_credentials" "example" {
  name      = "example-cloud-credentials"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--cloud_credentials--examples.md)
- [xcsh_cloud_credentials](../resources/cloud_credentials.md)
