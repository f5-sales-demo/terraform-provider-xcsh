---
page_title: "Property reference"
subcategory: "DNS"
description: "Property reference for xcsh_dns_load_balancer."
xcsh_docs: {"aliases": ["dns load balancer"], "body_bytes": 19604, "body_sha256": "sha256:a3a1ba01538dc2ecf03da3626d6a630aeea15db1b6003f5cf0993e5a2b4f193b", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_load_balancer:properties:fallback_pool", "xcsh-docs:data-sources:dns_load_balancer:properties:response_cache", "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_load_balancer:reference", "parent_id": "xcsh-docs:data-sources:dns_load_balancer:fundamentals", "path": "documentation/data-sources/dns_load_balancer/properties/index.md", "product": "distributed-cloud", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222", "registry_path": "docs/guides/data-sources--dns_load_balancer--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:dns_load_balancer:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:dns_load_balancer:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["fallback pool"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:dns_load_balancer:properties:fallback_pool", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["fallback_pool"], "syntax": "attribute", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:dns_load_balancer:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:dns_load_balancer:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:dns_load_balancer:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:dns_load_balancer:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["record type"], "anchor": "schema-record_type", "description": "Resource Record Type - A: A - AAAA: AAAA - MX: MX - CNAME: CNAME - SRV: SRV.", "document_id": "xcsh-docs:data-sources:dns_load_balancer:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["record_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["response cache"], "anchor": "section", "description": "Response Cache x-required.", "document_id": "xcsh-docs:data-sources:dns_load_balancer:properties:response_cache", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["response_cache"], "syntax": "attribute", "type": "object"}, {"aliases": ["rule list"], "anchor": "section", "description": "List of the Load Balancing Rules.", "document_id": "xcsh-docs:data-sources:dns_load_balancer:properties:rule_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rule_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_load_balancer/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_dns_load_balancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the DNSLoadBalancer.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

- [fallback_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/fallback_pool/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the DNSLoadBalancer.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Optional, Computed.

Namespace where the DNSLoadBalancer exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-record_type"></a>

### record_type property

Type: `"string"`. Computed.

\[Enum: A|AAAA|MX|CNAME|SRV\] Resource Record Type - A: A - AAAA: AAAA - MX: MX - CNAME: CNAME -
SRV: SRV. Possible values are \`A\`, \`AAAA\`, \`MX\`, \`CNAME\`, \`SRV\`. Defaults to \`A\`.

Upstream description:

Resource Record Type

&#8203;- A: A

&#8203;- AAAA: AAAA

&#8203;- MX: MX

&#8203;- CNAME: CNAME

&#8203;- SRV: SRV.

Receipt-pinned upstream constraints:

```json
{
  "default": "A",
  "enum": [
    "A",
    "AAAA",
    "MX",
    "CNAME",
    "SRV"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [response_cache](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/response_cache/): complete subsection reference.

- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/#schema-description) |
| `fallback_pool` | [fallback_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/fallback_pool/#section) |
| `fallback_pool.name` | [fallback_pool.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/fallback_pool/#schema-fallback_pool--name) |
| `fallback_pool.namespace` | [fallback_pool.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/fallback_pool/#schema-fallback_pool--namespace) |
| `fallback_pool.tenant` | [fallback_pool.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/fallback_pool/#schema-fallback_pool--tenant) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/#schema-namespace) |
| `record_type` | [record_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/#schema-record_type) |
| `response_cache` | [response_cache](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/response_cache/#section) |
| `response_cache.default_response_cache_parameters` | [response_cache.default_response_cache_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/response_cache/default_response_cache_parameters/#section) |
| `response_cache.disable_spec` | [response_cache.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/response_cache/disable_spec/#section) |
| `response_cache.response_cache_parameters` | [response_cache.response_cache_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/response_cache/response_cache_parameters/#section) |
| `response_cache.response_cache_parameters.cache_cidr_ipv4` | [response_cache.response_cache_parameters.cache_cidr_ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/response_cache/response_cache_parameters/#schema-response_cache--response_cache_parameters--cache_cidr_ipv4) |
| `response_cache.response_cache_parameters.cache_cidr_ipv6` | [response_cache.response_cache_parameters.cache_cidr_ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/response_cache/response_cache_parameters/#schema-response_cache--response_cache_parameters--cache_cidr_ipv6) |
| `response_cache.response_cache_parameters.cache_ttl` | [response_cache.response_cache_parameters.cache_ttl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/response_cache/response_cache_parameters/#schema-response_cache--response_cache_parameters--cache_ttl) |
| `rule_list` | [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/#section) |
| `rule_list.rules` | [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/#section) |
| `rule_list.rules.asn_list` | [rule_list.rules.asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/asn_list/#section) |
| `rule_list.rules.asn_list.as_numbers` | [rule_list.rules.asn_list.as_numbers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/asn_list/#schema-rule_list--rules--asn_list--as_numbers) |
| `rule_list.rules.asn_matcher` | [rule_list.rules.asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/asn_matcher/#section) |
| `rule_list.rules.asn_matcher.asn_sets` | [rule_list.rules.asn_matcher.asn_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/asn_matcher/asn_sets/#section) |
| `rule_list.rules.asn_matcher.asn_sets.kind` | [rule_list.rules.asn_matcher.asn_sets.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/asn_matcher/asn_sets/#schema-rule_list--rules--asn_matcher--asn_sets--kind) |
| `rule_list.rules.asn_matcher.asn_sets.name` | [rule_list.rules.asn_matcher.asn_sets.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/asn_matcher/asn_sets/#schema-rule_list--rules--asn_matcher--asn_sets--name) |
| `rule_list.rules.asn_matcher.asn_sets.namespace` | [rule_list.rules.asn_matcher.asn_sets.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/asn_matcher/asn_sets/#schema-rule_list--rules--asn_matcher--asn_sets--namespace) |
| `rule_list.rules.asn_matcher.asn_sets.tenant` | [rule_list.rules.asn_matcher.asn_sets.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/asn_matcher/asn_sets/#schema-rule_list--rules--asn_matcher--asn_sets--tenant) |
| `rule_list.rules.asn_matcher.asn_sets.uid` | [rule_list.rules.asn_matcher.asn_sets.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/asn_matcher/asn_sets/#schema-rule_list--rules--asn_matcher--asn_sets--uid) |
| `rule_list.rules.geo_location_label_selector` | [rule_list.rules.geo_location_label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/geo_location_label_selector/#section) |
| `rule_list.rules.geo_location_label_selector.expressions` | [rule_list.rules.geo_location_label_selector.expressions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/geo_location_label_selector/#schema-rule_list--rules--geo_location_label_selector--expressions) |
| `rule_list.rules.geo_location_set` | [rule_list.rules.geo_location_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/geo_location_set/#section) |
| `rule_list.rules.geo_location_set.name` | [rule_list.rules.geo_location_set.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/geo_location_set/#schema-rule_list--rules--geo_location_set--name) |
| `rule_list.rules.geo_location_set.namespace` | [rule_list.rules.geo_location_set.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/geo_location_set/#schema-rule_list--rules--geo_location_set--namespace) |
| `rule_list.rules.geo_location_set.tenant` | [rule_list.rules.geo_location_set.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/geo_location_set/#schema-rule_list--rules--geo_location_set--tenant) |
| `rule_list.rules.ip_prefix_list` | [rule_list.rules.ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/ip_prefix_list/#section) |
| `rule_list.rules.ip_prefix_list.invert_match` | [rule_list.rules.ip_prefix_list.invert_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/ip_prefix_list/#schema-rule_list--rules--ip_prefix_list--invert_match) |
| `rule_list.rules.ip_prefix_list.ip_prefixes` | [rule_list.rules.ip_prefix_list.ip_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/ip_prefix_list/#schema-rule_list--rules--ip_prefix_list--ip_prefixes) |
| `rule_list.rules.ip_prefix_set` | [rule_list.rules.ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/ip_prefix_set/#section) |
| `rule_list.rules.ip_prefix_set.invert_matcher` | [rule_list.rules.ip_prefix_set.invert_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/ip_prefix_set/#schema-rule_list--rules--ip_prefix_set--invert_matcher) |
| `rule_list.rules.ip_prefix_set.prefix_sets` | [rule_list.rules.ip_prefix_set.prefix_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/ip_prefix_set/prefix_sets/#section) |
| `rule_list.rules.ip_prefix_set.prefix_sets.kind` | [rule_list.rules.ip_prefix_set.prefix_sets.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/ip_prefix_set/prefix_sets/#schema-rule_list--rules--ip_prefix_set--prefix_sets--kind) |
| `rule_list.rules.ip_prefix_set.prefix_sets.name` | [rule_list.rules.ip_prefix_set.prefix_sets.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/ip_prefix_set/prefix_sets/#schema-rule_list--rules--ip_prefix_set--prefix_sets--name) |
| `rule_list.rules.ip_prefix_set.prefix_sets.namespace` | [rule_list.rules.ip_prefix_set.prefix_sets.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/ip_prefix_set/prefix_sets/#schema-rule_list--rules--ip_prefix_set--prefix_sets--namespace) |
| `rule_list.rules.ip_prefix_set.prefix_sets.tenant` | [rule_list.rules.ip_prefix_set.prefix_sets.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/ip_prefix_set/prefix_sets/#schema-rule_list--rules--ip_prefix_set--prefix_sets--tenant) |
| `rule_list.rules.ip_prefix_set.prefix_sets.uid` | [rule_list.rules.ip_prefix_set.prefix_sets.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/ip_prefix_set/prefix_sets/#schema-rule_list--rules--ip_prefix_set--prefix_sets--uid) |
| `rule_list.rules.pool` | [rule_list.rules.pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/pool/#section) |
| `rule_list.rules.pool.name` | [rule_list.rules.pool.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/pool/#schema-rule_list--rules--pool--name) |
| `rule_list.rules.pool.namespace` | [rule_list.rules.pool.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/pool/#schema-rule_list--rules--pool--namespace) |
| `rule_list.rules.pool.tenant` | [rule_list.rules.pool.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/pool/#schema-rule_list--rules--pool--tenant) |
| `rule_list.rules.score` | [rule_list.rules.score](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/rules/#schema-rule_list--rules--score) |

## Next pages

- [fallback_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/fallback_pool/)
- [response_cache](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/response_cache/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/rule_list/)
- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/)
