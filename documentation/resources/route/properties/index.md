---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_route."
xcsh_docs: {"aliases": ["route"], "body_bytes": 65519, "body_sha256": "sha256:c35b5bc0aba7907c7a006719422572bff5b84041e3be6a231dcb3dfe4038929a", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:route:properties:routes", "xcsh-docs:resources:route:properties:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:reference", "parent_id": "xcsh-docs:resources:route:fundamentals", "path": "documentation/resources/route/properties/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232", "registry_path": "docs/guides/resources--route--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:route:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:route:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:route:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:route:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:route:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:route:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:route:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes"], "anchor": "section", "description": "List of routes to match for incoming request.", "document_id": "xcsh-docs:resources:route:properties:routes", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:bot_defense_javascript_injection,inherited_bot_defense_javascript_injection", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:bot_defense_javascript_injection", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:bot_defense_javascript_injection,inherited_bot_defense_javascript_injection", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:inherited_bot_defense_javascript_injection", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:inherited_waf_exclusion,waf_exclusion_policy", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:inherited_waf_exclusion", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:route_destination,route_direct_response", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:route_destination,route_redirect", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_destination", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:route_destination,route_direct_response", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_direct_response", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:route_direct_response,route_redirect", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_direct_response", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:route_destination,route_redirect", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_redirect", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:route_direct_response,route_redirect", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:route_redirect", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes:ConflictingListObjectAttributes:inherited_waf_exclusion,waf_exclusion_policy", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:waf_exclusion_policy", "type": "conflicts"}], "schema_path": ["routes"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "lifecycle timeout", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:route:properties:timeouts", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Property reference for xcsh_route.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Additional upstream details:

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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

Additional upstream details:

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

Name of the Route. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Namespace where the Route is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/timeouts/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/#schema-disable) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/#schema-namespace) |
| `routes` | [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/#section) |
| `routes.bot_defense_javascript_injection` | [routes.bot_defense_javascript_injection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/bot_defense_javascript_injection/#section) |
| `routes.bot_defense_javascript_injection.javascript_location` | [routes.bot_defense_javascript_injection.javascript_location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/bot_defense_javascript_injection/#schema-routes--bot_defense_javascript_injection--javascript_location) |
| `routes.bot_defense_javascript_injection.javascript_tags` | [routes.bot_defense_javascript_injection.javascript_tags](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/bot_defense_javascript_injection/javascript_tags/#section) |
| `routes.bot_defense_javascript_injection.javascript_tags.javascript_url` | [routes.bot_defense_javascript_injection.javascript_tags.javascript_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/bot_defense_javascript_injection/javascript_tags/#schema-routes--bot_defense_javascript_injection--javascript_tags--javascript_url) |
| `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes` | [routes.bot_defense_javascript_injection.javascript_tags.tag_attributes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/bot_defense_javascript_injection/javascript_tags/tag_attributes/#section) |
| `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.javascript_tag` | [routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.javascript_tag](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/bot_defense_javascript_injection/javascript_tags/tag_attributes/#schema-routes--bot_defense_javascript_injection--javascript_tags--tag_attributes--javascript_tag) |
| `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.tag_value` | [routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.tag_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/bot_defense_javascript_injection/javascript_tags/tag_attributes/#schema-routes--bot_defense_javascript_injection--javascript_tags--tag_attributes--tag_value) |
| `routes.disable_location_add` | [routes.disable_location_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/#schema-routes--disable_location_add) |
| `routes.inherited_bot_defense_javascript_injection` | [routes.inherited_bot_defense_javascript_injection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/inherited_bot_defense_javascript_injection/#section) |
| `routes.inherited_waf_exclusion` | [routes.inherited_waf_exclusion](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/inherited_waf_exclusion/#section) |
| `routes.match` | [routes.match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/#section) |
| `routes.match.headers` | [routes.match.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/headers/#section) |
| `routes.match.headers.exact` | [routes.match.headers.exact](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/headers/#schema-routes--match--headers--exact) |
| `routes.match.headers.invert_match` | [routes.match.headers.invert_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/headers/#schema-routes--match--headers--invert_match) |
| `routes.match.headers.name` | [routes.match.headers.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/headers/#schema-routes--match--headers--name) |
| `routes.match.headers.presence` | [routes.match.headers.presence](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/headers/#schema-routes--match--headers--presence) |
| `routes.match.headers.regex` | [routes.match.headers.regex](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/headers/#schema-routes--match--headers--regex) |
| `routes.match.http_method` | [routes.match.http_method](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/#schema-routes--match--http_method) |
| `routes.match.incoming_port` | [routes.match.incoming_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/incoming_port/#section) |
| `routes.match.incoming_port.no_port_match` | [routes.match.incoming_port.no_port_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/incoming_port/no_port_match/#section) |
| `routes.match.incoming_port.port` | [routes.match.incoming_port.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/incoming_port/#schema-routes--match--incoming_port--port) |
| `routes.match.incoming_port.port_ranges` | [routes.match.incoming_port.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/incoming_port/#schema-routes--match--incoming_port--port_ranges) |
| `routes.match.path` | [routes.match.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/path/#section) |
| `routes.match.path.path` | [routes.match.path.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/path/#schema-routes--match--path--path) |
| `routes.match.path.prefix` | [routes.match.path.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/path/#schema-routes--match--path--prefix) |
| `routes.match.path.regex` | [routes.match.path.regex](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/path/#schema-routes--match--path--regex) |
| `routes.match.query_params` | [routes.match.query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/query_params/#section) |
| `routes.match.query_params.exact` | [routes.match.query_params.exact](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/query_params/#schema-routes--match--query_params--exact) |
| `routes.match.query_params.key` | [routes.match.query_params.key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/query_params/#schema-routes--match--query_params--key) |
| `routes.match.query_params.regex` | [routes.match.query_params.regex](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/query_params/#schema-routes--match--query_params--regex) |
| `routes.request_cookies_to_add` | [routes.request_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_cookies_to_add/#section) |
| `routes.request_cookies_to_add.name` | [routes.request_cookies_to_add.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_cookies_to_add/#schema-routes--request_cookies_to_add--name) |
| `routes.request_cookies_to_add.overwrite` | [routes.request_cookies_to_add.overwrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_cookies_to_add/#schema-routes--request_cookies_to_add--overwrite) |
| `routes.request_cookies_to_add.secret_value` | [routes.request_cookies_to_add.secret_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_cookies_to_add/secret_value/#section) |
| `routes.request_cookies_to_add.secret_value.blindfold_secret_info` | [routes.request_cookies_to_add.secret_value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_cookies_to_add/secret_value/blindfold_secret_info/#section) |
| `routes.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [routes.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_cookies_to_add/secret_value/blindfold_secret_info/#schema-routes--request_cookies_to_add--secret_value--blindfold_secret_info--decryption_provider) |
| `routes.request_cookies_to_add.secret_value.blindfold_secret_info.location` | [routes.request_cookies_to_add.secret_value.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_cookies_to_add/secret_value/blindfold_secret_info/#schema-routes--request_cookies_to_add--secret_value--blindfold_secret_info--location) |
| `routes.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [routes.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_cookies_to_add/secret_value/blindfold_secret_info/#schema-routes--request_cookies_to_add--secret_value--blindfold_secret_info--store_provider) |
| `routes.request_cookies_to_add.secret_value.clear_secret_info` | [routes.request_cookies_to_add.secret_value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_cookies_to_add/secret_value/clear_secret_info/#section) |
| `routes.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [routes.request_cookies_to_add.secret_value.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_cookies_to_add/secret_value/clear_secret_info/#schema-routes--request_cookies_to_add--secret_value--clear_secret_info--provider_ref) |
| `routes.request_cookies_to_add.secret_value.clear_secret_info.url` | [routes.request_cookies_to_add.secret_value.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_cookies_to_add/secret_value/clear_secret_info/#schema-routes--request_cookies_to_add--secret_value--clear_secret_info--url) |
| `routes.request_cookies_to_add.value` | [routes.request_cookies_to_add.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_cookies_to_add/#schema-routes--request_cookies_to_add--value) |
| `routes.request_cookies_to_remove` | [routes.request_cookies_to_remove](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/#schema-routes--request_cookies_to_remove) |
| `routes.request_headers_to_add` | [routes.request_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_headers_to_add/#section) |
| `routes.request_headers_to_add.append` | [routes.request_headers_to_add.append](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_headers_to_add/#schema-routes--request_headers_to_add--append) |
| `routes.request_headers_to_add.name` | [routes.request_headers_to_add.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_headers_to_add/#schema-routes--request_headers_to_add--name) |
| `routes.request_headers_to_add.secret_value` | [routes.request_headers_to_add.secret_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_headers_to_add/secret_value/#section) |
| `routes.request_headers_to_add.secret_value.blindfold_secret_info` | [routes.request_headers_to_add.secret_value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_headers_to_add/secret_value/blindfold_secret_info/#section) |
| `routes.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [routes.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_headers_to_add/secret_value/blindfold_secret_info/#schema-routes--request_headers_to_add--secret_value--blindfold_secret_info--decryption_provider) |
| `routes.request_headers_to_add.secret_value.blindfold_secret_info.location` | [routes.request_headers_to_add.secret_value.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_headers_to_add/secret_value/blindfold_secret_info/#schema-routes--request_headers_to_add--secret_value--blindfold_secret_info--location) |
| `routes.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [routes.request_headers_to_add.secret_value.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_headers_to_add/secret_value/blindfold_secret_info/#schema-routes--request_headers_to_add--secret_value--blindfold_secret_info--store_provider) |
| `routes.request_headers_to_add.secret_value.clear_secret_info` | [routes.request_headers_to_add.secret_value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_headers_to_add/secret_value/clear_secret_info/#section) |
| `routes.request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [routes.request_headers_to_add.secret_value.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_headers_to_add/secret_value/clear_secret_info/#schema-routes--request_headers_to_add--secret_value--clear_secret_info--provider_ref) |
| `routes.request_headers_to_add.secret_value.clear_secret_info.url` | [routes.request_headers_to_add.secret_value.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_headers_to_add/secret_value/clear_secret_info/#schema-routes--request_headers_to_add--secret_value--clear_secret_info--url) |
| `routes.request_headers_to_add.value` | [routes.request_headers_to_add.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/request_headers_to_add/#schema-routes--request_headers_to_add--value) |
| `routes.request_headers_to_remove` | [routes.request_headers_to_remove](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/#schema-routes--request_headers_to_remove) |
| `routes.response_cookies_to_add` | [routes.response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/#section) |
| `routes.response_cookies_to_add.add_domain` | [routes.response_cookies_to_add.add_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/#schema-routes--response_cookies_to_add--add_domain) |
| `routes.response_cookies_to_add.add_expiry` | [routes.response_cookies_to_add.add_expiry](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/#schema-routes--response_cookies_to_add--add_expiry) |
| `routes.response_cookies_to_add.add_httponly` | [routes.response_cookies_to_add.add_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/add_httponly/#section) |
| `routes.response_cookies_to_add.add_partitioned` | [routes.response_cookies_to_add.add_partitioned](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/add_partitioned/#section) |
| `routes.response_cookies_to_add.add_path` | [routes.response_cookies_to_add.add_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/#schema-routes--response_cookies_to_add--add_path) |
| `routes.response_cookies_to_add.add_secure` | [routes.response_cookies_to_add.add_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/add_secure/#section) |
| `routes.response_cookies_to_add.ignore_domain` | [routes.response_cookies_to_add.ignore_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/ignore_domain/#section) |
| `routes.response_cookies_to_add.ignore_expiry` | [routes.response_cookies_to_add.ignore_expiry](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/ignore_expiry/#section) |
| `routes.response_cookies_to_add.ignore_httponly` | [routes.response_cookies_to_add.ignore_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/ignore_httponly/#section) |
| `routes.response_cookies_to_add.ignore_max_age` | [routes.response_cookies_to_add.ignore_max_age](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/ignore_max_age/#section) |
| `routes.response_cookies_to_add.ignore_partitioned` | [routes.response_cookies_to_add.ignore_partitioned](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/ignore_partitioned/#section) |
| `routes.response_cookies_to_add.ignore_path` | [routes.response_cookies_to_add.ignore_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/ignore_path/#section) |
| `routes.response_cookies_to_add.ignore_samesite` | [routes.response_cookies_to_add.ignore_samesite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/ignore_samesite/#section) |
| `routes.response_cookies_to_add.ignore_secure` | [routes.response_cookies_to_add.ignore_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/ignore_secure/#section) |
| `routes.response_cookies_to_add.ignore_value` | [routes.response_cookies_to_add.ignore_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/ignore_value/#section) |
| `routes.response_cookies_to_add.max_age_value` | [routes.response_cookies_to_add.max_age_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/#schema-routes--response_cookies_to_add--max_age_value) |
| `routes.response_cookies_to_add.name` | [routes.response_cookies_to_add.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/#schema-routes--response_cookies_to_add--name) |
| `routes.response_cookies_to_add.overwrite` | [routes.response_cookies_to_add.overwrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/#schema-routes--response_cookies_to_add--overwrite) |
| `routes.response_cookies_to_add.samesite_lax` | [routes.response_cookies_to_add.samesite_lax](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/samesite_lax/#section) |
| `routes.response_cookies_to_add.samesite_none` | [routes.response_cookies_to_add.samesite_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/samesite_none/#section) |
| `routes.response_cookies_to_add.samesite_strict` | [routes.response_cookies_to_add.samesite_strict](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/samesite_strict/#section) |
| `routes.response_cookies_to_add.secret_value` | [routes.response_cookies_to_add.secret_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/secret_value/#section) |
| `routes.response_cookies_to_add.secret_value.blindfold_secret_info` | [routes.response_cookies_to_add.secret_value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/secret_value/blindfold_secret_info/#section) |
| `routes.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [routes.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/secret_value/blindfold_secret_info/#schema-routes--response_cookies_to_add--secret_value--blindfold_secret_info--decryption_provider) |
| `routes.response_cookies_to_add.secret_value.blindfold_secret_info.location` | [routes.response_cookies_to_add.secret_value.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/secret_value/blindfold_secret_info/#schema-routes--response_cookies_to_add--secret_value--blindfold_secret_info--location) |
| `routes.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [routes.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/secret_value/blindfold_secret_info/#schema-routes--response_cookies_to_add--secret_value--blindfold_secret_info--store_provider) |
| `routes.response_cookies_to_add.secret_value.clear_secret_info` | [routes.response_cookies_to_add.secret_value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/secret_value/clear_secret_info/#section) |
| `routes.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [routes.response_cookies_to_add.secret_value.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/secret_value/clear_secret_info/#schema-routes--response_cookies_to_add--secret_value--clear_secret_info--provider_ref) |
| `routes.response_cookies_to_add.secret_value.clear_secret_info.url` | [routes.response_cookies_to_add.secret_value.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/secret_value/clear_secret_info/#schema-routes--response_cookies_to_add--secret_value--clear_secret_info--url) |
| `routes.response_cookies_to_add.value` | [routes.response_cookies_to_add.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_cookies_to_add/#schema-routes--response_cookies_to_add--value) |
| `routes.response_cookies_to_remove` | [routes.response_cookies_to_remove](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/#schema-routes--response_cookies_to_remove) |
| `routes.response_headers_to_add` | [routes.response_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_headers_to_add/#section) |
| `routes.response_headers_to_add.append` | [routes.response_headers_to_add.append](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_headers_to_add/#schema-routes--response_headers_to_add--append) |
| `routes.response_headers_to_add.name` | [routes.response_headers_to_add.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_headers_to_add/#schema-routes--response_headers_to_add--name) |
| `routes.response_headers_to_add.secret_value` | [routes.response_headers_to_add.secret_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_headers_to_add/secret_value/#section) |
| `routes.response_headers_to_add.secret_value.blindfold_secret_info` | [routes.response_headers_to_add.secret_value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_headers_to_add/secret_value/blindfold_secret_info/#section) |
| `routes.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [routes.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_headers_to_add/secret_value/blindfold_secret_info/#schema-routes--response_headers_to_add--secret_value--blindfold_secret_info--decryption_provider) |
| `routes.response_headers_to_add.secret_value.blindfold_secret_info.location` | [routes.response_headers_to_add.secret_value.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_headers_to_add/secret_value/blindfold_secret_info/#schema-routes--response_headers_to_add--secret_value--blindfold_secret_info--location) |
| `routes.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [routes.response_headers_to_add.secret_value.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_headers_to_add/secret_value/blindfold_secret_info/#schema-routes--response_headers_to_add--secret_value--blindfold_secret_info--store_provider) |
| `routes.response_headers_to_add.secret_value.clear_secret_info` | [routes.response_headers_to_add.secret_value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_headers_to_add/secret_value/clear_secret_info/#section) |
| `routes.response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [routes.response_headers_to_add.secret_value.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_headers_to_add/secret_value/clear_secret_info/#schema-routes--response_headers_to_add--secret_value--clear_secret_info--provider_ref) |
| `routes.response_headers_to_add.secret_value.clear_secret_info.url` | [routes.response_headers_to_add.secret_value.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_headers_to_add/secret_value/clear_secret_info/#schema-routes--response_headers_to_add--secret_value--clear_secret_info--url) |
| `routes.response_headers_to_add.value` | [routes.response_headers_to_add.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/response_headers_to_add/#schema-routes--response_headers_to_add--value) |
| `routes.response_headers_to_remove` | [routes.response_headers_to_remove](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/#schema-routes--response_headers_to_remove) |
| `routes.route_destination` | [routes.route_destination](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/#section) |
| `routes.route_destination.auto_host_rewrite` | [routes.route_destination.auto_host_rewrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/#schema-routes--route_destination--auto_host_rewrite) |
| `routes.route_destination.buffer_policy` | [routes.route_destination.buffer_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/buffer_policy/#section) |
| `routes.route_destination.buffer_policy.disabled` | [routes.route_destination.buffer_policy.disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/buffer_policy/#schema-routes--route_destination--buffer_policy--disabled) |
| `routes.route_destination.buffer_policy.max_request_bytes` | [routes.route_destination.buffer_policy.max_request_bytes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/buffer_policy/#schema-routes--route_destination--buffer_policy--max_request_bytes) |
| `routes.route_destination.cors_policy` | [routes.route_destination.cors_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/cors_policy/#section) |
| `routes.route_destination.cors_policy.allow_credentials` | [routes.route_destination.cors_policy.allow_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/cors_policy/#schema-routes--route_destination--cors_policy--allow_credentials) |
| `routes.route_destination.cors_policy.allow_headers` | [routes.route_destination.cors_policy.allow_headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/cors_policy/#schema-routes--route_destination--cors_policy--allow_headers) |
| `routes.route_destination.cors_policy.allow_methods` | [routes.route_destination.cors_policy.allow_methods](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/cors_policy/#schema-routes--route_destination--cors_policy--allow_methods) |
| `routes.route_destination.cors_policy.allow_origin` | [routes.route_destination.cors_policy.allow_origin](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/cors_policy/#schema-routes--route_destination--cors_policy--allow_origin) |
| `routes.route_destination.cors_policy.allow_origin_regex` | [routes.route_destination.cors_policy.allow_origin_regex](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/cors_policy/#schema-routes--route_destination--cors_policy--allow_origin_regex) |
| `routes.route_destination.cors_policy.disabled` | [routes.route_destination.cors_policy.disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/cors_policy/#schema-routes--route_destination--cors_policy--disabled) |
| `routes.route_destination.cors_policy.expose_headers` | [routes.route_destination.cors_policy.expose_headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/cors_policy/#schema-routes--route_destination--cors_policy--expose_headers) |
| `routes.route_destination.cors_policy.maximum_age` | [routes.route_destination.cors_policy.maximum_age](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/cors_policy/#schema-routes--route_destination--cors_policy--maximum_age) |
| `routes.route_destination.csrf_policy` | [routes.route_destination.csrf_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/csrf_policy/#section) |
| `routes.route_destination.csrf_policy.all_load_balancer_domains` | [routes.route_destination.csrf_policy.all_load_balancer_domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/csrf_policy/all_load_balancer_domains/#section) |
| `routes.route_destination.csrf_policy.custom_domain_list` | [routes.route_destination.csrf_policy.custom_domain_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/csrf_policy/custom_domain_list/#section) |
| `routes.route_destination.csrf_policy.custom_domain_list.domains` | [routes.route_destination.csrf_policy.custom_domain_list.domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/csrf_policy/custom_domain_list/#schema-routes--route_destination--csrf_policy--custom_domain_list--domains) |
| `routes.route_destination.csrf_policy.disabled` | [routes.route_destination.csrf_policy.disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/csrf_policy/disabled/#section) |
| `routes.route_destination.destinations` | [routes.route_destination.destinations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/destinations/#section) |
| `routes.route_destination.destinations.cluster` | [routes.route_destination.destinations.cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/destinations/cluster/#section) |
| `routes.route_destination.destinations.cluster.kind` | [routes.route_destination.destinations.cluster.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/destinations/cluster/#schema-routes--route_destination--destinations--cluster--kind) |
| `routes.route_destination.destinations.cluster.name` | [routes.route_destination.destinations.cluster.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/destinations/cluster/#schema-routes--route_destination--destinations--cluster--name) |
| `routes.route_destination.destinations.cluster.namespace` | [routes.route_destination.destinations.cluster.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/destinations/cluster/#schema-routes--route_destination--destinations--cluster--namespace) |
| `routes.route_destination.destinations.cluster.tenant` | [routes.route_destination.destinations.cluster.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/destinations/cluster/#schema-routes--route_destination--destinations--cluster--tenant) |
| `routes.route_destination.destinations.cluster.uid` | [routes.route_destination.destinations.cluster.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/destinations/cluster/#schema-routes--route_destination--destinations--cluster--uid) |
| `routes.route_destination.destinations.endpoint_subsets` | [routes.route_destination.destinations.endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/destinations/endpoint_subsets/#section) |
| `routes.route_destination.destinations.priority` | [routes.route_destination.destinations.priority](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/destinations/#schema-routes--route_destination--destinations--priority) |
| `routes.route_destination.destinations.weight` | [routes.route_destination.destinations.weight](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/destinations/#schema-routes--route_destination--destinations--weight) |
| `routes.route_destination.do_not_retract_cluster` | [routes.route_destination.do_not_retract_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/do_not_retract_cluster/#section) |
| `routes.route_destination.endpoint_subsets` | [routes.route_destination.endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/endpoint_subsets/#section) |
| `routes.route_destination.hash_policy` | [routes.route_destination.hash_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/hash_policy/#section) |
| `routes.route_destination.hash_policy.cookie` | [routes.route_destination.hash_policy.cookie](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/hash_policy/cookie/#section) |
| `routes.route_destination.hash_policy.cookie.add_httponly` | [routes.route_destination.hash_policy.cookie.add_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/hash_policy/cookie/add_httponly/#section) |
| `routes.route_destination.hash_policy.cookie.add_secure` | [routes.route_destination.hash_policy.cookie.add_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/hash_policy/cookie/add_secure/#section) |
| `routes.route_destination.hash_policy.cookie.ignore_httponly` | [routes.route_destination.hash_policy.cookie.ignore_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/hash_policy/cookie/ignore_httponly/#section) |
| `routes.route_destination.hash_policy.cookie.ignore_samesite` | [routes.route_destination.hash_policy.cookie.ignore_samesite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/hash_policy/cookie/ignore_samesite/#section) |
| `routes.route_destination.hash_policy.cookie.ignore_secure` | [routes.route_destination.hash_policy.cookie.ignore_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/hash_policy/cookie/ignore_secure/#section) |
| `routes.route_destination.hash_policy.cookie.name` | [routes.route_destination.hash_policy.cookie.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/hash_policy/cookie/#schema-routes--route_destination--hash_policy--cookie--name) |
| `routes.route_destination.hash_policy.cookie.path` | [routes.route_destination.hash_policy.cookie.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/hash_policy/cookie/#schema-routes--route_destination--hash_policy--cookie--path) |
| `routes.route_destination.hash_policy.cookie.samesite_lax` | [routes.route_destination.hash_policy.cookie.samesite_lax](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/hash_policy/cookie/samesite_lax/#section) |
| `routes.route_destination.hash_policy.cookie.samesite_none` | [routes.route_destination.hash_policy.cookie.samesite_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/hash_policy/cookie/samesite_none/#section) |
| `routes.route_destination.hash_policy.cookie.samesite_strict` | [routes.route_destination.hash_policy.cookie.samesite_strict](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/hash_policy/cookie/samesite_strict/#section) |
| `routes.route_destination.hash_policy.cookie.ttl` | [routes.route_destination.hash_policy.cookie.ttl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/hash_policy/cookie/#schema-routes--route_destination--hash_policy--cookie--ttl) |
| `routes.route_destination.hash_policy.header_name` | [routes.route_destination.hash_policy.header_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/hash_policy/#schema-routes--route_destination--hash_policy--header_name) |
| `routes.route_destination.hash_policy.source_ip` | [routes.route_destination.hash_policy.source_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/hash_policy/#schema-routes--route_destination--hash_policy--source_ip) |
| `routes.route_destination.hash_policy.terminal` | [routes.route_destination.hash_policy.terminal](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/hash_policy/#schema-routes--route_destination--hash_policy--terminal) |
| `routes.route_destination.host_rewrite` | [routes.route_destination.host_rewrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/#schema-routes--route_destination--host_rewrite) |
| `routes.route_destination.mirror_policy` | [routes.route_destination.mirror_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/mirror_policy/#section) |
| `routes.route_destination.mirror_policy.cluster` | [routes.route_destination.mirror_policy.cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/mirror_policy/cluster/#section) |
| `routes.route_destination.mirror_policy.cluster.kind` | [routes.route_destination.mirror_policy.cluster.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/mirror_policy/cluster/#schema-routes--route_destination--mirror_policy--cluster--kind) |
| `routes.route_destination.mirror_policy.cluster.name` | [routes.route_destination.mirror_policy.cluster.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/mirror_policy/cluster/#schema-routes--route_destination--mirror_policy--cluster--name) |
| `routes.route_destination.mirror_policy.cluster.namespace` | [routes.route_destination.mirror_policy.cluster.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/mirror_policy/cluster/#schema-routes--route_destination--mirror_policy--cluster--namespace) |
| `routes.route_destination.mirror_policy.cluster.tenant` | [routes.route_destination.mirror_policy.cluster.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/mirror_policy/cluster/#schema-routes--route_destination--mirror_policy--cluster--tenant) |
| `routes.route_destination.mirror_policy.cluster.uid` | [routes.route_destination.mirror_policy.cluster.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/mirror_policy/cluster/#schema-routes--route_destination--mirror_policy--cluster--uid) |
| `routes.route_destination.mirror_policy.percent` | [routes.route_destination.mirror_policy.percent](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/mirror_policy/percent/#section) |
| `routes.route_destination.mirror_policy.percent.denominator` | [routes.route_destination.mirror_policy.percent.denominator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/mirror_policy/percent/#schema-routes--route_destination--mirror_policy--percent--denominator) |
| `routes.route_destination.mirror_policy.percent.numerator` | [routes.route_destination.mirror_policy.percent.numerator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/mirror_policy/percent/#schema-routes--route_destination--mirror_policy--percent--numerator) |
| `routes.route_destination.prefix_rewrite` | [routes.route_destination.prefix_rewrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/#schema-routes--route_destination--prefix_rewrite) |
| `routes.route_destination.priority` | [routes.route_destination.priority](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/#schema-routes--route_destination--priority) |
| `routes.route_destination.query_params` | [routes.route_destination.query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/query_params/#section) |
| `routes.route_destination.query_params.remove_all_params` | [routes.route_destination.query_params.remove_all_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/query_params/remove_all_params/#section) |
| `routes.route_destination.query_params.replace_params` | [routes.route_destination.query_params.replace_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/query_params/#schema-routes--route_destination--query_params--replace_params) |
| `routes.route_destination.query_params.retain_all_params` | [routes.route_destination.query_params.retain_all_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/query_params/retain_all_params/#section) |
| `routes.route_destination.regex_rewrite` | [routes.route_destination.regex_rewrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/regex_rewrite/#section) |
| `routes.route_destination.regex_rewrite.pattern` | [routes.route_destination.regex_rewrite.pattern](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/regex_rewrite/#schema-routes--route_destination--regex_rewrite--pattern) |
| `routes.route_destination.regex_rewrite.substitution` | [routes.route_destination.regex_rewrite.substitution](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/regex_rewrite/#schema-routes--route_destination--regex_rewrite--substitution) |
| `routes.route_destination.retract_cluster` | [routes.route_destination.retract_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/retract_cluster/#section) |
| `routes.route_destination.retry_policy` | [routes.route_destination.retry_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/retry_policy/#section) |
| `routes.route_destination.retry_policy.back_off` | [routes.route_destination.retry_policy.back_off](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/retry_policy/back_off/#section) |
| `routes.route_destination.retry_policy.back_off.base_interval` | [routes.route_destination.retry_policy.back_off.base_interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/retry_policy/back_off/#schema-routes--route_destination--retry_policy--back_off--base_interval) |
| `routes.route_destination.retry_policy.back_off.max_interval` | [routes.route_destination.retry_policy.back_off.max_interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/retry_policy/back_off/#schema-routes--route_destination--retry_policy--back_off--max_interval) |
| `routes.route_destination.retry_policy.num_retries` | [routes.route_destination.retry_policy.num_retries](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/retry_policy/#schema-routes--route_destination--retry_policy--num_retries) |
| `routes.route_destination.retry_policy.per_try_timeout` | [routes.route_destination.retry_policy.per_try_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/retry_policy/#schema-routes--route_destination--retry_policy--per_try_timeout) |
| `routes.route_destination.retry_policy.retriable_status_codes` | [routes.route_destination.retry_policy.retriable_status_codes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/retry_policy/#schema-routes--route_destination--retry_policy--retriable_status_codes) |
| `routes.route_destination.retry_policy.retry_condition` | [routes.route_destination.retry_policy.retry_condition](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/retry_policy/#schema-routes--route_destination--retry_policy--retry_condition) |
| `routes.route_destination.spdy_config` | [routes.route_destination.spdy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/spdy_config/#section) |
| `routes.route_destination.spdy_config.use_spdy` | [routes.route_destination.spdy_config.use_spdy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/spdy_config/#schema-routes--route_destination--spdy_config--use_spdy) |
| `routes.route_destination.timeout` | [routes.route_destination.timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/#schema-routes--route_destination--timeout) |
| `routes.route_destination.web_socket_config` | [routes.route_destination.web_socket_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/web_socket_config/#section) |
| `routes.route_destination.web_socket_config.use_websocket` | [routes.route_destination.web_socket_config.use_websocket](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_destination/web_socket_config/#schema-routes--route_destination--web_socket_config--use_websocket) |
| `routes.route_direct_response` | [routes.route_direct_response](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_direct_response/#section) |
| `routes.route_direct_response.response_body_encoded` | [routes.route_direct_response.response_body_encoded](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_direct_response/#schema-routes--route_direct_response--response_body_encoded) |
| `routes.route_direct_response.response_code` | [routes.route_direct_response.response_code](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_direct_response/#schema-routes--route_direct_response--response_code) |
| `routes.route_redirect` | [routes.route_redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_redirect/#section) |
| `routes.route_redirect.host_redirect` | [routes.route_redirect.host_redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_redirect/#schema-routes--route_redirect--host_redirect) |
| `routes.route_redirect.path_redirect` | [routes.route_redirect.path_redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_redirect/#schema-routes--route_redirect--path_redirect) |
| `routes.route_redirect.prefix_rewrite` | [routes.route_redirect.prefix_rewrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_redirect/#schema-routes--route_redirect--prefix_rewrite) |
| `routes.route_redirect.proto_redirect` | [routes.route_redirect.proto_redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_redirect/#schema-routes--route_redirect--proto_redirect) |
| `routes.route_redirect.remove_all_params` | [routes.route_redirect.remove_all_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_redirect/remove_all_params/#section) |
| `routes.route_redirect.replace_params` | [routes.route_redirect.replace_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_redirect/#schema-routes--route_redirect--replace_params) |
| `routes.route_redirect.response_code` | [routes.route_redirect.response_code](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_redirect/#schema-routes--route_redirect--response_code) |
| `routes.route_redirect.retain_all_params` | [routes.route_redirect.retain_all_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/route_redirect/retain_all_params/#section) |
| `routes.service_policy` | [routes.service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/service_policy/#section) |
| `routes.service_policy.disable_spec` | [routes.service_policy.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/service_policy/#schema-routes--service_policy--disable_spec) |
| `routes.waf_exclusion_policy` | [routes.waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_exclusion_policy/#section) |
| `routes.waf_exclusion_policy.name` | [routes.waf_exclusion_policy.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_exclusion_policy/#schema-routes--waf_exclusion_policy--name) |
| `routes.waf_exclusion_policy.namespace` | [routes.waf_exclusion_policy.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_exclusion_policy/#schema-routes--waf_exclusion_policy--namespace) |
| `routes.waf_exclusion_policy.tenant` | [routes.waf_exclusion_policy.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_exclusion_policy/#schema-routes--waf_exclusion_policy--tenant) |
| `routes.waf_type` | [routes.waf_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/#section) |
| `routes.waf_type.app_firewall` | [routes.waf_type.app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/app_firewall/#section) |
| `routes.waf_type.app_firewall.app_firewall` | [routes.waf_type.app_firewall.app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/app_firewall/app_firewall/#section) |
| `routes.waf_type.app_firewall.app_firewall.kind` | [routes.waf_type.app_firewall.app_firewall.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/app_firewall/app_firewall/#schema-routes--waf_type--app_firewall--app_firewall--kind) |
| `routes.waf_type.app_firewall.app_firewall.name` | [routes.waf_type.app_firewall.app_firewall.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/app_firewall/app_firewall/#schema-routes--waf_type--app_firewall--app_firewall--name) |
| `routes.waf_type.app_firewall.app_firewall.namespace` | [routes.waf_type.app_firewall.app_firewall.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/app_firewall/app_firewall/#schema-routes--waf_type--app_firewall--app_firewall--namespace) |
| `routes.waf_type.app_firewall.app_firewall.tenant` | [routes.waf_type.app_firewall.app_firewall.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/app_firewall/app_firewall/#schema-routes--waf_type--app_firewall--app_firewall--tenant) |
| `routes.waf_type.app_firewall.app_firewall.uid` | [routes.waf_type.app_firewall.app_firewall.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/app_firewall/app_firewall/#schema-routes--waf_type--app_firewall--app_firewall--uid) |
| `routes.waf_type.disable_waf` | [routes.waf_type.disable_waf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/disable_waf/#section) |
| `routes.waf_type.inherit_waf` | [routes.waf_type.inherit_waf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/waf_type/inherit_waf/#section) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/timeouts/#schema-timeouts--update) |
