---
page_title: "bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.path"
subcategory: "Load Balancing"
description: "Path match of the URI can be either be, Prefix match or exact match or regular expression match."
xcsh_docs: {"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection web only js insert all pages except exclude list path"], "body_bytes": 5010, "body_sha256": "sha256:7e7cd5227f020e66adf50517c863bb47e0ffc29590f7b7b04b34e0fb201a2139", "capabilities": ["load-balancing", "security.bot-defense"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insert_all_pages_except:exclude_list:path", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insert_all_pages_except:exclude_list", "path": "documentation/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/js_insert_all_pages_except/exclude_list/path/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3133321021213011-1113001333203220-2231301222213211-1030201200200330-2132110221330000-0111200133021133-0111011222210101-0002033010023121", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense_advanced_protection", "web_only", "js_insert_all_pages_except", "exclude_list", "path"], "schema_version": 1, "sections": [{"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection web only js insert all pages except exclude list path path"], "anchor": "schema-bot_defense_advanced_protection--web_only--js_insert_all_pages_except--exclude_list--path--path", "description": "Exclusive with Exact path value to match.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insert_all_pages_except:exclude_list:path", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense_advanced_protection", "web_only", "js_insert_all_pages_except", "exclude_list", "path", "path"], "syntax": "attribute", "type": "string"}, {"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection web only js insert all pages except exclude list path prefix"], "anchor": "schema-bot_defense_advanced_protection--web_only--js_insert_all_pages_except--exclude_list--path--prefix", "description": "Exclusive with Path prefix to match (e.g. The value / will match on all paths)", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insert_all_pages_except:exclude_list:path", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense_advanced_protection", "web_only", "js_insert_all_pages_except", "exclude_list", "path", "prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["advanced bot defense", "bot defense advanced protection", "bot defense advanced protection web only js insert all pages except exclude list path regex"], "anchor": "schema-bot_defense_advanced_protection--web_only--js_insert_all_pages_except--exclude_list--path--regex", "description": "Exclusive with Regular expression of path match (e.g. The value .* will match on all paths)", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insert_all_pages_except:exclude_list:path", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense_advanced_protection", "web_only", "js_insert_all_pages_except", "exclude_list", "path", "regex"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/js_insert_all_pages_except/exclude_list/path/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.path

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [bot_defense_advanced_protection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/)
- [bot_defense_advanced_protection.web_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/js_insert_all_pages_except/)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/js_insert_all_pages_except/exclude_list/)
- bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.path

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

<a id="schema-bot_defense_advanced_protection--web_only--js_insert_all_pages_except--exclude_list--path--path"></a>

### path property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="schema-bot_defense_advanced_protection--web_only--js_insert_all_pages_except--exclude_list--path--prefix"></a>

### prefix property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="schema-bot_defense_advanced_protection--web_only--js_insert_all_pages_except--exclude_list--path--regex"></a>

### regex property

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
