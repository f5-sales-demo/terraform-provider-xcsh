---
page_title: "cloudflare.manual_js_insert"
subcategory: ""
description: "cloudflare.manual_js_insert for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 2025, "body_sha256": "sha256:778089658fda13f0a8dece1642d94c5aa149a70d9e5b150d4969e03ecbd4f385", "canonical_id": "xcsh-docs:resources:protected_application:properties:cloudflare:manual_js_insert", "child_ids": [], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudflare:manual_js_insert", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudflare", "path": "docs/guides/resources--protected_application--properties--cloudflare--manual_js_insert.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudflare", "manual_js_insert"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudflare/manual_js_insert/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudflare.manual_js_insert for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# cloudflare.manual_js_insert

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md)
- [Property reference](resources--protected_application--reference.md)
- [cloudflare](resources--protected_application--properties--cloudflare.md)
- cloudflare.manual_js_insert

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Insert JavaScript Manually. Insert JavaScript manually.

Upstream description:

Insert JavaScript manually.

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
manual_js_insert {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-cloudflare--manual_js_insert--js_download_path"></a>

### js_download_path property

Type: `"string"`. Optional.

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths. If not specified, default to ‘/common.js’.

Upstream description:

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths.

If not specified, default to ‘/common.js’.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  }
}
```

## Next pages

- [cloudflare](resources--protected_application--properties--cloudflare.md)
- [xcsh_protected_application](../resources/protected_application.md)
