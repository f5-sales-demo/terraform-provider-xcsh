---
page_title: "xcsh_service_policy_set"
subcategory: ""
description: "Manages a Service Policy Set resource in F5 Distributed Cloud for get service_policy_set reads a given object from storage backend for metadata.namespace. configuration. (read-only data source)"
xcsh_docs: {"aliases": ["service policy set"], "body_bytes": 1500, "body_sha256": "sha256:90ab1019ad04019c283433ae523053aa55e2229f8c58af9a0819075e3cd88113", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:service_policy_set:reference", "xcsh-docs:data-sources:service_policy_set:examples"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:service_policy_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_set:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/service_policy_set/index.md", "product": "distributed-cloud", "provider_name": "service_policy_set", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1102322222212201-0323000120102213-0210100111122123-0013000010002101-1001031323000300-3302313333113330-3123230211030001-2322203121212312", "registry_path": "docs/data-sources/service_policy_set.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_set/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a Service Policy Set resource in F5 Distributed Cloud for get service_policy_set reads a given object from storage backend for metadata.namespace. configuration. (read-only data source)", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
