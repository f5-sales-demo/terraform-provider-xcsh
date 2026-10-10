---
page_title: "client_side_defense.policy.js_insert_all_pages_except"
subcategory: "Load Balancing"
description: "Insert Client-Side Defense JavaScript in all pages with the exceptions."
xcsh_docs: {"aliases": ["client side defense policy js insert all pages except"], "body_bytes": 1425, "body_sha256": "sha256:336ba97757c2cf1560b39691fa972415f5196cd64ff62741f5b76989b28714e6", "capabilities": ["load-balancing", "security.client-side-defense"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except:exclude_list"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy", "path": "documentation/resources/http_loadbalancer/properties/client_side_defense/policy/js_insert_all_pages_except/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3203113232311103-1203301023312113-1103112313002121-2300301213131000-2021012213013110-0200231011312133-0033320012121330-3322223131210231", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["client_side_defense", "policy", "js_insert_all_pages_except"], "schema_version": 1, "sections": [{"aliases": ["client side defense policy js insert all pages except exclude list"], "anchor": "section", "description": "Optional JavaScript insertions exclude list of domain and path matchers.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:client_side_defense:policy:js_insert_all_pages_except:exclude_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["client_side_defense", "policy", "js_insert_all_pages_except", "exclude_list"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/client_side_defense/policy/js_insert_all_pages_except/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Insert Client-Side Defense JavaScript in all pages with the exceptions.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_side_defense.policy.js_insert_all_pages_except

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [client_side_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/)
- [client_side_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/)
- client_side_defense.policy.js_insert_all_pages_except

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Insert Client-Side Defense JavaScript in all pages with the exceptions.

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
js_insert_all_pages_except {
  # Configure direct properties listed below.
}
```

## Direct properties

- [exclude_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/client_side_defense/policy/js_insert_all_pages_except/exclude_list/): complete subsection reference.
