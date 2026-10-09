---
page_title: "primary.dnssec_mode"
subcategory: "DNS"
description: "DNSSEC Mode."
xcsh_docs: {"aliases": ["primary dnssec mode"], "body_bytes": 1438, "body_sha256": "sha256:7f9dcb9388d6473033666f2d59697e82f6d0b3c2510fe3d4fecd0bf429e69e2b", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode:disable_spec", "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode:enable"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary", "path": "documentation/resources/dns_zone/properties/primary/dnssec_mode/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2200011203232322-0232022120232023-3232232023010300-3212102331200100-1003120100321102-2232232021023203-1013232302222223-0200033311033102", "registry_path": "docs/guides/resources--dns_zone--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "primary.dnssec_mode:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "primary.dnssec_mode:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode:enable", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "dnssec_mode"], "schema_version": 1, "sections": [{"aliases": ["primary dnssec mode disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode:disable_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "dnssec_mode", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["primary dnssec mode enable"], "anchor": "section", "description": "DNSSEC enable.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode:enable", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "dnssec_mode", "enable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/dnssec_mode/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "DNSSEC Mode.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.dnssec_mode

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/)
- primary.dnssec_mode

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

DNSSEC Mode.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable")}
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
  "x-ves-oneof-field-mode": "[\"disable\",\"enable\"]"
}
```

Terraform syntax:

```terraform
dnssec_mode {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/dnssec_mode/disable_spec/): complete subsection reference.

- [enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/dnssec_mode/enable/): complete subsection reference.
