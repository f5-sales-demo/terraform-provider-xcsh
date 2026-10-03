---
page_title: "site_acl.fast_acl_rules"
subcategory: ""
description: "Fast ACL rules to match."
xcsh_docs: {"aliases": ["site acl fast acl rules"], "body_bytes": 3275, "body_sha256": "sha256:6c7b7dba28ad721a0244f72a89682b4cc0e6a09de4e004a53d0d59b111e26427", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:action", "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:ip_prefix_set", "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:metadata", "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:port", "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:prefix"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules", "parent_id": "xcsh-docs:data-sources:fast_acl:properties:site_acl", "path": "documentation/data-sources/fast_acl/properties/site_acl/fast_acl_rules/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0033210120011020-2212331003113120-0231233101332301-3330033303222300-0232321100222210-2233201323133113-0131020302232121-1202112122300023", "registry_path": "docs/guides/data-sources--fast_acl--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["site_acl", "fast_acl_rules"], "schema_version": 1, "sections": [{"aliases": ["site acl fast acl rules action"], "anchor": "section", "description": "FastAclRuleAction specifies possible action to be applied on traffic, possible action include dropping, forwarding or ratelimiting the traffic.", "document_id": "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:action", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["site_acl", "fast_acl_rules", "action"], "syntax": "attribute", "type": "object"}, {"aliases": ["site acl fast acl rules ip prefix set"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:ip_prefix_set", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["site_acl", "fast_acl_rules", "ip_prefix_set"], "syntax": "attribute", "type": "object"}, {"aliases": ["site acl fast acl rules metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["site_acl", "fast_acl_rules", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["site acl fast acl rules port"], "anchor": "section", "description": "L4 port numbers to match.", "document_id": "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["site_acl", "fast_acl_rules", "port"], "syntax": "attribute", "type": "object"}, {"aliases": ["site acl fast acl rules prefix"], "anchor": "section", "description": "List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain mix of both IPv4 and IPv6 prefixes.", "document_id": "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:prefix", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["site_acl", "fast_acl_rules", "prefix"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl/properties/site_acl/fast_acl_rules/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Fast ACL rules to match.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["fast_aclCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_acl.fast_acl_rules

Breadcrumbs:

- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/)
- [site_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/)
- site_acl.fast_acl_rules

<a id="section"></a>

Type: `"list"`. Computed.

Rules. Fast ACL rules to match.

Upstream description:

Fast ACL rules to match.

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

## Direct properties

- [action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/fast_acl_rules/action/): complete subsection reference.

- [ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/fast_acl_rules/ip_prefix_set/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/fast_acl_rules/metadata/): complete subsection reference.

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/fast_acl_rules/port/): complete subsection reference.

- [prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/fast_acl_rules/prefix/): complete subsection reference.

## Next pages

- [site_acl.fast_acl_rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/fast_acl_rules/action/)
- [site_acl.fast_acl_rules.ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/fast_acl_rules/ip_prefix_set/)
- [site_acl.fast_acl_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/fast_acl_rules/metadata/)
- [site_acl.fast_acl_rules.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/fast_acl_rules/port/)
- [site_acl.fast_acl_rules.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/fast_acl_rules/prefix/)
- [site_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/site_acl/)
- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/)
