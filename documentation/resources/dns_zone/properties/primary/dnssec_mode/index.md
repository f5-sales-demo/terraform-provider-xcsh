---
page_title: "primary.dnssec_mode"
subcategory: "DNS"
description: "DNSSEC Mode."
xcsh_docs: {"aliases": ["primary dnssec mode"], "body_bytes": 1931, "body_sha256": "sha256:c748093294daab54b97d94bb5354f1d7056d013b92644cf475396eb7c0db3a48", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode:disable_spec", "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode:enable"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary", "path": "documentation/resources/dns_zone/properties/primary/dnssec_mode/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2200011203232322-0232022120232023-3232232023010300-3212102331200100-1003120100321102-2232232021023203-1013232302222223-0200033311033102", "registry_path": "docs/guides/resources--dns_zone--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "primary.dnssec_mode:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "primary.dnssec_mode:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode:enable", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "dnssec_mode"], "schema_version": 1, "sections": [{"aliases": ["disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode:disable_spec", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "dnssec_mode", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable"], "anchor": "section", "description": "DNSSEC enable.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:dnssec_mode:enable", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "dnssec_mode", "enable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/dnssec_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "DNSSEC Mode.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [primary.dnssec_mode.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/dnssec_mode/disable_spec/)
- [primary.dnssec_mode.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/dnssec_mode/enable/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
