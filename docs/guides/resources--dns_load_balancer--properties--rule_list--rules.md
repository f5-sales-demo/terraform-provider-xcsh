---
page_title: "rule_list.rules"
subcategory: "DNS"
description: "rule_list.rules for xcsh_dns_load_balancer."
xcsh_docs: {"aliases": [], "body_bytes": 5973, "body_sha256": "sha256:d83bba05ba79a69be2b005b0cc72d58e4a1fbaf232d96ce496a4b8a97ecf4c82", "canonical_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules", "child_ids": ["xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:asn_list", "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:asn_matcher", "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:geo_location_label_selector", "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:geo_location_set", "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_list", "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:ip_prefix_set", "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules:pool"], "collection_id": "xcsh-docs:resources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list:rules", "parent_id": "xcsh-docs:resources:dns_load_balancer:properties:rule_list", "path": "docs/guides/resources--dns_load_balancer--properties--rule_list--rules.md", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_load_balancer/properties/rule_list/rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules for xcsh_dns_load_balancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md)
- [Property reference](resources--dns_load_balancer--reference.md)
- [rule_list](resources--dns_load_balancer--properties--rule_list.md)
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

- [asn_list](resources--dns_load_balancer--properties--rule_list--rules--asn_list.md): complete subsection reference.

- [asn_matcher](resources--dns_load_balancer--properties--rule_list--rules--asn_matcher.md): complete subsection reference.

- [geo_location_label_selector](resources--dns_load_balancer--properties--rule_list--rules--geo_location_label_selector.md): complete subsection reference.

- [geo_location_set](resources--dns_load_balancer--properties--rule_list--rules--geo_location_set.md): complete subsection reference.

- [ip_prefix_list](resources--dns_load_balancer--properties--rule_list--rules--ip_prefix_list.md): complete subsection reference.

- [ip_prefix_set](resources--dns_load_balancer--properties--rule_list--rules--ip_prefix_set.md): complete subsection reference.

- [pool](resources--dns_load_balancer--properties--rule_list--rules--pool.md): complete subsection reference.

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

- [rule_list.rules.asn_list](resources--dns_load_balancer--properties--rule_list--rules--asn_list.md)
- [rule_list.rules.asn_matcher](resources--dns_load_balancer--properties--rule_list--rules--asn_matcher.md)
- [rule_list.rules.geo_location_label_selector](resources--dns_load_balancer--properties--rule_list--rules--geo_location_label_selector.md)
- [rule_list.rules.geo_location_set](resources--dns_load_balancer--properties--rule_list--rules--geo_location_set.md)
- [rule_list.rules.ip_prefix_list](resources--dns_load_balancer--properties--rule_list--rules--ip_prefix_list.md)
- [rule_list.rules.ip_prefix_set](resources--dns_load_balancer--properties--rule_list--rules--ip_prefix_set.md)
- [rule_list.rules.pool](resources--dns_load_balancer--properties--rule_list--rules--pool.md)
- [rule_list](resources--dns_load_balancer--properties--rule_list.md)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md)
