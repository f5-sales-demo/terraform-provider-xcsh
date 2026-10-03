---
page_title: "brute_force_detection"
subcategory: ""
description: "Configuration parameter for brute force detection."
xcsh_docs: {"aliases": ["brute force detection"], "body_bytes": 2251, "body_sha256": "sha256:c59116569efc55717004100ef1f5982e69962b8f8015018b35b07813862dc821", "capabilities": ["administration"], "category": "administration", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:resources:tenant_configuration:properties:brute_force_detection", "parent_id": "xcsh-docs:resources:tenant_configuration:reference", "path": "documentation/resources/tenant_configuration/properties/brute_force_detection/index.md", "product": "distributed-cloud", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2002011201213020-3120102323023032-3331212102222001-0213001223111132-0330203000010333-3320213330231102-1100132103102210-3220032120100020", "registry_path": "docs/guides/resources--tenant_configuration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["brute_force_detection"], "schema_version": 1, "sections": [{"aliases": ["brute force detection max login failures", "login", "login result", "sign in"], "anchor": "schema-brute_force_detection--max_login_failures", "description": "How many failures before wait is triggered. When login failure count is hit, user will be temporarily locked for a max duration of 15 minutes.", "document_id": "xcsh-docs:resources:tenant_configuration:properties:brute_force_detection", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["brute_force_detection", "max_login_failures"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tenant_configuration/properties/brute_force_detection/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for brute force detection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# brute_force_detection

Breadcrumbs:

- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/)
- brute_force_detection

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for brute force detection.

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
brute_force_detection {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-brute_force_detection--max_login_failures"></a>

### max_login_failures property

Type: `"number"`. Optional.

How many failures before wait is triggered. When login failure count is hit, user will be
temporarily locked for a max duration of 15 minutes.

Upstream description:

How many failures before wait is triggered. When login failure count is hit, user will be
temporarily locked for a max duration of 15 minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(30),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/)
- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/)
