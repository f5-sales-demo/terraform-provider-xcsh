---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_flow_anomaly."
xcsh_docs: {"aliases": [], "body_bytes": 1180, "body_sha256": "sha256:fc805c1d4b090ab41bd1020a4c6b08fd3e0c6411e88a6f0d9b6b9ca3c325cb33", "child_ids": [], "collection_id": "xcsh-docs:data-sources:flow_anomaly:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d39bc93b81bc691cec5167bb0d331b58de96886d72cd9a1fb46d22490728b1ca", "source_path": "examples/data-sources/xcsh_flow_anomaly/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:flow_anomaly:example:data-source", "parent_id": "xcsh-docs:data-sources:flow_anomaly:examples", "path": "documentation/data-sources/flow_anomaly/examples/data-source/index.md", "provider_name": "flow_anomaly", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/flow_anomaly/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_flow_anomaly.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_flow_anomaly](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/flow_anomaly/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/flow_anomaly/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_flow_anomaly/data-source.tf`; digest `sha256:d39bc93b81bc691cec5167bb0d331b58de96886d72cd9a1fb46d22490728b1ca`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/flow_anomaly/examples/)
- [xcsh_flow_anomaly](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/flow_anomaly/)
