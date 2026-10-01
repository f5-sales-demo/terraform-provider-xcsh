---
page_title: "bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list"
subcategory: "Load Balancing"
description: "bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3722, "body_sha256": "sha256:3b0e5651d54399efda2a50202e5e92b9a7490bab02c2fc349e251457005d787f", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insert_all_pages_except:exclude_list", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insert_all_pages_except:exclude_list:any_domain", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insert_all_pages_except:exclude_list:domain", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insert_all_pages_except:exclude_list:metadata", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insert_all_pages_except:exclude_list:path"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insert_all_pages_except:exclude_list", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:web_only:js_insert_all_pages_except", "path": "docs/guides/data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insert_all_pages_except--exclude_list.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense_advanced_protection", "web_only", "js_insert_all_pages_except", "exclude_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/web_only/js_insert_all_pages_except/exclude_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection.md)
- [bot_defense_advanced_protection.web_only](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only.md)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insert_all_pages_except.md)
- bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list

<a id="section"></a>

Type: `"list"`. Computed.

Optional JavaScript insertions exclude list of domain and path matchers.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [any_domain](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insert_all_pages_except--exclude_list--any_domain.md): complete subsection reference.

- [domain](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insert_all_pages_except--exclude_list--domain.md): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insert_all_pages_except--exclude_list--metadata.md): complete subsection reference.

- [path](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insert_all_pages_except--exclude_list--path.md): complete subsection reference.

## Next pages

- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.any_domain](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insert_all_pages_except--exclude_list--any_domain.md)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.domain](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insert_all_pages_except--exclude_list--domain.md)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.metadata](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insert_all_pages_except--exclude_list--metadata.md)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.path](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insert_all_pages_except--exclude_list--path.md)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection--web_only--js_insert_all_pages_except.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
