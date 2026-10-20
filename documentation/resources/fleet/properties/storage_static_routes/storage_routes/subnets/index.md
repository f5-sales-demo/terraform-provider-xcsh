---
page_title: "storage_static_routes.storage_routes.subnets"
subcategory: ""
description: "List of route prefixes."
xcsh_docs: {"aliases": ["storage static routes storage routes subnets"], "body_bytes": 2199, "body_sha256": "sha256:d2a6ba07f6b1a0a954fde2470bd7546f7005d306dc19cf2af1350d2b5ead0a6e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:subnets:ipv4", "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:subnets:ipv6"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:subnets", "parent_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes", "path": "documentation/resources/fleet/properties/storage_static_routes/storage_routes/subnets/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0011031323102333-2133013032021320-3023030012331100-2101022003022112-3302133213002111-2211232221112303-0333322213031131-1131101123022222", "registry_path": "docs/guides/resources--fleet--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "storage_static_routes.storage_routes.subnets:ConflictingListObjectAttributes:ipv4,ipv6", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:subnets:ipv4", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_static_routes.storage_routes.subnets:ConflictingListObjectAttributes:ipv4,ipv6", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:subnets:ipv6", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_static_routes", "storage_routes", "subnets"], "schema_version": 1, "sections": [{"aliases": ["storage static routes storage routes subnets ipv4"], "anchor": "section", "description": "IPv4 subnets specified as prefix and prefix-length. Prefix length must be <= 32.", "document_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:subnets:ipv4", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_static_routes", "storage_routes", "subnets", "ipv4"], "syntax": "block", "type": "object"}, {"aliases": ["storage static routes storage routes subnets ipv6"], "anchor": "section", "description": "IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be <= 128.", "document_id": "xcsh-docs:resources:fleet:properties:storage_static_routes:storage_routes:subnets:ipv6", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_static_routes", "storage_routes", "subnets", "ipv6"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_static_routes/storage_routes/subnets/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "List of route prefixes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["fleetCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_static_routes.storage_routes.subnets

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [storage_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/)
- [storage_static_routes.storage_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/)
- storage_static_routes.storage_routes.subnets

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Subnets. List of route prefixes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("ipv4",
    "ipv6")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-20T00:10:01+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

Terraform syntax:

```terraform
subnets {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/subnets/ipv4/): complete subsection reference.

- [ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_static_routes/storage_routes/subnets/ipv6/): complete subsection reference.
