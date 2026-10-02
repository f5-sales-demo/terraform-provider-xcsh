---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_discovery."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1281, "body_sha256": "sha256:76b34cebdac67502c4095c4b121c97d6045d5d39c0f6cd08b355e6efbb76c961", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:639a573707f4cc4152b42dbe4ad48b3edaa1017b066d38b4ee06ee6d80b4f03c", "source_path": "examples/resources/xcsh_discovery/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:discovery:example:resource", "parent_id": "xcsh-docs:resources:discovery:examples", "path": "documentation/resources/discovery/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3231332330332133-3012022130101021-0231011222223303-3001221030232022-3032130300020103-1220010012220120-2033233200021330-3222223122312033", "registry_path": "docs/guides/resources--discovery--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_discovery.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["discoveryCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_discovery/resource.tf`; digest `sha256:639a573707f4cc4152b42dbe4ad48b3edaa1017b066d38b4ee06ee6d80b4f03c`.

```terraform
# Discovery Resource Example
# Manages a Discovery resource in F5 Distributed Cloud for api to create discovery object for a site or virtual site in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Discovery configuration
resource "xcsh_discovery" "example" {
  name      = "example-discovery"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/examples/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
