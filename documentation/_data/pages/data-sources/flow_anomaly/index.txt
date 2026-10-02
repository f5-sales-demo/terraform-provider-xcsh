---
page_title: "xcsh_flow_anomaly"
subcategory: ""
description: "Manages a Flow Anomaly resource in F5 Distributed Cloud for flow anomaly specification. configuration. (read-only data source)"
xcsh_docs: {"aliases": ["flow anomaly"], "body_bytes": 1373, "body_sha256": "sha256:47ad98548a03572dbd4b60db161cc600ec87f484dfb1520a40b33a052a9e1a2f", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:flow_anomaly:reference", "xcsh-docs:data-sources:flow_anomaly:examples"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:flow_anomaly:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:flow_anomaly:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/flow_anomaly/index.md", "product": "distributed-cloud", "provider_name": "flow_anomaly", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0312230310200021-0201113122010203-1131300233200330-3312031303122031-2123011133112231-0220012002020113-3221033120022133-2323011011020301", "registry_path": "docs/data-sources/flow_anomaly.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/flow_anomaly/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages a Flow Anomaly resource in F5 Distributed Cloud for flow anomaly specification. configuration. (read-only data source)", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/flow_anomaly/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/flow_anomaly/examples/)
