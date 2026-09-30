---
page_title: "xcsh_flow_anomaly"
subcategory: ""
description: "xcsh_flow_anomaly for xcsh_flow_anomaly."
xcsh_docs: {"aliases": [], "body_bytes": 1189, "body_sha256": "sha256:d99ccfb20830951e0efdf2a126ed427a87344786285e89c1e46bf032b42c156f", "canonical_id": "xcsh-docs:data-sources:flow_anomaly:fundamentals", "child_ids": ["xcsh-docs:data-sources:flow_anomaly:reference", "xcsh-docs:data-sources:flow_anomaly:examples"], "collection_id": "xcsh-docs:data-sources:flow_anomaly:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:flow_anomaly:fundamentals", "parent_id": null, "path": "docs/data-sources/flow_anomaly.md", "provider_name": "flow_anomaly", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/flow_anomaly/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_flow_anomaly for xcsh_flow_anomaly.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_flow_anomaly

Breadcrumbs:

- xcsh_flow_anomaly

Manages a Flow Anomaly resource in F5 Distributed Cloud for flow anomaly specification.
configuration. (read-only data source)

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

## Next pages

- [Property reference](../guides/data-sources--flow_anomaly--reference.md)
- [Examples](../guides/data-sources--flow_anomaly--examples.md)
