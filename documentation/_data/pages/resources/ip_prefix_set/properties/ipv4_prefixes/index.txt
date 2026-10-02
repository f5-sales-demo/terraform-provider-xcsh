---
page_title: "ipv4_prefixes"
subcategory: ""
description: "List of IPv4 prefixes with description."
xcsh_docs: {"aliases": ["ipv4 prefixes"], "body_bytes": 3111, "body_sha256": "sha256:bd1ffc212f227f580c57bd6ba6515cd8bf25427c4afb81b7601d0cb98d4e7af7", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ip_prefix_set:collection", "completeness": "complete", "id": "xcsh-docs:resources:ip_prefix_set:properties:ipv4_prefixes", "parent_id": "xcsh-docs:resources:ip_prefix_set:reference", "path": "documentation/resources/ip_prefix_set/properties/ipv4_prefixes/index.md", "product": "distributed-cloud", "provider_name": "ip_prefix_set", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-2200333320220021-3021300112012033-2230031200012331-2110030133331300-3320332221003021-1020200323133031-0201022120222300-2112013311220230", "registry_path": "docs/guides/resources--ip_prefix_set--reference--group-001.md", "relationships": [{"anchor": "schema-ipv4_prefixes--ipv4_prefix", "enforcement": "provider-schema", "group": "ipv4_prefixes:RequiredListObjectAttributes:ipv4_prefix", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:ip_prefix_set:properties:ipv4_prefixes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ipv4_prefixes"], "schema_version": 1, "sections": [{"aliases": ["description spec"], "anchor": "schema-ipv4_prefixes--description_spec", "description": "Description. Human-readable description text", "document_id": "xcsh-docs:resources:ip_prefix_set:properties:ipv4_prefixes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipv4_prefixes", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["ipv4 prefix"], "anchor": "schema-ipv4_prefixes--ipv4_prefix", "description": "IP address configuration", "document_id": "xcsh-docs:resources:ip_prefix_set:properties:ipv4_prefixes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ipv4_prefixes", "ipv4_prefix"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ip_prefix_set/properties/ipv4_prefixes/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of IPv4 prefixes with description.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["ip_prefix_setCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipv4_prefixes

Breadcrumbs:

- [xcsh_ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ip_prefix_set/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ip_prefix_set/properties/)
- ipv4_prefixes

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

IPv4 Prefixes. List of IPv4 prefixes with description.

Upstream description:

List of IPv4 prefixes with description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ipv4_prefix")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1024,
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
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
ipv4_prefixes {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ipv4_prefixes--description_spec"></a>

### description_spec property

Type: `"string"`. Optional.

Description. Human-readable description text

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

<a id="schema-ipv4_prefixes--ipv4_prefix"></a>

### ipv4_prefix property

Type: `"string"`. Optional.

IPv4 Prefix. IP address configuration

Upstream description:

IP address configuration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ip_prefix_set/properties/)
- [xcsh_ip_prefix_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ip_prefix_set/)
