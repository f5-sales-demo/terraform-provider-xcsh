---
page_title: "snat_pool"
subcategory: "Networking"
description: "SNAT Pool configuration."
xcsh_docs: {"aliases": ["snat pool"], "body_bytes": 1340, "body_sha256": "sha256:db7dee3a4d26885fd49ee8da79fff1f80c68093523f2ddaae34682fcce3100a0", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:endpoint:properties:snat_pool:no_snat_pool", "xcsh-docs:resources:endpoint:properties:snat_pool:snat_pool"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:endpoint:collection", "completeness": "complete", "id": "xcsh-docs:resources:endpoint:properties:snat_pool", "parent_id": "xcsh-docs:resources:endpoint:reference", "path": "documentation/resources/endpoint/properties/snat_pool/index.md", "product": "distributed-cloud", "provider_name": "endpoint", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2320200211122210-3211302200313303-3020303021320320-1031000323132322-3113303203020011-3030011120020011-1330211210211133-1310330111120320", "registry_path": "docs/guides/resources--endpoint--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "snat_pool:ConflictingObjectAttributes:no_snat_pool,snat_pool", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:endpoint:properties:snat_pool:no_snat_pool", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "snat_pool:ConflictingObjectAttributes:no_snat_pool,snat_pool", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:endpoint:properties:snat_pool:snat_pool", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["snat_pool"], "schema_version": 1, "sections": [{"aliases": ["snat pool no snat pool"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:endpoint:properties:snat_pool:no_snat_pool", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["snat_pool", "no_snat_pool"], "syntax": "attribute", "type": "object"}, {"aliases": ["snat pool snat pool"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:resources:endpoint:properties:snat_pool:snat_pool", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["snat_pool", "snat_pool"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/endpoint/properties/snat_pool/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "SNAT Pool configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["endpointCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# snat_pool

Breadcrumbs:

- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/)
- snat_pool

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

## Direct properties

- [no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/snat_pool/no_snat_pool/): complete subsection reference.

- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/endpoint/properties/snat_pool/snat_pool/): complete subsection reference.
