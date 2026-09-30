---
page_title: "bot_defense.policy"
subcategory: "Load Balancing"
description: "bot_defense.policy for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 6464, "body_sha256": "sha256:e52c94074deb1b25fb9f138307deca46c26508c14395cd6acfff14ed8c95b466", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:disable_js_insert", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:disable_mobile_sdk", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insert_all_pages", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insertion_rules", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:mobile_sdk_config", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense", "path": "docs/guides/resources--http_loadbalancer--properties--bot_defense--policy.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense/policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# bot_defense.policy

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [bot_defense](resources--http_loadbalancer--properties--bot_defense.md)
- bot_defense.policy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines various configuration OPTIONS for Bot Defense policy.

Upstream description:

This defines various configuration OPTIONS for Bot Defense policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("protected_app_endpoints"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("disable_mobile_sdk",
    "mobile_sdk_config"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("js_insert_all_pages_except",
    "js_insertion_rules")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]",
  "x-ves-oneof-field-mobile_sdk_choice": "[\"disable_mobile_sdk\",\"mobile_sdk_config\"]"
}
```

Terraform syntax:

```terraform
policy {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_js_insert](resources--http_loadbalancer--properties--bot_defense--policy--disable_js_insert.md): complete subsection reference.

- [disable_mobile_sdk](resources--http_loadbalancer--properties--bot_defense--policy--disable_mobile_sdk.md): complete subsection reference.

<a id="schema-bot_defense--policy--javascript_mode"></a>

### javascript_mode property

Type: `"string"`. Optional.

\[Enum: ASYNC\_JS\_NO\_CACHING|ASYNC\_JS\_CACHING|SYNC\_JS\_NO\_CACHING|SYNC\_JS\_CACHING\] Web
Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is cacheable Bot Defense JavaScript for telemetry collection is requested.. Possible values
are \`ASYNC\_JS\_NO\_CACHING\`, \`ASYNC\_JS\_CACHING\`, \`SYNC\_JS\_NO\_CACHING\`,
\`SYNC\_JS\_CACHING\`. Defaults to \`ASYNC\_JS\_NO\_CACHING\`.

Upstream description:

Web Client JavaScript Mode.

Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is non-cacheable
Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is cacheable Bot
Defense JavaScript for telemetry collection is requested synchronously, and it is non-cacheable Bot
Defense JavaScript for telemetry collection is requested synchronously, and it is cacheable.

Provider validators and defaults (from schema source):

```go
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

<a id="schema-bot_defense--policy--js_download_path"></a>

### js_download_path property

Type: `"string"`. Optional.

Customize Bot Defense Client JavaScript path. If not specified, default

Upstream description:

Customize Bot Defense Client JavaScript path. If not specified, default \`/common.js\`

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

- [js_insert_all_pages](resources--http_loadbalancer--properties--bot_defense--policy--js_insert_all_pages.md): complete subsection reference.

- [js_insert_all_pages_except](resources--http_loadbalancer--properties--bot_defense--policy--js_insert_all_pages_except.md): complete subsection reference.

- [js_insertion_rules](resources--http_loadbalancer--properties--bot_defense--policy--js_insertion_rules.md): complete subsection reference.

- [mobile_sdk_config](resources--http_loadbalancer--properties--bot_defense--policy--mobile_sdk_config.md): complete subsection reference.

- [protected_app_endpoints](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md): complete subsection reference.

## Next pages

- [bot_defense.policy.disable_js_insert](resources--http_loadbalancer--properties--bot_defense--policy--disable_js_insert.md)
- [bot_defense.policy.disable_mobile_sdk](resources--http_loadbalancer--properties--bot_defense--policy--disable_mobile_sdk.md)
- [bot_defense.policy.js_insert_all_pages](resources--http_loadbalancer--properties--bot_defense--policy--js_insert_all_pages.md)
- [bot_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--properties--bot_defense--policy--js_insert_all_pages_except.md)
- [bot_defense.policy.js_insertion_rules](resources--http_loadbalancer--properties--bot_defense--policy--js_insertion_rules.md)
- [bot_defense.policy.mobile_sdk_config](resources--http_loadbalancer--properties--bot_defense--policy--mobile_sdk_config.md)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md)
- [bot_defense](resources--http_loadbalancer--properties--bot_defense.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
