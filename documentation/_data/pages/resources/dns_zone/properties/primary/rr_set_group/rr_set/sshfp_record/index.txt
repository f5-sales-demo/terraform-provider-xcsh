---
page_title: "primary.rr_set_group.rr_set.sshfp_record"
subcategory: "DNS"
description: "DNS SSHFP Record."
xcsh_docs: {"aliases": ["primary rr set group rr set sshfp record"], "body_bytes": 3304, "body_sha256": "sha256:eb8c1e51a1ec3923fbe7a173fc9bb56bc283f5f9aecb1a1877d2166e961ce869", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:sshfp_record:values"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:sshfp_record", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set", "path": "documentation/resources/dns_zone/properties/primary/rr_set_group/rr_set/sshfp_record/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3021303131230121-2130011003320213-1320312132022020-1030220220200211-1220030223000022-0033023310312020-1222101103301213-2133330022031311", "registry_path": "docs/guides/resources--dns_zone--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.sshfp_record:RequiredObjectAttributes:values", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:sshfp_record:values", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "rr_set_group", "rr_set", "sshfp_record"], "schema_version": 1, "sections": [{"aliases": ["primary rr set group rr set sshfp record name"], "anchor": "schema-primary--rr_set_group--rr_set--sshfp_record--name", "description": "SSHFP Record name, please provide only the specific subdomain or record name without the base domain.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:sshfp_record", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "rr_set_group", "rr_set", "sshfp_record", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["primary rr set group rr set sshfp record values"], "anchor": "section", "description": "Configuration parameter for values", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:sshfp_record:values", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.sshfp_record.values:ConflictingListObjectAttributes:sha1_fingerprint,sha256_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:sshfp_record:values:sha1_fingerprint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.sshfp_record.values:ConflictingListObjectAttributes:sha1_fingerprint,sha256_fingerprint", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:sshfp_record:values:sha256_fingerprint", "type": "conflicts"}], "schema_path": ["primary", "rr_set_group", "rr_set", "sshfp_record", "values"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/rr_set_group/rr_set/sshfp_record/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "DNS SSHFP Record.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.rr_set_group.rr_set.sshfp_record

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/)
- [primary.rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/)
- [primary.rr_set_group.rr_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/)
- primary.rr_set_group.rr_set.sshfp_record

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sshfp record.

Additional upstream details:

DNS SSHFP Record.

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
sshfp_record {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-primary--rr_set_group--rr_set--sshfp_record--name"></a>

### name property

Type: `"string"`. Optional.

SSHFP Record name, please provide only the specific subdomain or record name without the base
domain.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/sshfp_record/values/): complete subsection reference.
