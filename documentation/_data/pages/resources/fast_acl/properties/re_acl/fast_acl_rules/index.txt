---
page_title: "re_acl.fast_acl_rules"
subcategory: ""
description: "Fast ACL rules to match."
xcsh_docs: {"aliases": ["re acl fast acl rules"], "body_bytes": 3541, "body_sha256": "sha256:8a3162c0f65da84a3532c161a1cc9055c3293d2006a08d27a9be6af9de9088a4", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action", "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:ip_prefix_set", "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:metadata", "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port", "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:prefix"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules", "parent_id": "xcsh-docs:resources:fast_acl:properties:re_acl", "path": "documentation/resources/fast_acl/properties/re_acl/fast_acl_rules/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1130303320221201-1011302330133031-3133313221213022-0321130210023003-1131202032111011-0010132331100100-1212311011321102-2021011103320010", "registry_path": "docs/guides/resources--fast_acl--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules:ConflictingListObjectAttributes:ip_prefix_set,prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules:ConflictingListObjectAttributes:ip_prefix_set,prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:prefix", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["re_acl", "fast_acl_rules"], "schema_version": 1, "sections": [{"aliases": ["action"], "anchor": "section", "description": "FastAclRuleAction specifies possible action to be applied on traffic, possible action include dropping, forwarding or ratelimiting the traffic.", "document_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-re_acl--fast_acl_rules--action--simple_action", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.action:ConflictingObjectAttributes:policer_action,simple_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action", "type": "conflicts"}, {"anchor": "schema-re_acl--fast_acl_rules--action--simple_action", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.action:ConflictingObjectAttributes:protocol_policer_action,simple_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.action:ConflictingObjectAttributes:policer_action,protocol_policer_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action:policer_action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.action:ConflictingObjectAttributes:policer_action,simple_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action:policer_action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.action:ConflictingObjectAttributes:policer_action,protocol_policer_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action:protocol_policer_action", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.action:ConflictingObjectAttributes:protocol_policer_action,simple_action", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:action:protocol_policer_action", "type": "conflicts"}], "schema_path": ["re_acl", "fast_acl_rules", "action"], "syntax": "block", "type": "object"}, {"aliases": ["ip prefix set"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:ip_prefix_set", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["re_acl", "fast_acl_rules", "ip_prefix_set"], "syntax": "block", "type": "object"}, {"aliases": ["metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:metadata", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-re_acl--fast_acl_rules--metadata--name", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:metadata", "type": "requires"}], "schema_path": ["re_acl", "fast_acl_rules", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["port"], "anchor": "section", "description": "L4 port numbers to match.", "document_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-re_acl--fast_acl_rules--port--user_defined", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.port:ConflictingListObjectAttributes:all,user_defined", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port", "type": "conflicts"}, {"anchor": "schema-re_acl--fast_acl_rules--port--user_defined", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.port:ConflictingListObjectAttributes:dns,user_defined", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.port:ConflictingListObjectAttributes:all,dns", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port:all", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.port:ConflictingListObjectAttributes:all,user_defined", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port:all", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.port:ConflictingListObjectAttributes:all,dns", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port:dns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules.port:ConflictingListObjectAttributes:dns,user_defined", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:port:dns", "type": "conflicts"}], "schema_path": ["re_acl", "fast_acl_rules", "port"], "syntax": "block", "type": "object"}, {"aliases": ["prefix"], "anchor": "section", "description": "List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain mix of both IPv4 and IPv6 prefixes.", "document_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:prefix", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["re_acl", "fast_acl_rules", "prefix"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/properties/re_acl/fast_acl_rules/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Fast ACL rules to match.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["fast_aclCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# re_acl.fast_acl_rules

Breadcrumbs:

- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/)
- [re_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/)
- re_acl.fast_acl_rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Rules. Fast ACL rules to match. Defaults to \`\[\]\`. Server applies default when omitted.

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

- [action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/action/): complete subsection reference.

- [ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/ip_prefix_set/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/metadata/): complete subsection reference.

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/port/): complete subsection reference.

- [prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/prefix/): complete subsection reference.

## Next pages

- [re_acl.fast_acl_rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/action/)
- [re_acl.fast_acl_rules.ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/ip_prefix_set/)
- [re_acl.fast_acl_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/metadata/)
- [re_acl.fast_acl_rules.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/port/)
- [re_acl.fast_acl_rules.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/prefix/)
- [re_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/)
- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/)
