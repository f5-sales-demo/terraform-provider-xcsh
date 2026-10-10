---
page_title: "cloudfront.js_insertion_rules"
subcategory: ""
description: "This defines custom JavaScript insertion rules for Bot Defense Policy."
xcsh_docs: {"aliases": ["cloudfront js insertion rules"], "body_bytes": 5921, "body_sha256": "sha256:c39cde7a0f98670e0330e50f7f22a44ffd54097315b652f9548bb6be4b9b0978", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list", "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront", "path": "documentation/resources/protected_application/properties/cloudfront/js_insertion_rules/index.md", "product": "distributed-cloud", "provider_name": "protected_application", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210", "registry_path": "docs/guides/resources--protected_application--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules:RequiredObjectAttributes:rules", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cloudfront", "js_insertion_rules"], "schema_version": 1, "sections": [{"aliases": ["cloudfront js insertion rules exclude list"], "anchor": "section", "description": "Optional JavaScript insertions exclude list of domain and path matchers.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.exclude_list:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.exclude_list:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:exclude_list:domain", "type": "conflicts"}], "schema_path": ["cloudfront", "js_insertion_rules", "exclude_list"], "syntax": "block", "type": "object"}, {"aliases": ["cloudfront js insertion rules javascript location"], "anchor": "schema-cloudfront--js_insertion_rules--javascript_location", "description": "All inside networks. - JAVA_SCRIPT_LOCATION_UNDEFINED: JAVA_SCRIPT_LOCATION_UNDEFINED Undefined Insert JavaScript after <HEAD> tag Insert JavaScript after </title> tag. Insert JavaScript before first tag.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["AFTER_HEAD", "AFTER_TITLE_END", "BEFORE_SCRIPT", "JAVA_SCRIPT_LOCATION_UNDEFINED"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "js_insertion_rules", "javascript_location"], "syntax": "attribute", "type": "string"}, {"aliases": ["cloudfront js insertion rules javascript mode"], "anchor": "schema-cloudfront--js_insertion_rules--javascript_mode", "description": "Web Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is cacheable Bot Defense JavaScript for telemetry collection is requested synchronously, and it is", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["ASYNC_JS_CACHING", "ASYNC_JS_NO_CACHING", "SYNC_JS_CACHING", "SYNC_JS_NO_CACHING"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "js_insertion_rules", "javascript_mode"], "syntax": "attribute", "type": "string"}, {"aliases": ["cloudfront js insertion rules js download path"], "anchor": "schema-cloudfront--js_insertion_rules--js_download_path", "description": "Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any other website/application paths. If not specified, default to ‘/common.js’.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloudfront", "js_insertion_rules", "js_download_path"], "syntax": "attribute", "type": "string"}, {"aliases": ["cloudfront js insertion rules rules"], "anchor": "section", "description": "Required list of pages to insert Bot Defense client JavaScript.", "document_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-cloudfront--js_insertion_rules--rules--exact_path", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.rules:ConflictingListObjectAttributes:exact_path,glob", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--rules--exact_path", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.rules:ConflictingListObjectAttributes:exact_path,prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--rules--glob", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.rules:ConflictingListObjectAttributes:exact_path,glob", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--rules--glob", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.rules:ConflictingListObjectAttributes:glob,prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--rules--prefix", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.rules:ConflictingListObjectAttributes:exact_path,prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules", "type": "conflicts"}, {"anchor": "schema-cloudfront--js_insertion_rules--rules--prefix", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.rules:ConflictingListObjectAttributes:glob,prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.rules:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules:any_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloudfront.js_insertion_rules.rules:ConflictingListObjectAttributes:any_domain,domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:protected_application:properties:cloudfront:js_insertion_rules:rules:domain", "type": "conflicts"}], "schema_path": ["cloudfront", "js_insertion_rules", "rules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/js_insertion_rules/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This defines custom JavaScript insertion rules for Bot Defense Policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.js_insertion_rules

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/)
- cloudfront.js_insertion_rules

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

This defines custom JavaScript insertion rules for Bot Defense Policy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
```

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
js_insertion_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [exclude_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/js_insertion_rules/exclude_list/): complete subsection reference.

<a id="schema-cloudfront--js_insertion_rules--javascript_location"></a>

### javascript_location property

Type: `"string"`. Optional.

\[Enum: JAVA\_SCRIPT\_LOCATION\_UNDEFINED|AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside
networks. - JAVA\_SCRIPT\_LOCATION\_UNDEFINED: JAVA\_SCRIPT\_LOCATION\_UNDEFINED Undefined Insert
JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript
before first tag. Possible values are \`JAVA\_SCRIPT\_LOCATION\_UNDEFINED\`, \`AFTER\_HEAD\`,
\`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to \`JAVA\_SCRIPT\_LOCATION\_UNDEFINED\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["AFTER_HEAD","AFTER_TITLE_END","BEFORE_SCRIPT","JAVA_SCRIPT_LOCATION_UNDEFINED"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("JAVA_SCRIPT_LOCATION_UNDEFINED",
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

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

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ASYNC_JS_CACHING","ASYNC_JS_NO_CACHING","SYNC_JS_CACHING","SYNC_JS_NO_CACHING"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ASYNC_JS_NO_CACHING",
    "ASYNC_JS_CACHING",
    "SYNC_JS_NO_CACHING",
    "SYNC_JS_CACHING"),
}
```

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

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  }
}
```

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/js_insertion_rules/rules/): complete subsection reference.
