---
page_title: "xcsh_cminstance"
subcategory: ""
description: "Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["cminstance"], "body_bytes": 1651, "body_sha256": "sha256:46409f27847085721582abdc87f39115460ef569aa55d57b42de9d039898f544", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:cminstance:reference", "xcsh-docs:resources:cminstance:examples", "xcsh-docs:resources:cminstance:import", "xcsh-docs:resources:cminstance:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cminstance:collection", "completeness": "complete", "id": "xcsh-docs:resources:cminstance:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/cminstance/index.md", "product": "distributed-cloud", "provider_name": "cminstance", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-3030301003002223-1121312121232112-1232303010012000-0213330232003222-0030121133101123-1130200110322121-0013030301203132-0203031312323211", "registry_path": "docs/resources/cminstance.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cminstance/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["cminstanceCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_cminstance

Breadcrumbs:

- xcsh_cminstance

Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed
Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Cminstance Resource Example
# Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Cminstance configuration
resource "xcsh_cminstance" "example" {
  name      = "example-cminstance"
  namespace = "staging"

  port     = 1
  username = "example-value"
}
```

## Root configuration

Required root properties: `name`, `namespace`, `port`, `username`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cminstance/lifecycle/timeouts/)
