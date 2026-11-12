---
page_title: "xcsh_app_api_group"
subcategory: ""
description: "Reads App API Group information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["app api group"], "body_bytes": 1326, "body_sha256": "sha256:8d74b0fb0ef7286699182ccde4261acc9444b56afe756fd897a8ae6abfcdddde", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:data-sources:app_api_group:reference", "xcsh-docs:data-sources:app_api_group:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_api_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_api_group:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/app_api_group/index.md", "product": "distributed-cloud", "provider_name": "app_api_group", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "data-sources", "registry_anchor": "canonical-3011310102212300-2112210220330313-1330322333011033-1111313021020310-3121110131333221-1200033222203013-2121123032213301-2023132013121010", "registry_path": "docs/data-sources/app_api_group.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_api_group/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Reads App API Group information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["app_api_groupCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_app_api_group

Breadcrumbs:

- xcsh_app_api_group

Reads App API Group information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AppAPIGroup Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppAPIGroup by name
data "xcsh_app_api_group" "example" {
  name      = "example-app-api-group"
  namespace = "staging"
}

output "app_api_group_id" {
  value = data.xcsh_app_api_group.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/examples/)
