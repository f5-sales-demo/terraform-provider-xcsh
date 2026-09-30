---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 78945, "body_sha256": "sha256:39eaf1edaaf1bc924de82a76ea6baa59f261fe73534b656ce260b9b430bfbdc5", "canonical_id": "xcsh-docs:resources:protected_application:reference", "child_ids": ["xcsh-docs:resources:protected_application:properties:adobe_commerce_connector", "xcsh-docs:resources:protected_application:properties:big_ip_iapp", "xcsh-docs:resources:protected_application:properties:cloudflare", "xcsh-docs:resources:protected_application:properties:cloudfront", "xcsh-docs:resources:protected_application:properties:custom_connector", "xcsh-docs:resources:protected_application:properties:f5_big_ip", "xcsh-docs:resources:protected_application:properties:salesforce_commerce_connector", "xcsh-docs:resources:protected_application:properties:timeouts"], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:reference", "parent_id": "xcsh-docs:resources:protected_application:fundamentals", "path": "docs/guides/resources--protected_application--reference.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md)
- Property reference

## Direct properties

- [adobe_commerce_connector](resources--protected_application--properties--adobe_commerce_connector.md): complete subsection reference.

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

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
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [big_ip_iapp](resources--protected_application--properties--big_ip_iapp.md): complete subsection reference.

- [cloudflare](resources--protected_application--properties--cloudflare.md): complete subsection reference.

- [cloudfront](resources--protected_application--properties--cloudfront.md): complete subsection reference.

