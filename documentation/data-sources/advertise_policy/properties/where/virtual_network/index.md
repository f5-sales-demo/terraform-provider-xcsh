---
page_title: "where.virtual_network"
subcategory: ""
description: "This specifies a direct reference to a network configuration object."
xcsh_docs: {"aliases": ["where virtual network"], "body_bytes": 1485, "body_sha256": "sha256:e3f958b3574f2407ffecafd3380ecd5fa50619a26312ad528a88b8055371010c", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:advertise_policy:properties:where:virtual_network:ref"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:advertise_policy:properties:where:virtual_network", "parent_id": "xcsh-docs:data-sources:advertise_policy:properties:where", "path": "documentation/data-sources/advertise_policy/properties/where/virtual_network/index.md", "product": "distributed-cloud", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3022121311302203-3101121122331332-1030013301210030-1330202012133112-2103012211120022-3202303132311322-2231200013110312-3231223110331133", "registry_path": "docs/guides/data-sources--advertise_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["where", "virtual_network"], "schema_version": 1, "sections": [{"aliases": ["ref"], "anchor": "section", "description": "A virtual network direct reference.", "document_id": "xcsh-docs:data-sources:advertise_policy:properties:where:virtual_network:ref", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["where", "virtual_network", "ref"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/advertise_policy/properties/where/virtual_network/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This specifies a direct reference to a network configuration object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where.virtual_network

Breadcrumbs:

- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/)
- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/)
- where.virtual_network

<a id="section"></a>

Type: `"single"`. Computed.

Specifies a direct reference to a network configuration object.

Upstream description:

This specifies a direct reference to a network configuration object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

- [ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/virtual_network/ref/): complete subsection reference.

## Next pages

- [where.virtual_network.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/virtual_network/ref/)
- [where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/)
- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/)
