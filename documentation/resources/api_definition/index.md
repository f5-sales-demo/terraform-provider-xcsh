---
page_title: "xcsh_api_definition"
subcategory: "API Management"
description: "Manages API Definition in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["api definition"], "body_bytes": 1620, "body_sha256": "sha256:61bb8092ffe0f064a4ef48c3f791f7a66599ae5bbd414533401e793fde867e73", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:resources:api_definition:reference", "xcsh-docs:resources:api_definition:examples", "xcsh-docs:resources:api_definition:import", "xcsh-docs:resources:api_definition:timeouts"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:api_definition:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_definition:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/api_definition/index.md", "product": "distributed-cloud", "provider_name": "api_definition", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0122213120002012-3231302001220300-2001233123112233-1130330020321211-3113132132002121-2133030200111003-0210010013030021-1211030320012300", "registry_path": "docs/resources/api_definition.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_definition/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages API Definition in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [{"name": "api_endpoint", "source": "receipt-pinned-dependency:optional"}], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["api_definitionCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/lifecycle/timeouts/)
