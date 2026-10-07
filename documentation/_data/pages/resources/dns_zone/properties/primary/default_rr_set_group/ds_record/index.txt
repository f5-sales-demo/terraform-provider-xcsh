---
page_title: "primary.default_rr_set_group.ds_record"
subcategory: "DNS"
description: "DNS DS Record."
xcsh_docs: {"aliases": ["primary default rr set group ds record"], "body_bytes": 3096, "body_sha256": "sha256:6c9af269d4af1560ee1c74516fe530fcc972397975943992a92479eecaaf41d7", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group", "path": "documentation/resources/dns_zone/properties/primary/default_rr_set_group/ds_record/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-2221112200013312-2222002032211003-3222010212133221-2312211212122130-3033033011333033-1322303331020103-1003331022130023-3312202311121303", "registry_path": "docs/guides/resources--dns_zone--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.ds_record:RequiredObjectAttributes:values", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "default_rr_set_group", "ds_record"], "schema_version": 1, "sections": [{"aliases": ["primary default rr set group ds record name"], "anchor": "schema-primary--default_rr_set_group--ds_record--name", "description": "DS Record name, please provide only the specific subdomain or record name without the base domain.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "ds_record", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["primary default rr set group ds record values"], "anchor": "section", "description": "Configuration parameter for values", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.ds_record.values:ConflictingListObjectAttributes:sha1_digest,sha256_digest", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values:sha1_digest", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.ds_record.values:ConflictingListObjectAttributes:sha1_digest,sha384_digest", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values:sha1_digest", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.ds_record.values:ConflictingListObjectAttributes:sha1_digest,sha256_digest", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values:sha256_digest", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.ds_record.values:ConflictingListObjectAttributes:sha256_digest,sha384_digest", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values:sha256_digest", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.ds_record.values:ConflictingListObjectAttributes:sha1_digest,sha384_digest", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values:sha384_digest", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.ds_record.values:ConflictingListObjectAttributes:sha256_digest,sha384_digest", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values:sha384_digest", "type": "conflicts"}, {"anchor": "schema-primary--default_rr_set_group--ds_record--values--key_tag", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.ds_record.values:RequiredListObjectAttributes:key_tag", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:ds_record:values", "type": "requires"}], "schema_path": ["primary", "default_rr_set_group", "ds_record", "values"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/default_rr_set_group/ds_record/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "DNS DS Record.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.default_rr_set_group.ds_record

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/)
- [primary.default_rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/)
- primary.default_rr_set_group.ds_record

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

DNS DS Record. DNS DS Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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
ds_record {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-primary--default_rr_set_group--ds_record--name"></a>

### name property

Type: `"string"`. Optional.

DS Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/ds_record/values/): complete subsection reference.
