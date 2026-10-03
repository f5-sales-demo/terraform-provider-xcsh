---
page_title: "rule_list.rules"
subcategory: "DNS"
description: "Rules to perform load balancing."
xcsh_docs: {"aliases": ["rule list rules"], "body_bytes": 6916, "body_sha256": "sha256:bbff3c84f2d10b631f018131a7babfe2806b5d940a3adfdbd8fc1ef8d03e779c", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:asn_list", "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:asn_matcher", "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:geo_location_label_selector", "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:geo_location_set", "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_list", "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_set", "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:pool"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules", "parent_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list", "path": "documentation/resources/dns_load_balancer/properties/rule_list/rules/index.md", "product": "distributed-cloud", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2321132022133030-1332131221201231-2113100122130112-2110030202103300-2121322011013302-0211112102122231-1122101301213330-2223002010100121", "registry_path": "docs/guides/resources--dns_load_balancer--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:asn_list,asn_matcher", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:asn_list,geo_location_label_selector", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:asn_list,geo_location_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:asn_list,ip_prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:asn_list,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:asn_list,asn_matcher", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:asn_matcher,geo_location_label_selector", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:asn_matcher,geo_location_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:asn_matcher,ip_prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:asn_matcher,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:asn_list,geo_location_label_selector", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:geo_location_label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:asn_matcher,geo_location_label_selector", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:geo_location_label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:geo_location_label_selector,geo_location_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:geo_location_label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:geo_location_label_selector,ip_prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:geo_location_label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:geo_location_label_selector,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:geo_location_label_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:asn_list,geo_location_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:geo_location_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:asn_matcher,geo_location_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:geo_location_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:geo_location_label_selector,geo_location_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:geo_location_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:geo_location_set,ip_prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:geo_location_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:geo_location_set,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:geo_location_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:asn_list,ip_prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:asn_matcher,ip_prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:geo_location_label_selector,ip_prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:geo_location_set,ip_prefix_list", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:ip_prefix_list,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:asn_list,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:asn_matcher,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:geo_location_label_selector,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:geo_location_set,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules:ConflictingListObjectAttributes:ip_prefix_list,ip_prefix_set", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "schema-rule_list--rules--score", "enforcement": "provider-schema", "group": "rule_list.rules:RequiredListObjectAttributes:score", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules"], "schema_version": 1, "sections": [{"aliases": ["rule list rules asn list"], "anchor": "section", "description": "An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists for use in network policy or service policy. It can be used to create the allow list only for DNS Load Balancer.", "document_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:asn_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rule_list--rules--asn_list--as_numbers", "enforcement": "provider-schema", "group": "rule_list.rules.asn_list:RequiredObjectAttributes:as_numbers", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:asn_list", "type": "requires"}], "schema_path": ["rule_list", "rules", "asn_list"], "syntax": "block", "type": "object"}, {"aliases": ["rule list rules asn matcher"], "anchor": "section", "description": "Match any AS number contained in the list of bgp_asn_sets.", "document_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:asn_matcher", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.asn_matcher:RequiredObjectAttributes:asn_sets", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:asn_matcher:asn_sets", "type": "requires"}], "schema_path": ["rule_list", "rules", "asn_matcher"], "syntax": "block", "type": "object"}, {"aliases": ["rule list rules geo location label selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:geo_location_label_selector", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rule_list--rules--geo_location_label_selector--expressions", "enforcement": "provider-schema", "group": "rule_list.rules.geo_location_label_selector:RequiredObjectAttributes:expressions", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:geo_location_label_selector", "type": "requires"}], "schema_path": ["rule_list", "rules", "geo_location_label_selector"], "syntax": "block", "type": "object"}, {"aliases": ["rule list rules geo location set"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:geo_location_set", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rule_list--rules--geo_location_set--name", "enforcement": "provider-schema", "group": "rule_list.rules.geo_location_set:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:geo_location_set", "type": "requires"}], "schema_path": ["rule_list", "rules", "geo_location_set"], "syntax": "block", "type": "object"}, {"aliases": ["rule list rules ip prefix list"], "anchor": "section", "description": "List of IP Prefix strings to match against.", "document_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list", "rules", "ip_prefix_list"], "syntax": "block", "type": "object"}, {"aliases": ["rule list rules ip prefix set"], "anchor": "section", "description": "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true.", "document_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_set", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.ip_prefix_set:RequiredObjectAttributes:prefix_sets", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_set:prefix_sets", "type": "requires"}], "schema_path": ["rule_list", "rules", "ip_prefix_set"], "syntax": "block", "type": "object"}, {"aliases": ["rule list rules pool"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:pool", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rule_list--rules--pool--name", "enforcement": "provider-schema", "group": "rule_list.rules.pool:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:pool", "type": "requires"}], "schema_path": ["rule_list", "rules", "pool"], "syntax": "block", "type": "object"}, {"aliases": ["rule list rules score"], "anchor": "schema-rule_list--rules--score", "description": "When multiple load balancing rules match a query, the one with the highest score is chosen.", "document_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "score"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_load_balancer/properties/rule_list/rules/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Rules to perform load balancing.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules

Breadcrumbs:

- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/)
- rule_list.rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Load Balancing Rules. Rules to perform load balancing.

Upstream description:

Rules to perform load balancing.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("score"),
  validators.ConflictingListObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingListObjectAttributes("asn_list",
    "geo_location_label_selector"),
  validators.ConflictingListObjectAttributes("asn_list",
    "geo_location_set"),
  validators.ConflictingListObjectAttributes("asn_list",
    "ip_prefix_list"),
  validators.ConflictingListObjectAttributes("asn_list",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("asn_matcher",
    "geo_location_label_selector"),
  validators.ConflictingListObjectAttributes("asn_matcher",
    "geo_location_set"),
  validators.ConflictingListObjectAttributes("asn_matcher",
    "ip_prefix_list"),
  validators.ConflictingListObjectAttributes("asn_matcher",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("geo_location_label_selector",
    "geo_location_set"),
  validators.ConflictingListObjectAttributes("geo_location_label_selector",
    "ip_prefix_list"),
  validators.ConflictingListObjectAttributes("geo_location_label_selector",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("geo_location_set",
    "ip_prefix_list"),
  validators.ConflictingListObjectAttributes("geo_location_set",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("ip_prefix_list",
    "ip_prefix_set")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/asn_list/): complete subsection reference.

- [asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/asn_matcher/): complete subsection reference.

- [geo_location_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/geo_location_label_selector/): complete subsection reference.

- [geo_location_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/geo_location_set/): complete subsection reference.

- [ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/ip_prefix_list/): complete subsection reference.

- [ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/ip_prefix_set/): complete subsection reference.

- [pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/pool/): complete subsection reference.

<a id="schema-rule_list--rules--score"></a>

### score property

Type: `"number"`. Optional.

When multiple load balancing rules match a query, the one with the highest score is chosen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 32767),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32767,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32767"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32767"
  }
}
```

## Next pages

- [rule_list.rules.asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/asn_list/)
- [rule_list.rules.asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/asn_matcher/)
- [rule_list.rules.geo_location_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/geo_location_label_selector/)
- [rule_list.rules.geo_location_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/geo_location_set/)
- [rule_list.rules.ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/ip_prefix_list/)
- [rule_list.rules.ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/ip_prefix_set/)
- [rule_list.rules.pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/rules/pool/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/properties/rule_list/)
- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/)
