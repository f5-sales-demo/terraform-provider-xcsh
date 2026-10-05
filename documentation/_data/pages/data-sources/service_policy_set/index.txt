---
page_title: "xcsh_service_policy_set"
subcategory: ""
description: "Manages a Service Policy Set resource in F5 Distributed Cloud for get service_policy_set reads a given object from storage backend for metadata.namespace. configuration. (read-only data source)"
xcsh_docs: {"aliases": ["service policy set"], "body_bytes": 1500, "body_sha256": "sha256:90ab1019ad04019c283433ae523053aa55e2229f8c58af9a0819075e3cd88113", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:service_policy_set:reference", "xcsh-docs:data-sources:service_policy_set:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:service_policy_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_set:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/service_policy_set/index.md", "product": "distributed-cloud", "provider_name": "service_policy_set", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-1102322222212201-0323000120102213-0210100111122123-0013000010002101-1001031323000300-3302313333113330-3123230211030001-2322203121212312", "registry_path": "docs/data-sources/service_policy_set.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_set/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Manages a Service Policy Set resource in F5 Distributed Cloud for get service_policy_set reads a given object from storage backend for metadata.namespace. configuration. (read-only data source)", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_service_policy_set

Breadcrumbs:

- xcsh_service_policy_set

Manages a Service Policy Set resource in F5 Distributed Cloud for get service\_policy\_set reads a
given object from storage backend for metadata.namespace. configuration. (read-only data source)

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ServicePolicySet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ServicePolicySet by name
data "xcsh_service_policy_set" "example" {
  name      = "example-service-policy-set"
  namespace = "staging"
}

output "service_policy_set_id" {
  value = data.xcsh_service_policy_set.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_set/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_set/examples/)
