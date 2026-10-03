---
page_title: "xcsh_infraprotect_mitigation_ips"
subcategory: ""
description: "Resource retrieval operation."
xcsh_docs: {"aliases": ["infraprotect mitigation ips"], "body_bytes": 1370, "body_sha256": "sha256:63824e5fc448dc810766364130193286b71096d99e1ae18ae64fdaf61e74c07f", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:infraprotect_mitigation_ips:reference", "xcsh-docs:data-sources:infraprotect_mitigation_ips:examples"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/infraprotect_mitigation_ips/index.md", "product": "distributed-cloud", "provider_name": "infraprotect_mitigation_ips", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1223233033223000-1032001012013022-1301320302130033-2311101100202302-3130231002213230-3320300113130330-0113131111131313-3222111203302102", "registry_path": "docs/data-sources/infraprotect_mitigation_ips.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/infraprotect_mitigation_ips/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource retrieval operation.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_infraprotect_mitigation_ips

Breadcrumbs:

- xcsh_infraprotect_mitigation_ips

Resource retrieval operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# InfraprotectMitigationIps DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_infraprotect_mitigation_ips" "example" {
  mitigation_id = "example-value"
  namespace     = "example-value"
}

output "infraprotect_mitigation_ips_result" {
  value = data.xcsh_infraprotect_mitigation_ips.example
}
```

## Root configuration

Required root properties: `mitigation_id`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/examples/)
