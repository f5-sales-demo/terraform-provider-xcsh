---
page_title: "origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list"
subcategory: "Load Balancing"
description: "List of IP Prefix strings to match against."
xcsh_docs: {"aliases": ["origin server subset rule list origin server subset rules ip prefix list"], "body_bytes": 3511, "body_sha256": "sha256:6eeee8aa9d949f5fa65ae9d59fd40be6e18c76d1eb25579192066311fba8e800", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_prefix_list", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules", "path": "documentation/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/ip_prefix_list/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0111102320002302-0210101031332300-3113003023002203-3122211120031033-0021332113133230-0120303321101031-3133132212123213-3011103223302301", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-022.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "ip_prefix_list"], "schema_version": 1, "sections": [{"aliases": ["invert match"], "anchor": "schema-origin_server_subset_rule_list--origin_server_subset_rules--ip_prefix_list--invert_match", "description": "Invert the match result.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_prefix_list", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "ip_prefix_list", "invert_match"], "syntax": "attribute", "type": "bool"}, {"aliases": ["ip prefixes"], "anchor": "schema-origin_server_subset_rule_list--origin_server_subset_rules--ip_prefix_list--ip_prefixes", "description": "List of IPv4 prefix strings.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:origin_server_subset_rule_list:origin_server_subset_rules:ip_prefix_list", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_server_subset_rule_list", "origin_server_subset_rules", "ip_prefix_list", "ip_prefixes"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/ip_prefix_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of IP Prefix strings to match against.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [origin_server_subset_rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/)
- [origin_server_subset_rule_list.origin_server_subset_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/)
- origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

List of IP Prefix strings to match against.

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
ip_prefix_list {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-origin_server_subset_rule_list--origin_server_subset_rules--ip_prefix_list--invert_match"></a>

### invert_match property

Type: `"bool"`. Optional.

Invert Match Result. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="schema-origin_server_subset_rule_list--origin_server_subset_rules--ip_prefix_list--ip_prefixes"></a>

### ip_prefixes property

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
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
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [origin_server_subset_rule_list.origin_server_subset_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/origin_server_subset_rule_list/origin_server_subset_rules/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
