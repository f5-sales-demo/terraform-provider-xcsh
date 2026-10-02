---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_nfv_service."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1266, "body_sha256": "sha256:2f3a9d7f9395d92564fd153e1dab2600987ff1b3ff0919ec089186a4626be243", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e1bcf01b10f5271939ad8fda84b9d63610e3af5498aba69a8e2b34b2c60983a2", "source_path": "examples/data-sources/xcsh_nfv_service/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:nfv_service:example:data-source", "parent_id": "xcsh-docs:data-sources:nfv_service:examples", "path": "documentation/data-sources/nfv_service/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2323330232101000-2311330213121130-0232012023331321-0302310130122220-1322233212010303-0133031033103223-3002231002023331-1022303312123022", "registry_path": "docs/guides/data-sources--nfv_service--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_nfv_service.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_nfv_service/data-source.tf`; digest `sha256:e1bcf01b10f5271939ad8fda84b9d63610e3af5498aba69a8e2b34b2c60983a2`.

```terraform
# NfvService Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NfvService by name
data "xcsh_nfv_service" "example" {
  name      = "example-nfv-service"
  namespace = "staging"
}

output "nfv_service_id" {
  value = data.xcsh_nfv_service.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/examples/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
