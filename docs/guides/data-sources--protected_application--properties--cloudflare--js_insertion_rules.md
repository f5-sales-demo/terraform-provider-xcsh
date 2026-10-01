---
page_title: "cloudflare.js_insertion_rules"
subcategory: ""
description: "cloudflare.js_insertion_rules for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 3844, "body_sha256": "sha256:3d255cc1fa37471ae349269bfcf35e07ce8ee9d883de1174351689fb9fcc7db3", "canonical_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:js_insertion_rules", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudflare:js_insertion_rules:exclude_list", "xcsh-docs:data-sources:protected_application:properties:cloudflare:js_insertion_rules:rules"], "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudflare:js_insertion_rules", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudflare", "path": "docs/guides/data-sources--protected_application--properties--cloudflare--js_insertion_rules.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudflare", "js_insertion_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudflare/js_insertion_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudflare.js_insertion_rules for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.js_insertion_rules

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md)
- [Property reference](data-sources--protected_application--reference.md)
- [cloudflare](data-sources--protected_application--properties--cloudflare.md)
- cloudflare.js_insertion_rules

<a id="section"></a>

Type: `"single"`. Computed.

Defines custom JavaScript insertion rules for Bot Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Bot Defense Policy.

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

- [exclude_list](data-sources--protected_application--properties--cloudflare--js_insertion_rules--exclude_list.md): complete subsection reference.

<a id="schema-cloudflare--js_insertion_rules--javascript_location"></a>

### javascript_location property

Type: `"string"`. Computed.

\[Enum: JAVA\_SCRIPT\_LOCATION\_UNDEFINED|AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside
networks. - JAVA\_SCRIPT\_LOCATION\_UNDEFINED: JAVA\_SCRIPT\_LOCATION\_UNDEFINED Undefined Insert
JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript
before first tag. Possible values are \`JAVA\_SCRIPT\_LOCATION\_UNDEFINED\`, \`AFTER\_HEAD\`,
\`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to \`JAVA\_SCRIPT\_LOCATION\_UNDEFINED\`.

Upstream description:

All inside networks.

&#8203;- JAVA\_SCRIPT\_LOCATION\_UNDEFINED: JAVA\_SCRIPT\_LOCATION\_UNDEFINED

Undefined Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag.
Insert JavaScript before first tag.

Receipt-pinned upstream constraints:

```json
{
  "default": "JAVA_SCRIPT_LOCATION_UNDEFINED",
  "enum": [
    "JAVA_SCRIPT_LOCATION_UNDEFINED",
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-cloudflare--js_insertion_rules--js_download_path"></a>

### js_download_path property

Type: `"string"`. Computed.

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

- [rules](data-sources--protected_application--properties--cloudflare--js_insertion_rules--rules.md): complete subsection reference.

## Next pages

- [cloudflare.js_insertion_rules.exclude_list](data-sources--protected_application--properties--cloudflare--js_insertion_rules--exclude_list.md)
- [cloudflare.js_insertion_rules.rules](data-sources--protected_application--properties--cloudflare--js_insertion_rules--rules.md)
- [cloudflare](data-sources--protected_application--properties--cloudflare.md)
- [xcsh_protected_application](../data-sources/protected_application.md)
