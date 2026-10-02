---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_nfv_service."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1241, "body_sha256": "sha256:e5dbfb1d536da3cc906df42e9f72974f1b4eff08fc8180e13bbb8d4c804f4296", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:7b15b9a747e6ef39c623ab924043149ec709c96c07747307cc5c511433ad09f9", "source_path": "examples/resources/xcsh_nfv_service/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:nfv_service:example:resource", "parent_id": "xcsh-docs:resources:nfv_service:examples", "path": "documentation/resources/nfv_service/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-1213221123222201-0030223100121310-2210301222223030-2232002320212033-0031131003022211-2120300230211333-1333011030220133-3133220011303213", "registry_path": "docs/guides/resources--nfv_service--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_nfv_service.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_nfv_service/resource.tf`; digest `sha256:7b15b9a747e6ef39c623ab924043149ec709c96c07747307cc5c511433ad09f9`.

```terraform
# NfvService Resource Example
# Manages new NFV service with configured parameters in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NfvService configuration
resource "xcsh_nfv_service" "example" {
  name      = "example-nfv-service"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/examples/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
