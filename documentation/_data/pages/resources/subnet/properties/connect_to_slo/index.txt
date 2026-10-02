---
page_title: "connect_to_slo"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["connect to slo"], "body_bytes": 1074, "body_sha256": "sha256:e188a13ab4c57141e131f485d455999849242a30301e3b39d567beaecea52829", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:resources:subnet:properties:connect_to_slo", "parent_id": "xcsh-docs:resources:subnet:reference", "path": "documentation/resources/subnet/properties/connect_to_slo/index.md", "product": "distributed-cloud", "provider_name": "subnet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2221112002303323-1312003022203230-1110230122111201-1302202132320223-1132333021002002-1133133132221231-0220310313110131-0031223211313300", "registry_path": "docs/guides/resources--subnet--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["connect_to_slo"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/subnet/properties/connect_to_slo/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["subnetCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# connect_to_slo

Breadcrumbs:

- [xcsh_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/)
- connect_to_slo

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for connect to slo.

Upstream description:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
connect_to_slo = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/)
- [xcsh_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/)
