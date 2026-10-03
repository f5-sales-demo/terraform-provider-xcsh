---
page_title: "cloudflare.js_insertion_rules.exclude_list.metadata"
subcategory: ""
description: "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs."
xcsh_docs: {"aliases": ["cloudflare js insertion rules exclude list metadata"], "body_bytes": 4608, "body_sha256": "sha256:1cd26d7e35c0216ad73d5db7b1309b8f6e5b6607b11a43b294d71b7690156d95", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:metadata", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list", "path": "documentation/resources/protected_application/properties/cloudflare/js_insertion_rules/exclude_list/metadata/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2331301132110211-1013333210130302-3113310213010322-2200211101212021-2231331013311230-2003020110303201-3133210302223121-2000320220303022", "registry_path": "docs/guides/resources--protected_application--reference--group-001.md", "relationships": [{"anchor": "schema-cloudflare--js_insertion_rules--exclude_list--metadata--name", "enforcement": "provider-schema", "group": "cloudflare.js_insertion_rules.exclude_list.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:metadata", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudflare", "js_insertion_rules", "exclude_list", "metadata"], "schema_version": 1, "sections": [{"aliases": ["cloudflare js insertion rules exclude list metadata description spec"], "anchor": "schema-cloudflare--js_insertion_rules--exclude_list--metadata--description_spec", "description": "Description. Human readable description.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:metadata", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "js_insertion_rules", "exclude_list", "metadata", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["cloudflare js insertion rules exclude list metadata name"], "anchor": "schema-cloudflare--js_insertion_rules--exclude_list--metadata--name", "description": "This is the name of the message. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:metadata", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudflare", "js_insertion_rules", "exclude_list", "metadata", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudflare/js_insertion_rules/exclude_list/metadata/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.js_insertion_rules.exclude_list.metadata

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudflare](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/)
- [cloudflare.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/js_insertion_rules/)
- [cloudflare.js_insertion_rules.exclude_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/js_insertion_rules/exclude_list/)
- cloudflare.js_insertion_rules.exclude_list.metadata

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
metadata {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-cloudflare--js_insertion_rules--exclude_list--metadata--description_spec"></a>

### description_spec property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="schema-cloudflare--js_insertion_rules--exclude_list--metadata--name"></a>

### name property

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

## Next pages

- [cloudflare.js_insertion_rules.exclude_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudflare/js_insertion_rules/exclude_list/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
