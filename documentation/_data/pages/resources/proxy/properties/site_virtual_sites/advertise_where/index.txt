---
page_title: "site_virtual_sites.advertise_where"
subcategory: ""
description: "Where should this load balancer be available."
xcsh_docs: {"aliases": ["site virtual sites advertise where"], "body_bytes": 3629, "body_sha256": "sha256:fd172d26bd35439768bb8542dd3897f1a3b3d1a46b3d3177b833e01e1951c453", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:proxy:properties:site_virtual_sites:advertise_where:site", "xcsh-docs:resources:proxy:properties:site_virtual_sites:advertise_where:use_default_port", "xcsh-docs:resources:proxy:properties:site_virtual_sites:advertise_where:virtual_site"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:site_virtual_sites:advertise_where", "parent_id": "xcsh-docs:resources:proxy:properties:site_virtual_sites", "path": "documentation/resources/proxy/properties/site_virtual_sites/advertise_where/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0021313233203001-3231101003103230-3120120101021312-1113111010303102-2001002322132330-2330301333313313-1012201310232331-3232210300100002", "registry_path": "docs/guides/resources--proxy--reference--group-005.md", "relationships": [{"anchor": "schema-site_virtual_sites--advertise_where--port", "enforcement": "provider-schema", "group": "site_virtual_sites.advertise_where:ConflictingListObjectAttributes:port,use_default_port", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:site_virtual_sites:advertise_where", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_virtual_sites.advertise_where:ConflictingListObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:site_virtual_sites:advertise_where:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_virtual_sites.advertise_where:ConflictingListObjectAttributes:port,use_default_port", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:site_virtual_sites:advertise_where:use_default_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_virtual_sites.advertise_where:ConflictingListObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:site_virtual_sites:advertise_where:virtual_site", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["site_virtual_sites", "advertise_where"], "schema_version": 1, "sections": [{"aliases": ["site virtual sites advertise where port"], "anchor": "schema-site_virtual_sites--advertise_where--port", "description": "Exclusive with TCP port to Listen.", "document_id": "xcsh-docs:resources:proxy:properties:site_virtual_sites:advertise_where", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_virtual_sites", "advertise_where", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["site virtual sites advertise where site"], "anchor": "section", "description": "This defines a reference to a CE site along with network type and an optional IP address where a load balancer could be advertised.", "document_id": "xcsh-docs:resources:proxy:properties:site_virtual_sites:advertise_where:site", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["site_virtual_sites", "advertise_where", "site"], "syntax": "block", "type": "object"}, {"aliases": ["site virtual sites advertise where use default port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:proxy:properties:site_virtual_sites:advertise_where:use_default_port", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_virtual_sites", "advertise_where", "use_default_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["site virtual sites advertise where virtual site"], "anchor": "section", "description": "This defines a reference to a customer site virtual site along with network type where a load balancer could be advertised.", "document_id": "xcsh-docs:resources:proxy:properties:site_virtual_sites:advertise_where:virtual_site", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["site_virtual_sites", "advertise_where", "virtual_site"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/site_virtual_sites/advertise_where/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Where should this load balancer be available.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["proxyCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_virtual_sites.advertise_where

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [site_virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/site_virtual_sites/)
- site_virtual_sites.advertise_where

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Where should this load balancer be available.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("port",
    "use_default_port"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_site")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
advertise_where {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-site_virtual_sites--advertise_where--port"></a>

### port property

Type: `"number"`. Optional.

Exclusive with \[use\_default\_port\] TCP port to Listen.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/site_virtual_sites/advertise_where/site/): complete subsection reference.

- [use_default_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/site_virtual_sites/advertise_where/use_default_port/): complete subsection reference.

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/site_virtual_sites/advertise_where/virtual_site/): complete subsection reference.
