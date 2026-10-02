---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_application_profiles."
xcsh_docs: {"aliases": ["application profiles"], "body_bytes": 113971, "body_sha256": "sha256:909fbd5b4996915d78e0097e7e40372f82d75a080e567612e1b48c13ddb58c33", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:application_profiles:properties:advanced_tcp_profile", "xcsh-docs:data-sources:application_profiles:properties:ddos_profile", "xcsh-docs:data-sources:application_profiles:properties:irules", "xcsh-docs:data-sources:application_profiles:properties:virtual_server"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:application_profiles:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:application_profiles:reference", "parent_id": "xcsh-docs:data-sources:application_profiles:fundamentals", "path": "documentation/data-sources/application_profiles/properties/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231", "registry_path": "docs/guides/data-sources--application_profiles--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["advanced tcp profile"], "anchor": "section", "description": "BIG-IP Advanced TCP Profile.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:advanced_tcp_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["advanced_tcp_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:application_profiles:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["ddos profile"], "anchor": "section", "description": "BIG-IP DDoS Protection Rules.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:ddos_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ddos_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:application_profiles:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:application_profiles:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["irules"], "anchor": "section", "description": "OPTIONS for attaching iRules to BIG-IP Proxy.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:irules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["irules"], "syntax": "attribute", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:application_profiles:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:application_profiles:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:application_profiles:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["virtual server"], "anchor": "section", "description": "Specifies configuration related to virtual server.", "document_id": "xcsh-docs:data-sources:application_profiles:properties:virtual_server", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["virtual_server"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/application_profiles/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_application_profiles.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["application_profilesCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
- Property reference

## Direct properties

- [advanced_tcp_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/advanced_tcp_profile/): complete subsection reference.

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

- [ddos_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/ddos_profile/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the ApplicationProfiles.

Upstream description:

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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [irules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/irules/): complete subsection reference.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

Name of the ApplicationProfiles.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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

Namespace where the ApplicationProfiles exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `advanced_tcp_profile` | [advanced_tcp_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/advanced_tcp_profile/#section) |
| `advanced_tcp_profile.disable_tcp_advanced_profile` | [advanced_tcp_profile.disable_tcp_advanced_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/advanced_tcp_profile/disable_tcp_advanced_profile/#section) |
| `advanced_tcp_profile.enable_tcp_advanced_profile` | [advanced_tcp_profile.enable_tcp_advanced_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/advanced_tcp_profile/enable_tcp_advanced_profile/#section) |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/#schema-annotations) |
| `ddos_profile` | [ddos_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/ddos_profile/#section) |
| `ddos_profile.disable_ddos_mitigation` | [ddos_profile.disable_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/ddos_profile/disable_ddos_mitigation/#section) |
| `ddos_profile.enable_ddos_mitigation` | [ddos_profile.enable_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/ddos_profile/enable_ddos_mitigation/#section) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/#schema-id) |
| `irules` | [irules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/irules/#section) |
| `irules.kind` | [irules.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/irules/#schema-irules--kind) |
| `irules.name` | [irules.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/irules/#schema-irules--name) |
| `irules.namespace` | [irules.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/irules/#schema-irules--namespace) |
| `irules.tenant` | [irules.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/irules/#schema-irules--tenant) |
| `irules.uid` | [irules.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/irules/#schema-irules--uid) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/#schema-namespace) |
| `virtual_server` | [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/#section) |
| `virtual_server.access_profile` | [virtual_server.access_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/access_profile/#section) |
| `virtual_server.access_profile.kind` | [virtual_server.access_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/access_profile/#schema-virtual_server--access_profile--kind) |
| `virtual_server.access_profile.name` | [virtual_server.access_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/access_profile/#schema-virtual_server--access_profile--name) |
| `virtual_server.access_profile.namespace` | [virtual_server.access_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/access_profile/#schema-virtual_server--access_profile--namespace) |
| `virtual_server.access_profile.tenant` | [virtual_server.access_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/access_profile/#schema-virtual_server--access_profile--tenant) |
| `virtual_server.access_profile.uid` | [virtual_server.access_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/access_profile/#schema-virtual_server--access_profile--uid) |
| `virtual_server.address_translation` | [virtual_server.address_translation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/address_translation/#section) |
| `virtual_server.address_translation.address_translation_disable` | [virtual_server.address_translation.address_translation_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/address_translation/address_translation_disable/#section) |
| `virtual_server.address_translation.address_translation_enable` | [virtual_server.address_translation.address_translation_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/address_translation/address_translation_enable/#section) |
| `virtual_server.auto_last_hop` | [virtual_server.auto_last_hop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/auto_last_hop/#section) |
| `virtual_server.auto_last_hop.auto_last_hop_default` | [virtual_server.auto_last_hop.auto_last_hop_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/auto_last_hop/auto_last_hop_default/#section) |
| `virtual_server.auto_last_hop.auto_last_hop_disable` | [virtual_server.auto_last_hop.auto_last_hop_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/auto_last_hop/auto_last_hop_disable/#section) |
| `virtual_server.auto_last_hop.auto_last_hop_enable` | [virtual_server.auto_last_hop.auto_last_hop_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/auto_last_hop/auto_last_hop_enable/#section) |
| `virtual_server.clone_pool_client` | [virtual_server.clone_pool_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/clone_pool_client/#section) |
| `virtual_server.clone_pool_client.kind` | [virtual_server.clone_pool_client.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/clone_pool_client/#schema-virtual_server--clone_pool_client--kind) |
| `virtual_server.clone_pool_client.name` | [virtual_server.clone_pool_client.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/clone_pool_client/#schema-virtual_server--clone_pool_client--name) |
| `virtual_server.clone_pool_client.namespace` | [virtual_server.clone_pool_client.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/clone_pool_client/#schema-virtual_server--clone_pool_client--namespace) |
| `virtual_server.clone_pool_client.tenant` | [virtual_server.clone_pool_client.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/clone_pool_client/#schema-virtual_server--clone_pool_client--tenant) |
| `virtual_server.clone_pool_client.uid` | [virtual_server.clone_pool_client.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/clone_pool_client/#schema-virtual_server--clone_pool_client--uid) |
| `virtual_server.clone_pool_server` | [virtual_server.clone_pool_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/clone_pool_server/#section) |
| `virtual_server.clone_pool_server.kind` | [virtual_server.clone_pool_server.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/clone_pool_server/#schema-virtual_server--clone_pool_server--kind) |
| `virtual_server.clone_pool_server.name` | [virtual_server.clone_pool_server.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/clone_pool_server/#schema-virtual_server--clone_pool_server--name) |
| `virtual_server.clone_pool_server.namespace` | [virtual_server.clone_pool_server.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/clone_pool_server/#schema-virtual_server--clone_pool_server--namespace) |
| `virtual_server.clone_pool_server.tenant` | [virtual_server.clone_pool_server.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/clone_pool_server/#schema-virtual_server--clone_pool_server--tenant) |
| `virtual_server.clone_pool_server.uid` | [virtual_server.clone_pool_server.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/clone_pool_server/#schema-virtual_server--clone_pool_server--uid) |
| `virtual_server.connection_limit` | [virtual_server.connection_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/#schema-virtual_server--connection_limit) |
| `virtual_server.connection_rate_limit` | [virtual_server.connection_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/#schema-virtual_server--connection_rate_limit) |
| `virtual_server.connection_rate_limit_mode` | [virtual_server.connection_rate_limit_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/#section) |
| `virtual_server.connection_rate_limit_mode.per_destination_address` | [virtual_server.connection_rate_limit_mode.per_destination_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_destination_address/#section) |
| `virtual_server.connection_rate_limit_mode.per_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_destination_address.destination_mask](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_destination_address/#schema-virtual_server--connection_rate_limit_mode--per_destination_address--destination_mask) |
| `virtual_server.connection_rate_limit_mode.per_source_address` | [virtual_server.connection_rate_limit_mode.per_source_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_source_address/#section) |
| `virtual_server.connection_rate_limit_mode.per_source_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_source_address.source_mask](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_source_address/#schema-virtual_server--connection_rate_limit_mode--per_source_address--source_mask) |
| `virtual_server.connection_rate_limit_mode.per_source_destination_address` | [virtual_server.connection_rate_limit_mode.per_source_destination_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_source_destination_address/#section) |
| `virtual_server.connection_rate_limit_mode.per_source_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_source_destination_address.destination_mask](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_source_destination_address/#schema-virtual_server--connection_rate_limit_mode--per_source_destination_address--destination_mask) |
| `virtual_server.connection_rate_limit_mode.per_source_destination_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_source_destination_address.source_mask](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_source_destination_address/#schema-virtual_server--connection_rate_limit_mode--per_source_destination_address--source_mask) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server` | [virtual_server.connection_rate_limit_mode.per_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server/#section) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address` | [virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server_destination_address/#section) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address.destination_mask](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server_destination_address/#schema-virtual_server--connection_rate_limit_mode--per_virtual_server_destination_address--destination_mask) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_address` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server_source_address/#section) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_address.source_mask](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server_source_address/#schema-virtual_server--connection_rate_limit_mode--per_virtual_server_source_address--source_mask) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server_source_destination_address/#section) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.destination_mask](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server_source_destination_address/#schema-virtual_server--connection_rate_limit_mode--per_virtual_server_source_destination_address--destination_mask) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.source_mask](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/connection_rate_limit_mode/per_virtual_server_source_destination_address/#schema-virtual_server--connection_rate_limit_mode--per_virtual_server_source_destination_address--source_mask) |
| `virtual_server.default_persistence_profile` | [virtual_server.default_persistence_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/default_persistence_profile/#section) |
| `virtual_server.default_persistence_profile.kind` | [virtual_server.default_persistence_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/default_persistence_profile/#schema-virtual_server--default_persistence_profile--kind) |
| `virtual_server.default_persistence_profile.name` | [virtual_server.default_persistence_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/default_persistence_profile/#schema-virtual_server--default_persistence_profile--name) |
| `virtual_server.default_persistence_profile.namespace` | [virtual_server.default_persistence_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/default_persistence_profile/#schema-virtual_server--default_persistence_profile--namespace) |
| `virtual_server.default_persistence_profile.tenant` | [virtual_server.default_persistence_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/default_persistence_profile/#schema-virtual_server--default_persistence_profile--tenant) |
| `virtual_server.default_persistence_profile.uid` | [virtual_server.default_persistence_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/default_persistence_profile/#schema-virtual_server--default_persistence_profile--uid) |
| `virtual_server.default_pool` | [virtual_server.default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/default_pool/#section) |
| `virtual_server.default_pool.kind` | [virtual_server.default_pool.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/default_pool/#schema-virtual_server--default_pool--kind) |
| `virtual_server.default_pool.name` | [virtual_server.default_pool.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/default_pool/#schema-virtual_server--default_pool--name) |
| `virtual_server.default_pool.namespace` | [virtual_server.default_pool.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/default_pool/#schema-virtual_server--default_pool--namespace) |
| `virtual_server.default_pool.tenant` | [virtual_server.default_pool.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/default_pool/#schema-virtual_server--default_pool--tenant) |
| `virtual_server.default_pool.uid` | [virtual_server.default_pool.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/default_pool/#schema-virtual_server--default_pool--uid) |
| `virtual_server.fallback_persistence_profile` | [virtual_server.fallback_persistence_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/fallback_persistence_profile/#section) |
| `virtual_server.fallback_persistence_profile.kind` | [virtual_server.fallback_persistence_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/fallback_persistence_profile/#schema-virtual_server--fallback_persistence_profile--kind) |
| `virtual_server.fallback_persistence_profile.name` | [virtual_server.fallback_persistence_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/fallback_persistence_profile/#schema-virtual_server--fallback_persistence_profile--name) |
| `virtual_server.fallback_persistence_profile.namespace` | [virtual_server.fallback_persistence_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/fallback_persistence_profile/#schema-virtual_server--fallback_persistence_profile--namespace) |
| `virtual_server.fallback_persistence_profile.tenant` | [virtual_server.fallback_persistence_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/fallback_persistence_profile/#schema-virtual_server--fallback_persistence_profile--tenant) |
| `virtual_server.fallback_persistence_profile.uid` | [virtual_server.fallback_persistence_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/fallback_persistence_profile/#schema-virtual_server--fallback_persistence_profile--uid) |
| `virtual_server.fix_profile` | [virtual_server.fix_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/fix_profile/#section) |
| `virtual_server.fix_profile.kind` | [virtual_server.fix_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/fix_profile/#schema-virtual_server--fix_profile--kind) |
| `virtual_server.fix_profile.name` | [virtual_server.fix_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/fix_profile/#schema-virtual_server--fix_profile--name) |
| `virtual_server.fix_profile.namespace` | [virtual_server.fix_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/fix_profile/#schema-virtual_server--fix_profile--namespace) |
| `virtual_server.fix_profile.tenant` | [virtual_server.fix_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/fix_profile/#schema-virtual_server--fix_profile--tenant) |
| `virtual_server.fix_profile.uid` | [virtual_server.fix_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/fix_profile/#schema-virtual_server--fix_profile--uid) |
| `virtual_server.http` | [virtual_server.http](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/#section) |
| `virtual_server.http.client_ssl_profile` | [virtual_server.http.client_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/client_ssl_profile/#section) |
| `virtual_server.http.client_ssl_profile.kind` | [virtual_server.http.client_ssl_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/client_ssl_profile/#schema-virtual_server--http--client_ssl_profile--kind) |
| `virtual_server.http.client_ssl_profile.name` | [virtual_server.http.client_ssl_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/client_ssl_profile/#schema-virtual_server--http--client_ssl_profile--name) |
| `virtual_server.http.client_ssl_profile.namespace` | [virtual_server.http.client_ssl_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/client_ssl_profile/#schema-virtual_server--http--client_ssl_profile--namespace) |
| `virtual_server.http.client_ssl_profile.tenant` | [virtual_server.http.client_ssl_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/client_ssl_profile/#schema-virtual_server--http--client_ssl_profile--tenant) |
| `virtual_server.http.client_ssl_profile.uid` | [virtual_server.http.client_ssl_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/client_ssl_profile/#schema-virtual_server--http--client_ssl_profile--uid) |
| `virtual_server.http.http2_client_profile` | [virtual_server.http.http2_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/http2_client_profile/#section) |
| `virtual_server.http.http2_client_profile.kind` | [virtual_server.http.http2_client_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/http2_client_profile/#schema-virtual_server--http--http2_client_profile--kind) |
| `virtual_server.http.http2_client_profile.name` | [virtual_server.http.http2_client_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/http2_client_profile/#schema-virtual_server--http--http2_client_profile--name) |
| `virtual_server.http.http2_client_profile.namespace` | [virtual_server.http.http2_client_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/http2_client_profile/#schema-virtual_server--http--http2_client_profile--namespace) |
| `virtual_server.http.http2_client_profile.tenant` | [virtual_server.http.http2_client_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/http2_client_profile/#schema-virtual_server--http--http2_client_profile--tenant) |
| `virtual_server.http.http2_client_profile.uid` | [virtual_server.http.http2_client_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/http2_client_profile/#schema-virtual_server--http--http2_client_profile--uid) |
| `virtual_server.http.http2_server_profile` | [virtual_server.http.http2_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/http2_server_profile/#section) |
| `virtual_server.http.http2_server_profile.kind` | [virtual_server.http.http2_server_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/http2_server_profile/#schema-virtual_server--http--http2_server_profile--kind) |
| `virtual_server.http.http2_server_profile.name` | [virtual_server.http.http2_server_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/http2_server_profile/#schema-virtual_server--http--http2_server_profile--name) |
| `virtual_server.http.http2_server_profile.namespace` | [virtual_server.http.http2_server_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/http2_server_profile/#schema-virtual_server--http--http2_server_profile--namespace) |
| `virtual_server.http.http2_server_profile.tenant` | [virtual_server.http.http2_server_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/http2_server_profile/#schema-virtual_server--http--http2_server_profile--tenant) |
| `virtual_server.http.http2_server_profile.uid` | [virtual_server.http.http2_server_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/http2_server_profile/#schema-virtual_server--http--http2_server_profile--uid) |
| `virtual_server.http.http_client_profile` | [virtual_server.http.http_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/http_client_profile/#section) |
| `virtual_server.http.http_client_profile.kind` | [virtual_server.http.http_client_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/http_client_profile/#schema-virtual_server--http--http_client_profile--kind) |
| `virtual_server.http.http_client_profile.name` | [virtual_server.http.http_client_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/http_client_profile/#schema-virtual_server--http--http_client_profile--name) |
| `virtual_server.http.http_client_profile.namespace` | [virtual_server.http.http_client_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/http_client_profile/#schema-virtual_server--http--http_client_profile--namespace) |
| `virtual_server.http.http_client_profile.tenant` | [virtual_server.http.http_client_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/http_client_profile/#schema-virtual_server--http--http_client_profile--tenant) |
| `virtual_server.http.http_client_profile.uid` | [virtual_server.http.http_client_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/http_client_profile/#schema-virtual_server--http--http_client_profile--uid) |
| `virtual_server.http.http_server_profile` | [virtual_server.http.http_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/http_server_profile/#section) |
| `virtual_server.http.http_server_profile.kind` | [virtual_server.http.http_server_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/http_server_profile/#schema-virtual_server--http--http_server_profile--kind) |
| `virtual_server.http.http_server_profile.name` | [virtual_server.http.http_server_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/http_server_profile/#schema-virtual_server--http--http_server_profile--name) |
| `virtual_server.http.http_server_profile.namespace` | [virtual_server.http.http_server_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/http_server_profile/#schema-virtual_server--http--http_server_profile--namespace) |
| `virtual_server.http.http_server_profile.tenant` | [virtual_server.http.http_server_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/http_server_profile/#schema-virtual_server--http--http_server_profile--tenant) |
| `virtual_server.http.http_server_profile.uid` | [virtual_server.http.http_server_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/http_server_profile/#schema-virtual_server--http--http_server_profile--uid) |
| `virtual_server.http.ocsp_profile` | [virtual_server.http.ocsp_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/ocsp_profile/#section) |
| `virtual_server.http.ocsp_profile.kind` | [virtual_server.http.ocsp_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/ocsp_profile/#schema-virtual_server--http--ocsp_profile--kind) |
| `virtual_server.http.ocsp_profile.name` | [virtual_server.http.ocsp_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/ocsp_profile/#schema-virtual_server--http--ocsp_profile--name) |
| `virtual_server.http.ocsp_profile.namespace` | [virtual_server.http.ocsp_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/ocsp_profile/#schema-virtual_server--http--ocsp_profile--namespace) |
| `virtual_server.http.ocsp_profile.tenant` | [virtual_server.http.ocsp_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/ocsp_profile/#schema-virtual_server--http--ocsp_profile--tenant) |
| `virtual_server.http.ocsp_profile.uid` | [virtual_server.http.ocsp_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/ocsp_profile/#schema-virtual_server--http--ocsp_profile--uid) |
| `virtual_server.http.server_ssl_profile` | [virtual_server.http.server_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/server_ssl_profile/#section) |
| `virtual_server.http.server_ssl_profile.kind` | [virtual_server.http.server_ssl_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/server_ssl_profile/#schema-virtual_server--http--server_ssl_profile--kind) |
| `virtual_server.http.server_ssl_profile.name` | [virtual_server.http.server_ssl_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/server_ssl_profile/#schema-virtual_server--http--server_ssl_profile--name) |
| `virtual_server.http.server_ssl_profile.namespace` | [virtual_server.http.server_ssl_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/server_ssl_profile/#schema-virtual_server--http--server_ssl_profile--namespace) |
| `virtual_server.http.server_ssl_profile.tenant` | [virtual_server.http.server_ssl_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/server_ssl_profile/#schema-virtual_server--http--server_ssl_profile--tenant) |
| `virtual_server.http.server_ssl_profile.uid` | [virtual_server.http.server_ssl_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/server_ssl_profile/#schema-virtual_server--http--server_ssl_profile--uid) |
| `virtual_server.http.stream_profile` | [virtual_server.http.stream_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/stream_profile/#section) |
| `virtual_server.http.stream_profile.kind` | [virtual_server.http.stream_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/stream_profile/#schema-virtual_server--http--stream_profile--kind) |
| `virtual_server.http.stream_profile.name` | [virtual_server.http.stream_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/stream_profile/#schema-virtual_server--http--stream_profile--name) |
| `virtual_server.http.stream_profile.namespace` | [virtual_server.http.stream_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/stream_profile/#schema-virtual_server--http--stream_profile--namespace) |
| `virtual_server.http.stream_profile.tenant` | [virtual_server.http.stream_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/stream_profile/#schema-virtual_server--http--stream_profile--tenant) |
| `virtual_server.http.stream_profile.uid` | [virtual_server.http.stream_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/stream_profile/#schema-virtual_server--http--stream_profile--uid) |
| `virtual_server.http.tcp_client_profile` | [virtual_server.http.tcp_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/tcp_client_profile/#section) |
| `virtual_server.http.tcp_client_profile.kind` | [virtual_server.http.tcp_client_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/tcp_client_profile/#schema-virtual_server--http--tcp_client_profile--kind) |
| `virtual_server.http.tcp_client_profile.name` | [virtual_server.http.tcp_client_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/tcp_client_profile/#schema-virtual_server--http--tcp_client_profile--name) |
| `virtual_server.http.tcp_client_profile.namespace` | [virtual_server.http.tcp_client_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/tcp_client_profile/#schema-virtual_server--http--tcp_client_profile--namespace) |
| `virtual_server.http.tcp_client_profile.tenant` | [virtual_server.http.tcp_client_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/tcp_client_profile/#schema-virtual_server--http--tcp_client_profile--tenant) |
| `virtual_server.http.tcp_client_profile.uid` | [virtual_server.http.tcp_client_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/tcp_client_profile/#schema-virtual_server--http--tcp_client_profile--uid) |
| `virtual_server.http.tcp_server_profile` | [virtual_server.http.tcp_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/tcp_server_profile/#section) |
| `virtual_server.http.tcp_server_profile.kind` | [virtual_server.http.tcp_server_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/tcp_server_profile/#schema-virtual_server--http--tcp_server_profile--kind) |
| `virtual_server.http.tcp_server_profile.name` | [virtual_server.http.tcp_server_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/tcp_server_profile/#schema-virtual_server--http--tcp_server_profile--name) |
| `virtual_server.http.tcp_server_profile.namespace` | [virtual_server.http.tcp_server_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/tcp_server_profile/#schema-virtual_server--http--tcp_server_profile--namespace) |
| `virtual_server.http.tcp_server_profile.tenant` | [virtual_server.http.tcp_server_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/tcp_server_profile/#schema-virtual_server--http--tcp_server_profile--tenant) |
| `virtual_server.http.tcp_server_profile.uid` | [virtual_server.http.tcp_server_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/tcp_server_profile/#schema-virtual_server--http--tcp_server_profile--uid) |
| `virtual_server.http.websocket_client_profile` | [virtual_server.http.websocket_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/websocket_client_profile/#section) |
| `virtual_server.http.websocket_client_profile.kind` | [virtual_server.http.websocket_client_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/websocket_client_profile/#schema-virtual_server--http--websocket_client_profile--kind) |
| `virtual_server.http.websocket_client_profile.name` | [virtual_server.http.websocket_client_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/websocket_client_profile/#schema-virtual_server--http--websocket_client_profile--name) |
| `virtual_server.http.websocket_client_profile.namespace` | [virtual_server.http.websocket_client_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/websocket_client_profile/#schema-virtual_server--http--websocket_client_profile--namespace) |
| `virtual_server.http.websocket_client_profile.tenant` | [virtual_server.http.websocket_client_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/websocket_client_profile/#schema-virtual_server--http--websocket_client_profile--tenant) |
| `virtual_server.http.websocket_client_profile.uid` | [virtual_server.http.websocket_client_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/websocket_client_profile/#schema-virtual_server--http--websocket_client_profile--uid) |
| `virtual_server.http.websocket_server_profile` | [virtual_server.http.websocket_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/websocket_server_profile/#section) |
| `virtual_server.http.websocket_server_profile.kind` | [virtual_server.http.websocket_server_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/websocket_server_profile/#schema-virtual_server--http--websocket_server_profile--kind) |
| `virtual_server.http.websocket_server_profile.name` | [virtual_server.http.websocket_server_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/websocket_server_profile/#schema-virtual_server--http--websocket_server_profile--name) |
| `virtual_server.http.websocket_server_profile.namespace` | [virtual_server.http.websocket_server_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/websocket_server_profile/#schema-virtual_server--http--websocket_server_profile--namespace) |
| `virtual_server.http.websocket_server_profile.tenant` | [virtual_server.http.websocket_server_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/websocket_server_profile/#schema-virtual_server--http--websocket_server_profile--tenant) |
| `virtual_server.http.websocket_server_profile.uid` | [virtual_server.http.websocket_server_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http/websocket_server_profile/#schema-virtual_server--http--websocket_server_profile--uid) |
| `virtual_server.http3` | [virtual_server.http3](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/#section) |
| `virtual_server.http3.client_ssl_profile` | [virtual_server.http3.client_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/client_ssl_profile/#section) |
| `virtual_server.http3.client_ssl_profile.kind` | [virtual_server.http3.client_ssl_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/client_ssl_profile/#schema-virtual_server--http3--client_ssl_profile--kind) |
| `virtual_server.http3.client_ssl_profile.name` | [virtual_server.http3.client_ssl_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/client_ssl_profile/#schema-virtual_server--http3--client_ssl_profile--name) |
| `virtual_server.http3.client_ssl_profile.namespace` | [virtual_server.http3.client_ssl_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/client_ssl_profile/#schema-virtual_server--http3--client_ssl_profile--namespace) |
| `virtual_server.http3.client_ssl_profile.tenant` | [virtual_server.http3.client_ssl_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/client_ssl_profile/#schema-virtual_server--http3--client_ssl_profile--tenant) |
| `virtual_server.http3.client_ssl_profile.uid` | [virtual_server.http3.client_ssl_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/client_ssl_profile/#schema-virtual_server--http3--client_ssl_profile--uid) |
| `virtual_server.http3.http3_profile` | [virtual_server.http3.http3_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/http3_profile/#section) |
| `virtual_server.http3.http3_profile.kind` | [virtual_server.http3.http3_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/http3_profile/#schema-virtual_server--http3--http3_profile--kind) |
| `virtual_server.http3.http3_profile.name` | [virtual_server.http3.http3_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/http3_profile/#schema-virtual_server--http3--http3_profile--name) |
| `virtual_server.http3.http3_profile.namespace` | [virtual_server.http3.http3_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/http3_profile/#schema-virtual_server--http3--http3_profile--namespace) |
| `virtual_server.http3.http3_profile.tenant` | [virtual_server.http3.http3_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/http3_profile/#schema-virtual_server--http3--http3_profile--tenant) |
| `virtual_server.http3.http3_profile.uid` | [virtual_server.http3.http3_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/http3_profile/#schema-virtual_server--http3--http3_profile--uid) |
| `virtual_server.http3.http_client_profile` | [virtual_server.http3.http_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/http_client_profile/#section) |
| `virtual_server.http3.http_client_profile.kind` | [virtual_server.http3.http_client_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/http_client_profile/#schema-virtual_server--http3--http_client_profile--kind) |
| `virtual_server.http3.http_client_profile.name` | [virtual_server.http3.http_client_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/http_client_profile/#schema-virtual_server--http3--http_client_profile--name) |
| `virtual_server.http3.http_client_profile.namespace` | [virtual_server.http3.http_client_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/http_client_profile/#schema-virtual_server--http3--http_client_profile--namespace) |
| `virtual_server.http3.http_client_profile.tenant` | [virtual_server.http3.http_client_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/http_client_profile/#schema-virtual_server--http3--http_client_profile--tenant) |
| `virtual_server.http3.http_client_profile.uid` | [virtual_server.http3.http_client_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/http_client_profile/#schema-virtual_server--http3--http_client_profile--uid) |
| `virtual_server.http3.http_server_profile` | [virtual_server.http3.http_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/http_server_profile/#section) |
| `virtual_server.http3.http_server_profile.kind` | [virtual_server.http3.http_server_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/http_server_profile/#schema-virtual_server--http3--http_server_profile--kind) |
| `virtual_server.http3.http_server_profile.name` | [virtual_server.http3.http_server_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/http_server_profile/#schema-virtual_server--http3--http_server_profile--name) |
| `virtual_server.http3.http_server_profile.namespace` | [virtual_server.http3.http_server_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/http_server_profile/#schema-virtual_server--http3--http_server_profile--namespace) |
| `virtual_server.http3.http_server_profile.tenant` | [virtual_server.http3.http_server_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/http_server_profile/#schema-virtual_server--http3--http_server_profile--tenant) |
| `virtual_server.http3.http_server_profile.uid` | [virtual_server.http3.http_server_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/http_server_profile/#schema-virtual_server--http3--http_server_profile--uid) |
| `virtual_server.http3.quic_profile` | [virtual_server.http3.quic_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/quic_profile/#section) |
| `virtual_server.http3.quic_profile.kind` | [virtual_server.http3.quic_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/quic_profile/#schema-virtual_server--http3--quic_profile--kind) |
| `virtual_server.http3.quic_profile.name` | [virtual_server.http3.quic_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/quic_profile/#schema-virtual_server--http3--quic_profile--name) |
| `virtual_server.http3.quic_profile.namespace` | [virtual_server.http3.quic_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/quic_profile/#schema-virtual_server--http3--quic_profile--namespace) |
| `virtual_server.http3.quic_profile.tenant` | [virtual_server.http3.quic_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/quic_profile/#schema-virtual_server--http3--quic_profile--tenant) |
| `virtual_server.http3.quic_profile.uid` | [virtual_server.http3.quic_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/quic_profile/#schema-virtual_server--http3--quic_profile--uid) |
| `virtual_server.http3.server_ssl_profile` | [virtual_server.http3.server_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/server_ssl_profile/#section) |
| `virtual_server.http3.server_ssl_profile.kind` | [virtual_server.http3.server_ssl_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/server_ssl_profile/#schema-virtual_server--http3--server_ssl_profile--kind) |
| `virtual_server.http3.server_ssl_profile.name` | [virtual_server.http3.server_ssl_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/server_ssl_profile/#schema-virtual_server--http3--server_ssl_profile--name) |
| `virtual_server.http3.server_ssl_profile.namespace` | [virtual_server.http3.server_ssl_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/server_ssl_profile/#schema-virtual_server--http3--server_ssl_profile--namespace) |
| `virtual_server.http3.server_ssl_profile.tenant` | [virtual_server.http3.server_ssl_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/server_ssl_profile/#schema-virtual_server--http3--server_ssl_profile--tenant) |
| `virtual_server.http3.server_ssl_profile.uid` | [virtual_server.http3.server_ssl_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/server_ssl_profile/#schema-virtual_server--http3--server_ssl_profile--uid) |
| `virtual_server.http3.tcp_server_profile` | [virtual_server.http3.tcp_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/tcp_server_profile/#section) |
| `virtual_server.http3.tcp_server_profile.kind` | [virtual_server.http3.tcp_server_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/tcp_server_profile/#schema-virtual_server--http3--tcp_server_profile--kind) |
| `virtual_server.http3.tcp_server_profile.name` | [virtual_server.http3.tcp_server_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/tcp_server_profile/#schema-virtual_server--http3--tcp_server_profile--name) |
| `virtual_server.http3.tcp_server_profile.namespace` | [virtual_server.http3.tcp_server_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/tcp_server_profile/#schema-virtual_server--http3--tcp_server_profile--namespace) |
| `virtual_server.http3.tcp_server_profile.tenant` | [virtual_server.http3.tcp_server_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/tcp_server_profile/#schema-virtual_server--http3--tcp_server_profile--tenant) |
| `virtual_server.http3.tcp_server_profile.uid` | [virtual_server.http3.tcp_server_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/tcp_server_profile/#schema-virtual_server--http3--tcp_server_profile--uid) |
| `virtual_server.http3.udp_client_profile` | [virtual_server.http3.udp_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/udp_client_profile/#section) |
| `virtual_server.http3.udp_client_profile.kind` | [virtual_server.http3.udp_client_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/udp_client_profile/#schema-virtual_server--http3--udp_client_profile--kind) |
| `virtual_server.http3.udp_client_profile.name` | [virtual_server.http3.udp_client_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/udp_client_profile/#schema-virtual_server--http3--udp_client_profile--name) |
| `virtual_server.http3.udp_client_profile.namespace` | [virtual_server.http3.udp_client_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/udp_client_profile/#schema-virtual_server--http3--udp_client_profile--namespace) |
| `virtual_server.http3.udp_client_profile.tenant` | [virtual_server.http3.udp_client_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/udp_client_profile/#schema-virtual_server--http3--udp_client_profile--tenant) |
| `virtual_server.http3.udp_client_profile.uid` | [virtual_server.http3.udp_client_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/udp_client_profile/#schema-virtual_server--http3--udp_client_profile--uid) |
| `virtual_server.http3.udp_server_profile` | [virtual_server.http3.udp_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/udp_server_profile/#section) |
| `virtual_server.http3.udp_server_profile.kind` | [virtual_server.http3.udp_server_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/udp_server_profile/#schema-virtual_server--http3--udp_server_profile--kind) |
| `virtual_server.http3.udp_server_profile.name` | [virtual_server.http3.udp_server_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/udp_server_profile/#schema-virtual_server--http3--udp_server_profile--name) |
| `virtual_server.http3.udp_server_profile.namespace` | [virtual_server.http3.udp_server_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/udp_server_profile/#schema-virtual_server--http3--udp_server_profile--namespace) |
| `virtual_server.http3.udp_server_profile.tenant` | [virtual_server.http3.udp_server_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/udp_server_profile/#schema-virtual_server--http3--udp_server_profile--tenant) |
| `virtual_server.http3.udp_server_profile.uid` | [virtual_server.http3.udp_server_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/http3/udp_server_profile/#schema-virtual_server--http3--udp_server_profile--uid) |
| `virtual_server.https` | [virtual_server.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/#section) |
| `virtual_server.https.client_ssl_profile` | [virtual_server.https.client_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/client_ssl_profile/#section) |
| `virtual_server.https.client_ssl_profile.kind` | [virtual_server.https.client_ssl_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/client_ssl_profile/#schema-virtual_server--https--client_ssl_profile--kind) |
| `virtual_server.https.client_ssl_profile.name` | [virtual_server.https.client_ssl_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/client_ssl_profile/#schema-virtual_server--https--client_ssl_profile--name) |
| `virtual_server.https.client_ssl_profile.namespace` | [virtual_server.https.client_ssl_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/client_ssl_profile/#schema-virtual_server--https--client_ssl_profile--namespace) |
| `virtual_server.https.client_ssl_profile.tenant` | [virtual_server.https.client_ssl_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/client_ssl_profile/#schema-virtual_server--https--client_ssl_profile--tenant) |
| `virtual_server.https.client_ssl_profile.uid` | [virtual_server.https.client_ssl_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/client_ssl_profile/#schema-virtual_server--https--client_ssl_profile--uid) |
| `virtual_server.https.http2_client_profile` | [virtual_server.https.http2_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/http2_client_profile/#section) |
| `virtual_server.https.http2_client_profile.kind` | [virtual_server.https.http2_client_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/http2_client_profile/#schema-virtual_server--https--http2_client_profile--kind) |
| `virtual_server.https.http2_client_profile.name` | [virtual_server.https.http2_client_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/http2_client_profile/#schema-virtual_server--https--http2_client_profile--name) |
| `virtual_server.https.http2_client_profile.namespace` | [virtual_server.https.http2_client_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/http2_client_profile/#schema-virtual_server--https--http2_client_profile--namespace) |
| `virtual_server.https.http2_client_profile.tenant` | [virtual_server.https.http2_client_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/http2_client_profile/#schema-virtual_server--https--http2_client_profile--tenant) |
| `virtual_server.https.http2_client_profile.uid` | [virtual_server.https.http2_client_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/http2_client_profile/#schema-virtual_server--https--http2_client_profile--uid) |
| `virtual_server.https.http2_server_profile` | [virtual_server.https.http2_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/http2_server_profile/#section) |
| `virtual_server.https.http2_server_profile.kind` | [virtual_server.https.http2_server_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/http2_server_profile/#schema-virtual_server--https--http2_server_profile--kind) |
| `virtual_server.https.http2_server_profile.name` | [virtual_server.https.http2_server_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/http2_server_profile/#schema-virtual_server--https--http2_server_profile--name) |
| `virtual_server.https.http2_server_profile.namespace` | [virtual_server.https.http2_server_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/http2_server_profile/#schema-virtual_server--https--http2_server_profile--namespace) |
| `virtual_server.https.http2_server_profile.tenant` | [virtual_server.https.http2_server_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/http2_server_profile/#schema-virtual_server--https--http2_server_profile--tenant) |
| `virtual_server.https.http2_server_profile.uid` | [virtual_server.https.http2_server_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/http2_server_profile/#schema-virtual_server--https--http2_server_profile--uid) |
| `virtual_server.https.http_client_profile` | [virtual_server.https.http_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/http_client_profile/#section) |
| `virtual_server.https.http_client_profile.kind` | [virtual_server.https.http_client_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/http_client_profile/#schema-virtual_server--https--http_client_profile--kind) |
| `virtual_server.https.http_client_profile.name` | [virtual_server.https.http_client_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/http_client_profile/#schema-virtual_server--https--http_client_profile--name) |
| `virtual_server.https.http_client_profile.namespace` | [virtual_server.https.http_client_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/http_client_profile/#schema-virtual_server--https--http_client_profile--namespace) |
| `virtual_server.https.http_client_profile.tenant` | [virtual_server.https.http_client_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/http_client_profile/#schema-virtual_server--https--http_client_profile--tenant) |
| `virtual_server.https.http_client_profile.uid` | [virtual_server.https.http_client_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/http_client_profile/#schema-virtual_server--https--http_client_profile--uid) |
| `virtual_server.https.http_server_profile` | [virtual_server.https.http_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/http_server_profile/#section) |
| `virtual_server.https.http_server_profile.kind` | [virtual_server.https.http_server_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/http_server_profile/#schema-virtual_server--https--http_server_profile--kind) |
| `virtual_server.https.http_server_profile.name` | [virtual_server.https.http_server_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/http_server_profile/#schema-virtual_server--https--http_server_profile--name) |
| `virtual_server.https.http_server_profile.namespace` | [virtual_server.https.http_server_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/http_server_profile/#schema-virtual_server--https--http_server_profile--namespace) |
| `virtual_server.https.http_server_profile.tenant` | [virtual_server.https.http_server_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/http_server_profile/#schema-virtual_server--https--http_server_profile--tenant) |
| `virtual_server.https.http_server_profile.uid` | [virtual_server.https.http_server_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/http_server_profile/#schema-virtual_server--https--http_server_profile--uid) |
| `virtual_server.https.ocsp_profile` | [virtual_server.https.ocsp_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/ocsp_profile/#section) |
| `virtual_server.https.ocsp_profile.kind` | [virtual_server.https.ocsp_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/ocsp_profile/#schema-virtual_server--https--ocsp_profile--kind) |
| `virtual_server.https.ocsp_profile.name` | [virtual_server.https.ocsp_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/ocsp_profile/#schema-virtual_server--https--ocsp_profile--name) |
| `virtual_server.https.ocsp_profile.namespace` | [virtual_server.https.ocsp_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/ocsp_profile/#schema-virtual_server--https--ocsp_profile--namespace) |
| `virtual_server.https.ocsp_profile.tenant` | [virtual_server.https.ocsp_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/ocsp_profile/#schema-virtual_server--https--ocsp_profile--tenant) |
| `virtual_server.https.ocsp_profile.uid` | [virtual_server.https.ocsp_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/ocsp_profile/#schema-virtual_server--https--ocsp_profile--uid) |
| `virtual_server.https.server_ssl_profile` | [virtual_server.https.server_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/server_ssl_profile/#section) |
| `virtual_server.https.server_ssl_profile.kind` | [virtual_server.https.server_ssl_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/server_ssl_profile/#schema-virtual_server--https--server_ssl_profile--kind) |
| `virtual_server.https.server_ssl_profile.name` | [virtual_server.https.server_ssl_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/server_ssl_profile/#schema-virtual_server--https--server_ssl_profile--name) |
| `virtual_server.https.server_ssl_profile.namespace` | [virtual_server.https.server_ssl_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/server_ssl_profile/#schema-virtual_server--https--server_ssl_profile--namespace) |
| `virtual_server.https.server_ssl_profile.tenant` | [virtual_server.https.server_ssl_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/server_ssl_profile/#schema-virtual_server--https--server_ssl_profile--tenant) |
| `virtual_server.https.server_ssl_profile.uid` | [virtual_server.https.server_ssl_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/server_ssl_profile/#schema-virtual_server--https--server_ssl_profile--uid) |
| `virtual_server.https.stream_profile` | [virtual_server.https.stream_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/stream_profile/#section) |
| `virtual_server.https.stream_profile.kind` | [virtual_server.https.stream_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/stream_profile/#schema-virtual_server--https--stream_profile--kind) |
| `virtual_server.https.stream_profile.name` | [virtual_server.https.stream_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/stream_profile/#schema-virtual_server--https--stream_profile--name) |
| `virtual_server.https.stream_profile.namespace` | [virtual_server.https.stream_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/stream_profile/#schema-virtual_server--https--stream_profile--namespace) |
| `virtual_server.https.stream_profile.tenant` | [virtual_server.https.stream_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/stream_profile/#schema-virtual_server--https--stream_profile--tenant) |
| `virtual_server.https.stream_profile.uid` | [virtual_server.https.stream_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/stream_profile/#schema-virtual_server--https--stream_profile--uid) |
| `virtual_server.https.tcp_client_profile` | [virtual_server.https.tcp_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/tcp_client_profile/#section) |
| `virtual_server.https.tcp_client_profile.kind` | [virtual_server.https.tcp_client_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/tcp_client_profile/#schema-virtual_server--https--tcp_client_profile--kind) |
| `virtual_server.https.tcp_client_profile.name` | [virtual_server.https.tcp_client_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/tcp_client_profile/#schema-virtual_server--https--tcp_client_profile--name) |
| `virtual_server.https.tcp_client_profile.namespace` | [virtual_server.https.tcp_client_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/tcp_client_profile/#schema-virtual_server--https--tcp_client_profile--namespace) |
| `virtual_server.https.tcp_client_profile.tenant` | [virtual_server.https.tcp_client_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/tcp_client_profile/#schema-virtual_server--https--tcp_client_profile--tenant) |
| `virtual_server.https.tcp_client_profile.uid` | [virtual_server.https.tcp_client_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/tcp_client_profile/#schema-virtual_server--https--tcp_client_profile--uid) |
| `virtual_server.https.tcp_server_profile` | [virtual_server.https.tcp_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/tcp_server_profile/#section) |
| `virtual_server.https.tcp_server_profile.kind` | [virtual_server.https.tcp_server_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/tcp_server_profile/#schema-virtual_server--https--tcp_server_profile--kind) |
| `virtual_server.https.tcp_server_profile.name` | [virtual_server.https.tcp_server_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/tcp_server_profile/#schema-virtual_server--https--tcp_server_profile--name) |
| `virtual_server.https.tcp_server_profile.namespace` | [virtual_server.https.tcp_server_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/tcp_server_profile/#schema-virtual_server--https--tcp_server_profile--namespace) |
| `virtual_server.https.tcp_server_profile.tenant` | [virtual_server.https.tcp_server_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/tcp_server_profile/#schema-virtual_server--https--tcp_server_profile--tenant) |
| `virtual_server.https.tcp_server_profile.uid` | [virtual_server.https.tcp_server_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/tcp_server_profile/#schema-virtual_server--https--tcp_server_profile--uid) |
| `virtual_server.https.websocket_client_profile` | [virtual_server.https.websocket_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/websocket_client_profile/#section) |
| `virtual_server.https.websocket_client_profile.kind` | [virtual_server.https.websocket_client_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/websocket_client_profile/#schema-virtual_server--https--websocket_client_profile--kind) |
| `virtual_server.https.websocket_client_profile.name` | [virtual_server.https.websocket_client_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/websocket_client_profile/#schema-virtual_server--https--websocket_client_profile--name) |
| `virtual_server.https.websocket_client_profile.namespace` | [virtual_server.https.websocket_client_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/websocket_client_profile/#schema-virtual_server--https--websocket_client_profile--namespace) |
| `virtual_server.https.websocket_client_profile.tenant` | [virtual_server.https.websocket_client_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/websocket_client_profile/#schema-virtual_server--https--websocket_client_profile--tenant) |
| `virtual_server.https.websocket_client_profile.uid` | [virtual_server.https.websocket_client_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/websocket_client_profile/#schema-virtual_server--https--websocket_client_profile--uid) |
| `virtual_server.https.websocket_server_profile` | [virtual_server.https.websocket_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/websocket_server_profile/#section) |
| `virtual_server.https.websocket_server_profile.kind` | [virtual_server.https.websocket_server_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/websocket_server_profile/#schema-virtual_server--https--websocket_server_profile--kind) |
| `virtual_server.https.websocket_server_profile.name` | [virtual_server.https.websocket_server_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/websocket_server_profile/#schema-virtual_server--https--websocket_server_profile--name) |
| `virtual_server.https.websocket_server_profile.namespace` | [virtual_server.https.websocket_server_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/websocket_server_profile/#schema-virtual_server--https--websocket_server_profile--namespace) |
| `virtual_server.https.websocket_server_profile.tenant` | [virtual_server.https.websocket_server_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/websocket_server_profile/#schema-virtual_server--https--websocket_server_profile--tenant) |
| `virtual_server.https.websocket_server_profile.uid` | [virtual_server.https.websocket_server_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/https/websocket_server_profile/#schema-virtual_server--https--websocket_server_profile--uid) |
| `virtual_server.immediate_action_on_service_down` | [virtual_server.immediate_action_on_service_down](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/immediate_action_on_service_down/#section) |
| `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop` | [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/immediate_action_on_service_down/immediate_action_on_service_down_drop/#section) |
| `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none` | [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/immediate_action_on_service_down/immediate_action_on_service_down_none/#section) |
| `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset` | [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/immediate_action_on_service_down/immediate_action_on_service_down_reset/#section) |
| `virtual_server.last_hop_pool` | [virtual_server.last_hop_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/last_hop_pool/#section) |
| `virtual_server.last_hop_pool.kind` | [virtual_server.last_hop_pool.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/last_hop_pool/#schema-virtual_server--last_hop_pool--kind) |
| `virtual_server.last_hop_pool.name` | [virtual_server.last_hop_pool.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/last_hop_pool/#schema-virtual_server--last_hop_pool--name) |
| `virtual_server.last_hop_pool.namespace` | [virtual_server.last_hop_pool.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/last_hop_pool/#schema-virtual_server--last_hop_pool--namespace) |
| `virtual_server.last_hop_pool.tenant` | [virtual_server.last_hop_pool.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/last_hop_pool/#schema-virtual_server--last_hop_pool--tenant) |
| `virtual_server.last_hop_pool.uid` | [virtual_server.last_hop_pool.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/last_hop_pool/#schema-virtual_server--last_hop_pool--uid) |
| `virtual_server.nat64` | [virtual_server.nat64](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/nat64/#section) |
| `virtual_server.nat64.nat64_disable` | [virtual_server.nat64.nat64_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/nat64/nat64_disable/#section) |
| `virtual_server.nat64.nat64_enable` | [virtual_server.nat64.nat64_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/nat64/nat64_enable/#section) |
| `virtual_server.port_translation` | [virtual_server.port_translation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/port_translation/#section) |
| `virtual_server.port_translation.port_translation_disable` | [virtual_server.port_translation.port_translation_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/port_translation/port_translation_disable/#section) |
| `virtual_server.port_translation.port_translation_enable` | [virtual_server.port_translation.port_translation_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/port_translation/port_translation_enable/#section) |
| `virtual_server.request_logging_profile` | [virtual_server.request_logging_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/request_logging_profile/#section) |
| `virtual_server.request_logging_profile.kind` | [virtual_server.request_logging_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/request_logging_profile/#schema-virtual_server--request_logging_profile--kind) |
| `virtual_server.request_logging_profile.name` | [virtual_server.request_logging_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/request_logging_profile/#schema-virtual_server--request_logging_profile--name) |
| `virtual_server.request_logging_profile.namespace` | [virtual_server.request_logging_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/request_logging_profile/#schema-virtual_server--request_logging_profile--namespace) |
| `virtual_server.request_logging_profile.tenant` | [virtual_server.request_logging_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/request_logging_profile/#schema-virtual_server--request_logging_profile--tenant) |
| `virtual_server.request_logging_profile.uid` | [virtual_server.request_logging_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/request_logging_profile/#schema-virtual_server--request_logging_profile--uid) |
| `virtual_server.source_port` | [virtual_server.source_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/source_port/#section) |
| `virtual_server.source_port.source_port_change` | [virtual_server.source_port.source_port_change](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/source_port/source_port_change/#section) |
| `virtual_server.source_port.source_port_preserve` | [virtual_server.source_port.source_port_preserve](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/source_port/source_port_preserve/#section) |
| `virtual_server.source_port.source_port_preserve_strict` | [virtual_server.source_port.source_port_preserve_strict](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/source_port/source_port_preserve_strict/#section) |
| `virtual_server.statistics_profile` | [virtual_server.statistics_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/statistics_profile/#section) |
| `virtual_server.statistics_profile.kind` | [virtual_server.statistics_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/statistics_profile/#schema-virtual_server--statistics_profile--kind) |
| `virtual_server.statistics_profile.name` | [virtual_server.statistics_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/statistics_profile/#schema-virtual_server--statistics_profile--name) |
| `virtual_server.statistics_profile.namespace` | [virtual_server.statistics_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/statistics_profile/#schema-virtual_server--statistics_profile--namespace) |
| `virtual_server.statistics_profile.tenant` | [virtual_server.statistics_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/statistics_profile/#schema-virtual_server--statistics_profile--tenant) |
| `virtual_server.statistics_profile.uid` | [virtual_server.statistics_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/statistics_profile/#schema-virtual_server--statistics_profile--uid) |
| `virtual_server.tcp` | [virtual_server.tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/#section) |
| `virtual_server.tcp.client_ssl_profile` | [virtual_server.tcp.client_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/client_ssl_profile/#section) |
| `virtual_server.tcp.client_ssl_profile.kind` | [virtual_server.tcp.client_ssl_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/client_ssl_profile/#schema-virtual_server--tcp--client_ssl_profile--kind) |
| `virtual_server.tcp.client_ssl_profile.name` | [virtual_server.tcp.client_ssl_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/client_ssl_profile/#schema-virtual_server--tcp--client_ssl_profile--name) |
| `virtual_server.tcp.client_ssl_profile.namespace` | [virtual_server.tcp.client_ssl_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/client_ssl_profile/#schema-virtual_server--tcp--client_ssl_profile--namespace) |
| `virtual_server.tcp.client_ssl_profile.tenant` | [virtual_server.tcp.client_ssl_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/client_ssl_profile/#schema-virtual_server--tcp--client_ssl_profile--tenant) |
| `virtual_server.tcp.client_ssl_profile.uid` | [virtual_server.tcp.client_ssl_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/client_ssl_profile/#schema-virtual_server--tcp--client_ssl_profile--uid) |
| `virtual_server.tcp.ocsp_profile` | [virtual_server.tcp.ocsp_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/ocsp_profile/#section) |
| `virtual_server.tcp.ocsp_profile.kind` | [virtual_server.tcp.ocsp_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/ocsp_profile/#schema-virtual_server--tcp--ocsp_profile--kind) |
| `virtual_server.tcp.ocsp_profile.name` | [virtual_server.tcp.ocsp_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/ocsp_profile/#schema-virtual_server--tcp--ocsp_profile--name) |
| `virtual_server.tcp.ocsp_profile.namespace` | [virtual_server.tcp.ocsp_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/ocsp_profile/#schema-virtual_server--tcp--ocsp_profile--namespace) |
| `virtual_server.tcp.ocsp_profile.tenant` | [virtual_server.tcp.ocsp_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/ocsp_profile/#schema-virtual_server--tcp--ocsp_profile--tenant) |
| `virtual_server.tcp.ocsp_profile.uid` | [virtual_server.tcp.ocsp_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/ocsp_profile/#schema-virtual_server--tcp--ocsp_profile--uid) |
| `virtual_server.tcp.server_ssl_profile` | [virtual_server.tcp.server_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/server_ssl_profile/#section) |
| `virtual_server.tcp.server_ssl_profile.kind` | [virtual_server.tcp.server_ssl_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/server_ssl_profile/#schema-virtual_server--tcp--server_ssl_profile--kind) |
| `virtual_server.tcp.server_ssl_profile.name` | [virtual_server.tcp.server_ssl_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/server_ssl_profile/#schema-virtual_server--tcp--server_ssl_profile--name) |
| `virtual_server.tcp.server_ssl_profile.namespace` | [virtual_server.tcp.server_ssl_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/server_ssl_profile/#schema-virtual_server--tcp--server_ssl_profile--namespace) |
| `virtual_server.tcp.server_ssl_profile.tenant` | [virtual_server.tcp.server_ssl_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/server_ssl_profile/#schema-virtual_server--tcp--server_ssl_profile--tenant) |
| `virtual_server.tcp.server_ssl_profile.uid` | [virtual_server.tcp.server_ssl_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/server_ssl_profile/#schema-virtual_server--tcp--server_ssl_profile--uid) |
| `virtual_server.tcp.tcp_client_profile` | [virtual_server.tcp.tcp_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/tcp_client_profile/#section) |
| `virtual_server.tcp.tcp_client_profile.kind` | [virtual_server.tcp.tcp_client_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/tcp_client_profile/#schema-virtual_server--tcp--tcp_client_profile--kind) |
| `virtual_server.tcp.tcp_client_profile.name` | [virtual_server.tcp.tcp_client_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/tcp_client_profile/#schema-virtual_server--tcp--tcp_client_profile--name) |
| `virtual_server.tcp.tcp_client_profile.namespace` | [virtual_server.tcp.tcp_client_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/tcp_client_profile/#schema-virtual_server--tcp--tcp_client_profile--namespace) |
| `virtual_server.tcp.tcp_client_profile.tenant` | [virtual_server.tcp.tcp_client_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/tcp_client_profile/#schema-virtual_server--tcp--tcp_client_profile--tenant) |
| `virtual_server.tcp.tcp_client_profile.uid` | [virtual_server.tcp.tcp_client_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/tcp_client_profile/#schema-virtual_server--tcp--tcp_client_profile--uid) |
| `virtual_server.tcp.tcp_server_profile` | [virtual_server.tcp.tcp_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/tcp_server_profile/#section) |
| `virtual_server.tcp.tcp_server_profile.kind` | [virtual_server.tcp.tcp_server_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/tcp_server_profile/#schema-virtual_server--tcp--tcp_server_profile--kind) |
| `virtual_server.tcp.tcp_server_profile.name` | [virtual_server.tcp.tcp_server_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/tcp_server_profile/#schema-virtual_server--tcp--tcp_server_profile--name) |
| `virtual_server.tcp.tcp_server_profile.namespace` | [virtual_server.tcp.tcp_server_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/tcp_server_profile/#schema-virtual_server--tcp--tcp_server_profile--namespace) |
| `virtual_server.tcp.tcp_server_profile.tenant` | [virtual_server.tcp.tcp_server_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/tcp_server_profile/#schema-virtual_server--tcp--tcp_server_profile--tenant) |
| `virtual_server.tcp.tcp_server_profile.uid` | [virtual_server.tcp.tcp_server_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/tcp/tcp_server_profile/#schema-virtual_server--tcp--tcp_server_profile--uid) |
| `virtual_server.udp` | [virtual_server.udp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/#section) |
| `virtual_server.udp.client_ssl_profile` | [virtual_server.udp.client_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/client_ssl_profile/#section) |
| `virtual_server.udp.client_ssl_profile.kind` | [virtual_server.udp.client_ssl_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/client_ssl_profile/#schema-virtual_server--udp--client_ssl_profile--kind) |
| `virtual_server.udp.client_ssl_profile.name` | [virtual_server.udp.client_ssl_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/client_ssl_profile/#schema-virtual_server--udp--client_ssl_profile--name) |
| `virtual_server.udp.client_ssl_profile.namespace` | [virtual_server.udp.client_ssl_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/client_ssl_profile/#schema-virtual_server--udp--client_ssl_profile--namespace) |
| `virtual_server.udp.client_ssl_profile.tenant` | [virtual_server.udp.client_ssl_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/client_ssl_profile/#schema-virtual_server--udp--client_ssl_profile--tenant) |
| `virtual_server.udp.client_ssl_profile.uid` | [virtual_server.udp.client_ssl_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/client_ssl_profile/#schema-virtual_server--udp--client_ssl_profile--uid) |
| `virtual_server.udp.server_ssl_profile` | [virtual_server.udp.server_ssl_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/server_ssl_profile/#section) |
| `virtual_server.udp.server_ssl_profile.kind` | [virtual_server.udp.server_ssl_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/server_ssl_profile/#schema-virtual_server--udp--server_ssl_profile--kind) |
| `virtual_server.udp.server_ssl_profile.name` | [virtual_server.udp.server_ssl_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/server_ssl_profile/#schema-virtual_server--udp--server_ssl_profile--name) |
| `virtual_server.udp.server_ssl_profile.namespace` | [virtual_server.udp.server_ssl_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/server_ssl_profile/#schema-virtual_server--udp--server_ssl_profile--namespace) |
| `virtual_server.udp.server_ssl_profile.tenant` | [virtual_server.udp.server_ssl_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/server_ssl_profile/#schema-virtual_server--udp--server_ssl_profile--tenant) |
| `virtual_server.udp.server_ssl_profile.uid` | [virtual_server.udp.server_ssl_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/server_ssl_profile/#schema-virtual_server--udp--server_ssl_profile--uid) |
| `virtual_server.udp.udp_client_profile` | [virtual_server.udp.udp_client_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/udp_client_profile/#section) |
| `virtual_server.udp.udp_client_profile.kind` | [virtual_server.udp.udp_client_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/udp_client_profile/#schema-virtual_server--udp--udp_client_profile--kind) |
| `virtual_server.udp.udp_client_profile.name` | [virtual_server.udp.udp_client_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/udp_client_profile/#schema-virtual_server--udp--udp_client_profile--name) |
| `virtual_server.udp.udp_client_profile.namespace` | [virtual_server.udp.udp_client_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/udp_client_profile/#schema-virtual_server--udp--udp_client_profile--namespace) |
| `virtual_server.udp.udp_client_profile.tenant` | [virtual_server.udp.udp_client_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/udp_client_profile/#schema-virtual_server--udp--udp_client_profile--tenant) |
| `virtual_server.udp.udp_client_profile.uid` | [virtual_server.udp.udp_client_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/udp_client_profile/#schema-virtual_server--udp--udp_client_profile--uid) |
| `virtual_server.udp.udp_server_profile` | [virtual_server.udp.udp_server_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/udp_server_profile/#section) |
| `virtual_server.udp.udp_server_profile.kind` | [virtual_server.udp.udp_server_profile.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/udp_server_profile/#schema-virtual_server--udp--udp_server_profile--kind) |
| `virtual_server.udp.udp_server_profile.name` | [virtual_server.udp.udp_server_profile.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/udp_server_profile/#schema-virtual_server--udp--udp_server_profile--name) |
| `virtual_server.udp.udp_server_profile.namespace` | [virtual_server.udp.udp_server_profile.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/udp_server_profile/#schema-virtual_server--udp--udp_server_profile--namespace) |
| `virtual_server.udp.udp_server_profile.tenant` | [virtual_server.udp.udp_server_profile.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/udp_server_profile/#schema-virtual_server--udp--udp_server_profile--tenant) |
| `virtual_server.udp.udp_server_profile.uid` | [virtual_server.udp.udp_server_profile.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/udp/udp_server_profile/#schema-virtual_server--udp--udp_server_profile--uid) |
| `virtual_server.virtual_server_state` | [virtual_server.virtual_server_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/virtual_server_state/#section) |
| `virtual_server.virtual_server_state.state_disabled` | [virtual_server.virtual_server_state.state_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/virtual_server_state/state_disabled/#section) |
| `virtual_server.virtual_server_state.state_enabled` | [virtual_server.virtual_server_state.state_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/virtual_server_state/state_enabled/#section) |
| `virtual_server.vs_score` | [virtual_server.vs_score](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/#schema-virtual_server--vs_score) |

## Next pages

- [advanced_tcp_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/advanced_tcp_profile/)
- [ddos_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/ddos_profile/)
- [irules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/irules/)
- [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/properties/virtual_server/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/application_profiles/)
