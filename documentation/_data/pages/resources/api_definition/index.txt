---
page_title: "xcsh_api_definition"
subcategory: "API Management"
description: "Manages API Definition in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["api definition"], "body_bytes": 1633, "body_sha256": "sha256:e210c4c79f0f7f425ebb951efcd2dc31f729b241dd665bb889ba87bda14cb3a7", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:resources:api_definition:reference", "xcsh-docs:resources:api_definition:examples", "xcsh-docs:resources:api_definition:import", "xcsh-docs:resources:api_definition:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:api_definition:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_definition:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/api_definition/index.md", "product": "distributed-cloud", "provider_name": "api_definition", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-0122213120002012-3231302001220300-2001233123112233-1130330020321211-3113132132002121-2133030200111003-0210010013030021-1211030320012300", "registry_path": "docs/resources/api_definition.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_definition/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Manages API Definition in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [{"name": "api_endpoint", "source": "receipt-pinned-dependency:optional"}], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["api_definitionCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_api_definition

Breadcrumbs:

- xcsh_api_definition

Manages API Definition in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Optional integrations: `api_endpoint`.

- api_endpoint: Endpoints defined by this API

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# APIDefinition Resource Example
# Manages API Definition in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APIDefinition configuration
resource "xcsh_api_definition" "example" {
  name      = "example-api-definition"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/lifecycle/timeouts/)
