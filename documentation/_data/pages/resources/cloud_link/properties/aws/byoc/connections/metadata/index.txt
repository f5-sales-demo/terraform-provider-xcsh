---
page_title: "aws.byoc.connections.metadata"
subcategory: ""
description: "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs."
xcsh_docs: {"aliases": ["aws byoc connections metadata"], "body_bytes": 4277, "body_sha256": "sha256:464cb5652360f54c6affe34ab025105fb88e70edbb448570a3b1d6ed9f16cdb0", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections:metadata", "parent_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections", "path": "documentation/resources/cloud_link/properties/aws/byoc/connections/metadata/index.md", "product": "distributed-cloud", "provider_name": "cloud_link", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3101033123112001-2031201223322300-1122112301102320-1220233032023011-0133301031211120-0120122122331011-1012000112000221-0200101110202333", "registry_path": "docs/guides/resources--cloud_link--reference--group-001.md", "relationships": [{"anchor": "schema-aws--byoc--connections--metadata--name", "enforcement": "provider-schema", "group": "aws.byoc.connections.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections:metadata", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aws", "byoc", "connections", "metadata"], "schema_version": 1, "sections": [{"aliases": ["aws byoc connections metadata description spec"], "anchor": "schema-aws--byoc--connections--metadata--description_spec", "description": "Description. Human readable description.", "document_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections:metadata", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws", "byoc", "connections", "metadata", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws byoc connections metadata name"], "anchor": "schema-aws--byoc--connections--metadata--name", "description": "This is the name of the message. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:cloud_link:properties:aws:byoc:connections:metadata", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws", "byoc", "connections", "metadata", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_link/properties/aws/byoc/connections/metadata/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws.byoc.connections.metadata

Breadcrumbs:

- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/)
- [aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/)
- [aws.byoc](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/)
- [aws.byoc.connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/)
- aws.byoc.connections.metadata

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

<a id="schema-aws--byoc--connections--metadata--description_spec"></a>

### description_spec property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="schema-aws--byoc--connections--metadata--name"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [aws.byoc.connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/aws/byoc/connections/)
- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/)