- [custom_connector](resources--protected_application--properties--custom_connector.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="schema-disable"></a>

### disable property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

- [f5_big_ip](resources--protected_application--properties--f5_big_ip.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the Protected Application. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace where the Protected Application is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-region"></a>

### region property

Type: `"string"`. Optional, Computed.

\[Enum: US|EU|ASIA|CA\] Defines a selection for Bot Defense region - US: US United States of America
&#8203;- EU: EU European Union - ASIA: ASIA Asia - CA: CA Canada. Possible values are \`US\`, \`EU\`,
\`ASIA\`, \`CA\`. Defaults to \`US\`.

Upstream description:

Defines a selection for Bot Defense region

&#8203;- US: US

United States of America &#8203;- EU: EU

European Union &#8203;- ASIA: ASIA

Asia &#8203;- CA: CA

Canada.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("US",
    "EU",
    "ASIA",
    "CA"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "US",
  "enum": [
    "US",
    "EU",
    "ASIA",
    "CA"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [salesforce_commerce_connector](resources--protected_application--properties--salesforce_commerce_connector.md): complete subsection reference.

- [timeouts](resources--protected_application--properties--timeouts.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `adobe_commerce_connector` | [adobe_commerce_connector](resources--protected_application--properties--adobe_commerce_connector.md#section) |
| `annotations` | [annotations](resources--protected_application--reference.md#schema-annotations) |
| `big_ip_iapp` | [big_ip_iapp](resources--protected_application--properties--big_ip_iapp.md#section) |
| `cloudflare` | [cloudflare](resources--protected_application--properties--cloudflare.md#section) |
| `cloudflare.continue_mitigation_action_hdr` | [cloudflare.continue_mitigation_action_hdr](resources--protected_application--properties--cloudflare.md#schema-cloudflare--continue_mitigation_action_hdr) |
| `cloudflare.disable_js_insert` | [cloudflare.disable_js_insert](resources--protected_application--properties--cloudflare--disable_js_insert.md#section) |
| `cloudflare.disable_mobile_sdk` | [cloudflare.disable_mobile_sdk](resources--protected_application--properties--cloudflare--disable_mobile_sdk.md#section) |
| `cloudflare.js_insertion_rules` | [cloudflare.js_insertion_rules](resources--protected_application--properties--cloudflare--js_insertion_rules.md#section) |
| `cloudflare.js_insertion_rules.exclude_list` | [cloudflare.js_insertion_rules.exclude_list](resources--protected_application--properties--cloudflare--js_insertion_rules--exclude_list.md#section) |
| `cloudflare.js_insertion_rules.exclude_list.any_domain` | [cloudflare.js_insertion_rules.exclude_list.any_domain](resources--protected_application--properties--cloudflare--js_insertion_rules--exclude_list--any_domain.md#section) |
| `cloudflare.js_insertion_rules.exclude_list.domain` | [cloudflare.js_insertion_rules.exclude_list.domain](resources--protected_application--properties--cloudflare--js_insertion_rules--exclude_list--domain.md#section) |
| `cloudflare.js_insertion_rules.exclude_list.domain.exact_value` | [cloudflare.js_insertion_rules.exclude_list.domain.exact_value](resources--protected_application--properties--cloudflare--js_insertion_rules--exclude_list--domain.md#schema-cloudflare--js_insertion_rules--exclude_list--domain--exact_value) |
| `cloudflare.js_insertion_rules.exclude_list.domain.regex_value` | [cloudflare.js_insertion_rules.exclude_list.domain.regex_value](resources--protected_application--properties--cloudflare--js_insertion_rules--exclude_list--domain.md#schema-cloudflare--js_insertion_rules--exclude_list--domain--regex_value) |
| `cloudflare.js_insertion_rules.exclude_list.domain.suffix_value` | [cloudflare.js_insertion_rules.exclude_list.domain.suffix_value](resources--protected_application--properties--cloudflare--js_insertion_rules--exclude_list--domain.md#schema-cloudflare--js_insertion_rules--exclude_list--domain--suffix_value) |
| `cloudflare.js_insertion_rules.exclude_list.metadata` | [cloudflare.js_insertion_rules.exclude_list.metadata](resources--protected_application--properties--cloudflare--js_insertion_rules--exclude_list--metadata.md#section) |
| `cloudflare.js_insertion_rules.exclude_list.metadata.description_spec` | [cloudflare.js_insertion_rules.exclude_list.metadata.description_spec](resources--protected_application--properties--cloudflare--js_insertion_rules--exclude_list--metadata.md#schema-cloudflare--js_insertion_rules--exclude_list--metadata--description_spec) |
| `cloudflare.js_insertion_rules.exclude_list.metadata.name` | [cloudflare.js_insertion_rules.exclude_list.metadata.name](resources--protected_application--properties--cloudflare--js_insertion_rules--exclude_list--metadata.md#schema-cloudflare--js_insertion_rules--exclude_list--metadata--name) |
| `cloudflare.js_insertion_rules.exclude_list.path` | [cloudflare.js_insertion_rules.exclude_list.path](resources--protected_application--properties--cloudflare--js_insertion_rules--exclude_list--path.md#section) |
| `cloudflare.js_insertion_rules.exclude_list.path.path` | [cloudflare.js_insertion_rules.exclude_list.path.path](resources--protected_application--properties--cloudflare--js_insertion_rules--exclude_list--path.md#schema-cloudflare--js_insertion_rules--exclude_list--path--path) |
| `cloudflare.js_insertion_rules.exclude_list.path.prefix` | [cloudflare.js_insertion_rules.exclude_list.path.prefix](resources--protected_application--properties--cloudflare--js_insertion_rules--exclude_list--path.md#schema-cloudflare--js_insertion_rules--exclude_list--path--prefix) |
| `cloudflare.js_insertion_rules.exclude_list.path.regex` | [cloudflare.js_insertion_rules.exclude_list.path.regex](resources--protected_application--properties--cloudflare--js_insertion_rules--exclude_list--path.md#schema-cloudflare--js_insertion_rules--exclude_list--path--regex) |
| `cloudflare.js_insertion_rules.javascript_location` | [cloudflare.js_insertion_rules.javascript_location](resources--protected_application--properties--cloudflare--js_insertion_rules.md#schema-cloudflare--js_insertion_rules--javascript_location) |
| `cloudflare.js_insertion_rules.js_download_path` | [cloudflare.js_insertion_rules.js_download_path](resources--protected_application--properties--cloudflare--js_insertion_rules.md#schema-cloudflare--js_insertion_rules--js_download_path) |
| `cloudflare.js_insertion_rules.rules` | [cloudflare.js_insertion_rules.rules](resources--protected_application--properties--cloudflare--js_insertion_rules--rules.md#section) |
| `cloudflare.js_insertion_rules.rules.any_domain` | [cloudflare.js_insertion_rules.rules.any_domain](resources--protected_application--properties--cloudflare--js_insertion_rules--rules--any_domain.md#section) |
| `cloudflare.js_insertion_rules.rules.domain` | [cloudflare.js_insertion_rules.rules.domain](resources--protected_application--properties--cloudflare--js_insertion_rules--rules--domain.md#section) |
| `cloudflare.js_insertion_rules.rules.domain.exact_value` | [cloudflare.js_insertion_rules.rules.domain.exact_value](resources--protected_application--properties--cloudflare--js_insertion_rules--rules--domain.md#schema-cloudflare--js_insertion_rules--rules--domain--exact_value) |
| `cloudflare.js_insertion_rules.rules.domain.regex_value` | [cloudflare.js_insertion_rules.rules.domain.regex_value](resources--protected_application--properties--cloudflare--js_insertion_rules--rules--domain.md#schema-cloudflare--js_insertion_rules--rules--domain--regex_value) |
| `cloudflare.js_insertion_rules.rules.domain.suffix_value` | [cloudflare.js_insertion_rules.rules.domain.suffix_value](resources--protected_application--properties--cloudflare--js_insertion_rules--rules--domain.md#schema-cloudflare--js_insertion_rules--rules--domain--suffix_value) |
| `cloudflare.js_insertion_rules.rules.exact_path` | [cloudflare.js_insertion_rules.rules.exact_path](resources--protected_application--properties--cloudflare--js_insertion_rules--rules.md#schema-cloudflare--js_insertion_rules--rules--exact_path) |
| `cloudflare.js_insertion_rules.rules.glob` | [cloudflare.js_insertion_rules.rules.glob](resources--protected_application--properties--cloudflare--js_insertion_rules--rules.md#schema-cloudflare--js_insertion_rules--rules--glob) |
| `cloudflare.js_insertion_rules.rules.metadata` | [cloudflare.js_insertion_rules.rules.metadata](resources--protected_application--properties--cloudflare--js_insertion_rules--rules--metadata.md#section) |
| `cloudflare.js_insertion_rules.rules.metadata.description_spec` | [cloudflare.js_insertion_rules.rules.metadata.description_spec](resources--protected_application--properties--cloudflare--js_insertion_rules--rules--metadata.md#schema-cloudflare--js_insertion_rules--rules--metadata--description_spec) |
| `cloudflare.js_insertion_rules.rules.metadata.name` | [cloudflare.js_insertion_rules.rules.metadata.name](resources--protected_application--properties--cloudflare--js_insertion_rules--rules--metadata.md#schema-cloudflare--js_insertion_rules--rules--metadata--name) |
| `cloudflare.js_insertion_rules.rules.prefix` | [cloudflare.js_insertion_rules.rules.prefix](resources--protected_application--properties--cloudflare--js_insertion_rules--rules.md#schema-cloudflare--js_insertion_rules--rules--prefix) |
| `cloudflare.loglevel` | [cloudflare.loglevel](resources--protected_application--properties--cloudflare.md#schema-cloudflare--loglevel) |
| `cloudflare.manual_js_insert` | [cloudflare.manual_js_insert](resources--protected_application--properties--cloudflare--manual_js_insert.md#section) |
| `cloudflare.manual_js_insert.js_download_path` | [cloudflare.manual_js_insert.js_download_path](resources--protected_application--properties--cloudflare--manual_js_insert.md#schema-cloudflare--manual_js_insert--js_download_path) |
| `cloudflare.mobile_sdk_config` | [cloudflare.mobile_sdk_config](resources--protected_application--properties--cloudflare--mobile_sdk_config.md#section) |
| `cloudflare.mobile_sdk_config.mobile_identifier` | [cloudflare.mobile_sdk_config.mobile_identifier](resources--protected_application--properties--cloudflare--mobile_sdk_config--mobile_identifier.md#section) |
| `cloudflare.mobile_sdk_config.mobile_identifier.headers` | [cloudflare.mobile_sdk_config.mobile_identifier.headers](resources--protected_application--properties--cloudflare--mobile_sdk_config--mobile_identifier--headers.md#section) |
| `cloudflare.mobile_sdk_config.mobile_identifier.headers.exact` | [cloudflare.mobile_sdk_config.mobile_identifier.headers.exact](resources--protected_application--properties--cloudflare--mobile_sdk_config--mobile_identifier--headers.md#schema-cloudflare--mobile_sdk_config--mobile_identifier--headers--exact) |
| `cloudflare.mobile_sdk_config.mobile_identifier.headers.name` | [cloudflare.mobile_sdk_config.mobile_identifier.headers.name](resources--protected_application--properties--cloudflare--mobile_sdk_config--mobile_identifier--headers.md#schema-cloudflare--mobile_sdk_config--mobile_identifier--headers--name) |
| `cloudflare.mobile_sdk_config.mobile_identifier.headers.regex` | [cloudflare.mobile_sdk_config.mobile_identifier.headers.regex](resources--protected_application--properties--cloudflare--mobile_sdk_config--mobile_identifier--headers.md#schema-cloudflare--mobile_sdk_config--mobile_identifier--headers--regex) |
| `cloudflare.protected_endpoints` | [cloudflare.protected_endpoints](resources--protected_application--properties--cloudflare--protected_endpoints.md#section) |
| `cloudflare.protected_endpoints.any_domain` | [cloudflare.protected_endpoints.any_domain](resources--protected_application--properties--cloudflare--protected_endpoints--any_domain.md#section) |
| `cloudflare.protected_endpoints.domain` | [cloudflare.protected_endpoints.domain](resources--protected_application--properties--cloudflare--protected_endpoints--domain.md#section) |
| `cloudflare.protected_endpoints.domain.exact_value` | [cloudflare.protected_endpoints.domain.exact_value](resources--protected_application--properties--cloudflare--protected_endpoints--domain.md#schema-cloudflare--protected_endpoints--domain--exact_value) |
| `cloudflare.protected_endpoints.domain.regex_value` | [cloudflare.protected_endpoints.domain.regex_value](resources--protected_application--properties--cloudflare--protected_endpoints--domain.md#schema-cloudflare--protected_endpoints--domain--regex_value) |
| `cloudflare.protected_endpoints.domain.suffix_value` | [cloudflare.protected_endpoints.domain.suffix_value](resources--protected_application--properties--cloudflare--protected_endpoints--domain.md#schema-cloudflare--protected_endpoints--domain--suffix_value) |
| `cloudflare.protected_endpoints.http_methods` | [cloudflare.protected_endpoints.http_methods](resources--protected_application--properties--cloudflare--protected_endpoints.md#schema-cloudflare--protected_endpoints--http_methods) |
| `cloudflare.protected_endpoints.metadata` | [cloudflare.protected_endpoints.metadata](resources--protected_application--properties--cloudflare--protected_endpoints--metadata.md#section) |
| `cloudflare.protected_endpoints.metadata.description_spec` | [cloudflare.protected_endpoints.metadata.description_spec](resources--protected_application--properties--cloudflare--protected_endpoints--metadata.md#schema-cloudflare--protected_endpoints--metadata--description_spec) |
| `cloudflare.protected_endpoints.metadata.name` | [cloudflare.protected_endpoints.metadata.name](resources--protected_application--properties--cloudflare--protected_endpoints--metadata.md#schema-cloudflare--protected_endpoints--metadata--name) |
| `cloudflare.protected_endpoints.mobile_client` | [cloudflare.protected_endpoints.mobile_client](resources--protected_application--properties--cloudflare--protected_endpoints--mobile_client.md#section) |
| `cloudflare.protected_endpoints.mobile_client.block` | [cloudflare.protected_endpoints.mobile_client.block](resources--protected_application--properties--cloudflare--protected_endpoints--mobile_client--block.md#section) |
| `cloudflare.protected_endpoints.mobile_client.block.body` | [cloudflare.protected_endpoints.mobile_client.block.body](resources--protected_application--properties--cloudflare--protected_endpoints--mobile_client--block.md#schema-cloudflare--protected_endpoints--mobile_client--block--body) |
| `cloudflare.protected_endpoints.mobile_client.block.content_type` | [cloudflare.protected_endpoints.mobile_client.block.content_type](resources--protected_application--properties--cloudflare--protected_endpoints--mobile_client--block.md#schema-cloudflare--protected_endpoints--mobile_client--block--content_type) |
| `cloudflare.protected_endpoints.mobile_client.block.status` | [cloudflare.protected_endpoints.mobile_client.block.status](resources--protected_application--properties--cloudflare--protected_endpoints--mobile_client--block.md#schema-cloudflare--protected_endpoints--mobile_client--block--status) |
| `cloudflare.protected_endpoints.mobile_client.continue` | [cloudflare.protected_endpoints.mobile_client.continue](resources--protected_application--properties--cloudflare--protected_endpoints--mobile_client--continue.md#section) |
| `cloudflare.protected_endpoints.mobile_client.continue.add_header` | [cloudflare.protected_endpoints.mobile_client.continue.add_header](resources--protected_application--properties--cloudflare--protected_endpoints--mobile_client--continue--add_header.md#section) |
| `cloudflare.protected_endpoints.mobile_client.continue.no_header` | [cloudflare.protected_endpoints.mobile_client.continue.no_header](resources--protected_application--properties--cloudflare--protected_endpoints--mobile_client--continue--no_header.md#section) |
| `cloudflare.protected_endpoints.path` | [cloudflare.protected_endpoints.path](resources--protected_application--properties--cloudflare--protected_endpoints--path.md#section) |
| `cloudflare.protected_endpoints.path.caseinsensitive` | [cloudflare.protected_endpoints.path.caseinsensitive](resources--protected_application--properties--cloudflare--protected_endpoints--path.md#schema-cloudflare--protected_endpoints--path--caseinsensitive) |
| `cloudflare.protected_endpoints.path.path` | [cloudflare.protected_endpoints.path.path](resources--protected_application--properties--cloudflare--protected_endpoints--path.md#schema-cloudflare--protected_endpoints--path--path) |
| `cloudflare.protected_endpoints.query` | [cloudflare.protected_endpoints.query](resources--protected_application--properties--cloudflare--protected_endpoints.md#schema-cloudflare--protected_endpoints--query) |
| `cloudflare.protected_endpoints.web_client` | [cloudflare.protected_endpoints.web_client](resources--protected_application--properties--cloudflare--protected_endpoints--web_client.md#section) |
| `cloudflare.protected_endpoints.web_client.block` | [cloudflare.protected_endpoints.web_client.block](resources--protected_application--properties--cloudflare--protected_endpoints--web_client--block.md#section) |
| `cloudflare.protected_endpoints.web_client.block.body` | [cloudflare.protected_endpoints.web_client.block.body](resources--protected_application--properties--cloudflare--protected_endpoints--web_client--block.md#schema-cloudflare--protected_endpoints--web_client--block--body) |
| `cloudflare.protected_endpoints.web_client.block.content_type` | [cloudflare.protected_endpoints.web_client.block.content_type](resources--protected_application--properties--cloudflare--protected_endpoints--web_client--block.md#schema-cloudflare--protected_endpoints--web_client--block--content_type) |
| `cloudflare.protected_endpoints.web_client.block.status` | [cloudflare.protected_endpoints.web_client.block.status](resources--protected_application--properties--cloudflare--protected_endpoints--web_client--block.md#schema-cloudflare--protected_endpoints--web_client--block--status) |
| `cloudflare.protected_endpoints.web_client.continue` | [cloudflare.protected_endpoints.web_client.continue](resources--protected_application--properties--cloudflare--protected_endpoints--web_client--continue.md#section) |
| `cloudflare.protected_endpoints.web_client.continue.add_header` | [cloudflare.protected_endpoints.web_client.continue.add_header](resources--protected_application--properties--cloudflare--protected_endpoints--web_client--continue--add_header.md#section) |
| `cloudflare.protected_endpoints.web_client.continue.no_header` | [cloudflare.protected_endpoints.web_client.continue.no_header](resources--protected_application--properties--cloudflare--protected_endpoints--web_client--continue--no_header.md#section) |
| `cloudflare.protected_endpoints.web_client.redirect` | [cloudflare.protected_endpoints.web_client.redirect](resources--protected_application--properties--cloudflare--protected_endpoints--web_client--redirect.md#section) |
| `cloudflare.protected_endpoints.web_client.redirect.location` | [cloudflare.protected_endpoints.web_client.redirect.location](resources--protected_application--properties--cloudflare--protected_endpoints--web_client--redirect.md#schema-cloudflare--protected_endpoints--web_client--redirect--location) |
| `cloudflare.protected_endpoints.web_client.redirect.status` | [cloudflare.protected_endpoints.web_client.redirect.status](resources--protected_application--properties--cloudflare--protected_endpoints--web_client--redirect.md#schema-cloudflare--protected_endpoints--web_client--redirect--status) |
| `cloudflare.protected_endpoints.web_mobile_client` | [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client.md#section) |
| `cloudflare.protected_endpoints.web_mobile_client.block_mobile` | [cloudflare.protected_endpoints.web_mobile_client.block_mobile](resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--block_mobile.md#section) |
| `cloudflare.protected_endpoints.web_mobile_client.block_mobile.body` | [cloudflare.protected_endpoints.web_mobile_client.block_mobile.body](resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--block_mobile.md#schema-cloudflare--protected_endpoints--web_mobile_client--block_mobile--body) |
| `cloudflare.protected_endpoints.web_mobile_client.block_mobile.content_type` | [cloudflare.protected_endpoints.web_mobile_client.block_mobile.content_type](resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--block_mobile.md#schema-cloudflare--protected_endpoints--web_mobile_client--block_mobile--content_type) |
| `cloudflare.protected_endpoints.web_mobile_client.block_mobile.status` | [cloudflare.protected_endpoints.web_mobile_client.block_mobile.status](resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--block_mobile.md#schema-cloudflare--protected_endpoints--web_mobile_client--block_mobile--status) |
| `cloudflare.protected_endpoints.web_mobile_client.block_web` | [cloudflare.protected_endpoints.web_mobile_client.block_web](resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--block_web.md#section) |
| `cloudflare.protected_endpoints.web_mobile_client.block_web.body` | [cloudflare.protected_endpoints.web_mobile_client.block_web.body](resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--block_web.md#schema-cloudflare--protected_endpoints--web_mobile_client--block_web--body) |
| `cloudflare.protected_endpoints.web_mobile_client.block_web.content_type` | [cloudflare.protected_endpoints.web_mobile_client.block_web.content_type](resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--block_web.md#schema-cloudflare--protected_endpoints--web_mobile_client--block_web--content_type) |
| `cloudflare.protected_endpoints.web_mobile_client.block_web.status` | [cloudflare.protected_endpoints.web_mobile_client.block_web.status](resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--block_web.md#schema-cloudflare--protected_endpoints--web_mobile_client--block_web--status) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_mobile` | [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--continue_mobile.md#section) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header` | [cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header](resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--continue_mobile--add_header.md#section) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header` | [cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header](resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--continue_mobile--no_header.md#section) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_web` | [cloudflare.protected_endpoints.web_mobile_client.continue_web](resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--continue_web.md#section) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header` | [cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header](resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--continue_web--add_header.md#section) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header` | [cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header](resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--continue_web--no_header.md#section) |
| `cloudflare.protected_endpoints.web_mobile_client.redirect_web` | [cloudflare.protected_endpoints.web_mobile_client.redirect_web](resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--redirect_web.md#section) |
| `cloudflare.protected_endpoints.web_mobile_client.redirect_web.location` | [cloudflare.protected_endpoints.web_mobile_client.redirect_web.location](resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--redirect_web.md#schema-cloudflare--protected_endpoints--web_mobile_client--redirect_web--location) |
| `cloudflare.protected_endpoints.web_mobile_client.redirect_web.status` | [cloudflare.protected_endpoints.web_mobile_client.redirect_web.status](resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--redirect_web.md#schema-cloudflare--protected_endpoints--web_mobile_client--redirect_web--status) |
| `cloudflare.timeout` | [cloudflare.timeout](resources--protected_application--properties--cloudflare.md#schema-cloudflare--timeout) |
| `cloudflare.trusted_clients` | [cloudflare.trusted_clients](resources--protected_application--properties--cloudflare--trusted_clients.md#section) |
| `cloudflare.trusted_clients.http_header` | [cloudflare.trusted_clients.http_header](resources--protected_application--properties--cloudflare--trusted_clients--http_header.md#section) |
| `cloudflare.trusted_clients.http_header.headers` | [cloudflare.trusted_clients.http_header.headers](resources--protected_application--properties--cloudflare--trusted_clients--http_header--headers.md#section) |
| `cloudflare.trusted_clients.http_header.headers.exact` | [cloudflare.trusted_clients.http_header.headers.exact](resources--protected_application--properties--cloudflare--trusted_clients--http_header--headers.md#schema-cloudflare--trusted_clients--http_header--headers--exact) |
| `cloudflare.trusted_clients.http_header.headers.name` | [cloudflare.trusted_clients.http_header.headers.name](resources--protected_application--properties--cloudflare--trusted_clients--http_header--headers.md#schema-cloudflare--trusted_clients--http_header--headers--name) |
| `cloudflare.trusted_clients.http_header.headers.regex` | [cloudflare.trusted_clients.http_header.headers.regex](resources--protected_application--properties--cloudflare--trusted_clients--http_header--headers.md#schema-cloudflare--trusted_clients--http_header--headers--regex) |
| `cloudflare.trusted_clients.ip_prefix` | [cloudflare.trusted_clients.ip_prefix](resources--protected_application--properties--cloudflare--trusted_clients.md#schema-cloudflare--trusted_clients--ip_prefix) |
| `cloudflare.trusted_clients.metadata` | [cloudflare.trusted_clients.metadata](resources--protected_application--properties--cloudflare--trusted_clients--metadata.md#section) |
| `cloudflare.trusted_clients.metadata.description_spec` | [cloudflare.trusted_clients.metadata.description_spec](resources--protected_application--properties--cloudflare--trusted_clients--metadata.md#schema-cloudflare--trusted_clients--metadata--description_spec) |
| `cloudflare.trusted_clients.metadata.name` | [cloudflare.trusted_clients.metadata.name](resources--protected_application--properties--cloudflare--trusted_clients--metadata.md#schema-cloudflare--trusted_clients--metadata--name) |
| `cloudfront` | [cloudfront](resources--protected_application--properties--cloudfront.md#section) |
| `cloudfront.aws_configuration_id_selector` | [cloudfront.aws_configuration_id_selector](resources--protected_application--properties--cloudfront--aws_configuration_id_selector.md#section) |
| `cloudfront.aws_configuration_id_selector.ids` | [cloudfront.aws_configuration_id_selector.ids](resources--protected_application--properties--cloudfront--aws_configuration_id_selector.md#schema-cloudfront--aws_configuration_id_selector--ids) |
| `cloudfront.aws_configuration_tag_selector` | [cloudfront.aws_configuration_tag_selector](resources--protected_application--properties--cloudfront--aws_configuration_tag_selector.md#section) |
| `cloudfront.aws_configuration_tag_selector.tags` | [cloudfront.aws_configuration_tag_selector.tags](resources--protected_application--properties--cloudfront--aws_configuration_tag_selector.md#schema-cloudfront--aws_configuration_tag_selector--tags) |
| `cloudfront.continue_mitigation_action_hdr` | [cloudfront.continue_mitigation_action_hdr](resources--protected_application--properties--cloudfront.md#schema-cloudfront--continue_mitigation_action_hdr) |
| `cloudfront.data_sample` | [cloudfront.data_sample](resources--protected_application--properties--cloudfront.md#schema-cloudfront--data_sample) |
| `cloudfront.disable_aws_configuration` | [cloudfront.disable_aws_configuration](resources--protected_application--properties--cloudfront--disable_aws_configuration.md#section) |
| `cloudfront.disable_js_insert` | [cloudfront.disable_js_insert](resources--protected_application--properties--cloudfront--disable_js_insert.md#section) |
| `cloudfront.disable_mobile_sdk` | [cloudfront.disable_mobile_sdk](resources--protected_application--properties--cloudfront--disable_mobile_sdk.md#section) |
| `cloudfront.js_insertion_rules` | [cloudfront.js_insertion_rules](resources--protected_application--properties--cloudfront--js_insertion_rules.md#section) |
| `cloudfront.js_insertion_rules.exclude_list` | [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--properties--cloudfront--js_insertion_rules--exclude_list.md#section) |
| `cloudfront.js_insertion_rules.exclude_list.any_domain` | [cloudfront.js_insertion_rules.exclude_list.any_domain](resources--protected_application--properties--cloudfront--js_insertion_rules--exclude_list--any_domain.md#section) |
| `cloudfront.js_insertion_rules.exclude_list.domain` | [cloudfront.js_insertion_rules.exclude_list.domain](resources--protected_application--properties--cloudfront--js_insertion_rules--exclude_list--domain.md#section) |
| `cloudfront.js_insertion_rules.exclude_list.domain.exact_value` | [cloudfront.js_insertion_rules.exclude_list.domain.exact_value](resources--protected_application--properties--cloudfront--js_insertion_rules--exclude_list--domain.md#schema-cloudfront--js_insertion_rules--exclude_list--domain--exact_value) |
| `cloudfront.js_insertion_rules.exclude_list.domain.regex_value` | [cloudfront.js_insertion_rules.exclude_list.domain.regex_value](resources--protected_application--properties--cloudfront--js_insertion_rules--exclude_list--domain.md#schema-cloudfront--js_insertion_rules--exclude_list--domain--regex_value) |
| `cloudfront.js_insertion_rules.exclude_list.domain.suffix_value` | [cloudfront.js_insertion_rules.exclude_list.domain.suffix_value](resources--protected_application--properties--cloudfront--js_insertion_rules--exclude_list--domain.md#schema-cloudfront--js_insertion_rules--exclude_list--domain--suffix_value) |
| `cloudfront.js_insertion_rules.exclude_list.metadata` | [cloudfront.js_insertion_rules.exclude_list.metadata](resources--protected_application--properties--cloudfront--js_insertion_rules--exclude_list--metadata.md#section) |
| `cloudfront.js_insertion_rules.exclude_list.metadata.description_spec` | [cloudfront.js_insertion_rules.exclude_list.metadata.description_spec](resources--protected_application--properties--cloudfront--js_insertion_rules--exclude_list--metadata.md#schema-cloudfront--js_insertion_rules--exclude_list--metadata--description_spec) |
| `cloudfront.js_insertion_rules.exclude_list.metadata.name` | [cloudfront.js_insertion_rules.exclude_list.metadata.name](resources--protected_application--properties--cloudfront--js_insertion_rules--exclude_list--metadata.md#schema-cloudfront--js_insertion_rules--exclude_list--metadata--name) |
| `cloudfront.js_insertion_rules.exclude_list.path` | [cloudfront.js_insertion_rules.exclude_list.path](resources--protected_application--properties--cloudfront--js_insertion_rules--exclude_list--path.md#section) |
| `cloudfront.js_insertion_rules.exclude_list.path.path` | [cloudfront.js_insertion_rules.exclude_list.path.path](resources--protected_application--properties--cloudfront--js_insertion_rules--exclude_list--path.md#schema-cloudfront--js_insertion_rules--exclude_list--path--path) |
| `cloudfront.js_insertion_rules.exclude_list.path.prefix` | [cloudfront.js_insertion_rules.exclude_list.path.prefix](resources--protected_application--properties--cloudfront--js_insertion_rules--exclude_list--path.md#schema-cloudfront--js_insertion_rules--exclude_list--path--prefix) |
| `cloudfront.js_insertion_rules.exclude_list.path.regex` | [cloudfront.js_insertion_rules.exclude_list.path.regex](resources--protected_application--properties--cloudfront--js_insertion_rules--exclude_list--path.md#schema-cloudfront--js_insertion_rules--exclude_list--path--regex) |
| `cloudfront.js_insertion_rules.javascript_location` | [cloudfront.js_insertion_rules.javascript_location](resources--protected_application--properties--cloudfront--js_insertion_rules.md#schema-cloudfront--js_insertion_rules--javascript_location) |
| `cloudfront.js_insertion_rules.javascript_mode` | [cloudfront.js_insertion_rules.javascript_mode](resources--protected_application--properties--cloudfront--js_insertion_rules.md#schema-cloudfront--js_insertion_rules--javascript_mode) |
| `cloudfront.js_insertion_rules.js_download_path` | [cloudfront.js_insertion_rules.js_download_path](resources--protected_application--properties--cloudfront--js_insertion_rules.md#schema-cloudfront--js_insertion_rules--js_download_path) |
| `cloudfront.js_insertion_rules.rules` | [cloudfront.js_insertion_rules.rules](resources--protected_application--properties--cloudfront--js_insertion_rules--rules.md#section) |
| `cloudfront.js_insertion_rules.rules.any_domain` | [cloudfront.js_insertion_rules.rules.any_domain](resources--protected_application--properties--cloudfront--js_insertion_rules--rules--any_domain.md#section) |
| `cloudfront.js_insertion_rules.rules.domain` | [cloudfront.js_insertion_rules.rules.domain](resources--protected_application--properties--cloudfront--js_insertion_rules--rules--domain.md#section) |
| `cloudfront.js_insertion_rules.rules.domain.exact_value` | [cloudfront.js_insertion_rules.rules.domain.exact_value](resources--protected_application--properties--cloudfront--js_insertion_rules--rules--domain.md#schema-cloudfront--js_insertion_rules--rules--domain--exact_value) |
| `cloudfront.js_insertion_rules.rules.domain.regex_value` | [cloudfront.js_insertion_rules.rules.domain.regex_value](resources--protected_application--properties--cloudfront--js_insertion_rules--rules--domain.md#schema-cloudfront--js_insertion_rules--rules--domain--regex_value) |
| `cloudfront.js_insertion_rules.rules.domain.suffix_value` | [cloudfront.js_insertion_rules.rules.domain.suffix_value](resources--protected_application--properties--cloudfront--js_insertion_rules--rules--domain.md#schema-cloudfront--js_insertion_rules--rules--domain--suffix_value) |
| `cloudfront.js_insertion_rules.rules.exact_path` | [cloudfront.js_insertion_rules.rules.exact_path](resources--protected_application--properties--cloudfront--js_insertion_rules--rules.md#schema-cloudfront--js_insertion_rules--rules--exact_path) |
| `cloudfront.js_insertion_rules.rules.glob` | [cloudfront.js_insertion_rules.rules.glob](resources--protected_application--properties--cloudfront--js_insertion_rules--rules.md#schema-cloudfront--js_insertion_rules--rules--glob) |
| `cloudfront.js_insertion_rules.rules.metadata` | [cloudfront.js_insertion_rules.rules.metadata](resources--protected_application--properties--cloudfront--js_insertion_rules--rules--metadata.md#section) |
| `cloudfront.js_insertion_rules.rules.metadata.description_spec` | [cloudfront.js_insertion_rules.rules.metadata.description_spec](resources--protected_application--properties--cloudfront--js_insertion_rules--rules--metadata.md#schema-cloudfront--js_insertion_rules--rules--metadata--description_spec) |
| `cloudfront.js_insertion_rules.rules.metadata.name` | [cloudfront.js_insertion_rules.rules.metadata.name](resources--protected_application--properties--cloudfront--js_insertion_rules--rules--metadata.md#schema-cloudfront--js_insertion_rules--rules--metadata--name) |
| `cloudfront.js_insertion_rules.rules.prefix` | [cloudfront.js_insertion_rules.rules.prefix](resources--protected_application--properties--cloudfront--js_insertion_rules--rules.md#schema-cloudfront--js_insertion_rules--rules--prefix) |
| `cloudfront.loglevel` | [cloudfront.loglevel](resources--protected_application--properties--cloudfront.md#schema-cloudfront--loglevel) |
| `cloudfront.manual_js_insert` | [cloudfront.manual_js_insert](resources--protected_application--properties--cloudfront--manual_js_insert.md#section) |
| `cloudfront.manual_js_insert.javascript_mode` | [cloudfront.manual_js_insert.javascript_mode](resources--protected_application--properties--cloudfront--manual_js_insert.md#schema-cloudfront--manual_js_insert--javascript_mode) |
| `cloudfront.manual_js_insert.js_download_path` | [cloudfront.manual_js_insert.js_download_path](resources--protected_application--properties--cloudfront--manual_js_insert.md#schema-cloudfront--manual_js_insert--js_download_path) |
| `cloudfront.mobile_sdk_config` | [cloudfront.mobile_sdk_config](resources--protected_application--properties--cloudfront--mobile_sdk_config.md#section) |
| `cloudfront.mobile_sdk_config.mobile_identifier` | [cloudfront.mobile_sdk_config.mobile_identifier](resources--protected_application--properties--cloudfront--mobile_sdk_config--mobile_identifier.md#section) |
| `cloudfront.mobile_sdk_config.mobile_identifier.headers` | [cloudfront.mobile_sdk_config.mobile_identifier.headers](resources--protected_application--properties--cloudfront--mobile_sdk_config--mobile_identifier--headers.md#section) |
| `cloudfront.mobile_sdk_config.mobile_identifier.headers.exact` | [cloudfront.mobile_sdk_config.mobile_identifier.headers.exact](resources--protected_application--properties--cloudfront--mobile_sdk_config--mobile_identifier--headers.md#schema-cloudfront--mobile_sdk_config--mobile_identifier--headers--exact) |
| `cloudfront.mobile_sdk_config.mobile_identifier.headers.name` | [cloudfront.mobile_sdk_config.mobile_identifier.headers.name](resources--protected_application--properties--cloudfront--mobile_sdk_config--mobile_identifier--headers.md#schema-cloudfront--mobile_sdk_config--mobile_identifier--headers--name) |
| `cloudfront.mobile_sdk_config.mobile_identifier.headers.regex` | [cloudfront.mobile_sdk_config.mobile_identifier.headers.regex](resources--protected_application--properties--cloudfront--mobile_sdk_config--mobile_identifier--headers.md#schema-cloudfront--mobile_sdk_config--mobile_identifier--headers--regex) |
| `cloudfront.protected_endpoints` | [cloudfront.protected_endpoints](resources--protected_application--properties--cloudfront--protected_endpoints.md#section) |
| `cloudfront.protected_endpoints.any_domain` | [cloudfront.protected_endpoints.any_domain](resources--protected_application--properties--cloudfront--protected_endpoints--any_domain.md#section) |
| `cloudfront.protected_endpoints.domain` | [cloudfront.protected_endpoints.domain](resources--protected_application--properties--cloudfront--protected_endpoints--domain.md#section) |
| `cloudfront.protected_endpoints.domain.exact_value` | [cloudfront.protected_endpoints.domain.exact_value](resources--protected_application--properties--cloudfront--protected_endpoints--domain.md#schema-cloudfront--protected_endpoints--domain--exact_value) |
| `cloudfront.protected_endpoints.domain.regex_value` | [cloudfront.protected_endpoints.domain.regex_value](resources--protected_application--properties--cloudfront--protected_endpoints--domain.md#schema-cloudfront--protected_endpoints--domain--regex_value) |
| `cloudfront.protected_endpoints.domain.suffix_value` | [cloudfront.protected_endpoints.domain.suffix_value](resources--protected_application--properties--cloudfront--protected_endpoints--domain.md#schema-cloudfront--protected_endpoints--domain--suffix_value) |
| `cloudfront.protected_endpoints.flow_label` | [cloudfront.protected_endpoints.flow_label](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label.md#section) |
| `cloudfront.protected_endpoints.flow_label.account_management` | [cloudfront.protected_endpoints.flow_label.account_management](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--account_management.md#section) |
| `cloudfront.protected_endpoints.flow_label.account_management.create` | [cloudfront.protected_endpoints.flow_label.account_management.create](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--account_management--create.md#section) |
| `cloudfront.protected_endpoints.flow_label.account_management.password_reset` | [cloudfront.protected_endpoints.flow_label.account_management.password_reset](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--account_management--password_reset.md#section) |
| `cloudfront.protected_endpoints.flow_label.authentication` | [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication.md#section) |
| `cloudfront.protected_endpoints.flow_label.authentication.login` | [cloudfront.protected_endpoints.flow_label.authentication.login](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login.md#section) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result` | [cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login--disable_transaction_result.md#section) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login--transaction_result.md#section) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login--transaction_result--failure_conditions.md#section) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.name` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.name](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login--transaction_result--failure_conditions.md#schema-cloudfront--protected_endpoints--flow_label--authentication--login--transaction_result--failure_conditions--name) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.regex_values` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.regex_values](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login--transaction_result--failure_conditions.md#schema-cloudfront--protected_endpoints--flow_label--authentication--login--transaction_result--failure_conditions--regex_values) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.status` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.status](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login--transaction_result--failure_conditions.md#schema-cloudfront--protected_endpoints--flow_label--authentication--login--transaction_result--failure_conditions--status) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login--transaction_result--success_conditions.md#section) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.name` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.name](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login--transaction_result--success_conditions.md#schema-cloudfront--protected_endpoints--flow_label--authentication--login--transaction_result--success_conditions--name) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.regex_values` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.regex_values](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login--transaction_result--success_conditions.md#schema-cloudfront--protected_endpoints--flow_label--authentication--login--transaction_result--success_conditions--regex_values) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.status` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.status](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login--transaction_result--success_conditions.md#schema-cloudfront--protected_endpoints--flow_label--authentication--login--transaction_result--success_conditions--status) |
| `cloudfront.protected_endpoints.flow_label.authentication.login_mfa` | [cloudfront.protected_endpoints.flow_label.authentication.login_mfa](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login_mfa.md#section) |
| `cloudfront.protected_endpoints.flow_label.authentication.login_partner` | [cloudfront.protected_endpoints.flow_label.authentication.login_partner](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--login_partner.md#section) |
| `cloudfront.protected_endpoints.flow_label.authentication.logout` | [cloudfront.protected_endpoints.flow_label.authentication.logout](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--logout.md#section) |
| `cloudfront.protected_endpoints.flow_label.authentication.token_refresh` | [cloudfront.protected_endpoints.flow_label.authentication.token_refresh](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--authentication--token_refresh.md#section) |
| `cloudfront.protected_endpoints.flow_label.financial_services` | [cloudfront.protected_endpoints.flow_label.financial_services](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--financial_services.md#section) |
| `cloudfront.protected_endpoints.flow_label.financial_services.apply` | [cloudfront.protected_endpoints.flow_label.financial_services.apply](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--financial_services--apply.md#section) |
| `cloudfront.protected_endpoints.flow_label.financial_services.money_transfer` | [cloudfront.protected_endpoints.flow_label.financial_services.money_transfer](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--financial_services--money_transfer.md#section) |
| `cloudfront.protected_endpoints.flow_label.flight` | [cloudfront.protected_endpoints.flow_label.flight](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--flight.md#section) |
| `cloudfront.protected_endpoints.flow_label.flight.checkin` | [cloudfront.protected_endpoints.flow_label.flight.checkin](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--flight--checkin.md#section) |
| `cloudfront.protected_endpoints.flow_label.profile_management` | [cloudfront.protected_endpoints.flow_label.profile_management](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--profile_management.md#section) |
| `cloudfront.protected_endpoints.flow_label.profile_management.create` | [cloudfront.protected_endpoints.flow_label.profile_management.create](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--profile_management--create.md#section) |
| `cloudfront.protected_endpoints.flow_label.profile_management.update` | [cloudfront.protected_endpoints.flow_label.profile_management.update](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--profile_management--update.md#section) |
| `cloudfront.protected_endpoints.flow_label.profile_management.view` | [cloudfront.protected_endpoints.flow_label.profile_management.view](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--profile_management--view.md#section) |
| `cloudfront.protected_endpoints.flow_label.search` | [cloudfront.protected_endpoints.flow_label.search](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search.md#section) |
| `cloudfront.protected_endpoints.flow_label.search.flight_search` | [cloudfront.protected_endpoints.flow_label.search.flight_search](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search--flight_search.md#section) |
| `cloudfront.protected_endpoints.flow_label.search.product_search` | [cloudfront.protected_endpoints.flow_label.search.product_search](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search--product_search.md#section) |
| `cloudfront.protected_endpoints.flow_label.search.reservation_search` | [cloudfront.protected_endpoints.flow_label.search.reservation_search](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search--reservation_search.md#section) |
| `cloudfront.protected_endpoints.flow_label.search.room_search` | [cloudfront.protected_endpoints.flow_label.search.room_search](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search--room_search.md#section) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--shopping_gift_cards.md#section) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--shopping_gift_cards--gift_card_make_purchase_with_gift_card.md#section) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validation` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validation](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--shopping_gift_cards--gift_card_validation.md#section) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--shopping_gift_cards--shop_add_to_cart.md#section) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--shopping_gift_cards--shop_checkout.md#section) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--shopping_gift_cards--shop_choose_seat.md#section) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--shopping_gift_cards--shop_enter_drawing_submission.md#section) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--shopping_gift_cards--shop_make_payment.md#section) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--shopping_gift_cards--shop_order.md#section) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--shopping_gift_cards--shop_price_inquiry.md#section) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--shopping_gift_cards--shop_promo_code_validation.md#section) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--shopping_gift_cards--shop_purchase_gift_card.md#section) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quantity` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quantity](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--shopping_gift_cards--shop_update_quantity.md#section) |
| `cloudfront.protected_endpoints.http_methods` | [cloudfront.protected_endpoints.http_methods](resources--protected_application--properties--cloudfront--protected_endpoints.md#schema-cloudfront--protected_endpoints--http_methods) |
| `cloudfront.protected_endpoints.metadata` | [cloudfront.protected_endpoints.metadata](resources--protected_application--properties--cloudfront--protected_endpoints--metadata.md#section) |
| `cloudfront.protected_endpoints.metadata.description_spec` | [cloudfront.protected_endpoints.metadata.description_spec](resources--protected_application--properties--cloudfront--protected_endpoints--metadata.md#schema-cloudfront--protected_endpoints--metadata--description_spec) |
| `cloudfront.protected_endpoints.metadata.name` | [cloudfront.protected_endpoints.metadata.name](resources--protected_application--properties--cloudfront--protected_endpoints--metadata.md#schema-cloudfront--protected_endpoints--metadata--name) |
| `cloudfront.protected_endpoints.mobile_client` | [cloudfront.protected_endpoints.mobile_client](resources--protected_application--properties--cloudfront--protected_endpoints--mobile_client.md#section) |
| `cloudfront.protected_endpoints.mobile_client.block` | [cloudfront.protected_endpoints.mobile_client.block](resources--protected_application--properties--cloudfront--protected_endpoints--mobile_client--block.md#section) |
| `cloudfront.protected_endpoints.mobile_client.block.body` | [cloudfront.protected_endpoints.mobile_client.block.body](resources--protected_application--properties--cloudfront--protected_endpoints--mobile_client--block.md#schema-cloudfront--protected_endpoints--mobile_client--block--body) |
| `cloudfront.protected_endpoints.mobile_client.block.content_type` | [cloudfront.protected_endpoints.mobile_client.block.content_type](resources--protected_application--properties--cloudfront--protected_endpoints--mobile_client--block.md#schema-cloudfront--protected_endpoints--mobile_client--block--content_type) |
| `cloudfront.protected_endpoints.mobile_client.block.status` | [cloudfront.protected_endpoints.mobile_client.block.status](resources--protected_application--properties--cloudfront--protected_endpoints--mobile_client--block.md#schema-cloudfront--protected_endpoints--mobile_client--block--status) |
| `cloudfront.protected_endpoints.mobile_client.continue` | [cloudfront.protected_endpoints.mobile_client.continue](resources--protected_application--properties--cloudfront--protected_endpoints--mobile_client--continue.md#section) |
| `cloudfront.protected_endpoints.mobile_client.continue.add_header` | [cloudfront.protected_endpoints.mobile_client.continue.add_header](resources--protected_application--properties--cloudfront--protected_endpoints--mobile_client--continue--add_header.md#section) |
| `cloudfront.protected_endpoints.mobile_client.continue.no_header` | [cloudfront.protected_endpoints.mobile_client.continue.no_header](resources--protected_application--properties--cloudfront--protected_endpoints--mobile_client--continue--no_header.md#section) |
| `cloudfront.protected_endpoints.path` | [cloudfront.protected_endpoints.path](resources--protected_application--properties--cloudfront--protected_endpoints.md#schema-cloudfront--protected_endpoints--path) |
| `cloudfront.protected_endpoints.query` | [cloudfront.protected_endpoints.query](resources--protected_application--properties--cloudfront--protected_endpoints.md#schema-cloudfront--protected_endpoints--query) |
| `cloudfront.protected_endpoints.undefined_flow_label` | [cloudfront.protected_endpoints.undefined_flow_label](resources--protected_application--properties--cloudfront--protected_endpoints--undefined_flow_label.md#section) |
| `cloudfront.protected_endpoints.web_client` | [cloudfront.protected_endpoints.web_client](resources--protected_application--properties--cloudfront--protected_endpoints--web_client.md#section) |
| `cloudfront.protected_endpoints.web_client.block` | [cloudfront.protected_endpoints.web_client.block](resources--protected_application--properties--cloudfront--protected_endpoints--web_client--block.md#section) |
| `cloudfront.protected_endpoints.web_client.block.body` | [cloudfront.protected_endpoints.web_client.block.body](resources--protected_application--properties--cloudfront--protected_endpoints--web_client--block.md#schema-cloudfront--protected_endpoints--web_client--block--body) |
| `cloudfront.protected_endpoints.web_client.block.content_type` | [cloudfront.protected_endpoints.web_client.block.content_type](resources--protected_application--properties--cloudfront--protected_endpoints--web_client--block.md#schema-cloudfront--protected_endpoints--web_client--block--content_type) |
| `cloudfront.protected_endpoints.web_client.block.status` | [cloudfront.protected_endpoints.web_client.block.status](resources--protected_application--properties--cloudfront--protected_endpoints--web_client--block.md#schema-cloudfront--protected_endpoints--web_client--block--status) |
| `cloudfront.protected_endpoints.web_client.continue` | [cloudfront.protected_endpoints.web_client.continue](resources--protected_application--properties--cloudfront--protected_endpoints--web_client--continue.md#section) |
| `cloudfront.protected_endpoints.web_client.continue.add_header` | [cloudfront.protected_endpoints.web_client.continue.add_header](resources--protected_application--properties--cloudfront--protected_endpoints--web_client--continue--add_header.md#section) |
| `cloudfront.protected_endpoints.web_client.continue.no_header` | [cloudfront.protected_endpoints.web_client.continue.no_header](resources--protected_application--properties--cloudfront--protected_endpoints--web_client--continue--no_header.md#section) |
| `cloudfront.protected_endpoints.web_client.redirect` | [cloudfront.protected_endpoints.web_client.redirect](resources--protected_application--properties--cloudfront--protected_endpoints--web_client--redirect.md#section) |
| `cloudfront.protected_endpoints.web_client.redirect.location` | [cloudfront.protected_endpoints.web_client.redirect.location](resources--protected_application--properties--cloudfront--protected_endpoints--web_client--redirect.md#schema-cloudfront--protected_endpoints--web_client--redirect--location) |
| `cloudfront.protected_endpoints.web_client.redirect.status` | [cloudfront.protected_endpoints.web_client.redirect.status](resources--protected_application--properties--cloudfront--protected_endpoints--web_client--redirect.md#schema-cloudfront--protected_endpoints--web_client--redirect--status) |
| `cloudfront.protected_endpoints.web_mobile_client` | [cloudfront.protected_endpoints.web_mobile_client](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client.md#section) |
| `cloudfront.protected_endpoints.web_mobile_client.block_mobile` | [cloudfront.protected_endpoints.web_mobile_client.block_mobile](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--block_mobile.md#section) |
| `cloudfront.protected_endpoints.web_mobile_client.block_mobile.body` | [cloudfront.protected_endpoints.web_mobile_client.block_mobile.body](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--block_mobile.md#schema-cloudfront--protected_endpoints--web_mobile_client--block_mobile--body) |
| `cloudfront.protected_endpoints.web_mobile_client.block_mobile.content_type` | [cloudfront.protected_endpoints.web_mobile_client.block_mobile.content_type](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--block_mobile.md#schema-cloudfront--protected_endpoints--web_mobile_client--block_mobile--content_type) |
| `cloudfront.protected_endpoints.web_mobile_client.block_mobile.status` | [cloudfront.protected_endpoints.web_mobile_client.block_mobile.status](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--block_mobile.md#schema-cloudfront--protected_endpoints--web_mobile_client--block_mobile--status) |
| `cloudfront.protected_endpoints.web_mobile_client.block_web` | [cloudfront.protected_endpoints.web_mobile_client.block_web](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--block_web.md#section) |
| `cloudfront.protected_endpoints.web_mobile_client.block_web.body` | [cloudfront.protected_endpoints.web_mobile_client.block_web.body](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--block_web.md#schema-cloudfront--protected_endpoints--web_mobile_client--block_web--body) |
| `cloudfront.protected_endpoints.web_mobile_client.block_web.content_type` | [cloudfront.protected_endpoints.web_mobile_client.block_web.content_type](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--block_web.md#schema-cloudfront--protected_endpoints--web_mobile_client--block_web--content_type) |
| `cloudfront.protected_endpoints.web_mobile_client.block_web.status` | [cloudfront.protected_endpoints.web_mobile_client.block_web.status](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--block_web.md#schema-cloudfront--protected_endpoints--web_mobile_client--block_web--status) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_mobile` | [cloudfront.protected_endpoints.web_mobile_client.continue_mobile](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--continue_mobile.md#section) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_mobile.add_header` | [cloudfront.protected_endpoints.web_mobile_client.continue_mobile.add_header](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--continue_mobile--add_header.md#section) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_mobile.no_header` | [cloudfront.protected_endpoints.web_mobile_client.continue_mobile.no_header](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--continue_mobile--no_header.md#section) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_web` | [cloudfront.protected_endpoints.web_mobile_client.continue_web](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--continue_web.md#section) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_web.add_header` | [cloudfront.protected_endpoints.web_mobile_client.continue_web.add_header](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--continue_web--add_header.md#section) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_web.no_header` | [cloudfront.protected_endpoints.web_mobile_client.continue_web.no_header](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--continue_web--no_header.md#section) |
| `cloudfront.protected_endpoints.web_mobile_client.redirect_web` | [cloudfront.protected_endpoints.web_mobile_client.redirect_web](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--redirect_web.md#section) |
| `cloudfront.protected_endpoints.web_mobile_client.redirect_web.location` | [cloudfront.protected_endpoints.web_mobile_client.redirect_web.location](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--redirect_web.md#schema-cloudfront--protected_endpoints--web_mobile_client--redirect_web--location) |
| `cloudfront.protected_endpoints.web_mobile_client.redirect_web.status` | [cloudfront.protected_endpoints.web_mobile_client.redirect_web.status](resources--protected_application--properties--cloudfront--protected_endpoints--web_mobile_client--redirect_web.md#schema-cloudfront--protected_endpoints--web_mobile_client--redirect_web--status) |
| `cloudfront.timeout` | [cloudfront.timeout](resources--protected_application--properties--cloudfront.md#schema-cloudfront--timeout) |
| `cloudfront.trusted_clients` | [cloudfront.trusted_clients](resources--protected_application--properties--cloudfront--trusted_clients.md#section) |
| `cloudfront.trusted_clients.http_header` | [cloudfront.trusted_clients.http_header](resources--protected_application--properties--cloudfront--trusted_clients--http_header.md#section) |
| `cloudfront.trusted_clients.http_header.headers` | [cloudfront.trusted_clients.http_header.headers](resources--protected_application--properties--cloudfront--trusted_clients--http_header--headers.md#section) |
| `cloudfront.trusted_clients.http_header.headers.exact` | [cloudfront.trusted_clients.http_header.headers.exact](resources--protected_application--properties--cloudfront--trusted_clients--http_header--headers.md#schema-cloudfront--trusted_clients--http_header--headers--exact) |
| `cloudfront.trusted_clients.http_header.headers.name` | [cloudfront.trusted_clients.http_header.headers.name](resources--protected_application--properties--cloudfront--trusted_clients--http_header--headers.md#schema-cloudfront--trusted_clients--http_header--headers--name) |
| `cloudfront.trusted_clients.http_header.headers.regex` | [cloudfront.trusted_clients.http_header.headers.regex](resources--protected_application--properties--cloudfront--trusted_clients--http_header--headers.md#schema-cloudfront--trusted_clients--http_header--headers--regex) |
| `cloudfront.trusted_clients.ip_prefix` | [cloudfront.trusted_clients.ip_prefix](resources--protected_application--properties--cloudfront--trusted_clients.md#schema-cloudfront--trusted_clients--ip_prefix) |
| `cloudfront.trusted_clients.metadata` | [cloudfront.trusted_clients.metadata](resources--protected_application--properties--cloudfront--trusted_clients--metadata.md#section) |
| `cloudfront.trusted_clients.metadata.description_spec` | [cloudfront.trusted_clients.metadata.description_spec](resources--protected_application--properties--cloudfront--trusted_clients--metadata.md#schema-cloudfront--trusted_clients--metadata--description_spec) |
| `cloudfront.trusted_clients.metadata.name` | [cloudfront.trusted_clients.metadata.name](resources--protected_application--properties--cloudfront--trusted_clients--metadata.md#schema-cloudfront--trusted_clients--metadata--name) |
| `custom_connector` | [custom_connector](resources--protected_application--properties--custom_connector.md#section) |
| `description` | [description](resources--protected_application--reference.md#schema-description) |
| `disable` | [disable](resources--protected_application--reference.md#schema-disable) |
| `f5_big_ip` | [f5_big_ip](resources--protected_application--properties--f5_big_ip.md#section) |
| `id` | [id](resources--protected_application--reference.md#schema-id) |
| `labels` | [labels](resources--protected_application--reference.md#schema-labels) |
| `name` | [name](resources--protected_application--reference.md#schema-name) |
| `namespace` | [namespace](resources--protected_application--reference.md#schema-namespace) |
| `region` | [region](resources--protected_application--reference.md#schema-region) |
| `salesforce_commerce_connector` | [salesforce_commerce_connector](resources--protected_application--properties--salesforce_commerce_connector.md#section) |
| `timeouts` | [timeouts](resources--protected_application--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--protected_application--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--protected_application--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--protected_application--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--protected_application--properties--timeouts.md#schema-timeouts--update) |

## Next pages

- [adobe_commerce_connector](resources--protected_application--properties--adobe_commerce_connector.md)
- [big_ip_iapp](resources--protected_application--properties--big_ip_iapp.md)
- [cloudflare](resources--protected_application--properties--cloudflare.md)
- [cloudfront](resources--protected_application--properties--cloudfront.md)
- [custom_connector](resources--protected_application--properties--custom_connector.md)
- [f5_big_ip](resources--protected_application--properties--f5_big_ip.md)
- [salesforce_commerce_connector](resources--protected_application--properties--salesforce_commerce_connector.md)
- [timeouts](resources--protected_application--properties--timeouts.md)
- [xcsh_protected_application](../resources/protected_application.md)
