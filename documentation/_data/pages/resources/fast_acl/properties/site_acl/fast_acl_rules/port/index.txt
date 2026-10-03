---
page_title: "site_acl.fast_acl_rules.port"
subcategory: ""
description: "L4 port numbers to match."
xcsh_docs: {"aliases": ["site acl fast acl rules port"], "body_bytes": 3839, "body_sha256": "sha256:25bbebb94d65865fd9c3fb791d87752770f3b0a37e60f67810841bc108d9e300", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:port:all", "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:port:dns"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:port", "parent_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules", "path": "documentation/resources/fast_acl/properties/site_acl/fast_acl_rules/port/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2111312000322020-3101211201200202-2303133201012311-0031111313033003-1130300321231133-2313013031123003-2300132113301310-1033023230021002", "registry_path": "docs/guides/resources--fast_acl--reference--group-001.md", "relationships": [{"anchor": "schema-site_acl--fast_acl_rules--port--user_defined", "enforcement": "provider-schema", "group": "site_acl.fast_acl_rules.port:ConflictingListObjectAttributes:all,user_defined", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:port", "type": "conflicts"}, {"anchor": "schema-site_acl--fast_acl_rules--port--user_defined", "enforcement": "provider-schema", "group": "site_acl.fast_acl_rules.port:ConflictingListObjectAttributes:dns,user_defined", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_acl.fast_acl_rules.port:ConflictingListObjectAttributes:all,dns", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:port:all", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_acl.fast_acl_rules.port:ConflictingListObjectAttributes:all,user_defined", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:port:all", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_acl.fast_acl_rules.port:ConflictingListObjectAttributes:all,dns", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:port:dns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_acl.fast_acl_rules.port:ConflictingListObjectAttributes:dns,user_defined", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:port:dns", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["site_acl", "fast_acl_rules", "port"], "schema_version": 1, "sections": [{"aliases": ["site acl fast acl rules port all"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:port:all", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_acl", "fast_acl_rules", "port", "all"], "syntax": "attribute", "type": "object"}, {"aliases": ["site acl fast acl rules port dns"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:port:dns", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_acl", "fast_acl_rules", "port", "dns"], "syntax": "attribute", "type": "object"}, {"aliases": ["site acl fast acl rules port user defined"], "anchor": "schema-site_acl--fast_acl_rules--port--user_defined", "description": "Exclusive with Matches the user defined port.", "document_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_acl", "fast_acl_rules", "port", "user_defined"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/properties/site_acl/fast_acl_rules/port/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "L4 port numbers to match.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["fast_aclCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_acl.fast_acl_rules.port

Breadcrumbs:

- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/)
- [site_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/)
- [site_acl.fast_acl_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/)
- site_acl.fast_acl_rules.port

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
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

- [all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/port/all/): complete subsection reference.

- [dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/port/dns/): complete subsection reference.

<a id="schema-site_acl--fast_acl_rules--port--user_defined"></a>

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

- [site_acl.fast_acl_rules.port.all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/port/all/)
- [site_acl.fast_acl_rules.port.dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/port/dns/)
- [site_acl.fast_acl_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/)
- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/)
