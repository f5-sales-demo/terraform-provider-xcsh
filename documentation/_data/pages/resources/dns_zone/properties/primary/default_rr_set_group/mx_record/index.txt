---
page_title: "primary.default_rr_set_group.mx_record"
subcategory: "DNS"
description: "DNSMXResourceRecord."
xcsh_docs: {"aliases": ["primary default rr set group mx record"], "body_bytes": 3479, "body_sha256": "sha256:a21a765e2ef1f9606ae825097ee207b9d33f11dce3a26a5df890f379fc2cbbd7", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:mx_record:values"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:mx_record", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group", "path": "documentation/resources/dns_zone/properties/primary/default_rr_set_group/mx_record/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1002110103301213-0230131110201331-3223012003111321-0300033220002210-3000130132020202-0011011200330101-0120220030322210-3213201300203312", "registry_path": "docs/guides/resources--dns_zone--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.mx_record:RequiredObjectAttributes:values", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:mx_record:values", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "default_rr_set_group", "mx_record"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-primary--default_rr_set_group--mx_record--name", "description": "MX Record name, please provide only the specific subdomain or record name without the base domain.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:mx_record", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "mx_record", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["values"], "anchor": "section", "description": "Configuration parameter for values", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:mx_record:values", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["primary", "default_rr_set_group", "mx_record", "values"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/default_rr_set_group/mx_record/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "DNSMXResourceRecord.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.default_rr_set_group.mx_record

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/)
- [primary.default_rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/)
- primary.default_rr_set_group.mx_record

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

DNSMXResourceRecord.

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
mx_record {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-primary--default_rr_set_group--mx_record--name"></a>

### name property

Type: `"string"`. Optional.

MX Record name, please provide only the specific subdomain or record name without the base domain.

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
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

- [values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/mx_record/values/): complete subsection reference.

## Next pages

- [primary.default_rr_set_group.mx_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/mx_record/values/)
- [primary.default_rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
