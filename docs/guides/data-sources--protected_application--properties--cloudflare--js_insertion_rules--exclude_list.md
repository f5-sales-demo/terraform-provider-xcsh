---
page_title: "cloudflare.js_insertion_rules.exclude_list"
subcategory: ""
description: "cloudflare.js_insertion_rules.exclude_list for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 2914, "body_sha256": "sha256:5e58d7cf9448f8f50c7040c5f35c68867d1c7723a606999a036239ecb8480dd6", "canonical_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:any_domain", "xcsh-docs:data-sources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:domain", "xcsh-docs:data-sources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:metadata", "xcsh-docs:data-sources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list:path"], "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:js_insertion_rules", "path": "docs/guides/data-sources--protected_application--properties--cloudflare--js_insertion_rules--exclude_list.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudflare", "js_insertion_rules", "exclude_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudflare/js_insertion_rules/exclude_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudflare.js_insertion_rules.exclude_list for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.js_insertion_rules.exclude_list

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md)
- [Property reference](data-sources--protected_application--reference.md)
- [cloudflare](data-sources--protected_application--properties--cloudflare.md)
- [cloudflare.js_insertion_rules](data-sources--protected_application--properties--cloudflare--js_insertion_rules.md)
- cloudflare.js_insertion_rules.exclude_list

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

- [any_domain](data-sources--protected_application--properties--cloudflare--js_insertion_rules--exclude_list--any_domain.md): complete subsection reference.

- [domain](data-sources--protected_application--properties--cloudflare--js_insertion_rules--exclude_list--domain.md): complete subsection reference.

- [metadata](data-sources--protected_application--properties--cloudflare--js_insertion_rules--exclude_list--metadata.md): complete subsection reference.

- [path](data-sources--protected_application--properties--cloudflare--js_insertion_rules--exclude_list--path.md): complete subsection reference.

## Next pages

- [cloudflare.js_insertion_rules.exclude_list.any_domain](data-sources--protected_application--properties--cloudflare--js_insertion_rules--exclude_list--any_domain.md)
- [cloudflare.js_insertion_rules.exclude_list.domain](data-sources--protected_application--properties--cloudflare--js_insertion_rules--exclude_list--domain.md)
- [cloudflare.js_insertion_rules.exclude_list.metadata](data-sources--protected_application--properties--cloudflare--js_insertion_rules--exclude_list--metadata.md)
- [cloudflare.js_insertion_rules.exclude_list.path](data-sources--protected_application--properties--cloudflare--js_insertion_rules--exclude_list--path.md)
- [cloudflare.js_insertion_rules](data-sources--protected_application--properties--cloudflare--js_insertion_rules.md)
- [xcsh_protected_application](../data-sources/protected_application.md)
