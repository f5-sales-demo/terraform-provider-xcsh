---
page_title: "use_default_keylifetime"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["use default keylifetime"], "body_bytes": 873, "body_sha256": "sha256:c807ef8209694dcc6c6bdc3ab4a77ae886ead9a536f908aa576de4750ae41fe2", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike1:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike1:properties:use_default_keylifetime", "parent_id": "xcsh-docs:resources:ike1:reference", "path": "documentation/resources/ike1/properties/use_default_keylifetime/index.md", "product": "distributed-cloud", "provider_name": "ike1", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-1112213223110313-1203322131320102-0310210100213022-1123032232110031-1130131201010113-0102010333011230-0201023200023132-0221312331321122", "registry_path": "docs/guides/resources--ike1--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["use_default_keylifetime"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike1/properties/use_default_keylifetime/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["ike1CreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_default_keylifetime

Breadcrumbs:

- [xcsh_ike1](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike1/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike1/properties/)
- use_default_keylifetime

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use default keylifetime.

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
use_default_keylifetime = {}
```

This is an empty object or choice marker. It has no direct properties.
