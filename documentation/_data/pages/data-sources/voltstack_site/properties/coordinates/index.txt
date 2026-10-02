---
page_title: "coordinates"
subcategory: ""
description: "Coordinates of the site which provides the site physical location."
xcsh_docs: {"aliases": ["coordinates"], "body_bytes": 2161, "body_sha256": "sha256:c3522ba8b2a55fb3dc0ec65a5072302c8ac0ffb7a8abb8caeedec640d3417af6", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:coordinates", "parent_id": "xcsh-docs:data-sources:voltstack_site:reference", "path": "documentation/data-sources/voltstack_site/properties/coordinates/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0120100301310332-3321102221032322-3313312210320231-1230000302000320-0220120120321122-0322232023312223-3122020223200100-1323202230122303", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["coordinates"], "schema_version": 1, "sections": [{"aliases": ["latitude"], "anchor": "schema-coordinates--latitude", "description": "Latitude of the site location.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:coordinates", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["coordinates", "latitude"], "syntax": "attribute", "type": "number"}, {"aliases": ["longitude"], "anchor": "schema-coordinates--longitude", "description": "Longitude of site location.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:coordinates", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["coordinates", "longitude"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/coordinates/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Coordinates of the site which provides the site physical location.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# coordinates

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- coordinates

<a id="section"></a>

Type: `"single"`. Computed.

Coordinates of the site which provides the site physical location.

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

## Direct properties

<a id="schema-coordinates--latitude"></a>

### latitude property

Type: `"number"`. Computed.

Latitude. Latitude of the site location.

Upstream description:

Latitude of the site location.

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
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  }
}
```

<a id="schema-coordinates--longitude"></a>

### longitude property

Type: `"number"`. Computed.

Longitude. Longitude of site location.

Upstream description:

Longitude of site location.

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
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
