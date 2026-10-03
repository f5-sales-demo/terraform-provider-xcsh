---
page_title: "re_acl.fast_acl_rules.port"
subcategory: ""
description: "L4 port numbers to match."
xcsh_docs: {"aliases": ["re acl fast acl rules port"], "body_bytes": 3809, "body_sha256": "sha256:59895f01cc937e2d0e843b3521b9dfb3aa023de649e6f5831f04177d9f7dd3de", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port:all", "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port:dns"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port", "parent_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules", "path": "documentation/resources/fast_acl/properties/re_acl/fast_acl_rules/port/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0100122032000213-2012010313310033-1011023301101021-2123211122003310-0322300001013201-2323000322233123-1313222213331131-2232231202013120", "registry_path": "docs/guides/resources--fast_acl--reference--group-001.md", "relationships": [{"anchor": "schema-re_acl--fast_acl_rules--port--user_defined", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.port:ConflictingListObjectAttributes:all,user_defined", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port", "type": "conflicts"}, {"anchor": "schema-re_acl--fast_acl_rules--port--user_defined", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.port:ConflictingListObjectAttributes:dns,user_defined", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.port:ConflictingListObjectAttributes:all,dns", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port:all", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.port:ConflictingListObjectAttributes:all,user_defined", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port:all", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.port:ConflictingListObjectAttributes:all,dns", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port:dns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.port:ConflictingListObjectAttributes:dns,user_defined", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port:dns", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["re_acl", "fast_acl_rules", "port"], "schema_version": 1, "sections": [{"aliases": ["re acl fast acl rules port all"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port:all", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_acl", "fast_acl_rules", "port", "all"], "syntax": "attribute", "type": "object"}, {"aliases": ["re acl fast acl rules port dns"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port:dns", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_acl", "fast_acl_rules", "port", "dns"], "syntax": "attribute", "type": "object"}, {"aliases": ["re acl fast acl rules port user defined"], "anchor": "schema-re_acl--fast_acl_rules--port--user_defined", "description": "Exclusive with Matches the user defined port.", "document_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_acl", "fast_acl_rules", "port", "user_defined"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/properties/re_acl/fast_acl_rules/port/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "L4 port numbers to match.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["fast_aclCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# re_acl.fast_acl_rules.port

Breadcrumbs:

- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/)
- [re_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/)
- [re_acl.fast_acl_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/)
- re_acl.fast_acl_rules.port

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

- [all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/port/all/): complete subsection reference.

- [dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/port/dns/): complete subsection reference.

<a id="schema-re_acl--fast_acl_rules--port--user_defined"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [re_acl.fast_acl_rules.port.all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/port/all/)
- [re_acl.fast_acl_rules.port.dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/port/dns/)
- [re_acl.fast_acl_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/)
- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/)
