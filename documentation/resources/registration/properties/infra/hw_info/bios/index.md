---
page_title: "infra.hw_info.bios"
subcategory: ""
description: "BIOS information."
xcsh_docs: {"aliases": ["infra hw info bios"], "body_bytes": 3430, "body_sha256": "sha256:a6762c1a9fde39cc2996171a63ced16c4e034263ab921f47832f585608dcc5ac", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration:properties:infra:hw_info:bios", "parent_id": "xcsh-docs:resources:registration:properties:infra:hw_info", "path": "documentation/resources/registration/properties/infra/hw_info/bios/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3230213321222323-1333012313102022-1113123132022220-1010222000301033-2333020321320130-0121011220223330-3213220332222320-0203330323103323", "registry_path": "docs/guides/resources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "hw_info", "bios"], "schema_version": 1, "sections": [{"aliases": ["date"], "anchor": "schema-infra--hw_info--bios--date", "description": "Information from /sys/class/dmi/ID/bios_date.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:bios", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "bios", "date"], "syntax": "attribute", "type": "string"}, {"aliases": ["vendor"], "anchor": "schema-infra--hw_info--bios--vendor", "description": "Information from /sys/class/dmi/ID/bios_vendor.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:bios", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "bios", "vendor"], "syntax": "attribute", "type": "string"}, {"aliases": ["version"], "anchor": "schema-infra--hw_info--bios--version", "description": "Information from /sys/class/dmi/ID/bios_version.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info:bios", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hw_info", "bios", "version"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/properties/infra/hw_info/bios/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "BIOS information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.hw_info.bios

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/)
- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/)
- [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/)
- infra.hw_info.bios

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Bios Data. BIOS information.

Upstream description:

BIOS information.

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
bios {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-infra--hw_info--bios--date"></a>

### date property

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/bios\_date.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(10, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "temporal",
    "constraintType": "string",
    "deterministic": true,
    "format": "date",
    "formatDescription": "ISO 8601 date (e.g., 2026-01-19)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 10,
    "pattern": "^\\d{4}-\\d{2}-\\d{2}$",
    "validation": {
      "standard": "ISO 8601"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--hw_info--bios--vendor"></a>

### vendor property

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/bios\_vendor.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--hw_info--bios--version"></a>

### version property

Type: `"string"`. Optional.

Information from /sys/class/dmi/ID/bios\_version.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/)
- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
