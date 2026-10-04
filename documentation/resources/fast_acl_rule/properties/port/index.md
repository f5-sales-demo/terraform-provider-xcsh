---
page_title: "port"
subcategory: ""
description: "L4 port numbers to match."
xcsh_docs: {"aliases": ["port"], "body_bytes": 3494, "body_sha256": "sha256:0f230bf26f8f406357e76b38835c7c0ffd21829d9ee4e0e3efa24314dfa7e81e", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:fast_acl_rule:properties:port:all", "xcsh-docs:resources:fast_acl_rule:properties:port:dns"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl_rule:properties:port", "parent_id": "xcsh-docs:resources:fast_acl_rule:reference", "path": "documentation/resources/fast_acl_rule/properties/port/index.md", "product": "distributed-cloud", "provider_name": "fast_acl_rule", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1032003000303102-1033202210203210-3222020302212013-0321300133103321-3312031111132110-3202311220103302-2233113211303003-3331013000330130", "registry_path": "docs/guides/resources--fast_acl_rule--reference--group-001.md", "relationships": [{"anchor": "schema-port--user_defined", "enforcement": "provider-schema", "group": "port:ConflictingListObjectAttributes:all,user_defined", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl_rule:properties:port", "type": "conflicts"}, {"anchor": "schema-port--user_defined", "enforcement": "provider-schema", "group": "port:ConflictingListObjectAttributes:dns,user_defined", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl_rule:properties:port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "port:ConflictingListObjectAttributes:all,dns", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl_rule:properties:port:all", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "port:ConflictingListObjectAttributes:all,user_defined", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl_rule:properties:port:all", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "port:ConflictingListObjectAttributes:all,dns", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl_rule:properties:port:dns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "port:ConflictingListObjectAttributes:dns,user_defined", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl_rule:properties:port:dns", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["port"], "schema_version": 1, "sections": [{"aliases": ["port all"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:fast_acl_rule:properties:port:all", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["port", "all"], "syntax": "attribute", "type": "object"}, {"aliases": ["port dns"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:fast_acl_rule:properties:port:dns", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["port", "dns"], "syntax": "attribute", "type": "object"}, {"aliases": ["port user defined"], "anchor": "schema-port--user_defined", "description": "Exclusive with Matches the user defined port.", "document_id": "xcsh-docs:resources:fast_acl_rule:properties:port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["port", "user_defined"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl_rule/properties/port/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "L4 port numbers to match.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["fast_acl_ruleCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# port

Breadcrumbs:

- [xcsh_fast_acl_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/)
- port

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Source Ports. L4 port numbers to match.

Upstream description:

L4 port numbers to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("all",
    "dns"),
  validators.ConflictingListObjectAttributes("all",
    "user_defined"),
  validators.ConflictingListObjectAttributes("dns",
    "user_defined")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

## Direct properties

- [all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/port/all/): complete subsection reference.

- [dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/port/dns/): complete subsection reference.

<a id="schema-port--user_defined"></a>

### user_defined property

Type: `"number"`. Optional.

Exclusive with \[all DNS\] Matches the user defined port.

Upstream description:

Exclusive with \[all DNS\] Matches the user defined port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

## Next pages

- [port.all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/port/all/)
- [port.dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/port/dns/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/)
- [xcsh_fast_acl_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/)
