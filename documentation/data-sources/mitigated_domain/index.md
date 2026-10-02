---
page_title: "xcsh_mitigated_domain"
subcategory: ""
description: "Manages Mitigated Domain in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["mitigated domain"], "body_bytes": 1336, "body_sha256": "sha256:baf813c24f97b0e25fcff4f487aacbe64ec9fe100725b40f7ef917ade7e2b379", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:mitigated_domain:reference", "xcsh-docs:data-sources:mitigated_domain:examples"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:mitigated_domain:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:mitigated_domain:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/mitigated_domain/index.md", "product": "distributed-cloud", "provider_name": "mitigated_domain", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3233303323320020-3201003222322030-3103103212101100-1321230323023131-3310220303032323-0211212333301333-1330120302113020-2012202211332310", "registry_path": "docs/data-sources/mitigated_domain.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/mitigated_domain/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages Mitigated Domain in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["mitigated_domainCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_mitigated_domain

Breadcrumbs:

- xcsh_mitigated_domain

Manages Mitigated Domain in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# MitigatedDomain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing MitigatedDomain by name
data "xcsh_mitigated_domain" "example" {
  name      = "example-mitigated-domain"
  namespace = "staging"
}

output "mitigated_domain_id" {
  value = data.xcsh_mitigated_domain.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/mitigated_domain/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/mitigated_domain/examples/)
