---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_cloud_user_account."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1355, "body_sha256": "sha256:61249fa6d3e82df03b360685ff82b559783be96cb54a9dfc205b161b7ef1a83c", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_user_account:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:10ccdaeaddaf605a3535a274963418653193279aea9eade41b140aeacaaa4d59", "source_path": "examples/resources/xcsh_cloud_user_account/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:cloud_user_account:example:resource", "parent_id": "xcsh-docs:resources:cloud_user_account:examples", "path": "documentation/resources/cloud_user_account/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3320102201013303-0323232301333013-3022133202021203-2001003031233102-1100302300311030-2013113123333033-2021303002213021-1123010131030010", "registry_path": "docs/guides/resources--cloud_user_account--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_user_account/examples/resource/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource for xcsh_cloud_user_account.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_cloud_user_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cloud_user_account/resource.tf`; digest `sha256:10ccdaeaddaf605a3535a274963418653193279aea9eade41b140aeacaaa4d59`.

```terraform
# CloudUserAccount Resource Example
# Manages a Cloud User Account resource in F5 Distributed Cloud for cloud user account object create specifications.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudUserAccount configuration
resource "xcsh_cloud_user_account" "example" {
  name      = "example-cloud-user-account"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/examples/)
- [xcsh_cloud_user_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_user_account/)
