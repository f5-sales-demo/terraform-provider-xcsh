---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_flow_anomaly."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1279, "body_sha256": "sha256:547249ff16621554c044a9e34ba1f9c9c69b633960237402293aebb6d4b59c61", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:flow_anomaly:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d39bc93b81bc691cec5167bb0d331b58de96886d72cd9a1fb46d22490728b1ca", "source_path": "examples/data-sources/xcsh_flow_anomaly/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:flow_anomaly:example:data-source", "parent_id": "xcsh-docs:data-sources:flow_anomaly:examples", "path": "documentation/data-sources/flow_anomaly/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "flow_anomaly", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1102001120300013-1001321332012222-3302301310002213-0212023230212121-1322020333001221-0022323223011123-0122022322230232-1312203323330312", "registry_path": "docs/guides/data-sources--flow_anomaly--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/flow_anomaly/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_flow_anomaly.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
