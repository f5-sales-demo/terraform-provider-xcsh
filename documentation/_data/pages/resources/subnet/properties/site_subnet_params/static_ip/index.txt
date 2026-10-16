---
page_title: "site_subnet_params.static_ip"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["site subnet params static ip"], "body_bytes": 969, "body_sha256": "sha256:d150dc3ba5d58ec9a306326acb56965250adc41702d261e395026282f568a8f1", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:resources:subnet:properties:site_subnet_params:static_ip", "parent_id": "xcsh-docs:resources:subnet:properties:site_subnet_params", "path": "documentation/resources/subnet/properties/site_subnet_params/static_ip/index.md", "product": "distributed-cloud", "provider_name": "subnet", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-0000120300121031-2031001210121031-2323022232221120-0013310013233000-3222123111101312-0122021331012020-3112332211122123-1012123321213200", "registry_path": "docs/guides/resources--subnet--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["site_subnet_params", "static_ip"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/subnet/properties/site_subnet_params/static_ip/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["subnetCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_subnet_params.static_ip

Breadcrumbs:

- [xcsh_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/)
- [site_subnet_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/subnet/properties/site_subnet_params/)
- site_subnet_params.static_ip

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
static_ip = {}
```

This is an empty object or choice marker. It has no direct properties.
