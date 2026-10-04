---
page_title: "site_acl.fast_acl_rules"
subcategory: ""
description: "Fast ACL rules to match."
xcsh_docs: {"aliases": ["site acl fast acl rules"], "body_bytes": 3524, "body_sha256": "sha256:71895d81a9189d17118c8e027debe4f1babcfa1200032de556292d658e93d47d", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:action", "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:ip_prefix_set", "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:metadata", "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:port", "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:prefix"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules", "parent_id": "xcsh-docs:resources:fast_acl:properties:site_acl", "path": "documentation/resources/fast_acl/properties/site_acl/fast_acl_rules/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2033232322232021-1031002112122002-3013233303210121-1003332031133103-3133310212321032-1313112311020212-0310321003122312-3330103231133021", "registry_path": "docs/guides/resources--fast_acl--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "site_acl.fast_acl_rules:ConflictingListObjectAttributes:ip_prefix_set,prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_acl.fast_acl_rules:ConflictingListObjectAttributes:ip_prefix_set,prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:prefix", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["site_acl", "fast_acl_rules"], "schema_version": 1, "sections": [{"aliases": ["site acl fast acl rules action"], "anchor": "section", "description": "FastAclRuleAction specifies possible action to be applied on traffic, possible action include dropping, forwarding or ratelimiting the traffic.", "document_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:action", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-site_acl--fast_acl_rules--action--simple_action", "enforcement": "provider-schema", "group": "site_acl.fast_acl_rules.action:ConflictingObjectAttributes:policer_action,simple_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:action", "type": "conflicts"}, {"anchor": "schema-site_acl--fast_acl_rules--action--simple_action", "enforcement": "provider-schema", "group": "site_acl.fast_acl_rules.action:ConflictingObjectAttributes:protocol_policer_action,simple_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_acl.fast_acl_rules.action:ConflictingObjectAttributes:policer_action,protocol_policer_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:action:policer_action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_acl.fast_acl_rules.action:ConflictingObjectAttributes:policer_action,simple_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:action:policer_action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_acl.fast_acl_rules.action:ConflictingObjectAttributes:policer_action,protocol_policer_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:action:protocol_policer_action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_acl.fast_acl_rules.action:ConflictingObjectAttributes:protocol_policer_action,simple_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:action:protocol_policer_action", "type": "conflicts"}], "schema_path": ["site_acl", "fast_acl_rules", "action"], "syntax": "block", "type": "object"}, {"aliases": ["site acl fast acl rules ip prefix set"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:ip_prefix_set", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["site_acl", "fast_acl_rules", "ip_prefix_set"], "syntax": "block", "type": "object"}, {"aliases": ["site acl fast acl rules metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:metadata", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-site_acl--fast_acl_rules--metadata--name", "enforcement": "provider-schema", "group": "site_acl.fast_acl_rules.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:metadata", "type": "requires"}], "schema_path": ["site_acl", "fast_acl_rules", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["site acl fast acl rules port"], "anchor": "section", "description": "L4 port numbers to match.", "document_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:port", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-site_acl--fast_acl_rules--port--user_defined", "enforcement": "provider-schema", "group": "site_acl.fast_acl_rules.port:ConflictingListObjectAttributes:all,user_defined", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:port", "type": "conflicts"}, {"anchor": "schema-site_acl--fast_acl_rules--port--user_defined", "enforcement": "provider-schema", "group": "site_acl.fast_acl_rules.port:ConflictingListObjectAttributes:dns,user_defined", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_acl.fast_acl_rules.port:ConflictingListObjectAttributes:all,dns", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:port:all", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_acl.fast_acl_rules.port:ConflictingListObjectAttributes:all,user_defined", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:port:all", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_acl.fast_acl_rules.port:ConflictingListObjectAttributes:all,dns", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:port:dns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_acl.fast_acl_rules.port:ConflictingListObjectAttributes:dns,user_defined", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:port:dns", "type": "conflicts"}], "schema_path": ["site_acl", "fast_acl_rules", "port"], "syntax": "block", "type": "object"}, {"aliases": ["site acl fast acl rules prefix"], "anchor": "section", "description": "List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain mix of both IPv4 and IPv6 prefixes.", "document_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:prefix", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["site_acl", "fast_acl_rules", "prefix"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/properties/site_acl/fast_acl_rules/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Fast ACL rules to match.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["fast_aclCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_acl.fast_acl_rules

Breadcrumbs:

- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/)
- [site_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/)
- site_acl.fast_acl_rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Rules. Fast ACL rules to match.

Upstream description:

Fast ACL rules to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("ip_prefix_set",
    "prefix")}
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
fast_acl_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/action/): complete subsection reference.

- [ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/ip_prefix_set/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/metadata/): complete subsection reference.

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/port/): complete subsection reference.

- [prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/prefix/): complete subsection reference.

## Next pages

- [site_acl.fast_acl_rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/action/)
- [site_acl.fast_acl_rules.ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/ip_prefix_set/)
- [site_acl.fast_acl_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/metadata/)
- [site_acl.fast_acl_rules.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/port/)
- [site_acl.fast_acl_rules.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/prefix/)
- [site_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/)
- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/)
