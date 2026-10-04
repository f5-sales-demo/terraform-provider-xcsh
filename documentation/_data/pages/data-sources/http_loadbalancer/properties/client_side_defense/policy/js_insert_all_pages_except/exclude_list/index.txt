---
page_title: "client_side_defense.policy.js_insert_all_pages_except.exclude_list"
subcategory: "Load Balancing"
description: "Optional JavaScript insertions exclude list of domain and path matchers."
xcsh_docs: {"aliases": ["client side defense policy js insert all pages except exclude list"], "body_bytes": 4148, "body_sha256": "sha256:d0ac4ba1628d340c1366469ff7cbbdc08f126e86ea098694b77b4d4f4065eff8", "capabilities": ["load-balancing", "security.client-side-defense"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except:exclude_list:any_domain", "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except:exclude_list:domain", "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except:exclude_list:metadata", "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except:exclude_list:path"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except:exclude_list", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except", "path": "documentation/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insert_all_pages_except/exclude_list/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3001131021322100-1033021100001131-0032031130201230-1013300210330323-0120301310321321-2031303231323121-3301322213110123-1223221202211030", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["client_side_defense", "policy", "js_insert_all_pages_except", "exclude_list"], "schema_version": 1, "sections": [{"aliases": ["client side defense policy js insert all pages except exclude list any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except:exclude_list:any_domain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insert_all_pages_except", "exclude_list", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["client side defense policy js insert all pages except exclude list domain"], "anchor": "section", "description": "Domains names.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except:exclude_list:domain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insert_all_pages_except", "exclude_list", "domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["client side defense policy js insert all pages except exclude list metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except:exclude_list:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insert_all_pages_except", "exclude_list", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["client side defense policy js insert all pages except exclude list path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except:exclude_list:path", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insert_all_pages_except", "exclude_list", "path"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insert_all_pages_except/exclude_list/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Optional JavaScript insertions exclude list of domain and path matchers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_side_defense.policy.js_insert_all_pages_except.exclude_list

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [client_side_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/)
- [client_side_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/)
- [client_side_defense.policy.js_insert_all_pages_except](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insert_all_pages_except/)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insert_all_pages_except/exclude_list/any_domain/): complete subsection reference.

- [domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insert_all_pages_except/exclude_list/domain/): complete subsection reference.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insert_all_pages_except/exclude_list/metadata/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insert_all_pages_except/exclude_list/path/): complete subsection reference.

## Next pages

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insert_all_pages_except/exclude_list/any_domain/)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insert_all_pages_except/exclude_list/domain/)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insert_all_pages_except/exclude_list/metadata/)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insert_all_pages_except/exclude_list/path/)
- [client_side_defense.policy.js_insert_all_pages_except](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/client_side_defense/policy/js_insert_all_pages_except/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
