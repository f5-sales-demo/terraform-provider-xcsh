---
page_title: "xcsh_flow_anomaly"
subcategory: ""
description: "Reads Flow Anomaly information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["flow anomaly"], "body_bytes": 1317, "body_sha256": "sha256:05cc292b1881b5c208d59aec6ade4e68b1b0291bc0f9e2013d133863696fcd62", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:flow_anomaly:reference", "xcsh-docs:data-sources:flow_anomaly:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:flow_anomaly:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:flow_anomaly:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/flow_anomaly/index.md", "product": "distributed-cloud", "provider_name": "flow_anomaly", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-0312230310200021-0201113122010203-1131300233200330-3312031303122031-2123011133112231-0220012002020113-3221033120022133-2323011011020301", "registry_path": "docs/data-sources/flow_anomaly.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/flow_anomaly/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Reads Flow Anomaly information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_flow_anomaly

Breadcrumbs:

- xcsh_flow_anomaly

Reads Flow Anomaly information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# FlowAnomaly Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing FlowAnomaly by name
data "xcsh_flow_anomaly" "example" {
  name      = "example-flow-anomaly"
  namespace = "staging"
}

output "flow_anomaly_id" {
  value = data.xcsh_flow_anomaly.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/flow_anomaly/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/flow_anomaly/examples/)
