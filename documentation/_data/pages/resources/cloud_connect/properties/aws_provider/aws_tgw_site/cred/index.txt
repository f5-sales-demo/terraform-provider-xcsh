---
page_title: "aws_provider.aws_tgw_site.cred"
subcategory: ""
description: "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name."
xcsh_docs: {"aliases": ["aws provider aws tgw site cred"], "body_bytes": 6228, "body_sha256": "sha256:401696d28d6308b489d98d664c0d70ffbf452bbec15711f614f7d5ff724fff7c", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:cred", "parent_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site", "path": "documentation/resources/cloud_connect/properties/aws_provider/aws_tgw_site/cred/index.md", "product": "distributed-cloud", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3312332321213110-3213330110212003-3310122110111331-1231321322333232-3233202120123330-1133000031113111-1113123132210321-3223120103022131", "registry_path": "docs/guides/resources--cloud_connect--reference--group-001.md", "relationships": [{"anchor": "schema-aws_provider--aws_tgw_site--cred--name", "enforcement": "provider-schema", "group": "aws_provider.aws_tgw_site.cred:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:cred", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_provider", "aws_tgw_site", "cred"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-aws_provider--aws_tgw_site--cred--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:cred", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_tgw_site", "cred", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-aws_provider--aws_tgw_site--cred--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:cred", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_tgw_site", "cred", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["tenant"], "anchor": "schema-aws_provider--aws_tgw_site--cred--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:cred", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_tgw_site", "cred", "tenant"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_connect/properties/aws_provider/aws_tgw_site/cred/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_tgw_site.cred

Breadcrumbs:

- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/)
- [aws_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/)
- [aws_provider.aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/)
- aws_provider.aws_tgw_site.cred

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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
cred {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-aws_provider--aws_tgw_site--cred--name"></a>

### name property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="schema-aws_provider--aws_tgw_site--cred--namespace"></a>

### namespace property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="schema-aws_provider--aws_tgw_site--cred--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

## Next pages

- [aws_provider.aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/)
- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/)
