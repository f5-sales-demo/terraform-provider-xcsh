---
page_title: "cloudfront.js_insertion_rules"
subcategory: ""
description: "This defines custom JavaScript insertion rules for Bot Defense Policy."
xcsh_docs: {"aliases": ["cloudfront js insertion rules"], "body_bytes": 4559, "body_sha256": "sha256:4bb9299288b4e3f9a7334795f4663ce6beeb4e38dab08eebb78da4dfebc2c625", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list", "xcsh-docs:data-sources:protected_application:properties:cloudfront:js_insertion_rules:rules"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:js_insertion_rules", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront", "path": "documentation/data-sources/protected_application/properties/cloudfront/js_insertion_rules/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2130101011003221-3001021112313100-3232203110030212-2221323311332011-1300320223203120-2131032201302211-1110112322003011-2223021311232200", "registry_path": "docs/guides/data-sources--protected_application--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "js_insertion_rules"], "schema_version": 1, "sections": [{"aliases": ["cloudfront js insertion rules exclude list"], "anchor": "section", "description": "Optional JavaScript insertions exclude list of domain and path matchers.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["cloudfront", "js_insertion_rules", "exclude_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloudfront js insertion rules javascript location"], "anchor": "schema-cloudfront--js_insertion_rules--javascript_location", "description": "All inside networks. - JAVA_SCRIPT_LOCATION_UNDEFINED: JAVA_SCRIPT_LOCATION_UNDEFINED Undefined Insert JavaScript after <HEAD> tag Insert JavaScript after </title> tag. Insert JavaScript before first tag.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:js_insertion_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "js_insertion_rules", "javascript_location"], "syntax": "attribute", "type": "string"}, {"aliases": ["cloudfront js insertion rules javascript mode"], "anchor": "schema-cloudfront--js_insertion_rules--javascript_mode", "description": "Web Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is cacheable Bot Defense JavaScript for telemetry collection is requested synchronously, and it is", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:js_insertion_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "js_insertion_rules", "javascript_mode"], "syntax": "attribute", "type": "string"}, {"aliases": ["cloudfront js insertion rules js download path"], "anchor": "schema-cloudfront--js_insertion_rules--js_download_path", "description": "Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any other website/application paths. If not specified, default to ‘/common.js’.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:js_insertion_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "js_insertion_rules", "js_download_path"], "syntax": "attribute", "type": "string"}, {"aliases": ["cloudfront js insertion rules rules"], "anchor": "section", "description": "Required list of pages to insert Bot Defense client JavaScript.", "document_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:js_insertion_rules:rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["cloudfront", "js_insertion_rules", "rules"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/js_insertion_rules/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This defines custom JavaScript insertion rules for Bot Defense Policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.js_insertion_rules

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/)
- cloudfront.js_insertion_rules

<a id="section"></a>

Type: `"single"`. Computed.

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

- [exclude_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/js_insertion_rules/exclude_list/): complete subsection reference.

<a id="schema-cloudfront--js_insertion_rules--javascript_location"></a>

### javascript_location property

Type: `"string"`. Computed.

\[Enum: JAVA\_SCRIPT\_LOCATION\_UNDEFINED|AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside
networks. - JAVA\_SCRIPT\_LOCATION\_UNDEFINED: JAVA\_SCRIPT\_LOCATION\_UNDEFINED Undefined Insert
JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript
before first tag. Possible values are \`JAVA\_SCRIPT\_LOCATION\_UNDEFINED\`, \`AFTER\_HEAD\`,
\`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to \`JAVA\_SCRIPT\_LOCATION\_UNDEFINED\`.

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

<a id="schema-cloudfront--js_insertion_rules--javascript_mode"></a>

### javascript_mode property

Type: `"string"`. Computed.

\[Enum: ASYNC\_JS\_NO\_CACHING|ASYNC\_JS\_CACHING|SYNC\_JS\_NO\_CACHING|SYNC\_JS\_CACHING\] Web
Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is cacheable Bot Defense JavaScript for telemetry collection is requested.. Possible values
are \`ASYNC\_JS\_NO\_CACHING\`, \`ASYNC\_JS\_CACHING\`, \`SYNC\_JS\_NO\_CACHING\`,
\`SYNC\_JS\_CACHING\`. Defaults to \`ASYNC\_JS\_NO\_CACHING\`.

Additional upstream details:

Web Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested
asynchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested
asynchronously, and it is cacheable Bot Defense JavaScript for telemetry collection is requested
synchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested
synchronously, and it is cacheable.

Receipt-pinned upstream constraints:

```json
{
  "default": "ASYNC_JS_NO_CACHING",
  "enum": [
    "ASYNC_JS_NO_CACHING",
    "ASYNC_JS_CACHING",
    "SYNC_JS_NO_CACHING",
    "SYNC_JS_CACHING"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-cloudfront--js_insertion_rules--js_download_path"></a>

### js_download_path property

Type: `"string"`. Computed.

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths. If not specified, default to ‘/common.js’.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/properties/cloudfront/js_insertion_rules/rules/): complete subsection reference.
