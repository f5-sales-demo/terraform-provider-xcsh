---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_code_base_integration."
xcsh_docs: {"aliases": [], "body_bytes": 1120, "body_sha256": "sha256:7e23344a2d35f0f4ea250a85b53856b7b5f81690b513809420b1375c027fe014", "canonical_id": "xcsh-docs:resources:code_base_integration:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:code_base_integration:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:96e41779f04da03e1e24b87e34c40ee5955713fd86df36344fb3cbec779ab18f", "source_path": "examples/resources/xcsh_code_base_integration/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:code_base_integration:example:resource", "parent_id": "xcsh-docs:resources:code_base_integration:examples", "path": "docs/guides/resources--code_base_integration--example--resource.md", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/code_base_integration/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_code_base_integration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_code_base_integration](../resources/code_base_integration.md)
- [Examples](resources--code_base_integration--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_code_base_integration/resource.tf`; digest `sha256:96e41779f04da03e1e24b87e34c40ee5955713fd86df36344fb3cbec779ab18f`.

```terraform
# CodeBaseIntegration Resource Example
# Manages integration details in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CodeBaseIntegration configuration
resource "xcsh_code_base_integration" "example" {
  name      = "example-code-base-integration"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--code_base_integration--examples.md)
- [xcsh_code_base_integration](../resources/code_base_integration.md)
