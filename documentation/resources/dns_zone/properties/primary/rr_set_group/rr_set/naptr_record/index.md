---
page_title: "primary.rr_set_group.rr_set.naptr_record"
subcategory: "DNS"
description: "DNS NAPTR Record."
xcsh_docs: {"aliases": ["primary rr set group rr set naptr record"], "body_bytes": 3685, "body_sha256": "sha256:fa6c5f646f13cf15da9258e3e0ce43257287aa2113259c40b4373d3b128f9675", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:naptr_record:values"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:naptr_record", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set", "path": "documentation/resources/dns_zone/properties/primary/rr_set_group/rr_set/naptr_record/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0031001220121130-2232011132130122-3223223310210021-2102322233011223-2233030012200001-3232113133123013-2230020011101011-3203003021223223", "registry_path": "docs/guides/resources--dns_zone--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.naptr_record:RequiredObjectAttributes:values", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:naptr_record:values", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "rr_set_group", "rr_set", "naptr_record"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-primary--rr_set_group--rr_set--naptr_record--name", "description": "NAPTR Record name, please provide only the specific subdomain or record name without the base domain.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:naptr_record", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "rr_set_group", "rr_set", "naptr_record", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["values"], "anchor": "section", "description": "Configuration parameter for values", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:naptr_record:values", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-primary--rr_set_group--rr_set--naptr_record--values--flags", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.naptr_record.values:RequiredListObjectAttributes:flags,order,preference", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:naptr_record:values", "type": "requires"}, {"anchor": "schema-primary--rr_set_group--rr_set--naptr_record--values--order", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.naptr_record.values:RequiredListObjectAttributes:flags,order,preference", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:naptr_record:values", "type": "requires"}, {"anchor": "schema-primary--rr_set_group--rr_set--naptr_record--values--preference", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.naptr_record.values:RequiredListObjectAttributes:flags,order,preference", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:naptr_record:values", "type": "requires"}], "schema_path": ["primary", "rr_set_group", "rr_set", "naptr_record", "values"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/rr_set_group/rr_set/naptr_record/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "DNS NAPTR Record.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.rr_set_group.rr_set.naptr_record

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/)
- [primary.rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/)
- [primary.rr_set_group.rr_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/)
- primary.rr_set_group.rr_set.naptr_record

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for naptr record.

Upstream description:

DNS NAPTR Record.

Provider validators and defaults (from schema source):

```go
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
naptr_record {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-primary--rr_set_group--rr_set--naptr_record--name"></a>

### name property

Type: `"string"`. Optional.

NAPTR Record name, please provide only the specific subdomain or record name without the base
domain.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/naptr_record/values/): complete subsection reference.

## Next pages

- [primary.rr_set_group.rr_set.naptr_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/naptr_record/values/)
- [primary.rr_set_group.rr_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
