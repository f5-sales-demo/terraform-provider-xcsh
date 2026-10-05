---
page_title: "primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint"
subcategory: "DNS"
description: "Configuration parameter for sha256 fingerprint."
xcsh_docs: {"aliases": ["primary rr set group rr set sshfp record values sha256 fingerprint"], "body_bytes": 3487, "body_sha256": "sha256:90854001b312c1a13f3d744d218276d4543f63a703da15d7422795e0b946c810", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:sshfp_record:values:sha256_fingerprint", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:sshfp_record:values", "path": "documentation/resources/dns_zone/properties/primary/rr_set_group/rr_set/sshfp_record/values/sha256_fingerprint/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3103002202322302-2232010213112211-1102001223032112-3101301022320121-0102223210112131-3030021311311121-1133202232131312-2302110021303323", "registry_path": "docs/guides/resources--dns_zone--reference--group-003.md", "relationships": [{"anchor": "schema-primary--rr_set_group--rr_set--sshfp_record--values--sha256_fingerprint--fingerprint", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint:RequiredObjectAttributes:fingerprint", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:sshfp_record:values:sha256_fingerprint", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "rr_set_group", "rr_set", "sshfp_record", "values", "sha256_fingerprint"], "schema_version": 1, "sections": [{"aliases": ["primary rr set group rr set sshfp record values sha256 fingerprint fingerprint"], "anchor": "schema-primary--rr_set_group--rr_set--sshfp_record--values--sha256_fingerprint--fingerprint", "description": "The 'fingerprint' is the DS key and the actual contents of the DS record.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:sshfp_record:values:sha256_fingerprint", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "rr_set_group", "rr_set", "sshfp_record", "values", "sha256_fingerprint", "fingerprint"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/rr_set_group/rr_set/sshfp_record/values/sha256_fingerprint/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Configuration parameter for sha256 fingerprint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/)
- [primary.rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/)
- [primary.rr_set_group.rr_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/)
- [primary.rr_set_group.rr_set.sshfp_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/sshfp_record/)
- [primary.rr_set_group.rr_set.sshfp_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/sshfp_record/values/)
- primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha256 fingerprint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("fingerprint")}
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
sha256_fingerprint {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-primary--rr_set_group--rr_set--sshfp_record--values--sha256_fingerprint--fingerprint"></a>

### fingerprint property

Type: `"string"`. Optional.

The 'fingerprint' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(64, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[0-9a-fA-F]",
      "description": "Hexadecimal characters only"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 64,
    "pattern": "^[0-9a-fA-F]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

## Next pages

- [primary.rr_set_group.rr_set.sshfp_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/sshfp_record/values/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
