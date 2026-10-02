---
page_title: "client_side_defense.policy.js_insertion_rules.rules.path"
subcategory: "Load Balancing"
description: "Path match of the URI can be either be, Prefix match or exact match or regular expression match."
xcsh_docs: {"aliases": ["client side defense policy js insertion rules rules path"], "body_bytes": 5385, "body_sha256": "sha256:0edfeed315cfd56ed1fc69c5226dfb348fa4166e52f8bc06f1352a281533c607", "capabilities": ["load-balancing", "security.client-side-defense"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:path", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules", "path": "documentation/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/path/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2203320213003032-1030202233233330-3323210012103012-0220221000312111-2211210003121120-0023220111321323-2002221010202230-1313320201130312", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "rules", "path"], "schema_version": 1, "sections": [{"aliases": ["path"], "anchor": "schema-client_side_defense--policy--js_insertion_rules--rules--path--path", "description": "Exclusive with Exact path value to match.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:path", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "rules", "path", "path"], "syntax": "attribute", "type": "string"}, {"aliases": ["prefix"], "anchor": "schema-client_side_defense--policy--js_insertion_rules--rules--path--prefix", "description": "Exclusive with Path prefix to match (e.g. The value / will match on all paths)", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:path", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "rules", "path", "prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["regex"], "anchor": "schema-client_side_defense--policy--js_insertion_rules--rules--path--regex", "description": "Exclusive with Regular expression of path match (e.g. The value .* will match on all paths)", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insertion_rules:rules:path", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insertion_rules", "rules", "path", "regex"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/path/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_side_defense.policy.js_insertion_rules.rules.path

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [client_side_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/)
- [client_side_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/)
- [client_side_defense.policy.js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/)
- [client_side_defense.policy.js_insertion_rules.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/)
- client_side_defense.policy.js_insertion_rules.rules.path

<a id="section"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

## Direct properties

<a id="schema-client_side_defense--policy--js_insertion_rules--rules--path--path"></a>

### path property

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-client_side_defense--policy--js_insertion_rules--rules--path--prefix"></a>

### prefix property

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-client_side_defense--policy--js_insertion_rules--rules--path--regex"></a>

### regex property

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

## Next pages

- [client_side_defense.policy.js_insertion_rules.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insertion_rules/rules/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
