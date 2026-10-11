---
page_title: "xcsh_token"
subcategory: "Identity"
description: "Manages new token. Token object is used to manage site admission. User must generate token before provisioning and pass this token to site during it's registration in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["token"], "body_bytes": 1612, "body_sha256": "sha256:6d1bc0a34aac76cac094f1a34d7827babc45ffa3e4a59450f9c9ea8ce6060551", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:resources:token:reference", "xcsh-docs:resources:token:examples", "xcsh-docs:resources:token:import", "xcsh-docs:resources:token:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:token:collection", "completeness": "complete", "id": "xcsh-docs:resources:token:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/token/index.md", "product": "distributed-cloud", "provider_name": "token", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-1033011013002132-3200233320301222-0133100331111013-2201110031233021-3032032333322002-0033103111032332-0232210133220103-2232103023033020", "registry_path": "docs/resources/token.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/token/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Manages new token. Token object is used to manage site admission. User must generate token before provisioning and pass this token to site during it's registration in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["tokenCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_token

Breadcrumbs:

- xcsh_token

Manages new token. Token object is used to manage site admission. User must generate token before
provisioning and pass this token to site during it's registration in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Token Resource Example
# Manages new token.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Token configuration
resource "xcsh_token" "example" {
  name      = "example-token"
  namespace = "system"
  type      = 1
  site_name = "example-securemesh-site"
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/token/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/token/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/token/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/token/lifecycle/timeouts/)
