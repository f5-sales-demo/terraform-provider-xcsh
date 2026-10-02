---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_protocol_policer."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1384, "body_sha256": "sha256:8ceafe8b8c9fdedd65bd062ecf16bb0cbec25f388075f1ddceea4ff3c8b30419", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protocol_policer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a7c6d355644dc5a658884b8dba6d0261da188cbcd2691b54aca0019ad537569c", "source_path": "examples/resources/xcsh_protocol_policer/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:protocol_policer:example:resource", "parent_id": "xcsh-docs:resources:protocol_policer:examples", "path": "documentation/resources/protocol_policer/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2122230011130102-0023302203312021-0032100333210210-2203213210011030-0222321032023131-1321022121233233-0113001333120303-1212311031212332", "registry_path": "docs/guides/resources--protocol_policer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protocol_policer/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_protocol_policer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_protocol_policer/resource.tf`; digest `sha256:a7c6d355644dc5a658884b8dba6d0261da188cbcd2691b54aca0019ad537569c`.

```terraform
# ProtocolPolicer Resource Example
# Manages protocol_policer object, protocol_policer object contains list of L4 protocol match condition and corresponding traffic rate limits in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ProtocolPolicer configuration
resource "xcsh_protocol_policer" "example" {
  name      = "example-protocol-policer"
  namespace = "system"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/examples/)
- [xcsh_protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/)
