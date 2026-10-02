---
page_title: "f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site"
subcategory: ""
description: "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name."
xcsh_docs: {"aliases": ["f5 big ip aws service aws tgw site params aws tgw site"], "body_bytes": 6422, "body_sha256": "sha256:c739d7bb92bb890288acf27560922b517a6a941540397a54a05e16ebc6693f09", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:aws_tgw_site_params:aws_tgw_site", "parent_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:aws_tgw_site_params", "path": "documentation/resources/nfv_service/properties/f5_big_ip_aws_service/aws_tgw_site_params/aws_tgw_site/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1300102203120232-0232232122011313-1301020010010310-2332200230120321-1331000100232212-0031313231033012-2221132102223233-3313110210032323", "registry_path": "docs/guides/resources--nfv_service--reference--group-001.md", "relationships": [{"anchor": "schema-f5_big_ip_aws_service--aws_tgw_site_params--aws_tgw_site--name", "enforcement": "provider-schema", "group": "f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:aws_tgw_site_params:aws_tgw_site", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["f5_big_ip_aws_service", "aws_tgw_site_params", "aws_tgw_site"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-f5_big_ip_aws_service--aws_tgw_site_params--aws_tgw_site--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:aws_tgw_site_params:aws_tgw_site", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "aws_tgw_site_params", "aws_tgw_site", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-f5_big_ip_aws_service--aws_tgw_site_params--aws_tgw_site--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:aws_tgw_site_params:aws_tgw_site", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "aws_tgw_site_params", "aws_tgw_site", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["tenant"], "anchor": "schema-f5_big_ip_aws_service--aws_tgw_site_params--aws_tgw_site--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:resources:nfv_service:properties:f5_big_ip_aws_service:aws_tgw_site_params:aws_tgw_site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "aws_tgw_site_params", "aws_tgw_site", "tenant"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/f5_big_ip_aws_service/aws_tgw_site_params/aws_tgw_site/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [f5_big_ip_aws_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/)
- [f5_big_ip_aws_service.aws_tgw_site_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/aws_tgw_site_params/)
- f5_big_ip_aws_service.aws_tgw_site_params.aws_tgw_site

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
aws_tgw_site {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-f5_big_ip_aws_service--aws_tgw_site_params--aws_tgw_site--name"></a>

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

<a id="schema-f5_big_ip_aws_service--aws_tgw_site_params--aws_tgw_site--namespace"></a>

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

<a id="schema-f5_big_ip_aws_service--aws_tgw_site_params--aws_tgw_site--tenant"></a>

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

- [f5_big_ip_aws_service.aws_tgw_site_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/f5_big_ip_aws_service/aws_tgw_site_params/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
